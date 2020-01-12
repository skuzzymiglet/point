package main

import (
	//"fmt"
	"github.com/akamensky/argparse"
	"gitlab.com/golang-commonmark/markdown"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {

	parser := argparse.NewParser("point", "html presentation tool")
	inFile := parser.String("i", "in", &argparse.Options{Required: true, Help: "Input file"})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	inBytes, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatal(err)
	}

	in := string(inBytes)

	inParts := strings.Split(in, "\n\n\n")
	//fmt.Printf("%#v\n", inParts)

	md := markdown.New(markdown.XHTMLOutput(true))

	htmlParts := make([]string, 0)

	for _, e := range inParts {
		htmlParts = append(htmlParts, md.RenderToString([]byte(e)))
	}

	//fmt.Printf("%#v\n", htmlParts)

	var tl *template.Template
	tl = template.Must(template.ParseFiles("template.html"))
	err = tl.ExecuteTemplate(os.Stdout, "template.html", htmlParts)

}
