package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
)

var (
	destDir          *string
	includeEmptyKeys *bool
)

func init() {
	destDir = flag.String("d", ".", "destination directory")
	includeEmptyKeys = flag.Bool("z", false, "include empty keys")
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

	err = json.Unmarshal(bytes, &cm)
	for k, v := range cm.Data {
		if len(selected) > 0 && !contains(selected, k) {
			continue
		}

		if *includeEmptyKeys || v != "" {
			path := filepath.Join(*destDir, k)
			if err := os.WriteFile(path, []byte(v), 0644); err != nil {
				panic(err)
			}
		}
	}
}
