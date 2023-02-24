package main

import (
	"gorm.io/gorm"
	"rent-n-go-backend/models"
	"rent-n-go-backend/utils"
)

// migrateModel
// A simple wrapper around GORM Auto Migrate that will automatically complain if error
// based on app state.
func migrateModel(db *gorm.DB, model any) {
	err := db.AutoMigrate(model)

	if err != nil {
		utils.ShouldPanic(err)
	}
}

func seedByModule(args string, module string, callback func()) {
	if args == "" || args == module {
		callback()
	}
}

func processMigration(args []string) {
	if len(args) > 0 {
		if args[0] == "refresh" {
			tables, err := utils.GetDb().Migrator().GetTables()

			if err != nil {
				panic(err)
			}

			interfaces := make([]interface{}, len(tables))

			for i, v := range tables {
				interfaces[i] = v
			}

			utils.GetDb().Migrator().DropTable(interfaces...)
		}

		if args[0] == "migrate" || args[0] == "refresh" {
			// migrate all tables to database.
			migrate(utils.GetDb())
		}

		if args[0] == "seed" || args[0] == "refresh" {
			// check if user want to seed specific module
			arg := ""

			if len(args) > 1 {
				arg = args[1]
			}

			// seed the table based on arguments.
			seeder(utils.GetDb(), arg)
		}
	}
}

/*
*
Will be executed by GORM
in Before Hook
*/
func migrate(db *gorm.DB) {
	migrateModel(db, &models.User{})
	migrateModel(db, &models.Nik{})
	migrateModel(db, models.Sim{})
	migrateModel(db, models.RefreshToken{})
}

// Seed a data to a database
// produce fake data that will only be seeded under development state
// Will be executed in Before Hook.
func seeder(db *gorm.DB, args string) {
	seedByModule(args, "user", func() {
		password, _ := utils.HashPassword("admin12345")

		user := models.User{
			Name:     "Sang Admin",
			Email:    "admin@mail.com",
			Role:     "admin",
			Password: password,
		}

		db.Create(&user)
	})
}
