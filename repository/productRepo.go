package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AbdelilahOu/GoThingy/model"
)

type ProductRepo struct {
	DB *sql.DB
}

func (repo *ProductRepo) Insert(ctx context.Context, product model.Product) error {
	_, err := repo.DB.Exec("INSERT INTO products (id, name, description, price, tva) VALUES ($1, $2, $3, $4, $5)", product.Id, product.Name, product.Description, product.Price, product.Tva)
	if err != nil {
		fmt.Println("error inserting product", err)
		return err
	}
	return nil

}

func (repo *ProductRepo) Update(ctx context.Context, product model.Product, id string) error {
	return nil
}

func (repo *ProductRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func (repo *ProductRepo) Select(ctx context.Context, id string) (model.Product, error) {
	// execute
	row := repo.DB.QueryRow("SELECT * FROM products WHERE id = $1", id)
	// var
	var p model.Product
	// get product
	err := row.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Tva)
	// check for err
	if err != nil {
		fmt.Println("error scanning product", err)
		return model.Product{}, err
	}
	//
	return p, nil
}

type GetPAllResult struct {
	Products []model.Product
	Cursor   uint64
}

func (repo *ProductRepo) SelectAll(ctx context.Context, cursor uint64, size uint64) (GetPAllResult, error) {
	// get products
	rows, err := repo.DB.Query("SELECT * FROM products WHERE id > $1 LIMIT $2", cursor, size)
	if err != nil {
		fmt.Println("error getting products", err)
		return GetPAllResult{}, err
	}
	// close rows after
	defer rows.Close()
	// get products as model.client
	var products []model.Product
	for rows.Next() {
		var c model.Product
		// scane
		err := rows.Scan(&c.Id, &c.Name)
		if err != nil {
			fmt.Println("error scanning products", err)
			return GetPAllResult{}, err
		}
		//
		products = append(products, c)

	}
	// error while eterating
	err = rows.Err()
	if err != nil {
		fmt.Println("error eterating over rows")
	}
	// last result
	return GetPAllResult{
		Products: products,
	}, nil
}
