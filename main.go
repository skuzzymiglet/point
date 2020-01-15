package main

import (
	"github.com/akamensky/argparse"
	"gitlab.com/golang-commonmark/markdown"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {

	type Data struct {
		Slides []string
		Style  string
		Script string
	}

	parser := argparse.NewParser("point", "html presentation tool")
	inFile := parser.String("i", "in", &argparse.Options{Required: true, Help: "Input file"})
	style := parser.String("s", "style", &argparse.Options{Required: false, Default: "style.css", Help: "Stylesheet"})
	outFile := parser.String("o", "out", &argparse.Options{Required: false, Help: "Output file"})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	inBytes, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatal(err)
	}

	styleBytes, err := ioutil.ReadFile(*style)
	if err != nil {
		log.Fatal(err)
	}

	script, err := ioutil.ReadFile("script.js")
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
	var out io.Writer
	if *outFile == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(*outFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tl.ExecuteTemplate(out, "template.html", Data{htmlParts, string(styleBytes), string(script)})
}
