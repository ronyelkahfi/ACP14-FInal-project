package products

import (
	"context"
	"errors"
	_productDomain "final-project/business/products"
	"final-project/helpers"
	"time"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) _productDomain.Repository {
	return &ProductRepository{
		db: database,
	}
}

func (repo *ProductRepository) GetProduct(ctx context.Context) ([]_productDomain.Domain, error) {

	var products []Product
	result := repo.db.Find(&products)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_productDomain.Domain{}, helpers.ErrNotFound
		} else {
			return []_productDomain.Domain{}, helpers.ErrDbServer
		}
	}

	return ToListDomain(products), nil
}

func (repo *ProductRepository) CreateProduct(ctx context.Context, data _productDomain.Domain) (int, error) {
	product := Product{
		Name:       data.Name,
		Price:      data.Price,
		CategoryId: data.CategoryId,
	}
	// fmt.Println(data)
	result := repo.db.Create(&product)
	return int(result.RowsAffected), result.Error
}
func (repo *ProductRepository) DeleteProduct(ctx context.Context, id int) (int, error) {
	product := Product{}
	// fmt.Println(data)
	result := repo.db.Model(&product).Where("id", id).Update("deleted_at", time.Now())
	return int(result.RowsAffected), result.Error
}

func (repo *ProductRepository) GetProductById(ctx context.Context, id uint) (_productDomain.Domain, error) {
	product := Product{}

	result := repo.db.First(&product).Where("id", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _productDomain.Domain{}, errors.New("Not FOund")
		} else {
			return _productDomain.Domain{}, result.Error
		}
	} else {
		return _productDomain.Domain{
			Price: product.Price,
		}, nil
	}
}
