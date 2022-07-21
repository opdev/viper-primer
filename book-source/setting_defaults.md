# Setting Defaults

You'll eventually need to set a default value for a configuration item, and
doing so is a single line away.

Let's define an `os` key to our code. Don't add it to your config.yaml!

```go
func main() {
    // ... unchanged ...
	fmt.Println("The os is", viper.GetString("os"))
}

func init() {
    // ... unchanged ...
	viper.SetDefault("os", "centos")
}
```

Run the code and see that the default value is returned:

```shell
$ go run .
# ... unchanged ...
The os is centos
```
