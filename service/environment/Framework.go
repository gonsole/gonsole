package environment

type Framework struct {
	Path  Path  `json:"path"`
	Orm   Orm   `json:"orm"`
	Route Route `json:"route"`
}
