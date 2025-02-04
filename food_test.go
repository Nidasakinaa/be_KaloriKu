package bekaloriku_test

import (
	"fmt"
	"testing"

	module "github.com/Nidasakinaa/be_KaloriKu/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FUNCTION MENU ITEM
func TestGetMenuItemByID(t *testing.T) {
	_id := "6799d453b5803150f0a440e9"
	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	menu, err := module.GetMenuItemByID(objectID, module.MongoConn, "Menu")
	if err != nil {
		t.Fatalf("error calling GetMenuItemByID: %v", err)
	}
	fmt.Println(menu)
}

func TestGetMenuItemByCategory(t *testing.T) {
	category := "Diet"
	menu, err := module.GetMenuItemByCategory(category, module.MongoConn, "Menu")
	if err != nil {
		t.Fatalf("error calling GetMenuItemByCategory: %v", err)
	}
	fmt.Println(menu)
}

func TestGetAllMenu(t *testing.T) {
	data := module.GetAllMenuItem(module.MongoConn, "Menu")
	fmt.Println(data)
}

func TestInsertMenuItem(t *testing.T) {
    // Test data
	name := "Fruit Smoothie"
    ingredients := "Banana, Strawberry, Blueberry, Almond Milk, Honey"
    description := "A refreshing smoothie made with a blend of fruits and almond milk"
    calories := 200.0
    category := "Beverage"
    imageURL := "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQmTMSDdEMe7hCs2T1KBVj_1NKwYGb1lYqbKQ&s" // Valid URL
 
	 // Call the function
	 insertedID, err := module.InsertMenuItem(module.MongoConn, "Menu", name, ingredients, description, calories, category, imageURL)
	 if err != nil {
		 t.Fatalf("Error inserting menu item: %v", err)
	 }
 
	 // Print the result
	 fmt.Printf("Data berhasil disimpan dengan id %s\n", insertedID.Hex())
}

func TestDeleteMenuItemByID(t *testing.T) {
	id := "678a71310bb7a4334619cf8b"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteMenuItemByID(objectID, module.MongoConn, "Menu")
	if err != nil {
		t.Fatalf("error calling DeleteMenuItemByID: %v", err)
	}

	_, err = module.GetMenuItemByID(objectID, module.MongoConn, "Menu")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

//FUNCTION USER
//GetUserByID retrieves a user from the database by its ID
func TestGetUserByID(t *testing.T) {
	_id := "67a23885d8d58983179fe315"
	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	menu, err := module.GetUserByID(objectID, module.MongoConn, "User")
	if err != nil {
		t.Fatalf("error calling GetMenuItemByID: %v", err)
	}
	fmt.Println(menu)
}

func TestGetAllUsers(t *testing.T) {
	data, err := module.GetAllUser(module.MongoConn, "User")
	if err != nil {
		t.Fatalf("error calling GetAllUsers: %v", err)
	}
	fmt.Println(data)
}	

func TestInsertUser(t *testing.T) {
	name := "Nida Sakina"
    phone := "083174603834"
    username := "nida"
    password := "Nida150304"
    role := "admin"
    insertedID, err := module.InsertUser(module.MongoConn, "User", name, phone, username, password, role)
    if err != nil {
        t.Errorf("Error inserting data: %v", err)
    }
    fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestDeleteUserByID(t *testing.T) {
    id := "67a23885d8d58983179fe315"
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        t.Fatalf("error converting id to ObjectID: %v", err)
    }

    err = module.DeleteUserByID(objectID, module.MongoConn, "User")
    if err != nil {
        t.Fatalf("error calling DeleteUserByID: %v", err)
    }

    _, err = module.GetUserByID(objectID, module.MongoConn, "User")
    if err == nil {
        t.Fatalf("expected data to be deleted, but it still exists")
    }
}