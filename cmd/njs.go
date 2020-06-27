package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"nativejs/internal"
	"nativejs/internal/codegen"
	"nativejs/internal/parser"
	"os"
)

func main() {
	app := cli.App{
		Name:    "njs",
		Usage:   "A tool for managing NJS source code",
		Version: "0.0.0-beta.1",
		Authors: []*cli.Author{
			{
				Name:  "Ameer Jhan",
				Email: "ameerjhanprof@gmail.com",
			},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:        "build",
			Usage:       "Build a njs source code",
			Description: "Compiles njs source code to native binaries for distribution",
			ArgsUsage:   "[filename]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() == 1 {
					source, err := ioutil.ReadFile(ctx.Args().Get(0))

					if err != nil {
						return err
					}

					// parse the source code
					grammarParser := parser.NewParser(string(source))
					// generate LLVM IR
					llvmIR := codegen.GenerateIR(grammarParser.Program())

					print(llvmIR)

					return nil
				} else {
					return errors.New("no filename was provided, run njs help to know more")
				}
			},
		},
	}

	err := app.Run(os.Args)

	internal.PrintError(err)
}
