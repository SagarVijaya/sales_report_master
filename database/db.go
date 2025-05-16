package database

import (
	"database/sql"
	"fmt"
	"log"
	"salesproject/apps/models"
	"salesproject/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect Database initializes the DB and runs migrations
func ConnectDatabase() error {
	cf := config.GetConfig()
	fmt.Println(cf)

	lCreateDatabase := fmt.Sprintf("%s:%s@tcp(%s:%d)/?parseTime=true",
		cf.Database.User,
		cf.Database.Pass,
		cf.Database.Host,
		cf.Database.Port,
	)
	lCreateDbSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cf.Database.Name)
	lErr := CreateDatabase(lCreateDatabase, lCreateDbSql)
	if lErr != nil {
		return lErr
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cf.Database.User,
		cf.Database.Pass,
		cf.Database.Host,
		cf.Database.Port,
		cf.Database.Name,
	)

	// Open GORM DB
	db, lErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if lErr != nil {
		log.Fatalf("Failed to connect to GORM database: %v", lErr)
		return lErr
	}

	// Remove the existing table
	lErr = db.Migrator().DropTable(&models.OrderDetails{}, &models.Order{}, &models.Customer{}, &models.Product{})
	if lErr != nil {
		log.Fatalf("Failed to remove tables for auto migrate tables: %v", lErr)
		return lErr
	}
	// Run migrations
	lErr = db.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}, &models.OrderDetails{})
	if lErr != nil {
		log.Fatalf("Failed to auto migrate tables: %v", lErr)
		return lErr
	}

	DB = db

	lErr = CreateForeignKey()
	if lErr != nil {
		log.Fatalf("Failed to create foreign key: %v", lErr)
		return lErr
	}
	log.Println("GORM database setup completed successfully")
	return nil
}

// create database if not exists
func CreateDatabase(dsn string, createSQL string) error {
	sqlDB, lErr := sql.Open("mysql", dsn)
	if lErr != nil {
		log.Fatalf("Failed to connect to MariaDB server for DB creation: %v", lErr)
		return lErr
	}
	defer sqlDB.Close()

	if lErr = sqlDB.Ping(); lErr != nil {
		log.Fatalf("Ping failed: %v", lErr)
		return lErr
	}

	if _, lErr = sqlDB.Exec(createSQL); lErr != nil {
		log.Fatalf("Failed to create DB: %v", lErr)
		return lErr
	}

	return nil
}

func CreateForeignKey() error {
	err := DB.Exec(`
    ALTER TABLE orders
    ADD CONSTRAINT fk_orders_customer
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;
`).Error
	if err != nil {
		return err
	}

	err1 := DB.Exec(`
    ALTER TABLE order_details
    ADD CONSTRAINT fk_order_details_order
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;
`).Error
	if err1 != nil {
		return err1
	}

	err2 := DB.Exec(`
    ALTER TABLE order_details
    ADD CONSTRAINT fk_order_details_product
    FOREIGN KEY (product_id) REFERENCES products(product_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;
`).Error
	if err2 != nil {
		return err2
	}
	return nil
}
