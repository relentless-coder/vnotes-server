package notebooks

import "database/sql"

type notebooks struct {
	Id        int               `json:"id"`
	Title     string            `json:"title"`
	Config    map[string]string `json:"config"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

func (n *notebooks) create(db *sql.DB) error {
	_, err := db.Exec("insert into notebooks(title,config,created_at,updated_at) values($1,$2,$3,$4)", n.Title, n.Config, n.CreatedAt, n.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (n *notebooks) update(db *sql.DB) error {
	_, err := db.Exec("update notebooks set title=$1 where id=$2", n.Title, n.Id)
	if err != nil {
		return err
	}
	return nil
}

func (n *notebooks) get(db *sql.DB) ([]notebooks, error) {
	nbs := []notebooks{}
	rows, err := db.Query("select * from notebooks")
	if err != nil {
		return nbs, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec notebooks
		if err := rows.Scan(&rec.Title, &rec.Id, &rec.Config, &rec.CreatedAt, &rec.UpdatedAt); err != nil {
			return nbs, err
		}
		nbs = append(nbs, rec)
	}
	return nbs, nil
}

func (n *notebooks) getSingle(db *sql.DB) (notebooks, error) {
	var nbs notebooks
	rows, err := db.Query("select * from notebooks where id=$1", n.Id)
	if err != nil {
		return nbs, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&nbs.Title, &nbs.Id, &nbs.Config, &nbs.CreatedAt, &nbs.UpdatedAt); err != nil {
			return nbs, err
		}
	}
	return nbs, nil
}
