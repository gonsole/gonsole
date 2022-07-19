package environment

type Orm struct {
	Driver string `json:"driver"`
}

func (o *Orm) SetDriver(driver string) {
	o.Driver = driver
}
