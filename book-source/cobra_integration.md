# Integrating With Cobra

Note that this section leaves our previous `configprinter` project behind. The
referenced code is in the **code/snakes** directory of this
[repository](https://github.com/opdev/viper-primer).

---

[Cobra](https://github.com/spf13/cobra) and viper work really well together.
When building commandline interfaces with cobra, users can specify flags to
indicate behavioral changes in their program, and Viper can receive those values
and store them in configuration.

For this, I've scaffolded a new project using `cobra-cli`. When calling
`cobra-cli`, make sure to use the `--viper` flag to scaffold out the Viper
integrations.

```
$ cobra-cli init --viper=true
Your Cobra application is ready at
/Users/me/.go/src/github.com/opdev/viper-primer/code/snakes
```

As expected, we get some scaffolded code, and most importantly, we get our `cmd/root.go`

```
$ tree
.
├── LICENSE
├── cmd
│   └── root.go
├── go.mod
├── go.sum
└── main.go
```

Typically, the `rootCmd` is empty and is an organizational command (like calling
`cobra-cli` with no subcommands), but for this example, we'll make our changes
there so we don't have to scaffold out extra subcommands.

If you've gone through the [Cobra
Primer](https://opdev.github.com/cobra-primer/), you'll notice a few differences
in the `cmd/root.go` that's scaffolded when Viper support is enabled for your
application.

Notably, you'll see the `init()` function now calls an `initConfig` function.
And that `initConfig` function initializes a basic Viper configuration similar
to what we did in the **configprinter** example project. Here's what we get (a
few comments and blocks have been omitted).

```go
package cmd

// imports omitted


var cfgFile string

// ... cobra code omitted

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.snakes.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".snakes" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".snakes")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
```

Reading the `initConfig`, code, we'll see that that the user can set a config
file path, which is stored in the variable `cfgFile`, via a flag. If they do,
that's what's used by Viper. Alternatively, we will check `$HOME/.snakes.yaml`
for a configuration file.

Finally, `viper.AutomaticEnv()` is called, meaning that we can override values
using environment variables. Finally, `ReadInConfig` is called, and so the Viper
configuration file is read.

Uncomment the `rootCmd.Run` struct key and just fill in some placeholder logic
so that we can execute it without getting help output. I've also removed the
`Long` and `Short` descriptions just so that this text looks cleaner:

```go
var rootCmd = &cobra.Command{
	Use:   "snakes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run executed")
	},
}
```

At this point, we can run our command and get our demo print statements:

```
$ go run .
Run executed
```

`cobra-cli` scaffolded a `--toggle` boolean flag for us. Let's add another line
to print the value of that flag.

```go
// unchanged
var rootCmd = &cobra.Command{
	Use:   "snakes",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run executed")

		toggleValue, _ := cmd.Flags().GetBool("toggle")
		fmt.Println("Toggle is set to: ", toggleValue)
	},
}
// unchanged
```

So if we run this now, we see the toggle value printed (it has a default value
of false). We can set it to true by running something like the following:

```
$ go run . --toggle
Run executed
The toggle flag is set to:  true
```

But right now, this value never makes it to our Viper configuration. Try to access to `toggle` key using viper, and you'll see that regardless of whether the user runs the command with the `--toggle` flag, Viper never sees the change.


```go
// unchanged
var rootCmd = &cobra.Command{
	Use:   "snakes",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run executed")

		toggleValue, _ := cmd.Flags().GetBool("toggle")
		fmt.Println("The toggle flag is set to: ", toggleValue)
		fmt.Println("The toggle config in viper is set to:", viper.GetBool("toggle"))
	},
}
// unchanged
```

And here's what we see when we run this with the `--toggle` flag:

```
$ go run . --toggle
Run executed
The toggle flag is set to:  true
The toggle config in viper is set to: false
```

Right now, flag values aren't being stored in our Viper configuration. Luckily,
Viper provides a way to bind the flag's value to the configuration.

In our `init` function, where we define the `rootCmd`'s flag(s), we can also
bind an equivalent Viper configuration value:

```go
func init() {

    // ... unchanged ...

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.BindPFlag("toggle", rootCmd.Flags().Lookup("toggle"))
}
```

Note that _technically_ we can bind _another_ Viper configuration key to the
flag's key, but mapping the values similarly is common. 

Now, we see our Viper configuration contains the value of our flag.

```bash
$ go run . --toggle
Run executed
The toggle flag is set to:  true
The toggle config in viper is set to: true

$ go run . 
Run executed
The toggle flag is set to:  false
The toggle config in viper is set to: false
```

## Accessing configuration once you've integrated Cobra and Viper

Once you've bound your Viper configuration to Cobra, you can technically access
your user's configuration using either. With that said, it (subjectively) makes
sense to leverage your Viper configuration as your source of truth once your
cobra flags are bound. This is because you can store the values of your flags in
your Viper configuration, but it's not exactly a bi-directional relationship,
and your Viper configuration values don't get stored in your cobra flags.

To that end, if you're going to leverage viper, I would probably aim to use it
as your source of truth.

## A note on case discrepancies between viper and cobra flags.

Let's add a `--log-level` string flag, and corresponding Viper binding:

```go
func init() {

    // ... unchanged ...
	
	rootCmd.Flags().StringP("log-level", "l", "", "Help message for log level")
	viper.BindPFlag("logLevel", rootCmd.Flags().Lookup("log-level"))
}`
```

Viper configurations are typically configured using camelCase or snake_case, but
long flags are typically hyphenated. You'll likely want to bind your flag values
to appropriate Viper configuration values in cases where you have hyphenated
long flags by binding them to equivalent values in your preferred Viper-friendly case.

Here, I've decided to use camelCase. If the user provides a configuration file,
the key for log level would need to be `logLevel`.

## A note on precedence

It's important to understand the precedence of your configuration, especially
when you integrate with Cobra and mix flags in. I won't document it here, but
it's documented in the Viper documentation:

https://github.com/spf13/viper#why-viper

## A note on flag integration

While we demonstrate that Cobra and Viper play nicely together, it's also
possible to integrate with the flag and pflag packages in Golang directly.
