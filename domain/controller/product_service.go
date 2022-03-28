package controller

import (
	"errors"
	"fmt"
	"time"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"

	"github.com/bhanupbalusu/gomongoecomm1/domain/interface/repo"
	"github.com/bhanupbalusu/gomongoecomm1/domain/interface/service"
	"github.com/bhanupbalusu/gomongoecomm1/domain/model"
)

var (
	ErrProductNotFound = errors.New("Product Not Found")
	ErrProductInvalid  = errors.New("Product Invalid")
)

type productServices struct {
	ProductRepo repo.ProductRepoInterface
}

func NewProductServices(ProductRepo repo.ProductRepoInterface) service.ProductServiceInterface {
	return &productServices{ProductRepo}
}

func (p *productServices) Get() (model.ProductModelList, error) {
	return p.ProductRepo.Get()
}

func (p *productServices) GetByID(id string) (model.ProductModel, error) {
	fmt.Println("---- from inside domain.controller.product_service.GetByID -----")
	return p.ProductRepo.GetByID(id)
}

func (p *productServices) Create(pm model.ProductModel) (string, error) {
	fmt.Println("------- Inside ProductService.Create Before Calling validate.Validate -----------")
	if err := validate.Validate(pm); err != nil {
		return "", errs.Wrap(ErrProductInvalid, "domain.controller.product_service.Create")
	}
	fmt.Println("------- Inside ProductService.Create Before Calling ProductRepo.Create -----------")
	pm.CreatedAt = time.Now().UTC().Unix()
	return p.ProductRepo.Create(pm)
}

func (p *productServices) Update(pm model.ProductModel, pid string) error {
	if err := validate.Validate(pm); err != nil {
		return errs.Wrap(ErrProductInvalid, "domain.controller.product_service.Update")
	}
	pm.UpdatedAt = time.Now().UTC().Unix()
	return p.ProductRepo.Update(pm, pid)
}

func (p *productServices) Delete(pid string) error {
	return p.ProductRepo.Delete(pid)
}
