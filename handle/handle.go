package handle

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"main.go/link"
)

type FormData struct {
	Text   string
	Banner string
	Result string
	Error  string
}

var fileNum int = 0
var fileName string = ""
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
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
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

	if resultascii != "" {
		fileNum++

		fileName = "File_" + strconv.Itoa(fileNum) + ".txt"
		// Créer un fichier en écriture
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		// Copier les données dans le fichier
		_, err = file.Write([]byte(resultascii))
		if err != nil {
			fmt.Println(err)
		}

	}

	return "\n" + resultascii
}
func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	// Configuration de l'en-tête de la réponse HTTP
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Lentght", strconv.Itoa(len(resultascii)))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

	// Serveur du fichier en tant que réponse HTTP
	http.ServeFile(w, r, fileName)
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
