package main

import (
	"fmt"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/handler"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

func main() {
	p := handler.Params{}
	app := &cli.App{
		Name:  "gophkeeper",
		Usage: "password store",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "server",
				Aliases:     []string{"s"},
				Value:       "localhost:8000",
				Usage:       "Server address",
				EnvVars:     []string{"SERVER_ADDRESS"},
				Destination: &p.ServerAddress,
			},
			&cli.StringFlag{
				Name:        "login",
				Aliases:     []string{"l"},
				Usage:       "Login",
				Destination: &p.Login,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p"},
				Usage:       "Password",
				Destination: &p.Password,
			},
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "File config",
				EnvVars:     []string{"CONFIG"},
				Destination: &handler.ConfigPath,
			},
		},
		Action: func(*cli.Context) error {
			h, err := handler.NewHandler(p)
			if err != nil {
				return err
			}
			err = h.Register()
			return err
		},
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add info",
				Subcommands: []*cli.Command{
					{
						Name:  "credentials",
						Usage: "add a new pair login/password",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "meta",
								Aliases: []string{"m"},
								Value:   "",
							},
							&cli.StringFlag{
								Name:     "login",
								Aliases:  []string{"l"},
								Required: true,
							},
							&cli.StringFlag{
								Name:     "password",
								Aliases:  []string{"p"},
								Required: true,
							},
						},
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.AddCredentials(
								cCtx.String("login"),
								cCtx.String("password"),
								cCtx.String("meta"),
							)
							return err
						},
					},
					{
						Name:  "text",
						Usage: "add new text data",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "meta",
								Aliases: []string{"m"},
								Value:   "",
							},
							&cli.StringFlag{
								Name:     "path",
								Aliases:  []string{"p"},
								Required: true,
							},
						},
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.AddTextFromFile(cCtx.String("path"), cCtx.String("meta"))
							return err
						},
					},
					{
						Name:  "binary",
						Usage: "add new binary data",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "meta",
								Aliases: []string{"m"},
								Value:   "",
							},
							&cli.StringFlag{
								Name:     "path",
								Aliases:  []string{"p"},
								Required: true,
							},
						},
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.AddBinaryDataFromFile(cCtx.String("path"), cCtx.String("meta"))
							return err
						},
					},
					{
						Name:  "card",
						Usage: "add new bank card",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "meta",
								Aliases: []string{"m"},
								Value:   "",
							},
							&cli.StringFlag{
								Name:     "number",
								Aliases:  []string{"n"},
								Required: true,
							},
							&cli.StringFlag{
								Name:     "exp",
								Aliases:  []string{"e"},
								Required: true,
							},
							&cli.StringFlag{
								Name:     "cardholder",
								Aliases:  []string{"h"},
								Required: true,
							},
							&cli.StringFlag{
								Name:     "code",
								Aliases:  []string{"c"},
								Required: true,
							},
						},
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.AddBankCard(
								cCtx.String("number"),
								cCtx.String("exp"),
								cCtx.String("cardholder"),
								cCtx.String("code"),
								cCtx.String("meta"),
							)
							return err
						},
					},
				},
			},
			{
				Name:  "remove",
				Usage: "remove info",
				Subcommands: []*cli.Command{
					{
						Name:  "credentials",
						Usage: "remove credentials by login",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "login",
								Aliases:  []string{"l"},
								Required: true,
							},
						},
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.RemoveCredentialsByLogin(
								cCtx.String("login"),
							)
							return err
						},
					},
					{
						Name:  "text",
						Usage: "remove text data",
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.RemoveText()
							return err
						},
					},
					{
						Name:  "binary",
						Usage: "remove binary data",
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.RemoveBinaryData()
							return err
						},
					},
					{
						Name:  "card",
						Usage: "remove the bank card by number",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "number",
								Aliases:  []string{"n"},
								Required: true,
							},
						},
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.RemoveBankCardByNumber(
								cCtx.String("number"),
							)
							return err
						},
					},
				},
			},
			{
				Name:  "list",
				Usage: "show list",
				Subcommands: []*cli.Command{
					{
						Name:  "credentials",
						Usage: "show list login/password",
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.ShowCredentialsList()
							return err
						},
					},
					{
						Name:  "text",
						Usage: "show list text data",
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.ShowTextFilesList()
							return err
						},
					},
					{
						Name:  "binary",
						Usage: "show binary data list",
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.ShowBinaryDataList()
							return err
						},
					},
					{
						Name:  "card",
						Usage: "add new bank card",
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.ShowBankCardList()
							return err
						},
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	fmt.Println("end")
}
