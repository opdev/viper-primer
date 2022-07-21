# Setting Values

As we've seen already, Viper (when used with the library's singleton) can easily
be used to access configuration _across_ packages. What if a value needs to be
changed after reading the configuration due to a change in state internally in
your application? 

Viper also provides a setter (e.g. `viper.Set`) to do exactly that. Let's set a key `workers`. 

```go
func main() {
    // ... unchanged ...

	fmt.Println("the worker count is", viper.GetInt("workers"))
	fmt.Println("setting workers to 4")
	viper.Set("workers", 4)
	fmt.Println("the worker count is:", viper.GetInt("workers"))
}
```

This produces:

```shell
$ go run .
# ... unchanged ...
the worker count is 0
setting workers to 4
the worker count is: 4
```

Now the key `workers` can be accessed from any other package by accessing the
value from Viper.