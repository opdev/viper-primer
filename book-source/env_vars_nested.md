# Setting Nested Keys via Environment

Setting a nested key via the environment takes another configuration option.
Accessing nested keys in our configuration cannot be done in most shells using
dot notation, so we just tell Viper to replace any paths that contain a dot with
an underscore.

```go
func init() {
    // ... unchanged ...
	viper.SetEnvPrefix("CP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()
}
```

Now, the `metrics.listenAddress` path can be accessed at the environment. Don't forget the prefix!

```shell
$ CP_METRICS_LISTENADDRESS=localhost go run .
The log level is set to: debug
logging is enabled: true
The metrics endpoint is: localhost:9999         # this changed!
The first backend is:  192.168.10.01:8001
Here are all the backends: [192.168.10.01:8001 192.168.10.01:8002 192.168.10.01:8003 192.168.10.01:loca8004]
```

You can also change the `backends` array. Just separate your items with spaces:

```shell
$ CP_BACKENDS="10.0.0.1:8001 10.0.0.1:8002" go run .
The log level is set to: debug
logging is enabled: true
The metrics endpoint is: 127.0.0.1:9999
The first backend is:  
Here are all the backends: [10.0.0.1:8001 10.0.0.1:8002]    # two new items!
```