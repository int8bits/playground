package connectors

import "database/sql"

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(url string) (*MysqlRepository, error) {
	db, err := sql.Open("mysql", url)

	if err != nil {
		return nil, err
	}

	return &MysqlRepository{db: db}, nil
}

func (repo *MysqlRepository) Close() error {
	return repo.db.Close()
}
