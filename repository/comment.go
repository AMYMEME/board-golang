package repository

import (
	"time"

	"github.com/AMYMEME/board-golang/model"

	"github.com/pkg/errors"
)

func (d *DBConfig) AddComment(comment model.Comment) error {
	_, err := d.MyDB.Exec("INSERT INTO board.comment (post_id, member_id, contents, datetime) VALUES (?, ?, ?, ?)",
		comment.PostID, comment.MemberID, comment.Contents, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return err
	}
	return nil
}

func (d *DBConfig) GetAllCommentsByPostId(postID int) ([]model.Comment, error) {
	rows, err := d.MyDB.Query("SELECT id, post_id, member_id, contents, datetime FROM board.comment WHERE post_id = ?", postID)

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return nil, err
	}
	var results []model.Comment
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.MemberID, &comment.Contents, &comment.Datetime)
		if err != nil {
			err := errors.Wrap(err, "Fail loading sql results")
			return nil, err
		}
		results = append(results, comment)
	}
	return results, nil
}

func (d *DBConfig) GetComment(ID int) (model.Comment, error) {
	var comment model.Comment
	err := d.MyDB.QueryRow("SELECT id, post_id, member_id, contents, datetime FROM board.comment WHERE id = ?", ID).
		Scan(&comment.ID, &comment.PostID, &comment.MemberID, &comment.Contents, &comment.Datetime)

	if err != nil {
		err := errors.Wrap(err, "There is no such comment")
		return model.Comment{}, err
	}
	return comment, nil
}

func (d *DBConfig) DeleteComment(ID int) error {
	res, err := d.MyDB.Exec("DELETE FROM board.comment WHERE id = ?", ID)

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return err
	}

	n, _ := res.RowsAffected()
	if n != 1 {
		err := errors.Wrap(err, "There is no such comment")
		return err
	}
	return nil
}

func (d *DBConfig) UpdateComment(ID int, newComment model.Comment) error {
	res, err := d.MyDB.Exec("UPDATE board.comment SET contents = ? WHERE id = ?",
		newComment.Contents, ID)

	if err != nil {
		err := errors.Wrap(err, "Fail sql query by Invalid Input")
		return err
	}

	n, _ := res.RowsAffected()
	if n != 1 {
		err := errors.Wrap(err, "There is no such comment")
		return err
	}
	return nil
}
