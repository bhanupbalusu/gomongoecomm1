package utils

import (
	"fmt"
	"log"

	h "github.com/bhanupbalusu/gomongoecomm1/api/handler"
	"github.com/bhanupbalusu/gomongoecomm1/db"
	c "github.com/bhanupbalusu/gomongoecomm1/domain/controller"
	r "github.com/bhanupbalusu/gomongoecomm1/domain/interface/repo"
)

func Init() h.RedirectHandler {
	repo := getRepo()
	fmt.Println(repo)
	service := c.NewProductServices(repo)
	return h.NewHandler(service)
}

func getRepo() r.ProductRepoInterface {
	fmt.Println("............Getting mongodb repository .............")
	repo, err := db.NewMongoRepository("mongodb://localhost:27017", "ecomm-products", 30)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
