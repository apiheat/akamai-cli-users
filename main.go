package main

import (
	"os"
	"sort"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var (
	colorOn, raw, debug       bool
	version, appName          string
	configSection, configFile string
	edgeConfig                edgegrid.Config
)

// Constants
const (
	URL     = "/identity-management/v2"
	padding = 3
)

// User data
type User struct {
	UIIdentityID  string `json:"uiIdentityId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	UIUserName    string `json:"uiUserName"`
	Email         string `json:"email"`
	AccountID     string `json:"accountId"`
	LastLoginDate string `json:"lastLoginDate"`
	TfaEnabled    bool   `json:"tfaEnabled"`
	TfaConfigured bool   `json:"tfaConfigured"`
}

func main() {
	_, inCLI := os.LookupEnv("AKAMAI_CLI")

	appName = "akamai-users"
	if inCLI {
		appName = "akamai users"
	}

	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appName
	app.Usage = "A CLI to interact with Akamai Identity Management"
	app.Version = version
	app.Copyright = ""
	app.Authors = []cli.Author{
		{
			Name: "Petr Artamonov",
		},
		{
			Name: "Rafal Pieniazek",
		},
	}

	dir, _ := homedir.Dir()
	dir += string(os.PathSeparator) + ".edgerc"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "section, s",
			Value:       "default",
			Usage:       "`NAME` of section to use from credentials file",
			Destination: &configSection,
			EnvVar:      "AKAMAI_EDGERC_SECTION",
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       dir,
			Usage:       "Location of the credentials `FILE`",
			Destination: &configFile,
			EnvVar:      "AKAMAI_EDGERC",
		},
		cli.BoolFlag{
			Name:        "no-color",
			Usage:       "Disable color output",
			Destination: &colorOn,
		},
		cli.BoolFlag{
			Name:        "raw",
			Usage:       "Show raw output. It will be JSON format",
			Destination: &raw,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Debug info",
			Destination: &debug,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "Get a list of [subcommand]]",
			Subcommands: []cli.Command{
				{
					Name:   "users",
					Usage:  "... users",
					Action: cmdUsers,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "output",
							Value: "markdown",
							Usage: "Output format",
						},
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(c *cli.Context) error {
		if c.Bool("no-color") {
			color.NoColor = true
		}

		config(configFile, configSection)
		return nil
	}

	app.Run(os.Args)
}
