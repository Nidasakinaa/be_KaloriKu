package module

import (
	"context"
	"errors"
	"fmt"
	model "github.com/Nidasakinaa/be_KaloriKu/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func StaticAdminLogin(db *mongo.Database, col string, username, password string) (bool, error) {
	// Validasi input kosong
	if username == "" || password == "" {
		return false, errors.New("username and password cannot be empty")
	}
	
	// Placeholder untuk pengembangan ke database (jika diperlukan)
	// Implementasi statis sementara hanya memeriksa kecocokan dengan parameter
	return false, errors.New("invalid admin credentials")
}

//FUNCTION MENU ITEM
//GetMenuItemByID retrieves a menu item from the database by its ID
func GetMenuItemByID(_id primitive.ObjectID, db *mongo.Database, col string) (model.MenuItem, error) {
	var menu model.MenuItem
	collection := db.Collection("Menu")
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&menu)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return menu, fmt.Errorf("GetMenuItemByID: menu item dengan ID %s tidak ditemukan", _id.Hex())
		}
		return menu, fmt.Errorf("GetMenuItemByID: gagal mendapatkan menu item: %w", err)
	}
	return menu, nil
}

//GetAllMenuItem retrieves all menu items from the database
func GetAllMenuItem(db *mongo.Database, col string) (data []model.MenuItem) {
	menu := db.Collection(col)
	filter := bson.M{}
	cursor, err := menu.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllMenuItem :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// InsertMenuItem creates a new order in the database
func InsertMenuItem(db *mongo.Database, col string, name string, ingredients string, description string, calories float64, category string, image string) (insertedID primitive.ObjectID, err error) {
	menu := bson.M{
		"name":    		name,
		"ingredients":  ingredients,
		"description":  description,
		"calories":     calories,
		"category":   	category,
		"image":		image,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), menu)
	if err != nil {
		fmt.Printf("InsertMenuItem: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

//UpdateMenuItem updates an existing menu item in the database
func UpdateMenuItem(ctx context.Context, db *mongo.Database, col string, _id primitive.ObjectID, name string, ingredients string, description string, calories float64, category string, image string) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"name":    		name,
			"ingredients":  ingredients,
			"description":  description,
			"calories":     calories,
			"category":   	category,
			"image":		image,
		},
	}
	result, err := db.Collection(col).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateMenuItem: gagal memperbarui Menu Item: %w", err)
	}
	if result.MatchedCount == 0 {
		return errors.New("UpdateMenuItem: tidak ada data yang diubah dengan ID yang ditentukan")
	}
	return nil
}

// DeleteMenuItemByID deletes a menu item from the database by its ID
func DeleteMenuItemByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	menu := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := menu.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

//FUNCTION CUSTOMER
