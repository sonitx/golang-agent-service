package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type SessionStatusRepository struct {
	db *sqlx.DB
}

func NewSessionStatusRepository(db *sqlx.DB) *SessionStatusRepository {
	return &SessionStatusRepository{db: db}
}

func (s *SessionStatusRepository) GetSessionStatus(ctx context.Context, sqlQuery string) ([]map[string]interface{}, error) {
	var datas []map[string]interface{}
	err := s.db.SelectContext(ctx, &datas, sqlQuery)
	return datas, err
}
