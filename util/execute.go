package util

import (
	"fmt"
	gr "github.com/mikeshimura/goreport"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"io"
)
func transformEncoding( rawReader io.Reader, trans transform.Transformer) (string, error) {
    ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
    if err == nil {
        return string(ret), nil
    } else {
        return "", err
    }
}
func FromShiftJIS(str string) (string, error) {
    return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
}
func FromEUCJP(str string) (string, error) {
    return transformEncoding(strings.NewReader(str), japanese.EUCJP.NewDecoder())
}
func Execute(fonts []*gr.FontMap, infile string, outfile string, encoding string) {
	buf, err := ioutil.ReadFile(infile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input file:"+infile+" not found")
		os.Exit(2)
	}
	bufs:=string(buf)
	switch encoding {
	case "":
	
	case "ShiftJIS":bufs,err=FromShiftJIS(bufs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ShiftJIS Convert Error")
		os.Exit(2)
	}
	case "EUCJP":bufs,err=FromEUCJP(bufs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "EUCJP Convert Error")
		os.Exit(2)
	}
	default:
		fmt.Fprintf(os.Stderr, "encoding:"+encoding+
			" is illegal. ShiftJIS or EUCJP.")
		os.Exit(2)
	}

	//fmt.Printf("bufs %v\n",bufs)
	pos:=strings.Index(bufs,"\r\n")
	//fmt.Printf("pos %v\n",pos)
	if pos > -1{
		bufs=strings.Replace(bufs,"\r\n","\n",-1)
	}
	if len(outfile) == 0 {
		idx := strings.LastIndex(infile, ".")
		outfile = infile[0:idx] + ".pdf"
		//fmt.Printf("outfile :%v\n",outfile)
	}
	c := new(gr.Converter)
	c.Fonts = fonts
	c.Text = bufs
	c.Execute()
	c.Pdf.WritePdf(outfile)
	fmt.Printf("Pdf " + outfile + " generated")
}
func ReadFont(path string) []*gr.FontMap {
	if path == "" {
		path = "font.yml"
	}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "font file:"+path+" not found")
		os.Exit(2)
	}
	//fmt.Printf("%v\n", string(buf))
	m := make([]string, 0)
	err = yaml.Unmarshal(buf, &m)
	fmap := make([]*gr.FontMap, 0)
	for _, f := range m {
		pair := strings.Split(f, " ")
		font := new(gr.FontMap)
		for _, fs := range pair {
			fontpara := strings.Split(fs, ":")
			switch fontpara[0] {
			case "name":
				font.FontName = fontpara[1]

			case "file":
				font.FileName = fontpara[1]
			default:
				fmt.Fprintf(os.Stderr, "font parameter :"+fontpara[0]+" illegal")
				os.Exit(2)
			}
			fmap = append(fmap, font)
		}
	}
	return fmap
}
