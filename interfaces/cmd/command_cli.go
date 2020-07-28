package cmd

import (
	"fmt"
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/infrastructure/security"
	"go-rest-skeleton/infrastructure/util"
	"log"

	"github.com/urfave/cli/v2"
)

func NewCommand(dbServices *persistence.Repositories) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "create:secret",
			Usage: "make a base64 encoded string private key and public key for APP_PRIVATE_KEY and APP_PUBLIC_KEY",
			Action: func(c *cli.Context) error {
				secretPriPubKey, err := security.GenerateSecret()
				if err != nil {
					log.Println(err)
				}
				fmt.Println(util.PrettyJSON(secretPriPubKey))
				return nil
			},
		},
		{
			Name:  "db:seed",
			Usage: "run predefined database seeder",
			Action: func(c *cli.Context) error {
				err := dbServices.Seeds()
				if err != nil {
					log.Println(err)
				}
				return nil
			},
		},
	}
}
