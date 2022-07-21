# Reading other config formats

A given instance of Viper can only read a single configuration _format_ (e.g.
`YAML`), but Viper can read _various_ configuration formats.

We've told Viper directly that our configuration file would be an the YAML
format by calling this function:

```go
	viper.SetConfigType("yaml")
```

This actually isn't strictly required. What this does is it allows us to call
our configuration file simply `config` with no extension. It tells Viper exactly
what format to expect.

If you plan on using file extensions for your configuration file, then you can
simply leave this directive out. Then Viper will look for the file `config`
(which we defined) at any of the configured locations with common file
extensions for its various supported formats.

Here's the same configuration, but as a TOML file:

```TOML
backends = ['192.168.10.01:8001', '192.168.10.01:8002', '192.168.10.01:8003', '192.168.10.01:8004']
enablelogging = true
loglevel = 'debug'
os = 'centos'
workers = 4
[metrics]
listenaddress = '127.0.0.1'
listenport = 9999
```

Delete the config.yaml file you have and replace it with this contents in a file called config.toml.

Then, remove the `viper.SetConfigType` call and rerun the code. Everything should still be in place.

```
$ go run .
The log level is set to: debug
logging is enabled: true
The metrics endpoint is: 127.0.0.1:9999
The first backend is:  192.168.10.01:8001
Here are all the backends: [192.168.10.01:8001 192.168.10.01:8002 192.168.10.01:8003 192.168.10.01:8004]
metricsConf ListenAddress 127.0.0.1
metricsConf ListenPort 9999
Hello from the backends package!
0: 192.168.10.01:8001
1: 192.168.10.01:8002
2: 192.168.10.01:8003
3: 192.168.10.01:8004
The os is centos
the worker count is 4
setting workers to 4
the worker count is: 4
```

By the way, need a quick way to convert your config to a different format? Try setting:

```go
viper.SetConfigType("yourformat")
viper.SafeWriteConfigAs("path")
```

The list of config types supported by Viper are found [here](https://github.com/spf13/viper#reading-config-files).