package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type node struct {
	SubNodes []node
	Info     os.FileInfo
}

func (n node) Size() string {
	if n.Info.Size() > 0 {
		return fmt.Sprintf("%db", n.Info.Size())
	} else {
		return "empty"
	}
}

func (n node) Name() string {
	if n.Info.IsDir() {
		return n.Info.Name()
	} else {
		return fmt.Sprintf("%s (%s)", n.Info.Name(), n.Size())
	}
}

func get(path string, containsFiles bool) ([]node, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var nodes []node
	for _, file := range files {
		if !containsFiles && !file.IsDir() {
			continue
		}

		n := node{Info: file}

		if file.IsDir() {
			// recursively check subfiles
			subnodes, err := get(path + string(os.PathSeparator)+file.Name(), containsFiles)
			if err != nil {
				return nil, err
			}

			n.SubNodes = subnodes
		}

		nodes = append(nodes, n)
	}

	return nodes, nil
}

func print(out io.Writer, nodes []node, upLevelPrefix string) {
	var (
		prefix         = "├───"
		lowLevelPrefix = "│\t"
	)

	for i, n := range nodes {
		if i == len(nodes) - 1 {
			prefix = "└───"
			lowLevelPrefix = "\t"
		}

		fmt.Fprint(out, upLevelPrefix, prefix, n.Name(), "\n")

		if n.Info.IsDir() {
			print(out, n.SubNodes, upLevelPrefix+lowLevelPrefix)
		}
	}
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) (err error) {
	nodes, err := get(path, printFiles)
	if err != nil {
		return
	}

	print(out, nodes, "")
	return
}
