package repository

import (
	"context"
	"database/sql"
	"fmt"
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

func (o *ordersRepository) Listing(ctx context.Context, limit, offset, authID int) ([]entity.ListingOrderBuyer, error) {
	query := `
        SELECT
			o.id,
			p.product_name,
			p.description,
			u.id as id_seller,
			u.full_name as seller,
			us.id as id_buyer,
			us.full_name as buyer,
			us.email,
			o.delivery_destination_address,
			o.delivery_source_address,
			od.quantity,
			od.total_price,
			od.status,
			o.created_at
		FROM orders o
		JOIN order_details od on o.id = od.order_id
		JOIN products p on od.product_id = p.id
		JOIN users u on u.id = p.user_id
		JOIN users us on us.id = o.user_id
		WHERE o.user_id = ?
		ORDER BY created_at DESC LIMIT ? OFFSET ? `

	rows, err := o.db.QueryContext(ctx, query, authID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var orderBuyers []entity.ListingOrderBuyer

	for rows.Next() {
		var orderBuyer entity.ListingOrderBuyer
		if err := rows.Scan(
			&orderBuyer.Id,
			&orderBuyer.ProductName,
			&orderBuyer.Description,
			&orderBuyer.IdSeller,
			&orderBuyer.Seller,
			&orderBuyer.IdBuyer,
			&orderBuyer.Buyer,
			&orderBuyer.Email,
			&orderBuyer.DeliveryDestinationAddress,
			&orderBuyer.DeliverySourceAddress,
			&orderBuyer.Quantity,
			&orderBuyer.TotalPrice,
			&orderBuyer.Status,
			&orderBuyer.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		orderBuyers = append(orderBuyers, orderBuyer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration: %v", err)
	}

	return orderBuyers, nil
}

func (o *ordersRepository) Count(ctx context.Context, authID int) (int, error) {
	query := `SELECT COUNT(*) AS total_count FROM orders o
              JOIN order_details od ON o.id = od.order_id
              JOIN products p ON od.product_id = p.id
              JOIN users u ON u.id = p.user_id
              JOIN users us on us.id = o.user_id
              WHERE o.user_id = ?`

	var count int
	err := o.db.QueryRowContext(ctx, query, authID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count orders: %v", err)
	}

	return count, nil
}
