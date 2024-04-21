package repository

import (
	"context"
	"database/sql"
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
