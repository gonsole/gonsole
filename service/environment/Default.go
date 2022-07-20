package environment

func (s *Skeleton) PackageDefault() {
	s.Package = "main"
}
func (s *Skeleton) VersionDefault() {
	s.Version = "0.0.1"
}
func (s *Skeleton) AuthorDefault() {
	s.Author = "Gonsole"
}
func (s *Skeleton) EmailDefault() {
	s.Email = "info@gonsole.io"
}
func (s *Skeleton) LicenseDefault() {
	s.License = "MIT"
}
func (p *Path) ControllerDefault() {
	p.Controller = "controller"
}
func (p *Path) EntityDefault() {
	p.Entity = "entity"
}
func (p *Path) TemplateDefault() {
	p.Template = "template"
}
func (o *Orm) DriverDefault() {
	o.Driver = "entgo"
}
func (r *Route) DriverDefault() {
	r.Driver = "fiber"
}
