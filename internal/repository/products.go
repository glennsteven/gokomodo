package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (p *productsRepository) Listing(ctx context.Context, limit, offset, authId int) ([]entity.ListingProduct, error) {
	query := `
		SELECT
			p.id,
			p.user_id,
			u.full_name,
			u.address,
			p.product_name,
			p.description,
			p.price,
			p.created_at,
			p.updated_at
		FROM 
			products p
		JOIN 
			users AS u ON u.id = p.user_id
		WHERE user_id = ? 
		ORDER BY 
			p.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := p.db.QueryContext(ctx, query, authId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var products []entity.ListingProduct

	// Iterate over the rows and scan each result into a ListingProduct struct
	for rows.Next() {
		var product entity.ListingProduct
		err := rows.Scan(
			&product.Id,
			&product.UserId,
			&product.FullName,
			&product.Address,
			&product.ProductName,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		// Append the scanned product to the slice
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration: %v", err)
	}

	return products, nil
}

func (p *productsRepository) Count(ctx context.Context, authID int) (int, error) {
	query := `SELECT COUNT(*) FROM products WHERE user_id = ? `

	// Execute the query and scan the result
	var count int
	err := p.db.QueryRowContext(ctx, query, authID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count products: %v", err)
	}

	return count, nil
}

func (p *productsRepository) ListingOrderSeller(ctx context.Context, limit, offset, authId int) ([]entity.ListingOrderDetails, error) {
	query := `
		SELECT
			p.id as product_id,
			p.user_id as seller_id,
			u.full_name as seller,
			u.address as seller_address,
			us.id as buyer_id,
			us.full_name as buyer,
			us.address as buyer_address,
			od.status,
			od.quantity,
			od.total_price,
            product_name,
            description,
            price,
            p.created_at,
            p.updated_at
		FROM products p
		JOIN users AS u ON u.id = p.user_id
		JOIN order_details AS od ON od.product_id = p.id
		JOIN orders o on o.id = od.order_id
		JOIN users AS us ON us.id = o.user_id
		WHERE p.user_id = ? 
        ORDER BY created_at DESC LIMIT ? OFFSET ? 
	`

	rows, err := p.db.QueryContext(ctx, query, authId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	defer rows.Close()

	var orderDetails []entity.ListingOrderDetails

	// Iterate over the rows and scan each result into a ListingOrderDetails struct
	for rows.Next() {
		var orderDetail entity.ListingOrderDetails
		err := rows.Scan(
			&orderDetail.Id,
			&orderDetail.SellerId,
			&orderDetail.Seller,
			&orderDetail.SellerAddress,
			&orderDetail.BuyerId,
			&orderDetail.Buyer,
			&orderDetail.BuyerAddress,
			&orderDetail.Status,
			&orderDetail.Quantity,
			&orderDetail.TotalPrice,
			&orderDetail.ProductName,
			&orderDetail.Description,
			&orderDetail.Price,
			&orderDetail.CreatedAt,
			&orderDetail.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		// Append the scanned product to the slice
		orderDetails = append(orderDetails, orderDetail)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration: %v", err)
	}

	return orderDetails, nil
}

func (p *productsRepository) CountOrderSeller(ctx context.Context, authID int) (int, error) {
	query := `
		SELECT COUNT(*) AS count
		FROM products p
		JOIN users AS u ON u.id = p.user_id
		JOIN order_details AS od ON od.product_id = p.id
		JOIN orders o ON o.id = od.order_id
		JOIN users AS us ON us.id = o.user_id
		WHERE p.user_id = ?
		GROUP BY p.user_id
	`

	// Execute the query and scan the result
	var count int
	err := p.db.QueryRowContext(ctx, query, authID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count orders for seller: %v", err)
	}

	return count, nil
}
