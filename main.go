package main

import (
	"os"

	"github.com/gonsole/gonsole/service/environment"
)

func main() {
	skeleton := environment.Skeleton{}
	skeleton.Package = "main"
	skeleton.Version = "1.0.0"
	skeleton.Author = "gonsole"
	skeleton.Email = "info@gonsole.io"
	skeleton.License = "MIT"
	skeleton.Framework = environment.Framework{}
	skeleton.Framework.Path.Controller = "controller"
	skeleton.Framework.Path.Entity = "entity"
	skeleton.Framework.Path.Template = "template"
	skeleton.Framework.Orm = environment.Orm{}
	skeleton.Framework.Orm.Driver = "mysql"
	skeleton.Script = make(map[string]interface{})
	skeleton.Script["dev"] = "go run ."
	skeleton.Script["build"] = "go build ."
	f, _ := os.Create("gonsole.json")
	f.WriteString(skeleton.ToJson())
}
