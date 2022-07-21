# Reading other config files

It stands to reason that you may want to allow for multiple locations where a
configuration file might be found. An example might be if your application is
installed via a linux distribution's package manager, and its configuration file
is placed inside the `/etc` directory.

We've only told Viper to search relative to our current directory (i.e. `.`). So
let's also add `/etc/configprinter` and also the local `./conf` directory.

```go
func init() {
	viper.AddConfigPath("/etc/configprinter/")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath(".")
    // ... unchanged ...
}
```

Move your configuration to any of these locations and Viper should still find
your config file.