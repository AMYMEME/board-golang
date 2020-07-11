package repository

import (
	"time"

	"github.com/AMYMEME/board-golang/model"

	"github.com/pkg/errors"
)

func (d *DBConfig) AddPost(post model.Post) (int, error) {
	res, err := d.MyDB.Exec("INSERT INTO board.post (member_id, title, contents, datetime) VALUES (?, ?, ?, ?)",
		post.MemberID, post.Title, post.Contents, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return 0, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		err := errors.Wrap(err, "Fail getting added row's id")
		return 0, err
	}
	return int(ID), nil
}

func (d *DBConfig) GetAllPosts() ([]model.Post, error) {
	rows, err := d.MyDB.Query("SELECT id, member_id, title, datetime FROM board.post")

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return nil, err
	}
	var results []model.Post
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.MemberID, &post.Title, &post.Datetime)
		if err != nil {
			err := errors.Wrap(err, "Fail loading sql results")
			return nil, err
		}
		results = append(results, post)
	}
	return results, nil
}

func (d *DBConfig) GetPost(ID int) (model.Post, error) {
	var post model.Post
	err := d.MyDB.QueryRow("SELECT id, member_id, title, contents, datetime FROM board.post WHERE id = ?", ID).
		Scan(&post.ID, &post.MemberID, &post.Title, &post.Contents, &post.Datetime)

	if err != nil {
		err := errors.Wrap(err, "There is no such post")
		return model.Post{}, err
	}
	return post, nil
}

func (d *DBConfig) DeletePost(ID int) error {
	res, err := d.MyDB.Exec("DELETE FROM board.post WHERE id = ?", ID)

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return err
	}

	n, _ := res.RowsAffected()
	if n != 1 {
		err := errors.Wrap(err, "There is no such post")
		return err
	}
	return nil
}

func (d *DBConfig) UpdatePost(ID int, newPost model.Post) error {
	res, err := d.MyDB.Exec("UPDATE board.post SET title = ?, contents = ?, datetime = ? WHERE id = ?",
		newPost.Title, newPost.Contents, time.Now().Format("2006-01-02 15:04:05"), ID)

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return err
	}

	n, _ := res.RowsAffected()
	if n != 1 {
		err := errors.Wrap(err, "There is no such post")
		return err
	}
	return nil
}
