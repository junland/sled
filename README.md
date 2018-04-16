# Sled [![Build Status](https://travis-ci.org/junland/sled.svg?branch=master)](https://travis-ci.org/junland/sled) [![Open Source Helpers](https://www.codetriage.com/junland/sled/badges/users.svg)](https://www.codetriage.com/junland/sled)

A simple boiler-plate for web apps.

## Features

* Simple layout. No frameworks. No abstraction.

* Modular (Easy to swap out logger, router, and middleware.)

* Tries to adhere to the 12 factor application paradigm.

* TLS Enabled. (Make sure to generate your key and cert.)

## Get started

_Works best if you use `MacOS` or `Linux` for development._

1. Pull down the boiler plate repo. (With your Go environment already configured.)

    `go get github.com/junland/sled`

2. Build it!

    `go build`

3. Run it!

    `sudo ./sled`

4. Test it!

    `http://localhost:8080/hello` (Acknowledge the security warning)

    `http://localhost:8080/hello/Bob` (Acknowledge the security warning)

5. Hack it! Use it! Profit!

## Contents

`sled.go` - Entry file.

`cmd/cmd.go` - Command / flag file for interacting with the server binary.

`server/server.go` - Main logic for parsing environment variables and starting of web server.

`server/routes.go` - File for putting business logic for your endpoints.

`server/middleware.go` - File for placing middleware functions.

`utils/utils.go` - Generic file for putting custom utilities for your app. (Comes with functions to interacting with PID files and enviroment variables.)

## 3rd Party Libraries Used

`github.com/justinas/alice` - Simple middleware chaining library.

`github.com/sirupsen/logrus` -  Structured, pluggable logging for Go.

`github.com/julienschmidt/httprouter` - A high performance HTTP request router that scales well.

## License

Code is licensed under MIT which can be viewed in the `LICENSE` file.

_Please let me know through the issues tracker if you have any questions._

## TODO / Notes

* None here for now.
