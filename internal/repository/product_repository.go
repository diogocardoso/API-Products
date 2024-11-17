package repository

import (
	"API-Products/internal/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) *ProductRepository {
	return &ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	if pr.connection == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := "SELECT id, name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productList []model.Product
	for rows.Next() {
		var productObj model.Product
		if err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price); err != nil {
			return nil, err
		}
		productList = append(productList, productObj)
	}

	return productList, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = ?")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product
	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	// Prepara a consulta SQL
	query, err := pr.connection.Prepare("INSERT INTO product (name, price) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	// Executa a query
	_, err = query.Exec(product.Name, product.Price)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Retorna o ID do INSERT
	var id int
	err = pr.connection.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
