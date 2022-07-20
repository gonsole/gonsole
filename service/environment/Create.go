package environment

import "os"

func (s *Skeleton) Create() {
	CreateDir(s.Framework.Path)
}
func CreateDir(Path Path) {
	os.Mkdir(Path.Controller, 0755)
	os.Mkdir(Path.Entity, 0755)
	os.Mkdir(Path.Template, 0755)
}
