# selpg: Go implement
使用Go实现Linux命令行实用程序`selpg`

<!-- TOC -->

- [selpg: Go implement](#selpg-go-implement)
    - [概述](#概述)
    - [参数说明](#参数说明)
    - [代码解释](#代码解释)
    - [关键实现 & 坑](#关键实现--坑)
        - [flag处理参数](#flag处理参数)
        - [逐行从 文件/stdin 中读取](#逐行从-文件stdin-中读取)
        - [处理读入的行](#处理读入的行)
            - [突然想到的优化：逐字符读取](#突然想到的优化逐字符读取)
        - [环境变量的使用](#环境变量的使用)
        - [在程序内部调用 lp -d 打印](#在程序内部调用-lp--d-打印)

<!-- /TOC -->

## 概述
Go学习第一课，实现命令行实用程序`selpg`。功能是在输入（stdin或文件）的内容中，根据指定`-s [Number]` `-e [Number]`参数，选择指定页码范围内的内容输出，同时可以通过`-d [destPrinter]`将输出内容直接传送到目标打印机进行打印。

## 参数说明
`-s [Number]`: 指定开始页码

`-e [Number]`: 指定结束页码

**PS:** 页码从1开始，并且页码区间为包含开始页码和结束页码。

`-l [Number]`: 指定每页包含多少行，默认值为72

`-f`: 忽略`-l`选项，令程序在输入内容中寻找换行符(`\f`)，并作为页面分割的标志。

`-d [destPrinter]`: 指定目标打印机打印，程序将通过调用`lp -d[destPrinter]`将选择后的内容传至打印机

`[input_filename]`: 输入文件名。在所有`-`类型参数后指定。指定后`selpg`将查找并读取该文件。若不指定，则读取`stdin`中的内容作为输入。

## 代码解释
共3个主要函数。

`init()`: 用于初始化变量。go程序运行时会在`main()`之前先运行`init()`。

`parse()`: 处理输入参数。处理后的结果保存在全局变量中。

这里主要使用了官方提供的`flag`包。美中不足的是，`flag`包只支持3种类型的指定参数：
- `-x value`
- `-x=value`
- `-x` （仅`bool`类型）
其中`x`可以是单个或多个字母，而并没有我们常用的`-xValue` `--longParam`这种形式。

`run()`: 开始处理。

## 关键实现 & 坑

### flag处理参数
`flag`包用于处理命令行参数，api简单易用。这里全部使用了`flag.TypeVar()`风格的api。

```go
flag.IntVar(&lineCountPg, "l", 72, "specify line count per page. default to 72. if -f is used, this val will be ignored.")
flag.IntVar(&startpg, "s", -1, "must specify start page number, should be greater than 0.")
flag.IntVar(&endpg, "e", -1, "must specify end page number, should be greater than start number.")
flag.BoolVar(&findNewPageSign, "f", false, "set -f to let selpg find [new page mark] from input, if -f is used, -l will be ignored")
flag.StringVar(&destlp, "d", "", "specify the destination printer to print.")
flag.Parse()
```

[这里](http://blog.studygolang.com/2013/02/%E6%A0%87%E5%87%86%E5%BA%93-%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%8F%82%E6%95%B0%E8%A7%A3%E6%9E%90flag/)是一篇不错的介绍`flag`包用法的博客，仔细看一下代码就很容易明白。

### 逐行从 文件/stdin 中读取

**思考这部分代码用了我整个作业60%的时间**，最终想到了一个比较简短而且不容易出错的方案。

首先是读取方式的选择。`Go`处理IO的方式有很多，在按行读字符串上，主要有`bufio.Reader.ReadString()`, `bufio.Reader.ReadBytes()`, `bufio.Reader.ReadLine()`和`bufio.Scanner.Scan()`。其中，`ReadBytes()`和`ReadLine()`较为底层，需要手动处理buffer和拼接，并不推荐。

官方推荐的是`bufio.Scanner.Scan()`但谷歌了一番，发现它存在buffer最大值只有65536的限制，如果读取的一行长度超过了buffer限制就会报错。

这里采用的是`bufio.Reader.ReadString('\n')`。其中`\n`是指定以`\n`为换行符来读取每一行。`ReadString()`的内部实现是通过多次调用`ReadBytes()`拼接字符串，然后通过检查换行符来返回每一行。【不会出现`Scan()`的buffer限制问题。】【本句待确认】

PS: `Scan()`读入的行末尾没有换行符(`\n`)，`ReadString('\n')`读入的行末尾有换行符，所以后者更方便处理两种不同的情况（指定`-f`和不指定`-f`）。在指定`-f`时，只需要对读入的字符（包括换行符）逐个输出，遇到分页符(`\f`)就分页；不指定`-f`时，只需要使用`fmt.Print(line)`来直接将带有换行符的一行输出，无需对行末做额外处理。

```go
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
```

### 处理读入的行
这里的逻辑借鉴了`selpg.c`里的过程，不过做了一点改动，让代码短了一些。我是直接先用循环读入，然后在循环体内部判断是否指定了`-f`，和是否用`-d destPrinter`指定了打印机。

直观来看，这样的写法会比原来的写法每一行都多判断一次`-f`和`-d`，但是实际上每次if的判断结果都是一样的。那么在读取过程足够长（输入流足够大）的情况下，分支预测对这两个if的预测正确率可以无限接近`100%`，所以无须担心。

```go
// 这段代码省略了for循环的前一部分，所以直接复制会报错，复制请直接看文件
for {
    // ... : read each line
    // if EOF break
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
```

#### 突然想到的优化：逐字符读取
原来版本中，我是逐行读取，在指定`-f`的模式下，对每行的字符进行UTF-8的解析，在`strbyte = strbyte[n:]`这里应该会花很多的时间，而且似乎会造成流水线阻塞。然后发现有`bufio.Reader.ReadRune()`这个函数，可以逐字符读取。逐个读取并处理，相比之前每次处理完字符需要额外创建一个切片的操作，要快上很多。

PS: 流水线气泡的细节还是需要查阅《CSAPP》

修改后的代码如下。因为使用了`ReadRune()`，输出的代码只需要一份。所以代码也变短了。

```go
pagectr := 1
linectr := 0
for {
	ch, _, err := reader.ReadRune()
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
		if ch == '\f' {
			pagectr++
		}
	} else {
		if ch == '\n' {
			linectr++
			if linectr > lineCountPg {
				pagectr++
				linectr = 1
			}
		}
	}
	if pagectr >= startpg && pagectr <= endpg {
		if uselp {
			cmdinpipe.Write([]byte(string(ch)))
		} else {
			fmt.Print(string(ch))
		}
	}
}
```

### 环境变量的使用
Go语言处理环境变量是用到`os`包。`os.ExpandEnv()`即可对含有环境变量的字符串进行解析，将环境变量替换为实际地址。

PS: os包的解析只支持`$envName`的形式，并不支持Windows下的`%envName%`。【在Windows下用环境变量指定时，依然要使用`$envName`的形式。】【本句待确认】

```go
// process input file
readfile = true
filename = os.ExpandEnv(flag.Args()[0])
pwd, err := os.Getwd()
if err != nil {
	filename = pwd + filename
}
```

### 在程序内部调用 lp -d 打印

这里是使用了`os/exec`包中的`exec.Cmd.StdInPipe()`。在指定`-d`时，用此将标准输出重定向到`lp -d`的输入当中，实现调用。

这里需要注意的是，需要手动将子进程的`Stdout`和`Stderr`指定为`os.Stdout`和`os.Stderr`，这样调用`lp -d`的结果才能在终端中显示。
```go
// get the StdInPipe()
if destlp != "" {
	uselp = true
	lpcmd = exec.Command("lp", "-d", destlp)
	lpcmd.Stdout = os.Stdout // 使lp的结果在终端中显示
	lpcmd.Stderr = os.Stderr // 同上
	cmdinpipe, cmderr = lpcmd.StdinPipe()
	lpcmd.Start()
}
// ...

// output into Pipe
if uselp {
	cmdinpipe.Write([]byte(line))
} else {
	fmt.Print(line)
}
```