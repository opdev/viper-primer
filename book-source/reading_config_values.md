# Reading Config Values

At this point, we've already defined that the user can put in configuration into
a file **config.yaml** and the application will read it. But what config values
is it reading?

When using viper, access to your configuration values is pretty free-form, and
doesn't require the use of any structs or defined types in your code.

Let's add the `logLevel` key to our configuration.

```shell
echo logLevel: debug >> config.yaml
```

And at the very bottom of our `main()` function, we'll add a print statement
that queries from Viper the value of our `logLevel` key.

```go
package main

// ... unchanged ...

func main() {
    // ... unchanged ...

	fmt.Println("The log level is set to:", viper.GetString("logLevel"))
}
```

Running the command shows you the same value for `logLevel` as is in our configuration.

```
$ go run .
Start
The log level is set to: debug
end
```

And so you really only _need_ to know the _type_ of the value in your
configuration. This is because Go is strongly typed, but YAML has limited types
that can be represented in its markup. There are equivalent `Get<Type>` functions
for the various types that you can define in Golang.


It's important to note that if you use the **wrong** type getter function, you
get the **empty** value for that type, as `viper.Get<Type>` functions do not
return errors, or panic. So for example, running `viper.GetBool` for `logLevel`
would give us `false`. Running `viper.GetInt` would give us zero, etc.

As a practice: add another key to your config `enableLogging` with value set to
`true`, and get and print out its value. You can decide what the accompanying message says.

```
$ go run .
Start
The log level is set to: false
logging is enabled: true
end
```

Also, feel free to remove the Start and End statements at this point now that we
have something else printing out for us.