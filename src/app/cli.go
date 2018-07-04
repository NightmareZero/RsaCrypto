package app

import (
	cli "gopkg.in/urfave/cli.v1"
)

// CliApp ...
// cli路由表
func CliApp() *cli.App {
	var app = cli.NewApp()
	app.Name = "RasCripto"
	app.HelpName = "RasCripto"
	app.Version = "0.1.0"
	app.Description = "简易Rsa加密解密工具"
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "生成公钥和私钥",
			Action: func(c *cli.Context) error {
				return Generate()
			},
		},
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "加密指定文件名",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file,f",
					Usage: "加密的文件名",
				},
			},
			Action: func(c *cli.Context) error {
				fname := c.String("file")
				return EncryptFile(fname)
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "解密指定文件名",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file,f",
					Usage: "解密的文件名",
				},
			},
			Action: func(c *cli.Context) error {
				fname := c.String("file")
				return DecryptFile(fname)
			},
		},
	}
	return app
}
