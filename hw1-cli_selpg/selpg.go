package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"unicode/utf8"
)

var (
	startpg, endpg, lineCountPg      int
	findNewPageSign, readfile, uselp bool
	destlp, filename                 string
	warning                          *log.Logger
	lpcmd                            *exec.Cmd
	cmdinpipe                        io.WriteCloser
	cmderr                           error
)

func init() {
	warning = log.New(os.Stderr, "[Warning:]", log.Ldate|log.Ltime|log.Lshortfile)

	flag.IntVar(&lineCountPg, "l", 72, "specify line count per page. default to 72. if -f is used, this val will be ignored.")
	flag.IntVar(&startpg, "s", -1, "must specify start page number, should be greater than 0.")
	flag.IntVar(&endpg, "e", -1, "must specify end page number, should be greater than start number.")
	flag.BoolVar(&findNewPageSign, "f", false, "set -f to let selpg find [new page mark] from input, if -f is used, -l will be ignored")
	flag.StringVar(&destlp, "d", "", "specify the destination printer to print.")
	flag.Parse()
}

func parse() {
	if startpg < 0 || endpg < 0 || startpg > endpg {
		log.Fatalln(errors.New("Start Page Num must >= End Page Num, and should both > 0"))
	}
	if lineCountPg != 72 && findNewPageSign == true {
		warning.Println("-f is specified. [-l number] will be ignored.")
	}

	switch {
	case len(flag.Args()) > 1:
		warning.Println("cannot receive multi-file now. only the first filename will be received.")
		fallthrough
	case len(flag.Args()) == 1:
		// process input file
		readfile = true
		filename = os.ExpandEnv(flag.Args()[0])
		pwd, err := os.Getwd()
		if err != nil {
			filename = pwd + filename
		}
	case len(flag.Args()) == 0:
		// process stdin
		readfile = false
	}

	if destlp != "" {
		uselp = true
		lpcmd = exec.Command("lp", "-d", destlp)
		lpcmd.Stdout = os.Stdout
		lpcmd.Stderr = os.Stderr
		cmdinpipe, cmderr = lpcmd.StdinPipe()
		lpcmd.Start()
	}
}

func run() {
	var reader *bufio.Reader
	if readfile {
		inputFile, inputErr := os.Open(filename)
		if inputErr != nil {
			log.Fatal("An error occurred on opening the inputFile\nCheck if the file exists and access.\n")
		}
		defer inputFile.Close()
		reader = bufio.NewReader(inputFile)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	pagectr := 1
	linectr := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if uselp {
				cmdinpipe.Close()
				lpcmd.Wait()
			}
			break
		} else if err != nil {
			panic(err)
		}
		if findNewPageSign {
			strbyte := []byte(line)
			for len(strbyte) > 0 {
				r, n := utf8.DecodeRune(strbyte)
				if r == '\f' {
					pagectr++
				}
				if pagectr >= startpg && pagectr <= endpg {
					if uselp {
						cmdinpipe.Write([]byte(string(r)))
					} else {
						fmt.Print(string(r))
					}
				}
				strbyte = strbyte[n:]
			}
		} else {
			linectr++
			if linectr > lineCountPg {
				pagectr++
				linectr = 1
			}
			if pagectr >= startpg && pagectr <= endpg {
				if uselp {
					cmdinpipe.Write([]byte(line))
				} else {
					fmt.Print(line)
				}
			}
		}
	}
	if pagectr < startpg {
		warning.Printf("start_page (%v) greater than total pages (%v), no output written.\n", startpg, pagectr)
	} else if pagectr < endpg {
		warning.Printf("end_page (%v) greater than total pages (%d), less output than expected.\n", endpg, pagectr)
	}
}

func main() {
	parse()
	run()
}
