package repository

import (
	"context"
	"database/sql"
	"errors"
	"gokomodo_test/internal/entity"
	"log"
)

type orderDetailsRepository struct {
	db *sql.DB
}

func NewOrderDetails(db *sql.DB) OrderDetailsRepositories {
	return &orderDetailsRepository{db: db}
}

func (o *orderDetailsRepository) Store(ctx context.Context, payload entity.OrderDetails) (*entity.OrderDetails, error) {
	var (
		result entity.OrderDetails
		err    error
	)

	// Begin transaction
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			// Rollback the transaction if an error occurred
			tx.Rollback()
			return
		}
		// Commit the transaction if no error occurred
		err = tx.Commit()
	}()

	q := `INSERT INTO order_details(
            product_id,
            order_id,
            quantity,
            total_price,
            status,
            created_at,
            updated_at
			)
			VALUES (?,?,?,?,?,?,?)`

	qValues := []interface{}{
		payload.ProductId,
		payload.OrderId,
		payload.Quantity,
		payload.TotalPrice,
		payload.Status,
		payload.CreatedAt,
		payload.UpdatedAt,
	}

	res, err := o.db.ExecContext(ctx, q, qValues...)
	if err != nil {
		log.Printf("got error executing query order details: %v", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	result = entity.OrderDetails{
		Id:         int(id),
		ProductId:  payload.ProductId,
		OrderId:    payload.OrderId,
		Quantity:   payload.Quantity,
		TotalPrice: payload.TotalPrice,
		Status:     payload.Status,
		CreatedAt:  payload.CreatedAt,
		UpdatedAt:  payload.UpdatedAt,
	}

	return &result, nil
}

func (o *orderDetailsRepository) AcceptOrderBuyer(ctx context.Context, payload entity.OrderDetails) error {
	// Begin transaction
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			// Rollback the transaction if an error occurred
			tx.Rollback()
			return
		}
		// Commit the transaction if no error occurred
		err = tx.Commit()
	}()

	q := `UPDATE order_details 
			SET status = ?,
			    updated_at = ?
			WHERE id = ?`

	_, err = tx.ExecContext(ctx, q, payload.Status, payload.UpdatedAt, payload.Id)
	if err != nil {
		log.Printf("got error executing query order detail: %v", err)
		return err
	}

	return nil
}

func (o *orderDetailsRepository) FindOne(ctx context.Context, id int) (*entity.OrderDetails, error) {
	var (
		result entity.OrderDetails
		err    error
	)

	q := `SELECT 
			id,
			product_id,
            order_id,
            quantity,
            total_price,
            status,
            created_at,
            updated_at
		FROM order_details WHERE id = ?`

	rows, err := o.db.QueryContext(ctx, q, id)
	if err != nil {
		log.Printf("got error when find products %v", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&result.Id, &result.ProductId, &result.OrderId, &result.Quantity, &result.TotalPrice, &result.Status, &result.CreatedAt, &result.UpdatedAt)
		if err != nil {
			log.Printf("got error scan value %v", err)
			return nil, err
		}
		return &result, nil
	} else {
		return &result, errors.New("order detail is not found")
	}
}
