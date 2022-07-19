package environment

type Framework struct {
	Path struct {
		Controller string `json:"controller"`
		Entity     string `json:"entity"`
		Template   string `json:"template"`
	}
	Orm   Orm   `json:"orm"`
	Route Route `json:"route"`
}
