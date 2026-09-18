package seeder

import (
	"api/model"
	"log"
	"time"

	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	var userCount int64
	db.Model(&model.User{}).Count(&userCount)

	// if userCount == 0 {
	// 	users := []model.User{
	// 		{
	// 			HouseID:  100,
	// 			UserName: "100_house",
	// 			Name:     "John Doe",
	// 			Password: "555555",
	// 			Role:     "house",
	// 		},
	// 	}

	// 	visits := []model.Visit{
	// 		{
	// 			HouseID:      100,
	// 			Reason:       "Delivery",
	// 			Status:       "incoming",
	// 			ArrivalTime:  time.Now().Add(-time.Hour),
	// 			LicensePlate: "ABC-123",
	// 		},
	// 	}

	// 	for _, user := range users {
	// 		db.Create(&user)
	// 	}
	// 	for _, visit := range visits {
	// 		db.Create(&visit)
	// 	}
	// 	log.Println("Seed data inserted successfully")
	// } else {
	// 	log.Println("Seed data already exists, skipping")
	// }

	var adminUser model.User
	err := db.Where("username = ?", "admin").First(&adminUser).Error
	if err == gorm.ErrRecordNotFound {
		adminUser = model.User{
			HouseID:  999,
			UserName: "999_admin",
			Name:     "admin",
			Password: "123456",
			Role:     "admin",
		}
		db.Create(&adminUser)
		log.Println("Admin user inserted successfully")
	} else {
		log.Println("Admin user already exists, skipping")
	}

	visits := []model.Visit{
		{HouseID: 999, Reason: "First test", Status: "incoming", ArrivalTime: time.Now(), LicensePlate: "ABC123"},
		// {HouseID: 123, Reason: "Maintenance", Status: "completed", ArrivalTime: time.Now().Add(-24 * time.Hour), LicensePlate: "XYZ789"},
		// {HouseID: 157, Reason: "Family Visit", Status: "incoming", ArrivalTime: time.Now(), LicensePlate: "LMN456"},
	}

	for _, visit := range visits {
		var existingVisit model.Visit
		// Check if a visit with the same LicensePlate and HouseID exists
		err := db.Where("license_plate = ? AND house_id = ?", visit.LicensePlate, visit.HouseID).First(&existingVisit).Error
		if err == gorm.ErrRecordNotFound {
			// Insert the visit if it doesn't exist
			db.Create(&visit)
			log.Printf("Visit for HouseID %d and LicensePlate %s inserted successfully", visit.HouseID, visit.LicensePlate)
		} else {
			log.Printf("Visit for HouseID %d and LicensePlate %s already exists, skipping", visit.HouseID, visit.LicensePlate)
		}
	}
	log.Println("Visit seeding completed successfully.")
}
