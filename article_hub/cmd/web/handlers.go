package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ol-ilyassov/final/article_hub/articlepb"
	"github.com/ol-ilyassov/final/article_hub/notifypb"
	"github.com/ol-ilyassov/final/article_hub/pkg/forms"
	"github.com/ol-ilyassov/final/article_hub/pkg/models"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func StringToTime(value string) time.Time {
	result, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", value)
	return result
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	req := &articlepb.GetArticlesRequest{}
	stream, err := app.snippets.GetArticles(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetArticles RPC: %v", err)
	}
	defer stream.CloseSend()

	var articles []*models.Snippet

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break LOOP
			}
			log.Fatalf("Error while receiving from GetArticles RPC: %v", err)
		}
		log.Printf("Response from GetArticles RPC, ArticleID: %v \n", res.GetArticle().GetId())

		tempArticle := &models.Snippet{
			ID:      int(res.GetArticle().GetId()),
			Title:   res.GetArticle().GetTitle(),
			Content: res.GetArticle().GetContent(),
			Created: StringToTime(res.GetArticle().GetCreated()),
			Expires: StringToTime(res.GetArticle().GetExpires()),
		}
		articles = append(articles, tempArticle)
	}

	// Web Design
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: articles,
	})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	req := &articlepb.GetArticleRequest{
		Id: int32(id),
	}
	res, err := app.snippets.GetArticle(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetArticle RPC: %v", err)
	}
	log.Printf("Response from GetArticle RPC: %s, ArticleID: %v", res.GetResult(), res.GetArticle().GetId())

	article := &models.Snippet{
		ID:      int(res.GetArticle().GetId()),
		Title:   res.GetArticle().GetTitle(),
		Content: res.GetArticle().GetContent(),
		Created: StringToTime(res.GetArticle().GetCreated()),
		Expires: StringToTime(res.GetArticle().GetExpires()),
	}

	// Web Design
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: article,
	})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
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

	req := &articlepb.InsertArticleRequest{
		Article: &articlepb.Article{
			Title:   form.Get("title"),
			Content: form.Get("content"),
			Expires: form.Get("expires"),
		},
	}
	res, err := app.snippets.InsertArticle(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling InsertArticle RPC: %v", err)
	}
	log.Printf("Response from InsertArticle RPC: %v", res.GetResult())

	app.session.Put(r, "flash", "Snippet successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", res.GetId()), http.StatusSeeOther)
}

func (app *application) deleteSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	req := &articlepb.DeleteArticleRequest{
		Id: int32(id),
	}
	res, err := app.snippets.DeleteArticle(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling DeleteArticle RPC: %v", err)
	}
	log.Printf("Response from DeleteArticle RPC: %v", res.GetResult())

	http.Redirect(w, r, "/", http.StatusSeeOther)
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
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)
	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}
	// Try to create a new user record in the database. If the email already exists
	// add an error message to the form and re-display it.
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Otherwise add a confirmation flash message to the session confirming that
	// their signup worked and asking them to log in.
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
	// And redirect the user to the login page.
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("email", "password")
	if !form.Valid() {
		form.Errors.Add("generic", "Email and Password are required")
		app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		return
	}
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if err == models.ErrInvalidCredentials {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else if err == models.ErrNoRecord {
			form.Errors.Add("generic", "No user with given Email")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r, "authenticatedUserID", id)

	//http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")

	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) sendMail(w http.ResponseWriter, r *http.Request) {

	req := &notifypb.NotifyRequest{
		Email: &notifypb.Email{
			Address: "ol.ilyassov@gmail.com",
			Title:   "Golang Microservices",
			Time:    time.Now().Format("02 Jan 2006 at 15:04"),
		},
	}
	res, err := app.notifier.SendNotification(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling SendNotification RPC: %v", err)
	}
	log.Printf("Response from SendNotification RPC: %v", res.GetResult())

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
