# What's actually happening

Before we go on, I want to point out that what's happening here isn't magic.

The Viper library makes a
[Viper](https://github.com/spf13/viper/blob/v1.12.0/viper.go#L181) type, which
itself is responsible for implementing all of the logic we've called so far.

That includes:

- setting our configfile name to `config`

- setting our extension to `yaml`

- setting the path to `.` or "my current directory"

- reading in the configuration file at the the above path

- getting values from the configuration file read from that path


But we didn't create any instances of Viper! That's because the viper library
itself provides a "global" or "singleton" instance of the `Viper` struct (seen
[here](https://github.com/spf13/viper/blob/v1.12.0/viper.go#L63)) This is purely
out of convenience to callers, and it allows developers to simply import the
library, tell the library where the configuration file is, and be on their way
writing business logic

If you wanted to create your own instance of a `Viper` struct, you can
absolutely use the `viper.New()` function to instantiate one and use it
standalone. For the purpose of this text, however, we'll use the singleton.

Now, back on topic!