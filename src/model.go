package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]Product, error) {

	query := "select id, name, quantity, price from products"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	products := []Product{}

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (product *Product) getProduct(db *sql.DB) error {
	query := fmt.Sprintf("select id, name, quantity, price from products where id = %v", product.ID)
	row := db.QueryRow(query)
	err := row.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)

	if err != nil {
		return err
	}
	return nil
}

func (product *Product) createProduct(db *sql.DB) error {
	query := fmt.Sprintf("insert into products(name, quantity, price) values('%v', %v, %v)", product.Name, product.Quantity, product.Price)

	result, err := db.Exec(query)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	product.ID = int(id)
	return nil
}

func (product *Product) updateProduct(db *sql.DB) error {
	query := fmt.Sprintf("update products set name='%v', quantity=%v, price=%v where id=%v", product.Name, product.Quantity, product.Price, product.ID)

	result, err := db.Exec(query)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("No Such product exists")
	}

	return nil
}
