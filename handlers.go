package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slovnik/seznam"
	"github.com/slovnik/slovnik"
)

func translate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	lang := slovnik.DetectLanguage(word)
	translations, _ := seznam.Translate(word, lang)
	json.NewEncoder(w).Encode(translations)
}
