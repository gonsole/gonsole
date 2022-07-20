package environment

import "github.com/gonsole/gonsole/service/condition"

func (s *Skeleton) Check() {
	condition.IfThen(s.Package == "", func() {
		s.PackageDefault()
	})
	condition.IfThen(s.Version == "", func() {
		s.VersionDefault()
	})
	condition.IfThen(s.Author == "", func() {
		s.AuthorDefault()
	})
	condition.IfThen(s.Email == "", func() {
		s.EmailDefault()
	})
	condition.IfThen(s.License == "", func() {
		s.LicenseDefault()
	})
	condition.IfThen(s.Framework.Path.Controller == "", func() {
		s.Framework.Path.ControllerDefault()
	})
	condition.IfThen(s.Framework.Path.Entity == "", func() {
		s.Framework.Path.EntityDefault()
	})
	condition.IfThen(s.Framework.Path.Template == "", func() {
		s.Framework.Path.TemplateDefault()
	})
	condition.IfThen(s.Framework.Orm.Driver == "", func() {
		s.Framework.Orm.DriverDefault()
	})
	condition.IfThen(s.Framework.Route.Driver == "", func() {
		s.Framework.Route.DriverDefault()
	})
}
