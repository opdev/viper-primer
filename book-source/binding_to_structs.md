# Binding to structs

So while it's not strictly required, you can optionally bind your entire
configuration, or even just pieces of it, to golang struct types. Say we had a
new type `MetricsConfig` with our metrics-related keys.

```go
type MetricsConfig struct {
	ListenAddress string
	ListenPort    string
}
```

We can read (or "Unmarshal") the `metrics` key from our config directly into this struct:

```go
func main() {
    // ... unchanged ...

	var metricsConf MetricsConfig
	viper.UnmarshalKey("metrics", &metricsConf)
	fmt.Println("metricsConf ListenAddress", metricsConf.ListenAddress)
	fmt.Println("metricsConf ListenPort", metricsConf.ListenPort)
}
```

Your struct then contains the configuration values for use in whatever functions
you might have.

```shell
$ go run .
# ... previous output omitted ...
metricsConf ListenAddress 127.0.0.1
metricsConf ListenPort 9999
```

There is also an `Unmarshal` functionthat allows you to bind the entire
configuration to a struct, should you need it