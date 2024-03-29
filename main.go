package main

import (
	"flag"
	"fmt"
	"gopkg.in/flosch/pongo2.v3"
	"io"
	"os"
)

func main() {
	var config = flag.String("c", "", "parameter file path")
	var templatePath = flag.String("t", "", "template file path")
	var outputPath = flag.String("o", "", "render result output file path. if not specified, output to stdout.")
	var mode = flag.String("m", "json", "parameter file mode. json or yaml")
	flag.Parse()

	if *config == "" {
		fmt.Fprintln(os.Stderr, "must be specify -c option")
		os.Exit(1)
	}

	if *templatePath == "" {
		fmt.Fprintln(os.Stderr, "must be specify -t option")
		os.Exit(1)
	}

	var parser ParameterParser
	if *mode == "yaml" {
		parser = NewYamlParser()
	} else {
		parser = NewJsonParser()
	}

	f, err := os.Open(*config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	var context pongo2.Context
	context, err = parser.Parse(f)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	pt, err := pongo2.FromFile(*templatePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	output, err := pt.ExecuteBytes(context)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var outputWriter io.WriteCloser
	if *outputPath == "" {
		outputWriter = os.Stdout
	} else {
		outputWriter, err = os.Create(*outputPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	defer outputWriter.Close()
	if _, err := fmt.Fprintln(outputWriter, string(output)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
