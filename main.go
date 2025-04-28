// LandGut-App Backend (Böblingen Update) - main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Provider struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Products string  `json:"products"`
	Hours    string  `json:"hours"`
	Payment  string  `json:"payment"`
	Homepage string  `json:"homepage"`
}

type Product struct {
	Name  string `json:"name"`
	Price string `json:"price,omitempty"`
}

var providers = []Provider{
	{"1", "Bauernhof Müller", 48.3974, 9.9934, "Gemüse, Obst", "Mo-Sa 08:00-18:00", "Bar, Karte", "http://bauernhof-mueller.de"},
	{"2", "Metzgerei Schmidt", 48.6833, 9.4167, "Fleisch, Wurst", "Mo-Fr 09:00-18:00", "Bar", "http://metzgerei-schmidt.de"},
	{"3", "Bäckerei Huber", 48.05, 9.2, "Brot, Kuchen", "Di-So 07:00-17:00", "Bar, Karte", "http://baeckerei-huber.de"},
	{"4", "Bauernhof Weber", 48.6829, 9.0112, "Eier, Gemüse", "Mo-Sa 07:30-18:00", "Bar", "http://bauernhof-weber.de"},
	{"5", "Metzgerei Bauer", 48.6900, 9.0200, "Fleisch, Wurst", "Mo-Fr 08:00-18:00", "Bar, Karte", "http://metzgerei-bauer.de"},
	{"6", "Bäckerei Schäfer", 48.6870, 9.0060, "Brot, Brötchen, Kuchen", "Mo-So 06:30-17:00", "Bar, Karte", "http://baeckerei-schaefer.de"},
}

var products = map[string][]Product{
	"1": {{"Karotten", ""}, {"Kartoffeln", ""}, {"Äpfel", ""}},
	"2": {{"Schweinebraten", "12.99"}, {"Bratwurst", "3.49"}, {"Leberkäse", "2.99"}},
	"3": {{"Bauernbrot", "4.50"}, {"Apfelkuchen", "2.80"}, {"Brezen", "1.20"}},
	"4": {{"Eier", ""}, {"Zucchini", ""}, {"Tomaten", ""}},
	"5": {{"Rindersteak", "16.99"}, {"Schinken", "9.50"}, {"Wienerle", "4.00"}},
	"6": {{"Laugenbrötchen", "1.10"}, {"Dinkelbrot", "3.90"}, {"Käsekuchen", "3.20"}},
}

func providersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(providers)
}

func providerDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, p := range providers {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.NotFound(w, r)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	p, ok := products[id]
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func main() {
	http.HandleFunc("/api/providers", providersHandler)
	http.HandleFunc("/api/provider", providerDetailHandler)
	http.HandleFunc("/api/products", productsHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("Server läuft auf http://localhost:8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
