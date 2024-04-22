package repository

import (
	"context"
	"database/sql"
	"errors"
	"gokomodo_test/internal/entity"
	"log"
)

type productsRepository struct {
	db *sql.DB
}

func NewProducts(db *sql.DB) ProductRepositories {
	return &productsRepository{db: db}
}

func (p *productsRepository) Store(ctx context.Context, payload entity.Products) (*entity.Products, error) {
	var (
		result entity.Products
		err    error
	)

	// Begin transaction
	tx, err := p.db.BeginTx(ctx, nil)
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

	q := `INSERT INTO products(
            user_id,
            product_name,
            description,
            price,
            created_at,
            updated_at
			)
			VALUES (?,?,?,?,?,?)`

	qValues := []interface{}{
		payload.UserId,
		payload.ProductName,
		payload.Description,
		payload.Price,
		payload.CreatedAt,
		payload.UpdatedAt,
	}

	res, err := p.db.ExecContext(ctx, q, qValues...)
	if err != nil {
		log.Printf("got error executing query products: %v", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	result = entity.Products{
		Id:          int(id),
		UserId:      payload.UserId,
		ProductName: payload.ProductName,
		Description: payload.Description,
		Price:       payload.Price,
		CreatedAt:   payload.CreatedAt,
		UpdatedAt:   payload.UpdatedAt,
	}

	return &result, nil
}

func (p *productsRepository) FindOne(ctx context.Context, id int) (*entity.Products, error) {
	var (
		result entity.Products
		err    error
	)

	q := `SELECT 
			id,
			user_id,
            product_name,
            description,
            price,
            created_at,
            updated_at
		FROM products WHERE id = ?`

	rows, err := p.db.QueryContext(ctx, q, id)
	if err != nil {
		log.Printf("got error when find products %v", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&result.Id, &result.UserId, &result.ProductName, &result.Description, &result.Price, &result.CreatedAt, &result.UpdatedAt)
		if err != nil {
			log.Printf("got error scan value %v", err)
			return nil, err
		}
		return &result, nil
	} else {
		return &result, errors.New("product is not found")
	}
}
