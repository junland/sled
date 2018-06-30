# Sled [![Build Status](https://travis-ci.org/junland/sled.svg?branch=master)](https://travis-ci.org/junland/sled) [![Open Source Helpers](https://www.codetriage.com/junland/sled/badges/users.svg)](https://www.codetriage.com/junland/sled)

A simple boiler-plate for web apps.

## Features

* Simple layout. No frameworks. No abstraction.

* Modular (Easy to swap out logger, router, and middleware.)

* Tries to adhere to the 12 factor application paradigm.

* TLS Enabled. (Make sure to generate your key and cert.)

## Getting started

_Binary only builds on non-Windows systems, it's best if you use `MacOS` or `Linux` for development._

1. Pull down the boiler plate repo. (With your Go environment already configured.)

    `go get gitlab.com/junland/sled`

2. Build it!

    `go build`

3. Run it!

    `sudo ./sled`

4. Test it!

    `http://localhost:8080/hello` (Acknowledge the security warning)

    `http://localhost:8080/hello/Bob` (Acknowledge the security warning)

5. Hack it! Use it! Profit!

## Documentation

If you would like to know more about this software, you can visit the repository documentation which is located [here][docs].

## Built With

`github.com/justinas/alice` - Simple middleware chaining library.

`github.com/sirupsen/logrus` -  Structured, pluggable logging for Go.

`github.com/spf13/pflag` - Drop in replacement for the `flag` package. 

`github.com/julienschmidt/httprouter` - A high performance HTTP request router that scales well.

## Versioning

I use [SemVer 2.0.0](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/junland/pak-mule/tags).

## Authors

* **John Unland** - *Initial work* - [junland](https://github.com/junland)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project

## License

Code is licensed under MIT which can be viewed in the `LICENSE` file.

_Please let me know through the issues tracker if you have any questions._

## TODO / Notes

* None here for now.

[docs]: docs
