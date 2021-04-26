package main

import "database/sql"

type pages struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p *pages) createPage(db *sql.DB) error {
	_, err := db.Exec("insert into pages(content,created_at,updated_at) values($1,$2,$3,$4)", p.Content, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *pages) updatePage(db *sql.DB) error {
	_, err := db.Exec("update pages set content=$1, updated_at=$2 where id=$3", p.Content, p.UpdatedAt, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func getPages(db *sql.DB) ([]pages, error) {
	rows, err := db.Query("select content, id from pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	notePages := []pages{}
	for rows.Next() {
		var p pages
		if err := rows.Scan(&p.Content, &p.ID); err != nil {
			return nil, err
		}
		notePages = append(notePages, p)
	}
	return notePages, nil
}
