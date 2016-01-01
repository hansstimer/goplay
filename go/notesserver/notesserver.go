package main

import (
	"http"
	"fmt"
	"exec"
	"bytes"
)

func markdownHandler(w http.ResponseWriter, r *http.Request, title string) {
	cmd := "/bin/sh"
	args := []string{cmd, "-c", "ls passwd"}
	dir := "/etc"
	p, err := exec.Run(cmd, args, nil, dir,
		exec.DevNull, exec.Pipe, exec.PassThrough)
	if err != nil {
		return
	}
	var b bytes.Buffer
	_, err = b.ReadFrom(p.Stdout)
	if err != nil {
		return
	}
	err = p.Close()
	if err != nil {
		return
	}
	fmt.Println(b.String())
}

func main() {
	http.Handle("/", http.FileServer("/Users/marty/projects/notes", ""))
	http.ListenAndServe(":8080", nil)
}
