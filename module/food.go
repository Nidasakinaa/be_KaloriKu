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
func InsertMenuItem(db *mongo.Database, col string, name string, description string, category string, image string) (insertedID primitive.ObjectID, err error) {
	menu := bson.M{
		"name":    		name,
		"description":  description,
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
func UpdateMenuItem(ctx context.Context, db *mongo.Database, col string, _id primitive.ObjectID, name string, description string, price float64, category string, image string, stock float64) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"name":    		name,
			"description":  description,
			"category":   	category,
			"image"	:		image,
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
// GetCustomerByID retrieves a customer from the database by its ID
func GetCustomerByID(_id primitive.ObjectID, db *mongo.Database, col string) (model.Customer, error) {
	var customer model.Customer
	collection := db.Collection("Customer")
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&customer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return customer, fmt.Errorf("GetCostumerByID: costumer dengan ID %s tidak ditemukan", _id.Hex())
		}
		return customer, fmt.Errorf("GetCostumerByID: gagal mendapatkan costumer: %w", err)
	}
	return customer, nil
}

func GetAllCustomer(db *mongo.Database, col string) (data []model.Customer) {
	customer := db.Collection(col)
	filter := bson.M{}
	cursor, err := customer.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllCustomer :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertCustomer(db *mongo.Database, col string, name string, phone string) (insertedID primitive.ObjectID, err error) {
	customer := bson.M{
		"name":    name,
		"phone":   phone,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), customer)
	if err != nil {
		fmt.Printf("InsertCustomer: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateCustomer(ctx context.Context, db *mongo.Database, col string, _id primitive.ObjectID, name string, phone string) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"name":    name,
			"phone":   phone,
		},
	}
	result, err := db.Collection(col).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateCustomer: gagal memperbarui customer: %w", err)
	}
	if result.MatchedCount == 0 {
		return errors.New("UpdateCustomer: tidak ada data yang diubah dengan ID yang ditentukan")
	}
	return nil
}

func DeleteCustomerByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	customer := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := customer.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

//FUNCTION ADMIN
// GetCustomerByID retrieves a customer from the database by its ID
func GetAdminByID(_id primitive.ObjectID, db *mongo.Database, col string) (model.Admin, error) {
	var admin model.Admin
	collection := db.Collection("Admin")
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&admin)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return admin, fmt.Errorf("GetAdminByID: admin dengan ID %s tidak ditemukan", _id.Hex())
		}
		return admin, fmt.Errorf("GetAdminByID: gagal mendapatkan admin: %w", err)
	}
	return admin, nil
}

func GetAllAdmin(db *mongo.Database, col string) (data []model.Admin) {
	admin := db.Collection(col)
	filter := bson.M{}
	cursor, err := admin.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllAdmin :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertAdmin(db *mongo.Database, col string, username string, password string) (insertedID primitive.ObjectID, err error) {
	admin := bson.M{
		"username":   username,
		"password":   password,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), admin)
	if err != nil {
		fmt.Printf("InsertAdmin: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateAdmin(ctx context.Context, db *mongo.Database, col string, _id primitive.ObjectID, username string, password string) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"username":   username,
			"password":   password,
		},
	}
	result, err := db.Collection(col).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateAdmin: gagal memperbarui data admin: %w", err)
	}
	if result.MatchedCount == 0 {
		return errors.New("UpdateAdmin: tidak ada data yang diubah dengan ID yang ditentukan")
	}
	return nil
}

func DeleteAdminByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	admin := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := admin.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
