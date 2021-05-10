package pages

import "database/sql"

type pages struct {
	Id      int `json:"id"`
	Content []struct {
		Step   string `json:"step"`
		Points []int  `json:"points",omitempty`
	} `json:"content",omitempty`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p *pages) create(db *sql.DB) error {
	_, err := db.Exec("insert into pages(content,created_at,updated_at) values($1,$2,$3)", p.Content, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *pages) update(db *sql.DB) error {
	_, err := db.Exec("update pages set content=$1, updated_at=$2 where id=$3", p.Content, p.UpdatedAt, p.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *pages) get(db *sql.DB) ([]pages, error) {
	rows, err := db.Query("select content, id from pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	notePages := []pages{}
	for rows.Next() {
		var p pages
		if err := rows.Scan(&p.Content, &p.Id); err != nil {
			return nil, err
		}
		notePages = append(notePages, p)
	}
	return notePages, nil
}
