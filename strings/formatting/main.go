package main

import "fmt"

type Config struct {
	Key   string
	Value interface{}
	IsSet bool
}

func (c Config) String() string {
	return fmt.Sprintf("key: %s", "value: %s", c.Key, c.Value)
}

func main() {
	x := 10
	p := &x
	*p = 20
	fmt.Println(x)  // 10
	fmt.Println(*p) // 20
}
