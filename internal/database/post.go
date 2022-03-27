package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DanielMartin96/posts/internal/post"
	uuid "github.com/satori/go.uuid"
)

type PostRow struct {
	ID string
	Slug sql.NullString
	Author sql.NullString
	Body sql.NullString
}

func convertPostRowToPost(p PostRow) post.Post {
	return post.Post{
		ID: p.ID,
		Slug: p.Slug.String,
		Author: p.Author.String,
		Body: p.Body.String,
	}
}

func (d *Database) GetPost(ctx context.Context, uuid string) (post.Post, error) {
	var pstRow PostRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, author, body
		FROM posts
		WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&pstRow.ID, &pstRow.Slug, &pstRow.Author, &pstRow.Body)
	if err != nil {
		return post.Post{}, fmt.Errorf("an error occurred fetching a comment by uuid: %w", err)
	}

	return convertPostRowToPost(pstRow), nil
}

func (d *Database) CreatePost(ctx context.Context, pst post.Post) (post.Post, error)  {
	pst.ID = uuid.NewV4().String()
	postRow := PostRow{
		ID: pst.ID,
		Slug: sql.NullString{String: pst.Slug, Valid: true},
		Author: sql.NullString{String: pst.Author, Valid: true},
		Body: sql.NullString{String: pst.Slug, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO posts 
		(id, slug, author, body) VALUES
		(:id, :slug, :author, :body)`,
		postRow,
	)
	if err != nil {
		return post.Post{}, fmt.Errorf("failed to insert post: %w", err)
	}
	if err := rows.Close(); err != nil {
		return post.Post{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return pst, nil
}

func (d *Database) UpdatePost(ctx context.Context, id string, pst post.Post) (post.Post, error) {
	pstRow := PostRow{
		ID:     id,
		Slug:   sql.NullString{String: pst.Slug, Valid: true},
		Author:   sql.NullString{String: pst.Author, Valid: true},
		Body: sql.NullString{String: pst.Body, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE posts SET
		slug = :slug,
		author = :author,
		body = :body 
		WHERE id = :id`,
		pstRow,
	)
	if err != nil {
		return post.Post{}, fmt.Errorf("failed to insert post: %w", err)
	}
	if err := rows.Close(); err != nil {
		return post.Post{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertPostRowToPost(pstRow), nil
}

func (d *Database) DeletePost(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM posts where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete post from the database: %w", err)
	}

	return nil
}