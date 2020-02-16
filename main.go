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
	"regexp"
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

func SplitAfter(s string, re *regexp.Regexp) []string {
	var (
		r []string
		p int
	)
	is := re.FindAllStringIndex(s, -1)
	if is == nil {
		return append(r, s)
	}
	for index := range is {
		if index == len(is)-1 {
			r = append(r, s[p:])
		} else {
			r = append(r, s[p:is[index+1][0]])
			p = is[index+1][0]
		}

	}
	return append(r, s[p:])
}

func Slides(m []byte) []string {
	md := markdown.New(markdown.XHTMLOutput(true))
	rendered := md.RenderToString(m)
	rendered = Base64Images(rendered)
	splitExpr := regexp.MustCompile("<h1>.+</h1>")
	return SplitAfter(rendered, splitExpr)
}

func main() {

	type Data struct {
		Slides []string
		Style  string
		Script string
	}

	// Where our files are
	base := build.Default.GOPATH + "/src/github.com/skuzzymiglet/point/"

	// Parse args
	parser := argparse.NewParser("point", "html presentation tool")
	inFile := parser.String("i", "in", &argparse.Options{Required: true, Help: "Input file"})
	style := parser.String("s", "style", &argparse.Options{Required: false, Default: "", Help: "Stylesheet"})
	outFile := parser.String("o", "out", &argparse.Options{Required: false, Help: "Output file. Default: stdout"})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// Read files
	inBytes, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatal(err)
	}

	script, err := ioutil.ReadFile(base + "script.js")
	if err != nil {
		log.Fatal(err)
	}

	// Read style
	var styleBytes []byte
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

	// Render markdown, embed images and split into slides
	parts := Slides(inBytes)

	tl := template.Must(template.ParseFiles(base + "template.html"))

	// Output
	var out io.Writer
	if *outFile == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(*outFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Insert slides, style and script into template
	err = tl.Execute(out, Data{parts, string(styleBytes), string(script)})
	if err != nil {
		log.Fatal(err)
	}
}
