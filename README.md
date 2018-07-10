# Akamai CLI for Users Audit
*NOTE:* This tool is intended to be installed via the Akamai CLI package manager, which can be retrieved from the releases page of the [Akamai CLI](https://github.com/akamai/cli) tool.

### Local Install, if you choose not to use the akamai package manager
If you want to compile it from source, you will need Go 1.9 or later, and the [Glide](https://glide.sh) package manager installed:
1. Fetch the package:
   `go get https://github.com/partamonov/akamai-cli-users`
1. Change to the package directory:
   `cd $GOPATH/src/github.com/partamonov/akamai-cli-users`
1. Install dependencies using Glide:
   `glide install`
1. Compile the binary:
   `go build -ldflags="-s -w -X main.version=X.X.X" -o akamai-users`

### Credentials
In order to use this configuration, you need to:
* Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Expects `default` section in .edgerc, can be changed via --section parameter

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

## Overview
The Akamai Users will give you users overview in your account

## Main Command Usage
```shell
NAME:
   akamai users - A CLI to interact with Akamai Identity Management

USAGE:
   akamai users [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     list, ls  Get a list of [subcommand]]
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: [$AKAMAI_EDGERC])
   --debug                  Debug info
   --no-color               Disable color output
   --raw                    Show raw output. It will be JSON format
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```

### Raw output. (JSON)
Any command can run with `--raw` parameter. Output will be as received from Akamai in JSON format

```shell
akamai users --section custom-name --raw list users
[
  {
    "uiIdentityId": "...",
    "firstName": "Name",
    "lastName": "LastName",
    "uiUserName": "a@a.com",
    "email": "bas.grabowsky@ing.nl",
    "accountId": "...",
    "lastLoginDate": "...",
    "tfaEnabled": true,
    "tfaConfigured": true
  }
]
```

### List commands

#### Users

You can list all users

There are 2 types of output:

* markdown (default)
* table

To specify desired output, please use `--output` key

```shell
> akamai users list users --output table
...
```
