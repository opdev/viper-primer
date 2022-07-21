# The init function

So now that we know a little bit about what's happening under the hood, we're
going to change things around just a little bit more before we get back into
Viper's features.

In Go, the [init](https://go.dev/doc/effective_go#init) function is a special
function that allows you to establish any kind of state you need to manage for
your code to operate _before_ your `main` function is executed. The link above
describes in more detail the exact semantics of when the `init()` function runs,
if you're interested.

We currently have our Viper configuration taking place in our `main()` function,
but it's more common to have this take place in an `init()` function, or
something init-adjacent (like another function that is called by a package's
`init` function). Because Viper tends to represent the global configuration of
an application, it's fairly common to include its configuration in the `init`
function of your `main.go`, or as close to your entrpoint as possible. That
ensures that configuration parsing _always_ happens when your entrypoint is
called.

Let's move most of our Viper-related function calls over to the `init` function
before we go further.

```go
package main

import // ... unchanged ... 

func main() {
	if err := viper.ReadInConfig(); err != nil {
		// handle this error if you desire
		fmt.Println("config file not found")
	}

	fmt.Println("The log level is set to:", viper.GetBool("logLevel"))
	fmt.Println("logging is enabled:", viper.GetBool("enableLogging"))
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
}
```

You might notice that I didn't move the `viper.ReadInConfig` call into the init
function. You absolutely can do this as well - just make sure to make it the
last thing called. For this example, I've left it in `main` as it makes things
easier to represent.

Now back to the config hacking!