package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product model.Product) error {
	result := p.db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (p *ProductRepository) ReadProducts() ([]model.Product, error) {
	result := []model.Product{}
	resp := p.db.Raw("SELECT * FROM products WHERE deleted_at is null").Scan(&result)
	if resp.Error != nil {
		return []model.Product{}, resp.Error
	}
	return result, nil // TODO: replace this
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	result := p.db.Where("id = ?", id).Delete(&model.Product{})
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (p *ProductRepository) UpdateProduct(id uint, product model.Product) error {
	result := p.db.Model(&model.Product{}).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
