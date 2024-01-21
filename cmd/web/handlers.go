package main

import (
	"aidoskanatbay.net/snippetbox/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the new render helper.
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Use the new render helper.
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"
	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle form submission

		// Access form values
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		// Do something with the form data, e.g., send an email, save to database, etc.

		// For now, let's just print the form data
		fmt.Printf("Name: %s\nEmail: %s\nMessage: %s\n", name, email, message)

		// Redirect to a thank you page or display a success message
		http.Redirect(w, r, "/thank-you", http.StatusSeeOther)
		return
	}

	// Render the contact form page for GET requests
	app.render(w, r, "contact.page.tmpl", nil)
}
func (app *application) students(w http.ResponseWriter, r *http.Request) {
	// Retrieve snippets with the category "students"
	snippets, err := app.snippets.Student("students")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, "student.page.tmpl", &templateData{
		Snippets: snippets,
	})
}

func (app *application) staff(w http.ResponseWriter, r *http.Request) {
	// Retrieve snippets with the category "students"
	snippets, err := app.snippets.Staff("staff")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, "staff.page.tmpl", &templateData{
		Snippets: snippets,
	})
}
func (app *application) applicant(w http.ResponseWriter, r *http.Request) {
	// Retrieve snippets with the category "students"
	snippets, err := app.snippets.Applicant("staff")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, "applicant.page.tmpl", &templateData{
		Snippets: snippets,
	})
}

func (app *application) researcher(w http.ResponseWriter, r *http.Request) {
	// Retrieve snippets with the category "students"
	snippets, err := app.snippets.Researcher("researcher")
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, "researcher.page.tmpl", &templateData{
		Snippets: snippets,
	})
}
