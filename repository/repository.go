package repository

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"telkomsel-technical-test.com/config"
	"telkomsel-technical-test.com/dto"
)

type Repository interface {
	FetchProduct(brand, variety string, id string) ([]*dto.Product, error)
	Create(new_product *dto.Product) error
	Update(new_product *dto.Product) error
	Delete(product_id uuid.UUID) error
}

type ProductRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FetchProduct(brand, variety string, id string) ([]*dto.Product, error) {
	resp := []*dto.Product{}
	query := "SELECT * FROM product"
	var err error

	if len(id) > 0 {
		query += ` WHERE product_id = $1`
		err = config.DB.Select(&resp, query, id)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	if len(variety) > 0 && len(brand) > 0 {
		query += ` WHERE lower(product_variety) LIKE $1 AND lower(product_brand) LIKE $2`
		err = config.DB.Select(&resp, query, "%"+variety+"%", "%"+brand+"%")
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	if len(variety) > 0 {
		query += ` WHERE lower(product_variety) LIKE $1`
		err = config.DB.Select(&resp, query, "%"+variety+"%")
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	if len(brand) > 0 {
		query += ` WHERE lower(product_brand) LIKE $1`
		err = config.DB.Select(&resp, query, "%"+brand+"%")
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	err = config.DB.Select(&resp, query)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ProductRepository) Create(new_product *dto.Product) error {
	query := `INSERT INTO product (
	product_name, product_description, 
	product_price, product_variety, product_rating, 
	product_stock, product_url, product_brand) 
	VALUES(:product_name, :product_description, :product_price,
	:product_variety, :product_rating, :product_stock, :product_url, :product_brand);`

	_, err := config.DB.NamedExec(query, new_product)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Update(new_product *dto.Product) error {
	query := `UPDATE product SET 
	product_name = :product_name, product_description = :product_description, 
	product_price = :product_price, product_variety = :product_variety, product_rating = :product_rating, 
	product_stock = :product_stock, product_url = :product_url, product_brand = :product_brand 
	WHERE product_id = :product_id;`

	res, err := config.DB.NamedExec(query, new_product)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("no product found")
	}

	fmt.Println(affected)

	return nil
}

func (r *ProductRepository) Delete(product_id uuid.UUID) error {
	query := `DELETE FROM product WHERE product_id = $1;`

	res, err := config.DB.Exec(query, product_id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("no product found")
	}

	fmt.Println(affected)

	return nil
}
