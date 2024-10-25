package models

import (
	"database/sql"
	"fmt"
)

type Product struct {
	Id         int64   `json:"id"`
	NamaProduk string  `json:"nama_produk"`
	Deskripsi  string  `json:"deskripsi"`
	Harga      float32 `json:"harga"`
	Stok       int     `json:"stok"`
}

func GetProducts() ([]Product, error) {

	rows, err := DB.Query("SELECT * from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]Product, 0)

	for rows.Next() {
		singleProduct := Product{}
		err = rows.Scan(&singleProduct.Id, &singleProduct.NamaProduk, &singleProduct.Deskripsi, &singleProduct.Harga, &singleProduct.Stok)

		if err != nil {
			return nil, err
		}

		products = append(products, singleProduct)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return products, err
}
func GetProductById(id string) (Product, error) {

	stmt, err := DB.Prepare("SELECT * from products WHERE id = ?")

	if err != nil {
		return Product{}, err
	}

	product := Product{}

	sqlErr := stmt.QueryRow(id).Scan(&product.Id, &product.NamaProduk, &product.Deskripsi, &product.Harga, &product.Stok)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Product{}, nil
		}
		return Product{}, sqlErr
	}
	return product, nil
}
func AddProduct(newProduct Product) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO products (id,nama_produk,deskripsi,harga,stok) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newProduct.Id, newProduct.NamaProduk, newProduct.Deskripsi, newProduct.Harga, newProduct.Stok)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
func PutProductById(newProduct Product, id string) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE products SET nama_produk = ?, deskripsi = ?, harga = ?, stok = ? WHERE id = ?")
	fmt.Println(err)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newProduct.NamaProduk, newProduct.Deskripsi, newProduct.Harga, newProduct.Stok, id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
func DeleteProduct(id string) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("DELETE FROM products WHERE id = ?")
	fmt.Println(err)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
