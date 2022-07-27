# Reading Simple Configuration

Say we've got an application that has a simple configuration file that looks like this:

```yaml
logLevel: debug
magicFeatureEnabled: true
somethingNeatEnabled: false
```

That's a pretty simple YAML file that we can easily represent as a Golang struct like so:

```go
type Configuration struct {
	LogLevel             string `yaml:"logLevel"`
	MagicFeatureEnabled  bool   `yaml:"magicFeatureEnabled"`
	SomethingNeatEnabled bool   `yaml:"somethingNeatEnabled"`
}
```

And these struct tags on each key of the `Configuration` struct are read by tools like the [go-yaml](https://github.com/go-yaml/yaml.git) library to map the data in our YAML file to the struct in our code. Reading this file is as easy as this:

```Go
package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	d, _ := os.ReadFile("config.yaml")
	var cfg Configuration
	yaml.Unmarshal(d, &cfg)

	fmt.Println("The Log Level is: ", cfg.LogLevel)
	fmt.Println("We are using the Magic Feature: ", cfg.MagicFeatureEnabled)
	fmt.Println("We are using Something Neat ", cfg.SomethingNeatEnabled)
}

type Configuration struct {
	LogLevel             string `yaml:"logLevel"`
	MagicFeatureEnabled  bool   `yaml:"magicFeatureEnabled"`
	SomethingNeatEnabled bool   `yaml:"somethingNeatEnabled"`
}
```

Running this with our previous configuration file gives us this result:

```
The Log Level is:  debug
We are using the Magic Feature:  true
We are using Something Neat  false
```

So far so good.
