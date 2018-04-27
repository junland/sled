# Download and build

## System requirements
Fill in your system requirements here.

## Download pre-built binary from releases

You can grab a pre compiled binary in the [releases](https://github.com/junland/sled/releases) section of this repository which is is the fastest way to get this software.

## Building binary from source

### Prerequisites

`go` - 1.8.x or higher, with `GOPATH` set up.

`make` - For using the `Makefile`.

Once your prerequisites have been installed follow these steps:

1. `git clone https://github.com/junland/sled.git`

2. `cd sled`

3. `make binary`

These steps should produce a binary called `sled` which you can execute from your current shell.

## Verifying the binary

To verify that your downloaded or built binary executes correct is to issue this command:

```
./sled version
```

This should display version text of the compiled binary.
