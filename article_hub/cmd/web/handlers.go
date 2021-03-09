package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ol-ilyassov/final/article_hub/articlepb"
	"github.com/ol-ilyassov/final/article_hub/authpb"
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
	stream, err := app.articles.GetArticles(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetArticles RPC: %v", err)
	}
	defer stream.CloseSend()

	var articles []*models.Article

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

		tempArticle := &models.Article{
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
		Articles: articles,
	})
}

func (app *application) showArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	req := &articlepb.GetArticleRequest{
		Id: int32(id),
	}
	res, err := app.articles.GetArticle(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetArticle RPC: %v", err)
	}
	log.Printf("Response from GetArticle RPC: %s, ArticleID: %v", res.GetResult(), res.GetArticle().GetId())

	article := &models.Article{
		ID:      int(res.GetArticle().GetId()),
		Title:   res.GetArticle().GetTitle(),
		Content: res.GetArticle().GetContent(),
		Created: StringToTime(res.GetArticle().GetCreated()),
		Expires: StringToTime(res.GetArticle().GetExpires()),
	}

	// Web Design
	app.render(w, r, "show.page.tmpl", &templateData{
		Article: article,
	})
}

func (app *application) createArticleForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createArticle(w http.ResponseWriter, r *http.Request) {
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
	res, err := app.articles.InsertArticle(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling InsertArticle RPC: %v", err)
	}
	log.Printf("Response from InsertArticle RPC: %v", res.GetResult())

	app.session.Put(r, "flash", "Article successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/article/%d", res.GetId()), http.StatusSeeOther)
}

func (app *application) deleteArticle(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	req := &articlepb.DeleteArticleRequest{
		Id: int32(id),
	}
	res, err := app.articles.DeleteArticle(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling DeleteArticle RPC: %v", err)
	}
	log.Printf("Response from DeleteArticle RPC: %v", res.GetResult())

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

	req := &authpb.CreateUserRequest{
		User: &authpb.User{
			Name:     form.Get("name"),
			Email:    form.Get("email"),
			Password: form.Get("password"),
		},
	}

	res, _ := app.auth.CreateUser(context.Background(), req)

	if !res.GetStatus() {
		if res.GetResult() == models.ErrDuplicateEmail.Error() {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, errors.New(res.GetResult()))
		}
		return
	}

	app.session.Put(r, "flash", "Your signup was successful. Please log in.")

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

	req := &authpb.AuthUserRequest{
		User: &authpb.User{
			Email:    form.Get("email"),
			Password: form.Get("password"),
		},
	}

	res, _ := app.auth.AuthUser(context.Background(), req)

	if !res.GetStatus() {
		if res.GetResult() == models.ErrInvalidCredentials.Error() {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else if res.GetResult() == models.ErrNoRecord.Error() {
			form.Errors.Add("generic", "No user with given Email")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, errors.New(res.GetResult()))
		}
		return
	}

	app.session.Put(r, "authenticatedUserID", int(res.GetId()))

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")

	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) searchForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "search.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) search(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title")
	if !form.Valid() {
		app.render(w, r, "search.page.tmpl", &templateData{Form: form})
		return
	}
	req := &articlepb.SearchArticlesRequest{
		Title: form.Get("title"),
	}
	stream, err := app.articles.SearchArticles(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling SearchArticles RPC: %v", err)
	}
	defer stream.CloseSend()
	var articles []*models.Article

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break LOOP
			}
			log.Fatalf("Error while receiving from SearchArticles RPC: %v", err)
		}
		log.Printf("Response from SearchArticles RPC, ArticleTitle: %v \n", res.GetArticle().GetTitle())

		tempArticle := &models.Article{
			ID:      int(res.GetArticle().GetId()),
			Title:   res.GetArticle().GetTitle(),
			Content: res.GetArticle().GetContent(),
			Created: StringToTime(res.GetArticle().GetCreated()),
			Expires: StringToTime(res.GetArticle().GetExpires()),
		}
		articles = append(articles, tempArticle)
	}
	app.render(w, r, "search.page.tmpl", &templateData{
		Articles: articles,
	})
}
