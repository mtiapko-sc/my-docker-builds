package main

/*
#include <stdlib.h>

// Trivial C helper so the build needs a C toolchain (cgo).
static int meaning_of_life() { return 42; }
*/
import "C"

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		fmt.Fprintf(w, "Hello from Go! request-id=%s answer=%d\n", id, int(C.meaning_of_life()))
	})
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
