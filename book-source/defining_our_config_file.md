# Defining our Config File

You should have `main.go` open at this point with just the `package main`
declaration. To get started, we'll add the Viper import as well as the `fmt`
package and a basic `main` funcion that does nothing.

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("end")
}
```

The very first thing we'll do is define that we want our users to be able to
define a YAML config file in the running directory. For that, we'll use three
functions from the Cobra library

```go
package main

import (
	// ... unchanged ...
)

func main() {
	// ... unchanged ...

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		// handle this error if you desire
		fmt.Println("config file not found")
	}
}
```

And at this point, running your code would just print out the start/end messages
and a message indicating that a config file was not found. Simple enough.

```
$ go run .
Start
config file not found
end
```

Create a config file with the name **config.yaml**. Leave it empty, and then run
program again.

```
touch config.yaml
```

The application should stop complaining about not having a configuration file.
Even with no values, the file exists and it has the right extension (.yaml) so
we know that this file is being read by our program.

```
$ go run .
Start
end
```