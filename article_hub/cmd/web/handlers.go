package main

import (
	"context"
	"fmt"
	"github.com/ol-ilyassov/final/article_hub/articlepb"
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
	log.Printf("Response from GetArticle RPC, ArticleID: %v", res.GetResult())

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
	log.Printf("Response from InsertArticle RPC: %v", res.Result)

	app.session.Put(r, "flash", "Snippet successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", res.GetId()), http.StatusSeeOther)
}

func (app *application) deleteSnippet(w http.ResponseWriter, r *http.Request) {
	//err := r.ParseForm()
	//if err != nil {
	//	app.clientError(w, http.StatusBadRequest)
	//	return
	//}
	//
	//form := forms.New(r.)

}
