/*
  Panic
    Occur when program cannot continue at all
      Don't use wjen file cannot be opened, unless it is critical
      Use for unrecoverable events - cannot obtain TCP port for web server
    Function will stop executing
      Deferred functions will still fire
    If nothing handles panic, program will exit
*/

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
