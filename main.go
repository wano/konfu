package main

import (
	"encoding/json"
	"fmt"
	"flag"
	"gopkg.in/flosch/pongo2.v3"
	"io/ioutil"
	"os"
)

func main() {
	var config = flag.String("c", "", "json parameter file path")
	var templatePath = flag.String("t", "", "template file path")
	var outputPath = flag.String("o", "", "render result output file path")
	flag.Parse()

	if *config == "" {
		fmt.Fprintln(os.Stderr, "must be specify -c option")
		os.Exit(1)
	}

	if *templatePath == "" {
		fmt.Fprintln(os.Stderr, "must be specify -t option")
		os.Exit(1)
	}


	buf, err := ioutil.ReadFile(*config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var context pongo2.Context
	if err := json.Unmarshal(buf, &context); err != nil {
		fmt.Fprintln(os.Stderr, err)
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

	if *outputPath == "" {
		fmt.Println(string(output))
	} else {
		if err := ioutil.WriteFile(*outputPath, output, 0664); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}


}
