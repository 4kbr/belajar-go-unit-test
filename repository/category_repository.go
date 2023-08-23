package repository

import "belajar-go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
