# Sled [![Build Status](https://travis-ci.org/junland/sled.svg?branch=master)](https://travis-ci.org/junland/sled) [![Open Source Helpers](https://www.codetriage.com/junland/sled/badges/users.svg)](https://www.codetriage.com/junland/sled)

A simple boiler-plate for web apps.

## Features

* Simple layout. No frameworks. No abstraction.

* Modular (Easy to swap out logger, router, and middleware.)

* Tries to adhere to the 12 factor application paradigm.

* Secure from the start. HTTPS on default. (Make sure to generate your key and cert.)

## Get started

_Works best if you use `MacOS` or `Linux` for development._

1. Pull down the boiler plate repo. (With your Go environment already configured.)

    `go get github.com/junland/sled`

2. Generate test certs.

    `make test-tls`

3. Build it!

    `go build`

4. Run it!

    `sudo ./sled`

5. Test it!

    `https://localhost:443/hello` (Acknowledge the security warning)

    `https://localhost:443/hello/Bob` (Acknowledge the security warning)

6. Hack it! Use it! Profit!

## Contents

`cmd.go` - Command / flag file for interacting with the server binary.

`main.go` - Main file where everything comes together.

`routes.go` - File for putting business logic for your endpoints.

`utils.go` - Generic file for putting custom utilities for your app. (Comes with functions to interacting with PID files and enviroment variables.)

`middleware.go` - File for placing middleware functions.

## 3rd Party Libraries Used

`github.com/justinas/alice` - Simple middleware chaning library.

`github.com/sirupsen/logrus` -  Structured, pluggable logging for Go.

`github.com/julienschmidt/httprouter` - A high performance HTTP request router that scales well.

## License

Code is licensed under MIT which can be viewed in the `LICENSE` file.

_Please let me know through the issues tracker if you have any questions._

## TODO / Notes

* None here for now.
