package main

import (
	"assignment4/database"
	"assignment4/models"
	"assignment4/routes"
 

	"github.com/go-pg/pg/v10/orm"

)

func init() {
    orm.RegisterTable((*models.CustomerToAccount)(nil))
}

func main() {
    LoadDatabase()
    routes.Router()
    //DeleteDatabase()
}

func LoadDatabase() error {
	
	database.Connect()

	models := []interface{}{
        (*models.Bank)(nil),
        (*models.Branch)(nil),
        (*models.Customer)(nil),
        (*models.Account)(nil),
        (*models.CustomerToAccount)(nil),
		(*models.Transaction)(nil),
    }

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
        FKConstraints: true,
	}

    for _, model := range models {
        err := database.Db.Model(model).CreateTable(opts)
        if err != nil {
            println(err.Error())
        }
    }
    return nil

}


func DeleteDatabase() error {
    database.Connect()

    models := []interface{}{
        (*models.Transaction)(nil),
        (*models.CustomerToAccount)(nil),
        (*models.Account)(nil),
        (*models.Customer)(nil),
        (*models.Branch)(nil),
        (*models.Bank)(nil),
    }

    for  _, model := range models {
        err := database.Db.Model(model).DropTable(&orm.DropTableOptions{
            IfExists: true,
            Cascade:  true, // Drop dependent objects
        })
        if err != nil {
            println(err.Error())
        }
    }
    return nil
}
