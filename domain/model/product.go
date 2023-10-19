package model

import (
	"hash/crc32"
	"math"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/asaskevich/govalidator"
)

type ProductRepositoryInterface interface {
	FindAll() (*[]Product, error)
	Create(product *Product) (*Product, error)
}

type Product struct {
	Base        `valid:"required"`
	Name        string  `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	Price       float32 `json:"price" gorm:"default:0" valid:"-"`
}

func (product *Product) isValid() error {
	_, err := govalidator.ValidateStruct(product)
	return err
}

func NewProduct(name string, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	uuid := uuid.NewV4()
	product.ID = int32(math.Abs(float64(crc32.ChecksumIEEE(uuid[:]))))
	product.CreatedAt = time.Now()

	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
