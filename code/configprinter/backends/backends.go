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
