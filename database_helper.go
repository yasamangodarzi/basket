package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// define  const
const (
	host     = "localhost"
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_database_name"
)

// Basket Model for the shopping basket
type Basket struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Data      string    `json:"data" gorm:"size:2048"`
	State     string    `json:"state"`
	UserID    uint      `json:"user_id"`
}

// User Model for users
type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Age       uint   `json:"age"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

func connectionDatabase() (db *gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tehran",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Basket{}, &User{})
	return db
}

func SelectById(tableName string, id uint) (interface{}, error) {
	var record interface{}
	db := connectionDatabase()
	switch tableName {
	case "users":
		var user User
		result := db.First(&user, id)
		record = &user
		if result.Error != nil {
			return nil, result.Error
		}
	case "baskets":
		var basket Basket
		result := db.First(&basket, id)
		record = &basket
		if result.Error != nil {
			return nil, result.Error
		}
	default:
		return nil, fmt.Errorf("unsupported table name")
	}

	return record, nil
}

func DeleteRecordByID(tableName string, id uint) (uint, error) {
	db := connectionDatabase()
	strId := 00000000
	switch tableName {
	case "baskets":
		result := db.Where("id = ?", id).Delete(&Basket{})
		strId = result.id

	case "users":
		result := db.Where("id = ?", id).Delete(&users{})
		strId = result.id

	default:
		return strId, fmt.Errorf("unsupported table name")
	}

	if result.Error != nil {
		return strId, result.Error
	}
	return strId, nil

}

func GetBasketsByUserID(userID uint) ([]Basket, error) {
	var baskets []Basket
	db := connectionDatabase()
	result := db.Where("user_id = ?", userID).Find(&baskets)
	if result.Error != nil {
		return nil, result.Error
	}
	return baskets, nil
}

func AddBasket(basket *Basket) (uint, error) {
	db := connectionDatabase()
	result := db.Create(&basket)
	return result.id, result.Error
}

func AddUser(user *User) (unit, error) {
	db := connectionDatabase()
	result := db.Create(&user)
	return result.id, result.Error
}

//func main() {
//	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tehran",
//		host, port, user, password, dbname)
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// Auto migrate tables
//	db.AutoMigrate(&Basket{}, &User{})
//
//	// Create a user
//	user := User{Name: "John Doe", Age: 30, Email: "john@example.com", Address: "123 Main St", Phone: "123-456-7890"}
//	db.Create(&user)
//
//	// Create a basket
//	basket := Basket{CreatedAt: time.Now(), UpdatedAt: time.Now(), Data: "Some data", State: "PENDING", UserID: user.ID}
//	InsertBasket(db, &basket)
//
//	// Select basket by ID
//	selectedBasket, err := SelectBasketByID(db, basket.ID)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Selected Basket:", selectedBasket)
//
//	// Update basket state
//	selectedBasket.State = "COMPLETED"
//	err = UpdateBasket(db, selectedBasket)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Updated Basket:", selectedBasket)
//}
