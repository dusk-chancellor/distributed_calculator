package storage

import (
	"context"
)

type Expression struct {
	ID 		   int64
	UserID 	   int64
	Expression string
	Answer 	   string
	Date 	   string
	Status 	   string
}

// Status types: stored, processing, done, error

func (s *Storage) InsertExpression(ctx context.Context, expr *Expression) (int64, error) {

	var q = `
	INSERT INTO expressions (userid, expression, date, status) values ($1, $2, $3, $4)
	`
	
	res, err := s.db.ExecContext(ctx, q, expr.UserID, expr.Expression, expr.Date, expr.Status)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) SelectExpressions(ctx context.Context, userID int64) ([]Expression, error) {

	var (
		expressions []Expression
		q = `SELECT id, expression, answer, date, status FROM expressions WHERE userid = $1`
	)

	rows, err := s.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		e := Expression{}
		err := rows.Scan(&e.ID, &e.Expression, &e.Answer, &e.Date, &e.Status)
		if err != nil {
			return nil, err
		}

		expressions = append(expressions, e)
	}

	return expressions, nil
}

func (s *Storage) SelectExpressionByID(ctx context.Context, userID int64, id int64) (Expression, error) {

	e := Expression{}
	var q = `SELECT id, expression, answer, date, status FROM expressions WHERE userid = $1, id = $2`

	err := s.db.QueryRowContext(ctx, q, userID, id).Scan(&e.ID, &e.Expression, &e.Answer, &e.Date, &e.Status)
	if err != nil {
		return e, err
	}

	return e, nil
}

func (s *Storage) UpdateExpression(
	ctx context.Context, answer, status string, id int64,
	) error {

	var q = `UPDATE expressions SET answer = $1, status = $2 WHERE id = $3`

	_, err := s.db.ExecContext(ctx, q, answer, status, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteExpression(ctx context.Context, id int64) error {

	var q = `DELETE FROM expressions WHERE id = ?`

	_, err := s.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}
