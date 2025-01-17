package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Ingredients string             `bson:"ingredients,omitempty" json:"ingredients,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Calories    float64            `bson:"calories,omitempty" json:"calories,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"` 
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
}

type Customer struct {
	ID        	 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      	 string             `bson:"name,omitempty" json:"name,omitempty"`
	Phone     	 string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}

// type Category struct {
// 	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
// 	Name string             `bson:"name,omitempty" json:"name,omitempty"`
// }

// type Order struct {
// 	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
// 	CustomerID      primitive.ObjectID `bson:"customer_id,omitempty" json:"customer_id,omitempty"`
// 	OrderItems      []OrderItem        `bson:"order_item,omitempty" json:"order_item,omitempty"`
// 	OrderDate       string             `bson:"order_date,omitempty" json:"order_date,omitempty"`
// 	TotalAmount     float64            `bson:"total_amount,omitempty" json:"total_amount,omitempty"`
// 	Status          string             `bson:"status,omitempty" json:"status,omitempty"` 
// 	DeliveryDate    string             `bson:"delivery_date,omitempty" json:"delivery_date,omitempty"`
// 	DeliveryAddress string            `bson:"delivery_address,omitempty" json:"delivery_address,omitempty"`
// }

// type OrderItem struct {
// 	MenuItemID primitive.ObjectID `bson:"menu_item_id,omitempty" json:"menu_item_id,omitempty"`
// 	Quantity   int                `bson:"quantity,omitempty" json:"quantity,omitempty"`
// 	Price      float64            `bson:"price,omitempty" json:"price,omitempty"`
// }
