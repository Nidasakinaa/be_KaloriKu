package bekaloriku_test

import (
	"fmt"
	"testing"

	module "github.com/Nidasakinaa/be_KaloriKu/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FUNCTION MENU ITEM
func TestGetMenuItemByID(t *testing.T) {
	_id := "678a72468c7c04668e9b40e5"
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
	category := "Salad"
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
	 name := "Grilled Chicken with Quinoa"
	 ingredients := "Chicken Breast 200 gram, Quinoa 100 gram, Olive Oil, Spices"
	 description := "Grilled chicken breast served with quinoa and seasoned with olive oil and spices"
	 calories := 350.0
	 category := "High Protein"
	 imageURL := "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRq-eqd1MdChsAz5ThM2ltfaEoeLv4psphvxg&s" // Valid URL
 
	 // Call the function
	 insertedID, err := module.InsertMenuItem(module.MongoConn, "Menu", name, ingredients, description, calories, category, imageURL)
	 if err != nil {
		 t.Fatalf("Error inserting menu item: %v", err)
	 }
 
	 // Print the result
	 fmt.Printf("Data berhasil disimpan dengan id %s\n", insertedID.Hex())
 
	 // Verify the inserted data and display the image
	 menuItem, err := module.GetMenuItemByIDAndDisplayImage(insertedID, module.MongoConn, "Menu")
	 if err != nil {
		 t.Fatalf("Error retrieving menu item: %v", err)
	 }
	 if menuItem.Image == "" {
		 t.Errorf("Expected image path to be set, but got empty string")
	 }
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
	_id := "678ba051c7522337e180b946"
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
	data := module.GetAllUser(module.MongoConn, "User")
	fmt.Println(data)
}

func TestInsertUser(t *testing.T) {
	name := "Admin"
    phone := "1234567890"
    username := "admin"
    password := "admin12345"
    role := "Admin"
    insertedID, err := module.InsertUser(module.MongoConn, "User", name, phone, username, password, role)
    if err != nil {
        t.Errorf("Error inserting data: %v", err)
    }
    fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestDeleteUserByID(t *testing.T) {
    id := "678b9ffac895eeea0d5144b1a"
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