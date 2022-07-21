package main

import (
	"fmt"
	"strings"

	"example.com/configprinter/backends"
	"github.com/spf13/viper"
)

func main() {
	if err := viper.ReadInConfig(); err != nil {
		// handle this error if you desire
		fmt.Println("config file not found")
	}

	fmt.Println("The log level is set to:", viper.GetString("logLevel"))
	fmt.Println("logging is enabled:", viper.GetBool("enableLogging"))

	fmt.Println("The metrics endpoint is:",
		fmt.Sprintf("%s:%s",
			viper.GetString("metrics.listenAddress"),
			viper.GetString("metrics.listenPort"),
		),
	)

	fmt.Println("The first backend is: ", viper.GetString("backends.0"))

	fmt.Println("Here are all the backends:", viper.GetStringSlice("backends"))

	var metricsConf MetricsConfig
	viper.UnmarshalKey("metrics", &metricsConf)
	fmt.Println("metricsConf ListenAddress", metricsConf.ListenAddress)
	fmt.Println("metricsConf ListenPort", metricsConf.ListenPort)

	backends.Enumerate()

	fmt.Println("The os is", viper.GetString("os"))

	fmt.Println("the worker count is", viper.GetInt("workers"))
	fmt.Println("setting workers to 4")
	viper.Set("workers", 4)
	fmt.Println("the worker count is:", viper.GetInt("workers"))
}

func init() {
	viper.AddConfigPath("/etc/configprinter/")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	viper.SetEnvPrefix("CP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()
	viper.SetDefault("os", "centos")
}

type MetricsConfig struct {
	ListenAddress string
	ListenPort    string
}
