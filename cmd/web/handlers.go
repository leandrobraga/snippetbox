package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/leandrobraga/snippetbox/pkg/forms"
	"github.com/leandrobraga/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Só é preciso com o mux nativo do Go. Usando o pacote Pat não é preciso
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
	// data := &templateData{Snippets: s}

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})

	// data := &templateData{Snippet: s}

	// files := []string{
	// 	"./ui/html/show.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Não é necessário checar o método com o pacote Pat
	// if r.Method != "POST" {
	// 	w.Header().Set("Allow", "POST")
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }
	// title := "O snail"
	// content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	// expires := "7"
	// id, err := app.snippets.Insert(title, content, expires)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
	// http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
	err := r.ParseForm()

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))

	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Snippet successfuly created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}

	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err == models.ErrDuplicateEmail {
		form.Errors.Add("email", "Address is already in use")
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Your signup was successfull. Please log in")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form ...")
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user ...")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logou the user ...")
}
