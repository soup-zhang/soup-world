/**
 * @Time : 2019-08-13 14:29 
 * @Author : soupzhb@gmail.com
 * @File : main.go
 * @Software: GoLand
 */

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("main start ...")
	fmt.Println("main start 222 ...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Soup World, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))


}
