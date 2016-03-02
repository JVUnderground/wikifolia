package main

import (
	"net/http"
)

/* Index:
Página principal.
*/
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/developers", http.StatusFound)
}
