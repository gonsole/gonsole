package main

import (
	"github.com/gonsole/gonsole/service/environment"
)

func main() {
	skeleton := new(environment.Skeleton)
	skeleton.Include()
	skeleton.Check()
	skeleton.Create()
}
