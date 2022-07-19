package environment

type Skeleton struct {
	Package   string     `json:"package"`
	Version   string     `json:"version"`
	Author    string     `json:"author"`
	Email     string     `json:"email"`
	License   string     `json:"license"`
	Framework Framework  `json:"framework"`
	Script    []struct{} `json:"script"`
}
