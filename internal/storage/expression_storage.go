package storage

import (
	"context"
)

// Model methods for expressions table in database

type Expression struct {
	ID         int64
	UserID     int64
	Expression string
	Answer     string // Types: null, {answer}
	Date       string
	Status     string // Types: stored, done, error
}

// InsertExpression - putting new expression into database table
func (s *Storage) InsertExpression(ctx context.Context, expr *Expression) (int64, error) {

	var q = `
	INSERT INTO expressions (userid, expression, answer, date, status) values ($1, $2, $3, $4, $5)
	`

	res, err := s.db.ExecContext(ctx, q, expr.UserID, expr.Expression, expr.Answer, expr.Date, expr.Status)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// SelectExpressions - returns all expressions slice from database table
func (s *Storage) SelectAllExpressions(ctx context.Context) ([]Expression, error) {

	var (
		expressions []Expression
		q           = `SELECT id, expression, answer, date, status FROM expressions`
	)

	rows, err := s.db.QueryContext(ctx, q)
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

// SelectExpressionByID - guess yourself :)
func (s *Storage) SelectExpressionsByID(ctx context.Context, userID int64) ([]Expression, error) {

	var ( 
		q = `SELECT id, expression, answer, date, status FROM expressions WHERE userid = $1`
		expressions []Expression
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

// UpdateExpression - updates data about expression, specifically its answer and/or status
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

// DeleteExpression - deletes expression row from database table
func (s *Storage) DeleteExpression(ctx context.Context, id int64) error {

	var q = `DELETE FROM expressions WHERE id = ?`

	_, err := s.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}
