package handle

import (
	"fmt"
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

var resultascii string

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
		if r.URL.Path == "ascii-art" {
			w.WriteHeader(405)
			http.ServeFile(w, r, "templates/405.html")

		} else {
			w.WriteHeader(400)
			http.ServeFile(w, r, "templates/404.html")

		}

		return
	} else if r.Method != "GET" {
		w.WriteHeader(405)
		http.ServeFile(w, r, "templates/405.html")

	}
	varinternal := "templates/index.html"
	if varinternal == "templates/index.html" {
		tmpl := template.Must(template.ParseFiles(varinternal))
		tmpl.Execute(w, nil)

	} else {
		w.WriteHeader(500)
		http.ServeFile(w, r, "templates/500.html")

	}

}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		http.ServeFile(w, r, "templates/405.html")
		return
	} else if r.Method == "POST" && r.URL.Path != "/ascii-art" {
		w.WriteHeader(400)
		http.ServeFile(w, r, "templates/400.html")
		return
	}

	// Obtenir les données du formulaire
	text := r.FormValue("text")
	text = strings.ReplaceAll(text, "\\r", "\\n")
	banner := r.FormValue("banner")

	// Vérifier si le type de bannière est valide
	if !isValidBanner(banner) {
		w.WriteHeader(404)
		http.ServeFile(w, r, "templates/404.html")
		return
	}

	// Vérifier si le texte est vide
	if text == "" {
		handleError(w, http.StatusBadRequest, "Le champ de texte ne peut pas être vide")
		http.ServeFile(w, r, "templates/400.html")
		return
	}

	// Générer l'art ASCII
	asciiArt := generateASCIIArt(text, banner)

	// Vérifier si l'art ASCII est généré avec succès
	if len(asciiArt) == 0 || asciiArt == "\nError" {
		w.WriteHeader(400)
		http.ServeFile(w, r, "templates/400.html")
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

	resultascii = link.PrintAscii(text, tabmax)
	//Condition de validation

	return resultascii
}
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	result := r.URL.Query().Get("result")

	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(result)))

	fmt.Fprint(w, result)
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
