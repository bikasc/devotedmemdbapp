##   Golang devotedbapp CLI

Demo console app for the Golang In-memory Data Management CLI Project.
This project uses hashicorp/go-memdb in memory database (https://github.com/hashicorp/go-memdb)

## Running app

A proper Go environment is required in order to run this project.

To run the project

`go build devotedbapp`
`./devotedbapp`
The console app is a commandline app. If app launch with command line arguments, it will do the job and exit.
It the app is launched without command line arguments, it will be interactive command line app and accept the 
following commands until END command is entered.
END                 
SET  [key] [value]
GET  [key]
COUNT [value]
DELETE [key]
BEGIN
ROLLBACK
COMMIT

Once setup, tests can be run with the following command:

`go test -v ./data/`

### Running with Docker

To build the image from the Dockerfile, run:

`docker build -t devote-db-app .`

To start an interactive shell, run:

`docker run -it --rm --name devote-memdb devote-db-app`

From inside the shell, run the app with:
`/go/bin/devotedbapp`
From inside the shell, chnage directory to /src/app and run the app by following
`devotedbapp`
