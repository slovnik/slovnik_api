package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slovnik/seznam"
	"github.com/slovnik/slovnik"
)

type FastTrans struct {
	Word        string `json:"word"`
	Translation string `json:"translation"`
}

type SuggestionResponse struct {
	Result []SuggestionObject `json:"result"`
}

type SuggestionObject struct {
	Suggest []Suggestion `json:"suggest"`
}

type Suggestion struct {
	Value     string `json:"value"`
	Relevance int    `json:"relevance"`
}

func translate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	lang := slovnik.DetectLanguage(word)
	translations, err := seznam.Translate(word, lang)

	if err != nil {
		log.Println(err)
	}

	err = json.NewEncoder(w).Encode(translations)

	if err != nil {
		log.Println(err)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	term := vars["term"]

	response, err := http.Get("https://slovnik.seznam.cz/suggest/mix_cz_ru?result=json&count=10&phrase=" + term)
	if err != nil {
		log.Panic(err)
	}

	defer response.Body.Close()

	var suggestions SuggestionResponse

	err = json.NewDecoder(response.Body).Decode(&suggestions)
	if err != nil {
		log.Panic(err)
	}

	result := make([]FastTrans, 0, 10)
	for _, s := range suggestions.Result[0].Suggest {
		result = append(result, FastTrans{Word: s.Value, Translation: ""})
	}

	json.NewEncoder(w).Encode(result)
}
