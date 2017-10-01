package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// define global vars
var location, listen, sysPath, host string
var mdFiles []string

// get values fom env or set defaults
func init() {
	// metrics url
	location = os.Getenv("MD_URL")
	if location == "" {
		location = "/metrics"
	}
	// listen port
	listen = os.Getenv("MD_LISTEN")
	if listen == "" {
		listen = ":8080"
	}
	// mount path of /sys filesystem
	sysPath = os.Getenv("MD_SYSPATH")
	if sysPath == "" {
		sysPath = "/sys"
	}
	// set hostname label
	host = os.Getenv("MD_HOST")
	if host == "" {
		host, _ = os.Hostname()
	}
	// slice of devices filepaths ["/sys/block/md0", "/sys/block/md1", ...]
	mdFiles, _ = filepath.Glob(fmt.Sprintf("%s/block/md*", sysPath))
}

// path to array state /sys/block/mdN/md/array_state
func serveMetrics(w http.ResponseWriter, r *http.Request) {
	var state []byte
	var stateCode int
	var dev string
	fmt.Fprintf(w, "# metric values: 1 - clean, 0 - degraded\n")
	for _, file := range mdFiles {
		state, _ = ioutil.ReadFile(fmt.Sprintf("%s/md/array_state", file))
		dev = filepath.Base(file)
		if strings.TrimSpace(string(state)) == "clean" {
			stateCode = 1
		} else {
			stateCode = 0
		}
		fmt.Fprintf(w, "md_state{host=\"%s\",dev=\"%s\"} %d\n", host, dev, stateCode)
	}
}

func main() {
	fmt.Printf("Listening %s%s\n", listen, location)
	http.HandleFunc(location, serveMetrics)
	log.Fatal(http.ListenAndServe(listen, nil))
}
