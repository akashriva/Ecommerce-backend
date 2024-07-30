package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Product modle `products` table

type Product struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Price float64 `bson:"price" json:"price"`
	Image []string`bson:"image" json:"image"`
	Metainfo  map[string]interface{} `bson:"metainfo,omitempty" json:"metainfo"`
}
