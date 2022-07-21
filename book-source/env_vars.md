# Adding Environment Variables

One of the more painful things in our introductory example was handling
environment variables, and I promised Viper would make this easier.

Viper actually has a feature called `AutomaticEnv` which will automatically read
your configuration keys and look for a standardized environment variable string
that matches.

The [docs](https://github.com/spf13/viper#working-with-environment-variables) go
into more detail about the semantics of working with environment variables, but
suffice to say that `AutomaticEnv` will take a key `logLevel` and look for the
equivalent `LOGLEVEL` from the environment.

Let's enable `AutomaticEnv` by adding it to our `init()` function.

```go
func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
}
```

The run your code while overridding the `debug` value for `LOGLEVEL` with `info`.

```
$ LOGLEVEL=info go run .
The log level is set to: info
```

However, it's possible (and almost likely) for another application to also
expect the `LOGLEVEL` environment variable, so you can add a prefix to your
environment variables to disambiguate them from others.

```go
func init() {
    // ... unchanged ...
	viper.SetEnvPrefix("cp")
	viper.AutomaticEnv()
}
```

Now our environment variables must have the `CP_` prefix, e.g. `CP_LOGLEVEL`.

```shell
# nothing happened
$ LOGLEVEL=info go run .
The log level is set to: debug
...

# it changed
$ CP_LOGLEVEL=info go run .
The log level is set to: info
...
```