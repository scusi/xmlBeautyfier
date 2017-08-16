// xmlBeautyfier - prints an xml file indented
//
package main

import (
	"flag"
	"fmt"
	"github.com/clbanning/mxj"
	"io/ioutil"
	"log"
)

var prefix string
var indent string

func init() {
	flag.StringVar(&prefix, "p", "", "prefix to be used; when set to FILENAME it uses the current filename as prefix")
	flag.StringVar(&indent, "i", "  ", "indent to be used")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	xmlFile := flag.Arg(0)
	if prefix == "FILENAME" {
		prefix = xmlFile
	}

	data, err := readXMLFile(xmlFile)
	check(err)

	data, err = remarshalIndentXML(data)
	check(err)

	// printout
	fmt.Printf("%s\n", data)
}

func readXMLFile(fileName string) (data []byte, err error) {
	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		return data, err
	}
	return data, err
}

func remarshalIndentXML(in []byte) (out []byte, err error) {
	// Unmarshall
	mv, err := mxj.NewMapXml(in)
	if err != nil {
		return out, err
	}
	// Marshall with Indent
	out, err = mv.XmlIndent(prefix, indent)
	if err != nil {
		return out, err
	}
	return out, nil
}
