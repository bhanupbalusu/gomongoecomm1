package repo

import (
	"github.com/bhanupbalusu/gomongoecomm1/domain/model"
)

type ProductRepoInterface interface {
	Get() (model.ProductModelList, error)
	GetByID(id string) (model.ProductModel, error)
	Create(pm model.ProductModel) (string, error)
	Update(pm model.ProductModel, pid string) error
	Delete(pid string) error
}
