package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

type ByteUnit int

const (
	K ByteUnit = 1024
	M ByteUnit = 1024 * K
	G ByteUnit = 1024 * M
)

func (c ByteUnit) String() string {
	if c < G {
		return fmt.Sprintf("%dM", c/M)
	}
	return fmt.Sprintf("%dG", c/G)
}

var testSuites = []struct {
	NumFiles int
	NumBytes ByteUnit
}{
	{100, 100 * M},
	{100, 1 * G},
	{100, 5 * G},
	{100, 10 * G},
	{1000, 1 * G},
	{1000, 2 * G},
	{1000, 5 * G},
	{1000, 10 * G},
	{10000, 10 * G},
	{20000, 10 * G},
}

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "genFiles":
		// go run ./ genFiles tmp_100_100M 100 100M
		// go run ./ genFiles tmp_1000_100M 1000 100M

		// go run ./ genFiles tmp_1000_1G 1000 1G
		// go run ./ genFiles tmp_2000_1G 2000 1G
		// go run ./ genFiles tmp_5000_1G 5000 1G

		// go run ./ genFiles tmp_1000_5G 3000 5G
		err := GenFiles(os.Args[2], mustInt(os.Args[3]), mustSize(os.Args[4]))
		if err != nil {
			log.Fatalf("gen files: %v", err)
		}
	case "showGenCommands":
		for _, c := range testSuites {
			caseName := fmt.Sprintf("%d_%s", c.NumFiles, c.NumBytes.String())
			fmt.Printf("go run ./ genFiles tmp_%s %d %d\n", caseName, c.NumFiles, c.NumBytes)
		}
	case "diffDir":
		diffDir(os.Args[2], os.Args[3])
	default:
		panic(fmt.Errorf("unkonwn command"))
	}
}

func mustInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Errorf("invalid int: %s %v", s, err))
	}
	return int(i)
}
func mustSize(s string) ByteUnit {
	if s == "" {
		panic(fmt.Errorf("empty string"))
	}
	n, u := s[:len(s)-1], s[len(s)-1:]
	i := mustInt(n)
	switch u {
	case "K":
		return K * ByteUnit(i)
	case "M":
		return M * ByteUnit(i)
	case "G":
		return G * ByteUnit(i)
	}
	panic(fmt.Errorf("unreachable"))
}

func diffDir(a string, b string) {
	log.Printf("diff: %s %s", a, b)
	aFiles, err := ioutil.ReadDir(a)
	if err != nil {
		panic(err)
	}
	bFiles, err := ioutil.ReadDir(b)
	if err != nil {
		panic(err)
	}

	if len(aFiles) != len(bFiles) {
		panic(fmt.Errorf("files not match: %d vs %d", len(aFiles), len(bFiles)))
	}

	sort.Slice(aFiles, func(i, j int) bool {
		return aFiles[i].Name() < aFiles[j].Name()
	})

	sort.Slice(bFiles, func(i, j int) bool {
		return bFiles[i].Name() < bFiles[j].Name()
	})

	// compare size
	aSize := 0
	bSize := 0
	for i := 0; i < len(aFiles); i++ {
		aFile := aFiles[i]
		bFile := bFiles[i]

		if aFile.Name() != bFile.Name() {
			panic(fmt.Errorf("file name mismatch: %d %s vs %s", i, aFile.Name(), bFile.Name()))
		}

		if aFile.Size() != bFile.Size() {
			panic(fmt.Errorf("file size mismatch: %s %d vs %d", aFile.Name(), aFile.Size(), bFile.Size()))
		}
		aSize += int(aFile.Size())
		bSize += int(bFile.Size())
	}
	log.Printf("sizes: %v %v", aSize, bSize)
	// compare contents
	for i := 0; i < len(aFiles); i++ {
		aFile := aFiles[i]
		bFile := bFiles[i]

		aContent, err := ioutil.ReadFile(filepath.Join(a, aFile.Name()))
		if err != nil {
			panic(err)
		}
		bContent, err := ioutil.ReadFile(filepath.Join(b, bFile.Name()))
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(aContent, bContent) {
			panic(fmt.Errorf("file content mismatch: %s", aFile.Name()))
		}
	}
}
