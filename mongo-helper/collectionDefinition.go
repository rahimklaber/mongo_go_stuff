package mongo_helper

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	OrderId   primitive.ObjectID `bson:"order_id"`
	Paid      bool               `bson:"paid"`
	Items     []string           `bson:"items"`
	UserId    primitive.ObjectID `bson:"user_id"`
	TotalCost float64            `bson:"total_cost"`
}
