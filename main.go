package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
)

var (
	destDir          *string
	includeEmptyKeys *bool
	verbose          *bool
	base64decode     *bool
)

func init() {
	destDir = flag.String("d", ".", "destination directory")
	includeEmptyKeys = flag.Bool("z", false, "include empty keys")
	verbose = flag.Bool("v", false, "log filenames")
	base64decode = flag.Bool("b", false, "base64 decode content before writing")
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()

	selected := flag.Args()

	var cm corev1.ConfigMap
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(bytes, &cm); err != nil {
		panic(err)
	}

	if cm.Kind == "Secret" {
		log.Printf("will base64 decode values for secrets")
		*base64decode = true
	}

	for k, v := range cm.Data {
		if len(selected) > 0 && !contains(selected, k) {
			continue
		}

		if *includeEmptyKeys || v != "" {
			var buf []byte
			var err error

			path := filepath.Join(*destDir, k)
			if *verbose {
				log.Printf("writing %s", path)
			}
			if *base64decode {
				buf, err = base64.StdEncoding.DecodeString(v)
				if err != nil {
					panic(err)
				}
			} else {
				buf = []byte(v)
			}

			if err := os.WriteFile(path, buf, 0644); err != nil {
				panic(err)
			}
		}
	}
}
