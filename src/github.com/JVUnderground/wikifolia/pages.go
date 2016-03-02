package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
)

var DB_DRIVER string

func init() {
	sql.Register(DB_DRIVER, &sqlite3.SQLiteDriver{})
}

/* Index:
Página principal.
*/
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/developers", http.StatusFound)
}
