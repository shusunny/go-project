package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tmpls/*"))
}

func main() {
	http.HandleFunc("/", index)
	// route to serve pictures
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	// process from submission
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		// create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		hash := sha1.New()
		io.Copy(hash, mf)
		fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext
		// create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		// copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		// add filename to this user's cookie
		c = appendValue(w, c, fname)
	}
	xs := strings.Split(c.Value, "|")
	// slice cookie to show only images
	tpl.ExecuteTemplate(w, "index.html", xs[1:])
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
