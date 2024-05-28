package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	ID               int
	Name             string
	Job              string
	FavoriteLanguage string
	Email            string
	Employed         bool
}

var people = []Person{
	{ID: 1, Name: "Cy Ganderton", Job: "Quality Control Specialist", FavoriteLanguage: "Go", Email: "cy@example.com", Employed: true},
	{ID: 2, Name: "Hart Hagerty", Job: "Desktop Support Technician", FavoriteLanguage: "Python", Email: "hart@example.com", Employed: false},
	{ID: 3, Name: "Brice Swyre", Job: "Tax Accountant", FavoriteLanguage: "JavaScript", Email: "brice@example.com", Employed: true},
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	job := r.FormValue("job")
	favoriteLanguage := r.FormValue("favorite_language")
	email := r.FormValue("email")
	employed := r.FormValue("employed") == "true"

	newPerson := Person{
		ID:               len(people) + 1,
		Name:             name,
		Job:              job,
		FavoriteLanguage: favoriteLanguage,
		Email:            email,
		Employed:         employed,
	}

	people = append(people, newPerson)

	tmpl := template.Must(template.ParseFiles("views/people-table.html"))
	err := tmpl.ExecuteTemplate(w, "people-table", people)
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, p := range people {
		if p.ID == idInt {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}

	tmpl := template.Must(template.ParseFiles("views/people-table.html"))
	err = tmpl.ExecuteTemplate(w, "people-table", people)
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}
}

func main() {
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/modal-form.html",
			"views/footer.html",
			"views/people-table.html",
		))
		err := tmpl.ExecuteTemplate(w, "index.html", people)
		if err != nil {
			log.Printf("Template execution failed: %s", err)
		}
	})

	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist")))) // Serve static files

	log.Fatal(http.ListenAndServe(":8080", nil))
}
