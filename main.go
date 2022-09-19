package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
)

func main() {
	destDir := flag.String("d", ".", "destination directory")
	flag.Parse()

	var cm corev1.ConfigMap
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &cm)
	for k, v := range cm.Data {
		if v != "" {
			path := filepath.Join(*destDir, k)
			if err := os.WriteFile(path, []byte(v), 0644); err != nil {
				panic(err)
			}
		}
	}
}
