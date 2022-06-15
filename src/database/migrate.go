package database

import (
	"log"

	"github.com/depri11/vehicle/src/database/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "migrate up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "migrate down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := SetupDB()
	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "001",
			Migrate: func(tx *gorm.DB) error {

				return tx.AutoMigrate(&models.User{}, &models.Historys{}, &models.Vehicle{}, &models.VehicleImage{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&models.User{}, &models.Historys{}, &models.Vehicle{}, "vehicle_images")
			},
		},
	})

	if migUp {
		if err = m.Migrate(); err != nil {
			return err
		}
		log.Fatal("Migration did run successfully")
		return nil
	}

	if migDown {
		if err = m.RollbackLast(); err != nil {
			return err
		}
		log.Fatal("Rollback did run successfully")
		return nil
	}

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.User{},
			&models.Vehicle{},
			&models.VehicleImage{},
			&models.Historys{},
		)
		if err != nil {
			return err
		}

		return nil
	})

	if err = m.Migrate(); err != nil {
		return err
	}

	log.Fatal("init schema successfully")
	return nil

}
