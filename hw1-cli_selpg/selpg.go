package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	startpg, endpg, lineCountPg int
	findNewPageSign, readfile   bool
	destlp, filename            string
	warning                     *log.Logger
	inputText                   []string
	outputText                  string
)

func init() {
	warning = log.New(os.Stderr, "[Warning:]", log.Ldate|log.Ltime|log.Lshortfile)
	inputText = make([]string, 4, 4)

	flag.IntVar(&lineCountPg, "l", 72, "specify line count per page. default to 72. if -f is used, this val will be ignored.")
	flag.IntVar(&startpg, "s", -1, "must specify start page number, should be greater than 0.")
	flag.IntVar(&endpg, "e", -1, "must specify end page number, should be greater than start number.")
	flag.BoolVar(&findNewPageSign, "f", false, "set -f to let selpg find [new page mark] from input, if -f is used, -l will be ignored")
	flag.StringVar(&destlp, "d", "", "specify the destination printer to print.")
	flag.Parse()
}

func parse() {
	/**
	 * fmt.Printf("[test arguments]:\n")
	 * fmt.Printf("-l num is %v\n", lineCountPg)
	 * fmt.Printf("-f is %v\n", findNewPageSign)
	 * fmt.Printf("-s = %v\n", startpg)
	 * fmt.Printf("-e = %v\n", endpg)
	 * fmt.Printf("-d = \"%+v\"\n", destlp)
	 */
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

}

func run() {
	if readfile {
		inputFile, inputErr := os.Open(filename)
		if inputErr != nil {
			log.Fatal("An error occurred on opening the inputFile\nCheck if the file exists and access.\n")
		}
		defer inputFile.Close()

		inputReader := bufio.NewReader(inputFile)
		for {
			inputString, readerErr := inputReader.ReadString('\n')
			// fmt.Printf("line: %s", inputString)
			inputText = append(inputText, inputString)
			if readerErr == io.EOF {
				break
			}
		}
	} else {
		inputReader := bufio.NewReader(os.Stdin)
		for {
			inputString, inputErr := inputReader.ReadString('\n')
			if inputErr == io.EOF {
				break
			} else if inputErr != nil {
				panic(inputErr)
			}
			inputText = append(inputText, inputString)
		}
	}
	// select page
	if findNewPageSign {

	} else {
		if (startpg-1)*lineCountPg > len(inputText)-1 {
			log.Fatal("Start page number out of range")
		}
		if (endpg-1)*lineCountPg > len(inputText)-1 {
			log.Fatal("End page number out of range")
		}
		startLineIdx := (startpg - 1) * lineCountPg
		endLineIdx := endpg * lineCountPg
		if endLineIdx > len(inputText) {
			endLineIdx = len(inputText)
		}
		fmt.Printf("startIdx: %v, endIdx: %v, len(inputText): %v\n", startLineIdx, endLineIdx, len(inputText))

		for i, line := range inputText[startLineIdx:endLineIdx] {
			fmt.Printf("%v,", i)
			_, err := fmt.Println(line)
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	parse()
	run()
}
