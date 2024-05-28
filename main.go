package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	ID            int
	Name          string
	Job           string
	FavoriteColor string
}

func main() {
	// Sample data
	people := []Person{
		{ID: 1, Name: "Cy Ganderton", Job: "Quality Control Specialist", FavoriteColor: "Blue"},
		{ID: 2, Name: "Hart Hagerty", Job: "Desktop Support Technician", FavoriteColor: "Purple"},
		{ID: 3, Name: "Brice Swyre", Job: "Tax Accountant", FavoriteColor: "Red"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/modal-form.html",
			"views/footer.html",
		))
		err := tmpl.ExecuteTemplate(w, "index.html", people)
		if err != nil {
			log.Printf("Template execution failed: %s", err)
		}
	})

	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist")))) // Serve static files

	log.Fatal(http.ListenAndServe(":8080", nil))
}
