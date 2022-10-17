package main

import (
	"fmt"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/handler"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

func main() {
	// разобраться что с zip
	// написать тесты и readme
	// добавить шифрование tls
	// проверить все эндпоинты
	// отформатировать проект
	// написать все комментарии
	// на клиенте интерфейсы - написать тесты 80%
	// https://github.com/Netflix/go-expect
	// log.SetLevel(log.PanicLevel)
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
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "File config",
				EnvVars:     []string{"CONFIG"},
				Destination: &handler.ConfigPath,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "register",
				Usage: "register new user",
				Flags: []cli.Flag{
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

					err = h.Register(
						cCtx.Context,
						cCtx.String("login"),
						cCtx.String("password"),
					)
					return err
				},
			},
			{
				Name:  "login",
				Usage: "user authorization",
				Flags: []cli.Flag{
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
					err = h.Login(
						cCtx.Context,
						cCtx.String("login"),
						cCtx.String("password"),
					)
					return err
				},
			},
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
								cCtx.Context,
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
							err = h.AddTextFromFile(
								cCtx.Context,
								cCtx.String("path"),
								cCtx.String("meta"),
							)
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
							err = h.AddBinaryDataFromFile(
								cCtx.Context,
								cCtx.String("path"), cCtx.String("meta"),
							)
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
								cCtx.Context,
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
								cCtx.Context,
								cCtx.String("login"),
							)
							return err
						},
					},
					{
						Name:  "text",
						Usage: "remove text data",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Aliases:  []string{"n"},
								Required: true,
							},
						},
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.RemoveTextByName(cCtx.Context, cCtx.String("name"))
							return err
						},
					},
					{
						Name:  "binary",
						Usage: "remove binary data",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Aliases:  []string{"n"},
								Required: true,
							},
						},
						Action: func(cCtx *cli.Context) error {
							h, err := handler.NewHandler(p)
							if err != nil {
								return err
							}
							err = h.RemoveBinaryDataByName(cCtx.Context, cCtx.String("name"))
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
								cCtx.Context,
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
							h.ShowCredentialsList(cCtx.Context)
							return nil
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
							h.ShowTextFilesList(cCtx.Context)
							return nil
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
							h.ShowBinaryDataList(cCtx.Context)
							return nil
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
							h.ShowBankCardList(cCtx.Context)
							return nil
						},
					},
				},
			},
			{
				Name:  "get",
				Usage: "get item",
				Subcommands: []*cli.Command{
					{
						Name:  "credentials",
						Usage: "get login/password",
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
							h.GetCredentialsByLogin(cCtx.Context, cCtx.String("login"))
							return nil
						},
					},
					{
						Name:  "text",
						Usage: "get text data",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Aliases:  []string{"n"},
								Required: true,
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
							return h.GetTextFileByName(
								cCtx.Context,
								cCtx.String("name"),
								cCtx.String("path"),
							)
						},
					},
					{
						Name:  "binary",
						Usage: "get binary data",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Aliases:  []string{"n"},
								Required: true,
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
							return h.GetBinaryDataByName(
								cCtx.Context,
								cCtx.String("name"),
								cCtx.String("path"),
							)
						},
					},
					{
						Name:  "card",
						Usage: "get bank card",
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
							h.GetBankCardByNumber(cCtx.Context, cCtx.String("number"))
							return nil
						},
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}
}
