
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/liuzl/ling"
)

var (
	input  = flag.String("input", "", "file of original text to read")
	output = flag.String("output", "", "file of tokenized text to write")
	typ    = flag.String("type", "token", "type: token, span, all")
	full   = flag.Bool("full", false, "full output")
)

func main() {
	flag.Parse()
	var err error
	var in, out *os.File //io.Reader
	var nlp *ling.Pipeline
	if *input == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer in.Close()
	}
	br := bufio.NewReader(in)

	if *output == "" {
		out = os.Stdout
	} else {
		out, err = os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	}
	if nlp, err = ling.DefaultNLP(); err != nil {
		log.Fatal(err)
	}
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		if c != nil {
			log.Fatal(c)
		}
		d := ling.NewDocument(strings.TrimSpace(line))
		if err = nlp.Annotate(d); err != nil {
			log.Fatal(err)
		}
		if *typ == "token" || *typ == "all" {
			var ret []string
			for _, t := range d.Tokens {
				if t.Type == ling.Space {
					continue
				}
				if *full {
					ret = append(ret, t.String()+"\t"+t.Type.String())
				} else {
					ret = append(ret, t.String())
				}
			}
			fmt.Fprintf(out, "tokens:\n==========\n%s\n", strings.Join(ret, "\n"))
		}
		if *typ == "span" || *typ == "all" {
			var ret []string
			for _, s := range d.Spans {
				if *full {
					ret = append(ret,
						s.String()+"\t"+fmt.Sprintf("%+v", s.Annotations))
				} else {
					ret = append(ret, s.String())
				}
			}
			fmt.Fprintf(out, "spans:\n==========\n%s\n", strings.Join(ret, "\n"))
		}
		fmt.Fprintf(out, "\n")
	}
}