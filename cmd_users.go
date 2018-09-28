package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	common "github.com/apiheat/akamai-cli-common"
)

func cmdUsers(c *cli.Context) error {
	return listUsers(c)
}

func listUsers(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/user-admin/ui-identities", URL)

	if debug {
		println(urlStr)
	}

	data := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := userRespParse(data)
	common.ErrorCheck(err)

	if raw {
		common.OutputJSON(result)
		return nil
	}

	switch c.String("output") {
	case "markdown":
		generateMD(result)
	case "table":
		printUsers(result)
	}

	return nil
}

func printUsers(users []User) {
	color.Set(color.FgGreen)
	fmt.Println("# Akamai Users:")
	color.Unset()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, fmt.Sprint("# ID\tName\tE-Mail\tUI Name\tAccount ID\tLast Login\t2FA enabled?"))
	for _, u := range users {
		fmt.Fprintln(w, fmt.Sprintf("%v\t%s %s\t%s\t%s\t%s\t%s\t%t",
			u.UIIdentityID, u.FirstName, u.LastName, u.Email, u.UIUserName, u.AccountID, u.LastLoginDate, u.TfaEnabled))
	}
	w.Flush()
}

func generateMD(users []User) {
	fmt.Println("## ING Users")
	fmt.Println()
	fmt.Println("Name|E-Mail|Last Login|2 Factor Auth enabled?")
	fmt.Println(":-:|:-:|:-:|:-:")
	for _, user := range users {
		if !strings.Contains(user.Email, "schubergphilis.com") {
			fmt.Printf("%s %s|%s|%s|%t\n", user.FirstName, user.LastName, user.Email, user.LastLoginDate, user.TfaEnabled)
		}
	}
	fmt.Println()
	fmt.Println("## Schuberg Philis Users")
	fmt.Println("Name|E-Mail|Last Login|2 Factor Auth enabled?")
	fmt.Println(":-:|:-:|:-:|:-:")
	for _, user := range users {
		if strings.Contains(user.Email, "schubergphilis.com") {
			fmt.Printf("%s %s|%s|%s|%t\n", user.FirstName, user.LastName, user.Email, user.LastLoginDate, user.TfaEnabled)
		}
	}

}

func userRespParse(in string) (users []User, err error) {
	if err = json.Unmarshal([]byte(in), &users); err != nil {
		return
	}
	return
}
