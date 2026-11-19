// Command preview serves the site at http://127.0.0.1:4174 for a person to
// look at before pushing. It is what `game run preview` runs.
//
// It differs from a plain file server in two ways, both so that what the
// browser shows is always what is on disk now:
//
//   - every answer says Cache-Control: no-store, and the browser's "has this
//     changed?" questions (If-Modified-Since, If-None-Match, If-Range) are
//     dropped before the file is served, so a file replaced under the same
//     name is never answered with 304 Not Modified;
//   - a path with no extension is tried as path.html, as Vercel's cleanUrls
//     does, so /v1 is v1.html here as it is there.
//
// It listens on the loopback address only and takes no arguments: it serves
// the directory it is started in.
package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// address is where the preview listens: loopback only, so nobody else on
// the network can reach it.
const address = "127.0.0.1:4174"

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	files := http.FileServer(http.Dir(root))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("If-Modified-Since")
		r.Header.Del("If-None-Match")
		r.Header.Del("If-Range")
		r.Header.Del("If-Unmodified-Since")
		r.Header.Del("If-Match")
		w.Header().Set("Cache-Control", "no-store")
		clean := path.Clean(r.URL.Path)
		if clean != "/" && !strings.HasSuffix(r.URL.Path, "/") && path.Ext(clean) == "" {
			if _, err := os.Stat(root + clean + ".html"); err == nil {
				r.URL.Path = clean + ".html"
			}
		}
		log.Printf("%s %s", r.Method, r.URL.Path)
		files.ServeHTTP(w, r)
	})
	log.Printf("serving %s at http://%s", root, address)
	log.Fatal(http.ListenAndServe(address, handler))
}
