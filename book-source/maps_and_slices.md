# Maps and Slices

One of the huge benefits of using a tool like Viper is the ability to use dot
notation to get various config values in your tree. So let's extend our
config.yaml to look like this:

```yaml
logLevel: debug
enableLogging: true
metrics:
  listenAddress: 127.0.0.1
  listenPort: 9999
backends:
  - 192.168.10.01:8001
  - 192.168.10.01:8002
  - 192.168.10.01:8003
  - 192.168.10.01:8004
```

Now we've got a `metrics` entry with various keys, and a `backends` array with
various entries. And we can access them using Viper getter functions.

The `metrics` end effectively is an "object", or a "map" in Go terms. Accessing
the listenAddress key can be done using dot notation, similar to what you might
see using tools like `jq` or `yq`.

```go
func main() {
    // ... unchanged ...

	fmt.Println("The metrics endpoint is:",
		fmt.Sprintf("%s:%s",
			viper.GetString("metrics.listenAddress"),
			viper.GetString("metrics.listenPort"),
		),
	)
}
```

An item in the `backends` key can be accessed the same way:

```go
func main() {
    // ... unchanged ...
	fmt.Println("The first backend is: ", viper.GetString("backends.0"))
}
```

But what if you need the entire slice of `backends`? There's a getter for that, too:

```go
func main() {
    // ... unchanged ...
fmt.Println("Here are all the backends:", viper.GetStringSlice("backends"))
}
```

After adding all of these, here's what we see in our output:

```
The log level is set to: false
logging is enabled: true
The metrics endpoint is: 127.0.0.1:9999
The first backend is:  192.168.10.01:8001
Here are all the backends: [192.168.10.01:8001 192.168.10.01:8002 192.168.10.01:8003 192.168.10.01:8004]
```
