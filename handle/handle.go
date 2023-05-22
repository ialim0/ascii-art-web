package handle

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"main.go/link"
)

type FormData struct {
	Text   string
	Banner string
	Result string
	Error  string
}

var validBanners = map[string]bool{
	"standard":   true,
	"shadow":     true,
	"thinkertoy": true,
	"graffiti":   true,
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if len(os.Args) > 1 {

	}
	if r.Method != "POST" && r.URL.Path != "/" {
		w.WriteHeader(404)
		http.ServeFile(w, r, "templates/error.html")

		return
	} else if r.Method != "GET" {
		w.WriteHeader(404)
		http.ServeFile(w, r, "templates/error.html")

	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		http.ServeFile(w, r, "templates/error.html")
		return
	} else if r.Method == "POST" && r.URL.Path != "/ascii-art" {
		w.WriteHeader(400)
		http.ServeFile(w, r, "templates/error.html")
		return
	}

	// Obtenir les données du formulaire
	text := r.FormValue("text")
	text = strings.ReplaceAll(text, "\\r", "\\n")
	banner := r.FormValue("banner")

	// Vérifier si le type de bannière est valide
	if !isValidBanner(banner) {
		w.WriteHeader(404)
		http.ServeFile(w, r, "templates/error.html")
		return
	}

	// Vérifier si le texte est vide
	if text == "" {
		handleError(w, http.StatusBadRequest, "Le champ de texte ne peut pas être vide")
		http.ServeFile(w, r, "templates/badrequest.html")
		return
	}

	// Générer l'art ASCII
	asciiArt := generateASCIIArt(text, banner)

	// Vérifier si l'art ASCII est généré avec succès
	if len(asciiArt) == 0 || asciiArt == "\nError" {
		w.WriteHeader(500)
		http.ServeFile(w, r, "templates/internalserver.html")
		return

	}

	// Afficher le résultat sur la page
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, FormData{Text: text, Banner: banner, Result: asciiArt})
}

func generateASCIIArt(text, banner string) string {

	banner = "static/" + banner + ".txt"
	var tabmax [][]string
	if bitdata, err := ioutil.ReadFile(banner); err == nil {

		lines := strings.Split(string(bitdata), "\n")
		var tabmin []string
		for _, donne := range lines[1:] {
			if len(donne) != 0 {
				tabmin = append(tabmin, donne)
			} else {
				tabmax = append(tabmax, tabmin)
				tabmin = []string{}
			}
		}

	}

	return "\n" + link.PrintAscii(text, tabmax)
}

func handleError(w http.ResponseWriter, statusCode int, errMsg string) {
	w.WriteHeader(statusCode)
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, FormData{Error: errMsg})
}

func isValidBanner(banner string) bool {
	_, ok := validBanners[banner]
	return ok
}
