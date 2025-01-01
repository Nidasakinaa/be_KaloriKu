// package bekaloriku_test

// import (
// 	"fmt"
// 	"testing"

// 	module "github.com/Nidasakinaa/be_KaloriKu/module"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// //FUNCTION CATEGORY
// // func TestGetCategoryByID
// func TestGetCategoryByID(t *testing.T) {
// 	_id := "67619243c9f2a5f803110bde"
// 	objectID, err := primitive.ObjectIDFromHex(_id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}
// 	category, err := module.GetOrderByID(objectID, module.MongoConn, "Category")
// 	if err != nil {
// 		t.Fatalf("error calling GetOrderByID: %v", err)
// 	}
// 	fmt.Println(category)
// }

// // func TestGetAllCategory
// func TestGetAllCategory(t *testing.T) {
// 	data := module.GetAllCategory(module.MongoConn, "Category")
// 	fmt.Println(data)
// }

// // func TestInsertCategory
// func TestInsertCategory(t *testing.T) {
// 	name := "Salad"
// 	insertedID, err := module.InsertCategory(module.MongoConn, "Category", name)
// 	if err != nil {
// 		t.Errorf("Error inserting data: %v", err)
// 	}
// 	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
// }

// func TestDeleteCategoryID(t *testing.T) {
// 	id := "67619243c9f2a5f803110bde"
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}

// 	err = module.DeleteCategoryByID(objectID, module.MongoConn, "Category")
// 	if err != nil {
// 		t.Fatalf("error calling DeleteCategoryByID: %v", err)
// 	}

// 	_, err = module.GetCategoryByID(objectID, module.MongoConn, "Category")
// 	if err == nil {
// 		t.Fatalf("expected data to be deleted, but it still exists")
// 	}
// }


// //FUNCTION MENU ITEM
// func TestGetMenuItemByID(t *testing.T) {
// 	_id := "67619243c9f2a5f803110bde"
// 	objectID, err := primitive.ObjectIDFromHex(_id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}
// 	menu, err := module.GetMenuItemByID(objectID, module.MongoConn, "MenuItem")
// 	if err != nil {
// 		t.Fatalf("error calling GetMenuItemByID: %v", err)
// 	}
// 	fmt.Println(menu)
// }

// func TestGetAllMenu(t *testing.T) {
// 	data := module.GetAllMenuItem(module.MongoConn, "MenuItem")
// 	fmt.Println(data)
// }

// func TestInsertMenuItem(t *testing.T) {
// 	name := "Vegetable Salad with Salted Egg"
// 	description := "Fresh vegetables with salted egg"
// 	price := 45.000
// 	category := "Salad"
// 	insertedID, err := module.InsertMenuItem(module.MongoConn, "MenuItem", name, description, price, category)
// 	if err != nil {
// 		t.Errorf("Error inserting data: %v", err)
// 	}
// 	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
// }

// func TestDeleteMenuItemByID(t *testing.T) {
// 	id := "67619243c9f2a5f803110bde"
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}

// 	err = module.DeleteMenuItemByID(objectID, module.MongoConn, "MenuItem")
// 	if err != nil {
// 		t.Fatalf("error calling DeleteMenuItemByID: %v", err)
// 	}

// 	_, err = module.GetMenuItemByID(objectID, module.MongoConn, "DataPasien")
// 	if err == nil {
// 		t.Fatalf("expected data to be deleted, but it still exists")
// 	}
// }

// //FUNCTION ORDER
// // func TestGetOrderByID
// func TestGetOrderByID(t *testing.T) {
// 	_id := "67619243c9f2a5f803110bde"
// 	objectID, err := primitive.ObjectIDFromHex(_id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}
// 	order, err := module.GetOrderByID(objectID, module.MongoConn, "Order")
// 	if err != nil {
// 		t.Fatalf("error calling GetOrderByID: %v", err)
// 	}
// 	fmt.Println(order)
// }

// // func TestGetAllOrder
// func TestGetAllOrder(t *testing.T) {
// 	data := module.GetAllOrder(module.MongoConn, "Order")
// 	fmt.Println(data)
// }

// // func TestInsertOrder
// func TestInsertOrder(t *testing.T) {
// 	orderItems := []model.OrderItem{
// 		{
// 			MenuItemID: primitive.NewObjectID(),
// 			Quantity:   1,
// 			Price:      45.000,
// 		},
// 	}
// 	orderDate := "2021-07-01"
// 	totalAmount := 45.000
// 	status := "Pending"
// 	deliveryDate := "2021-07-02"
// 	deliveryAddress := "Jalan Sariasih No. 12"

// 	// Ensure MongoConn is not nil
// 	if module.MongoConn == nil {
// 		t.Fatal("MongoConn is not initialized")
// 	}

// 	insertedID, err := module.InsertOrder(module.MongoConn, "Order", orderItems, orderDate, totalAmount, status, deliveryDate, deliveryAddress)
// 	if err != nil {
// 		t.Fatalf("Error inserting data: %v", err)
// 	}

// 	// Check if insertedID is valid
// 	if insertedID.IsZero() {
// 		t.Fatal("Inserted ID is zero, insertion may have failed")
// 	}

// 	fmt.Printf("Data berhasil disimpan dengan id %s\n", insertedID.Hex())
// }

// func TestDeleteCategoryByID(t *testing.T) {
// 	id := "668e2b1540bdb1d47710a316"
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}

// 	err = module.DeleteCustomerByID(objectID, module.MongoConn, "Category")
// 	if err != nil {
// 		t.Fatalf("error calling DeleteCategoryByID: %v", err)
// 	}

// 	_, err = module.GetCategoryByID(objectID, module.MongoConn, "Category")
// 	if err == nil {
// 		t.Fatalf("expected data to be deleted, but it still exists")
// 	}
// }

// //FUNCTION COSTUMER
// func TestGetCustomerByID(t *testing.T) {
// 	_id := "669534d2af52bee3d2606c34"
// 	objectID, err := primitive.ObjectIDFromHex(_id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}
// 	customer, err := module.GetCustomerByID(objectID, module.MongoConn, "Customer")
// 	if err != nil {
// 		t.Fatalf("error calling GetCustomerFromID: %v", err)
// 	}
// 	fmt.Println(customer)
// }

// func TestGetAllCustomer(t *testing.T) {
// 	data := module.GetAllCustomer(module.MongoConn, "Customer")
// 	fmt.Println(data)
// }

// func TestInsertCustomer(t *testing.T) {
// 	name := "Sari Endah"
// 	phone := "0831654321"
// 	address := "Sariasih Street No. 12"
// 	insertedID, err := module.InsertCustomer(module.MongoConn, "Customer", name, phone, address)
// 	if err != nil {
// 		t.Errorf("Error inserting data: %v", err)
// 	}
// 	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
// }

// func TestDeleteCustomerByID(t *testing.T) {
// 	id := "668e2b1540bdb1d47710a316"
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}

// 	err = module.DeleteCustomerByID(objectID, module.MongoConn, "Customer")
// 	if err != nil {
// 		t.Fatalf("error calling DeleteCustomerByID: %v", err)
// 	}

// 	_, err = module.GetCustomerByID(objectID, module.MongoConn, "Customer")
// 	if err == nil {
// 		t.Fatalf("expected data to be deleted, but it still exists")
// 	}
// }

// //FUNCTION ADMIN
// func TestGetAdminByID(t *testing.T) {
// 	_id := "669534d2af52bee3d2606c34"
// 	objectID, err := primitive.ObjectIDFromHex(_id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}
// 	admin, err := module.GetAdminByID(objectID, module.MongoConn, "Admin")
// 	if err != nil {
// 		t.Fatalf("error calling GetAdminFromID: %v", err)
// 	}
// 	fmt.Println(admin)
// }

// func TestGetAllAdmin(t *testing.T) {
// 	data := module.GetAllAdmin(module.MongoConn, "Admin")
// 	fmt.Println(data)
// }

// func TestInsertAdmin(t *testing.T) {
// 	username := "Sari Endah"
// 	password := "0831654321"
// 	insertedID, err := module.InsertCustomer(module.MongoConn, "Admin", username, password, "admin@example.com")
// 	if err != nil {
// 		t.Errorf("Error inserting data: %v", err)
// 	}
// 	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
// }

// func TestDeleteAdminByID(t *testing.T) {
// 	id := "668e2b1540bdb1d47710a316"
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Fatalf("error converting id to ObjectID: %v", err)
// 	}

// 	err = module.DeleteAdminByID(objectID, module.MongoConn, "Admin")
// 	if err != nil {
// 		t.Fatalf("error calling DeleteAdminByID: %v", err)
// 	}

// 	_, err = module.GetAdminByID(objectID, module.MongoConn, "Admin")
// 	if err == nil {
// 		t.Fatalf("expected data to be deleted, but it still exists")
// 	}
// }
