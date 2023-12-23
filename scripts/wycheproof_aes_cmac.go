package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	if err := download(); err != nil {
		panic(err)
	}
	if err := gen(); err != nil {
		panic(err)
	}
}

var jsonFile = filepath.Join("wycheproof_cache", "aes_cmac_test.json")

func download() error {
	// check if file is cached
	{
		fi, err := os.Stat(jsonFile)
		if err == nil {
			if fi.IsDir() {
				return fmt.Errorf("%s is dir", jsonFile)
			}
			fmt.Println("cached => no download")
			return nil
		}
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
	}

	const url = "https://raw.githubusercontent.com/google/wycheproof/master/testvectors_v1/aes_cmac_test.json"

	fmt.Println("download:", url)

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	_ = os.MkdirAll(filepath.Dir(jsonFile), 0700)

	f, err := os.Create(jsonFile)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println("save:", jsonFile)

	_, err = io.Copy(f, r.Body)

	return err
}

func gen() error {
	fmt.Println("read", jsonFile)

	f, err := os.Open(jsonFile)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println("decode", jsonFile)

	var d Data
	if err := json.NewDecoder(f).Decode(&d); err != nil {
		return err
	}

	fmt.Println("parse template")

	t, err := template.New("foo").Parse(tmpl)
	if err != nil {
		return err
	}

	dstFile := filepath.Join("internal", "crypto", "aes", "cmac", "wycheproof_test.go")

	fmt.Println("create", dstFile)

	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	var tests []Test
	for _, tg := range d.TestGroups {
		tests = append(tests, tg.Tests...)
	}

	fmt.Println("generate code to", dstFile)

	return t.Execute(dst, tests)
}

type Data struct {
	TestGroups []TestGroup `json:"testGroups"`
}

type TestGroup struct {
	KeySize int    `json:"keySize"`
	TagSize int    `json:"tagSize"`
	Tests   []Test `json:"tests"`
}

type Test struct {
	TcId   int      `json:"tcId"`
	Flags  []string `json:"flags"`
	Key    string   `json:"key"`
	Msg    string   `json:"msg"`
	Tag    string   `json:"tag"`
	Result string   `json:"result"`
}

const tmpl = `
// GENERATED FILE, DO NOT EDIT!

package cmac

type wycheproofVector struct {
	TcId   int
	Flag  string
	Key    string
	Msg    string
	Tag    string
	Result string
}

var wycheproofVectors = []wycheproofVector{
{{range .}}
{TcId: {{.TcId}}, Flag: "{{index .Flags 0}}", Key: "{{.Key}}", Msg: "{{.Msg}}", Tag: "{{.Tag}}", Result: "{{.Result}}"},
{{end}}
}
`
