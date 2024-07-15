package migrations

import (
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"

	"github.com/sethvargo/go-password/password"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		settings, _ := dao.FindSettings()
		settings.Meta.AppName = "Project Serendib"
		settings.Backups.Cron = "@hourly"
		settings.Backups.CronMaxKeep = 24

		admin := &models.Admin{}
		admin.Email = "admin@projectserendib.com"
		password, err := password.Generate(16, 4, 4, false, false)
		if err != nil {
			return err
		}
		admin.SetPassword(password)

		log.Printf("Admin email: %s\n", admin.Email)
		log.Printf("Admin password: %s\n", password)
		
		admerr := dao.SaveAdmin(admin)
		seterr := dao.SaveSettings(settings)
		if seterr != nil {
			return seterr
		}
		if admerr != nil {
			return admerr
		}
		return nil
	}, func(db dbx.Builder) error {
		// add down queries...

		return nil
	})
}
