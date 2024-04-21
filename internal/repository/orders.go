package repository

import (
	"context"
	"database/sql"
	"gokomodo_test/internal/entity"
	"log"
)

type ordersRepository struct {
	db *sql.DB
}

func NewOrders(db *sql.DB) OrderRepositories {
	return &ordersRepository{db: db}
}

func (o *ordersRepository) Store(ctx context.Context, payload entity.Orders) (*entity.Orders, error) {
	var (
		result entity.Orders
		err    error
	)

	q := `INSERT INTO orders(
            user_id,
            delivery_source_address,
            delivery_destination_address,
            created_at,
            updated_at
			)
			VALUES (?,?,?,?,?)`

	qValues := []interface{}{
		payload.UserId,
		payload.DeliverySourceAddress,
		payload.DeliveryDestinationAddress,
		payload.CreatedAt,
		payload.UpdatedAt,
	}

	res, err := o.db.ExecContext(ctx, q, qValues...)
	if err != nil {
		log.Printf("got error executing query products: %v", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	result = entity.Orders{
		Id:                         int(id),
		UserId:                     payload.UserId,
		DeliverySourceAddress:      payload.DeliverySourceAddress,
		DeliveryDestinationAddress: payload.DeliveryDestinationAddress,
		CreatedAt:                  payload.CreatedAt,
		UpdatedAt:                  payload.UpdatedAt,
	}

	return &result, nil
}
