package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"gopkg.in/urfave/cli.v2"
)

var goos string
var goarch string

func main() {

	config, err := config.Create()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}

	////we're going to load the config file manually, since we need to validate it.
	//err = config.ReadConfig("/etc/mediadepot.yaml")          // Find and read the config file
	//if _, ok := err.(errors.ConfigFileMissingError); ok { // Handle errors reading the config file
	//	//ignore "could not find config file"
	//} else if err != nil {
	//	os.Exit(1)
	//}

	//createFlags, err := createFlags(config)
	//if err != nil {
	//	fmt.Printf("FATAL: %+v\n", err)
	//	os.Exit(1)
	//}

	cli.CommandHelpTemplate = `NAME:
   {{.HelpName}} - {{.Usage}}
USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}
CATEGORY:
   {{.Category}}{{end}}{{if .Description}}
DESCRIPTION:
   {{.Description}}{{end}}{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`

	cli.AppHelpTemplate = fmt.Sprintf("%s %s", CustomizeHelpTemplate(), cli.AppHelpTemplate)

	app := &cli.App{
		Name:     "mediadepot",
		Usage:    "helping you build the Ultimate Home Media Server",
		Version:  version.VERSION,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Jason Kulatunga",
				Email: "jason@thesparktree.com",
			},
		},
		Flags: []cli.Flag {
			&cli.BoolFlag{
				Name: "debug",
				Value: false,
				Usage: "Enable Debug mode, with extra logging",
			},
		},

		Commands: []*cli.Command{
			{
				Name:  "install",
				Usage: "Install MediaDepot on your Home Server",
				Action: func(c *cli.Context) error {
					fmt.Fprintln(c.App.Writer, c.Command.Usage)

					projectList, err := project.CreateProjectListFromProvidedAnswers(config)
					if err != nil {
						return err
					}

					answerData := map[string]interface{}{}
					if projectList.Length() > 0 && utils.StdinQueryBoolean(fmt.Sprintf("Would you like to create a Drawbridge config using preconfigured answers? (%v available). [yes/no]", projectList.Length())) {

						answerData, err = projectList.Prompt("Enter number to base your configuration from")
						if err != nil {
							return err
						}
					}

					//extend current answerData with CLI provided options.
					cliAnswers, err := createFlagHandler(config, answerData, c.FlagNames(), c)
					if err != nil {
						return err
					}

					createAction := actions.CreateAction{Config: config}
					return createAction.Start(cliAnswers, c.Bool("dryrun"))
				},

				Flags: createFlags,
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(color.HiRedString("ERROR: %v", err))
	}
}


func CustomizeHelpTemplate() string {
	tentacle := "github.com/mediadepot/cli"

	var versionInfo string
	if len(goos) > 0 && len(goarch) > 0 {
		versionInfo = fmt.Sprintf("%s.%s-%s", goos, goarch, version.VERSION)
	} else {
		versionInfo = fmt.Sprintf("dev-%s", version.VERSION)
	}

	subtitle := tentacle + utils.LeftPad2Len(versionInfo, " ", 65-len(tentacle))

	return fmt.Sprintf(utils.StripIndent(
		`
		 ____  ____  __ _  ____  __    ___  __    ____
		(_  _)(  __)(  ( \(_  _)/ _\  / __)(  )  (  __)
		  )(   ) _) /    /  )( /    \( (__ / (_/\ ) _)
		 (__) (____)\_)__) (__)\_/\_/ \___)\____/(____)
		%s
	
		`), subtitle)
}