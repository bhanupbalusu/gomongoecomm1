package db

import (
	"context"
	"fmt"
	"log"

	contr "github.com/bhanupbalusu/gomongoecomm1/domain/controller"
	"github.com/bhanupbalusu/gomongoecomm1/domain/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get all products
func (r *MongoRepository) Get() (model.ProductModelList, error) {
	var results model.ProductModelList

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_coll1")
	fmt.Println(coll)

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments || err == mongo.ErrNilValue {
			return nil, errors.Wrap(contr.ErrProductNotFound, "db.repository.Get")
		}
		return nil, errors.Wrap(err, "db.repository.Get")
	}

	if err = cursor.All(ctx, &results); err != nil {
		errors.Wrap(err, "db.repository.Get.cursor.All")
		log.Fatal(err)
	}
	return results, nil
}

// Get single product using id
func (r *MongoRepository) GetByID(id string) (model.ProductModel, error) {
	var result model.ProductModel
	fmt.Println("------- Inside db/repository.GetByID Before Calling r.GetCollection -----------")
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_coll1")
	fmt.Println(coll)

	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": newId}

	fmt.Println(newId)
	fmt.Println("------- Inside db/repository.GetByID Before Calling coll.FindOne -----------")
	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, errors.Wrap(contr.ErrProductNotFound, "domain.repository.mongo.repository.GetProductByID")
		}
		return result, errors.Wrap(err, "domain.repository.mongo.repository.GetProductByID")
	}
	return result, nil
}

// Create or insert a new product
func (r *MongoRepository) Create(pm model.ProductModel) (string, error) {
	fmt.Println("------- Inside db/repository.Create Before Calling r.GetCollection -----------")

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_coll1")
	fmt.Println(coll)

	fmt.Println("------- Inside db/repository.Create Before Calling coll.InsertOne -----------")
	result, err := coll.InsertOne(
		ctx,
		bson.M{
			"pre_order_request_id": pm.PreOrderRequestId,
			"customer_id":          pm.CustomerId,
			"product_details": bson.M{
				"product_name": pm.ProductDetails.ProductName,
				"description":  pm.ProductDetails.Description,
				"ImageUrl":     pm.ProductDetails.ImageUrl,
			},
			"quantity_details": bson.M{
				"bulk_quantity": bson.M{
					"volume": pm.QuantityDetails.BulkQuantity.Volume,
					"units":  pm.QuantityDetails.BulkQuantity.Units,
				},
				"price": bson.M{
					"amount":   pm.QuantityDetails.Price.Amount,
					"currency": pm.QuantityDetails.Price.Currency,
					"per_unit": pm.QuantityDetails.Price.PerUnit,
					"units":    pm.QuantityDetails.Price.Units,
				},
			},
			"schedular": bson.M{
				"start_date": pm.Schedular.StartDate,
				"end_date":   pm.Schedular.EndDate,
			},
			"created_at": pm.CreatedAt,
			"updated_at": pm.UpdatedAt,
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "db.repository.Create")
	}
	pid := (result.InsertedID).(primitive.ObjectID).Hex()
	return pid, nil
}

// Update existing product
func (r *MongoRepository) Update(pm model.ProductModel, id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_coll1")
	fmt.Println(coll, ctx)

	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": pid}

	update := bson.M{
		"$set": bson.M{
			"pre_order_request_id":                  pm.PreOrderRequestId,
			"customer_id":                           pm.CustomerId,
			"product_details.product_name":          pm.ProductDetails.ProductName,
			"product_details.description":           pm.ProductDetails.Description,
			"product_details.ImageUrl":              pm.ProductDetails.ImageUrl,
			"quantity_details.bulk_quantity.volume": pm.QuantityDetails.BulkQuantity.Volume,
			"quantity_details.bulk_quantity.units":  pm.QuantityDetails.BulkQuantity.Units,
			"quantity_details.price.amount":         pm.QuantityDetails.Price.Amount,
			"quantity_details.price.currency":       pm.QuantityDetails.Price.Currency,
			"quantity_details.price.per_unit":       pm.QuantityDetails.Price.PerUnit,
			"quantity_details.price.units":          pm.QuantityDetails.Price.Units,
			"schedular.start_date":                  pm.Schedular.StartDate,
			"schedular.end_date":                    pm.Schedular.EndDate,
		},
	}

	fmt.Println(update)

	_, err = coll.UpdateOne(ctx, filter, update)

	return err
}

// Delete an existing product
func (r *MongoRepository) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_coll1")
	fmt.Println(coll)

	pid, err := primitive.ObjectIDFromHex(id)
	_, err = coll.DeleteOne(ctx, bson.M{"_id": pid})
	if err != nil {
		if err == mongo.ErrNoDocuments || err == mongo.ErrNilValue {
			return errors.Wrap(contr.ErrProductNotFound, "domain.repository.mongo.repository.DeleteProduct")
		}
		return errors.Wrap(err, "domain.repository.mongo.repository.DeleteProduct")
	}
	return nil
}
