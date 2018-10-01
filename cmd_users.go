package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/urfave/cli"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
)

func cmdUsers(c *cli.Context) error {
	return listUsers(c)
}

func listUsers(c *cli.Context) error {
	data, response, err := apiClient.IdentityManagement.ListUsers()
	common.ErrorCheck(err)

	switch c.String("output") {
	case "markdown":
		generateMD(data)
	case "table":
		printUsers(data)
	case "json":
		common.PrintJSON(response.Body)
	}

	return nil
}

func printUsers(users *[]edgegrid.AkamaiUser) {
	fmt.Println("# Akamai Users:")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, fmt.Sprint("# ID\tName\tE-Mail\tUI Name\tAccount ID\tLast Login\t2FA enabled?"))
	for _, u := range *users {
		fmt.Fprintln(w, fmt.Sprintf("%v\t%s %s\t%s\t%s\t%s\t%s\t%t",
			u.UIIdentityID, u.FirstName, u.LastName, u.Email, u.UIUserName, u.AccountID, u.LastLoginDate, u.TfaEnabled))
	}
	w.Flush()
}

func generateMD(users *[]edgegrid.AkamaiUser) {
	fmt.Println("## ING Users")
	fmt.Println()
	fmt.Println("Name|E-Mail|Last Login|2 Factor Auth enabled?")
	fmt.Println(":-:|:-:|:-:|:-:")
	for _, user := range *users {
		if !strings.Contains(user.Email, "schubergphilis.com") {
			fmt.Printf("%s %s|%s|%s|%t\n", user.FirstName, user.LastName, user.Email, user.LastLoginDate, user.TfaEnabled)
		}
	}
	fmt.Println()
	fmt.Println("## Schuberg Philis Users")
	fmt.Println("Name|E-Mail|Last Login|2 Factor Auth enabled?")
	fmt.Println(":-:|:-:|:-:|:-:")
	for _, user := range *users {
		if strings.Contains(user.Email, "schubergphilis.com") {
			fmt.Printf("%s %s|%s|%s|%t\n", user.FirstName, user.LastName, user.Email, user.LastLoginDate, user.TfaEnabled)
		}
	}

}
