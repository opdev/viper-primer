# Reading configuration across packages

One of the nicer features of having Viper manage your configuration in its own
singleton is that you can read your configuration across packages, so long as
you've already executed your viper initialization when you work with your other packages.

So as an example, I'll create a new directory (to represent a new package, or
module) called **backends**, and I'll write a small function that enumerates my
`backends`.

I'm running this from my project's base directory.

```bash
mkdir backends
echo package backends >> backends/backends.go
```

I have this sample function `Enumerate`. The logic isn't super important.

```go
package backends

import (
	"fmt"

	"github.com/spf13/viper"
)

func Enumerate() {
	fmt.Println("Hello from the backends package!")
	backends := viper.GetStringSlice("backends")
	for i, backend := range backends {
		fmt.Printf("%d: %s\n", i, backend)
	}
}


```

What is important is that you notice that I've simply
called `viper.GetStringSlice` directly. Access to configuration entries is
vastly simplified using tools like Viper. This is also made possible because
we've leveraged our `main.go`'s `init` function to configure Viper, and that
doing so causes that configuration to take place well before any of this code is
called.

Add this to your `main` function and see that it works as expected.

```go
func main() {
    // ... unchanged ...

	backends.Enumerate()
}
```

And the result is as expected!

```bash
$ go run .
# ... unchanged ...
Hello from the backends package!
0: 192.168.10.01:8001
1: 192.168.10.01:8002
2: 192.168.10.01:8003
3: 192.168.10.01:8004
```