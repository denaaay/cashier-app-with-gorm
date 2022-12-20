package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	result := []model.JoinCart{}
	c.db.Table("products").Select("carts.id, carts.product_id, products.name, carts.quantity, carts.total_price").Joins("left join carts on carts.product_id = products.id").Scan(&result)
	return result, nil // TODO: replace this
}

func (c *CartRepository) AddCart(product model.Product) error {
	totalHarga := product.Price - ((product.Discount / 100) * product.Price)
	cart := model.Cart{
		ProductID:  product.ID,
		Quantity:   1,
		TotalPrice: totalHarga,
	}

	carts := model.Cart{}
	resp := c.db.Raw("SELECT * FROM carts WHERE id = ?", product.ID).Scan(&carts)
	product.Stock--
	if resp.RowsAffected == 0 {
		c.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&cart).Error; err != nil {
				return err
			}
			return nil
		})
	}

	c.db.Transaction(func(tx *gorm.DB) error {
		if product.Stock == 0 {
			return errors.New("Product 0")
		}

		tx.Model(&model.Product{}).Where("id = ?", product.ID).Updates(product)
		cart.Quantity += carts.Quantity
		cart.TotalPrice += carts.TotalPrice
		tx.Model(&model.Cart{}).Where("id = ?", product.ID).Updates(cart)

		return nil
	})

	return nil // TODO: replace this
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {
	c.db.Transaction(func(tx *gorm.DB) error {
		cart := model.Cart{}
		tx.Raw("SELECT * FROM carts WHERE id = ?", id).Scan(&cart)

		result := tx.Where("id = ?", id).Delete(&model.Cart{})
		if result.Error != nil {
			return result.Error
		}

		product := model.Product{}
		tx.Raw("SELECT * FROM products WHERE id = ?", productID).Scan(&product)

		product.Stock += int(cart.Quantity)

		tx.Model(&model.Product{}).Where("id = ?", productID).Update("stock", product.Stock)
		return nil
	})
	return nil // TODO: replace this
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	result := c.db.Model(&model.Cart{}).Where("id = ?", id).Updates(cart)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
