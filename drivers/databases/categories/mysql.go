package categories

import (
	"context"
	_categoryDomain "final-project/business/categories"
	"final-project/helpers"
	"time"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(database *gorm.DB) _categoryDomain.Repository {
	return &CategoryRepository{
		db: database,
	}
}

func (repo *CategoryRepository) GetCategory(ctx context.Context) ([]_categoryDomain.Domain, error) {

	var categories []Category
	result := repo.db.Find(&categories)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_categoryDomain.Domain{}, helpers.ErrNotFound
		} else {
			return []_categoryDomain.Domain{}, helpers.ErrDbServer
		}
	}

	return ToListDomain(categories), nil
}

func (repo *CategoryRepository) CreateCategory(ctx context.Context, data _categoryDomain.Domain) (int, error) {
	category := Category{
		Name: data.Name,
	}
	// fmt.Println(data)
	result := repo.db.Create(&category)
	return int(result.RowsAffected), result.Error
}
func (repo *CategoryRepository) DeleteCategory(ctx context.Context, id int) (int, error) {
	category := Category{}
	// fmt.Println(data)
	result := repo.db.Model(&category).Where("id", id).Update("deleted_at", time.Now())
	return int(result.RowsAffected), result.Error
}
