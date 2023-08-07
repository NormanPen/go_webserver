package main

import (
	"fmt"
	"net/http"
	"path"
)

func main() {
	// Statische Dateien aus dem Build-Ordner deiner React-Anwendung servieren
	fs := http.FileServer(http.Dir(path.Join(".", "web")))
	http.Handle("/", fs)

	port := 8181
	fmt.Printf("Server läuft auf Port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
