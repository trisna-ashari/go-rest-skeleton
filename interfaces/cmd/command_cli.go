package cmd

import (
	"fmt"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/pkg/encoder"
	"go-rest-skeleton/pkg/security"
	"log"

	"github.com/urfave/cli/v2"
)

// NewCommand construct a CLI commands.
func NewCommand(dbService *persistence.Repositories) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "create:secret",
			Usage: "make a base64 encoded string private key and public key for APP_PRIVATE_KEY and APP_PUBLIC_KEY",
			Action: func(c *cli.Context) error {
				secretPriPubKey, err := security.GenerateSecret()
				if err != nil {
					log.Println(err)
				}
				fmt.Println(encoder.PrettyJSONWithIndent(secretPriPubKey))
				return nil
			},
		},
		{
			Name:  "db:migrate",
			Usage: "run database migration",
			Action: func(c *cli.Context) error {
				err := dbService.AutoMigrate()
				if err != nil {
					log.Println(err)
				}
				return nil
			},
		},
		{
			Name:  "db:init",
			Usage: "run predefined database initial seeder",
			Action: func(c *cli.Context) error {
				err := dbService.InitialSeeds()
				if err != nil {
					log.Println(err)
				}
				return nil
			},
		},
		{
			Name:  "db:seed",
			Usage: "run predefined database seeder",
			Action: func(c *cli.Context) error {
				err := dbService.Seeds()
				if err != nil {
					log.Println(err)
				}
				return nil
			},
		},
	}
}
