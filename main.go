package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/akamensky/argparse"
	"gitlab.com/golang-commonmark/markdown"
	"go/build"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"strings"
	"text/template"
)

func EncodeBase64(src string) string {
	f, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	var out strings.Builder
	encoder := base64.NewEncoder(base64.StdEncoding, &out)
	encoder.Write(f)
	encoder.Close()

	extParts := strings.Split(src, ".")
	ext := "." + extParts[1]

	return fmt.Sprintf("data:%v;base64,%s", mime.TypeByExtension(ext), out.String())
}

func Base64Images(document string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(document))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		s.Each(func(i int, sel *goquery.Selection) {
			src, _ := sel.Attr("src")
			src = EncodeBase64(src)
			sel.SetAttr("src", src)
		})
	})
	var b bytes.Buffer
	html.Render(&b, doc.Selection.Nodes[0])
	return b.String()
}

func ParseStyle(s, base string) string {
	_, err := os.Open(s)
	if err != nil {
		path := base + "styles/" + s + ".css"
		_, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		} else {
			return path
		}
	} else {
		return s
	}
	return ""
}

func main() {

	type Data struct {
		Slides []string
		Style  string
		Script string
	}

	base := build.Default.GOPATH + "/src/github.com/skuzzymiglet/point/"

	parser := argparse.NewParser("point", "html presentation tool")
	inFile := parser.String("i", "in", &argparse.Options{Required: true, Help: "Input file"})
	style := parser.String("s", "style", &argparse.Options{Required: false, Default: "", Help: "Stylesheet"})
	outFile := parser.String("o", "out", &argparse.Options{Required: false, Help: "Output file. Default: stdout"})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(base, *style)

	inBytes, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatal(err)
	}

	var styleBytes []byte
	fmt.Println(*style == "")
	if *style == "" {
		styleBytes, err = ioutil.ReadFile(base + "style.css")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		styleBytes, err = ioutil.ReadFile(ParseStyle(*style, base))
		if err != nil {
			log.Fatal(err)
		}
	}

	script, err := ioutil.ReadFile(base + "script.js")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(styleBytes))

	in := string(inBytes)

	inParts := strings.Split(in, "\n\n\n")
	//fmt.Printf("%#v\n", inParts)

	md := markdown.New(markdown.XHTMLOutput(true))

	htmlParts := make([]string, 0)

	for _, e := range inParts {
		htmlParts = append(htmlParts, md.RenderToString([]byte(e)))
	}

	for i, e := range htmlParts {
		htmlParts[i] = Base64Images(e)
	}

	var tl *template.Template
	tl = template.Must(template.ParseFiles(base + "template.html"))
	var out io.Writer
	if *outFile == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(*outFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tl.Execute(out, Data{htmlParts, string(styleBytes), string(script)})
	if err != nil {
		log.Fatal(err)
	}
}
