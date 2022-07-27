# Hands on with Viper

Viper is a library, and as such, we just need to import it or `go get` the
module into our application. For this cheat sheet, we're mostly interested in
the Viper library calls and so we won't really do much business logic other than
just printing out our values. If you want to follow along, create a directory
called `configprinter` and install Viper into this path.

```bash
# create the directory and change into it
mkdir configprinter
cd configprinter

# create an empty main.go
echo package main >> main.go

# initialize the module and install viper
go mod init example.com/configprinter
go get github.com/spf13/viper
go mod tidy
```

Then open up `main.go`.

NOTE: Viper can be very tightly integrated with Cobra. We'll look at that later.