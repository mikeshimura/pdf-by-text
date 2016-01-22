package main

import (
	"flag"
	"fmt"
	"github.com/mikeshimura/pdf-by-text/util"
	"os"
)

func main() {
	var fontfile string
	var outfile string
	var encoding string
	flag.StringVar(&fontfile, "f", "", "fonts file path default font.yml")
	flag.StringVar(&outfile, "o", "", "output file path")
	flag.StringVar(&encoding, "e", "", "encoding default UTF-8, accept ShiftJIS, EUCJP")
	flag.Parse()
	//fmt.Printf("f:%v o:%v\n", fontfile, outfile)
	//fmt.Printf("args %v\n", flag.Args())
	if len(flag.Args())==0{
		fmt.Fprintf(os.Stderr,"Input file not specified\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	fonts := util.ReadFont(fontfile)
	util.Execute(fonts,flag.Args()[0],outfile,encoding)
}
