package main

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strconv"
	"time"
)

var m = &ArticleModel{}

type ArticleModel struct {
	DB *pgxpool.Pool
}

func (m *ArticleModel) Insert(title, content, expires string, authorid int) (int, error) {

	stmt := "INSERT INTO articles (title, content, created, expires, authorid) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	var id uint64
	days, err := strconv.Atoi(expires)
	if err != nil {
		return 0, err
	}
	row := m.DB.QueryRow(context.Background(), stmt, title, content, time.Now(), time.Now().AddDate(0, 0, days), authorid)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *ArticleModel) Get(id int) (*Article, error) {

	stmt := "SELECT id, title, content, created, expires, authorid FROM articles WHERE expires > $1 AND id = $2"

	row := m.DB.QueryRow(context.Background(), stmt, time.Now(), id)
	s := &Article{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires, &s.AuthorID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *ArticleModel) Latest() ([]*Article, error) {

	stmt := "SELECT id, title, content, created, expires, authorid FROM articles WHERE expires > $1 ORDER BY created DESC"

	rows, err := m.DB.Query(context.Background(), stmt, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*Article

	for rows.Next() {
		s := &Article{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires, &s.AuthorID)
		if err != nil {
			return nil, err
		}
		articles = append(articles, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

func (m *ArticleModel) Delete(id int) error {

	stmt := "DELETE FROM articles WHERE id = $1"

	ct, err := m.DB.Exec(context.Background(), stmt, id)
	if err != nil {
		log.Printf("Unable to DELETE: %v\n", err)
		return err
	}

	if ct.RowsAffected() == 0 {
		return errors.New("no affected rows")
	}

	return nil
}

func (m *ArticleModel) Search(title string) ([]*Article, error) {

	stmt := "SELECT id, title, content, created, expires, authorid FROM articles WHERE expires > $1 AND LOWER(title) LIKE '%' || $2 || '%'"

	rows, err := m.DB.Query(context.Background(), stmt, time.Now(), title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*Article

	for rows.Next() {
		s := &Article{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires, &s.AuthorID)
		if err != nil {
			return nil, err
		}
		articles = append(articles, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

func (m *ArticleModel) Edit(title, content string, id int) error {
	stmt := "UPDATE articles SET title = $1, content = $2 WHERE id = $3"
	ct, err := m.DB.Exec(context.Background(), stmt, title, content, id)
	if err != nil {
		log.Printf("Unable to EDIT: %v\n", err)
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("no affected rows")
	}

	return nil
}
