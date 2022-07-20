package environment

import (
	"encoding/json"
	"os"
)

var gonsole string = "gonsole.json"

type Framework struct {
	Path  Path  `json:"path"`
	Orm   Orm   `json:"orm"`
	Route Route `json:"route"`
}
type Path struct {
	Controller string `json:"controller"`
	Entity     string `json:"entity"`
	Template   string `json:"template"`
}
type Orm struct {
	Driver string `json:"driver"`
}
type Route struct {
	Driver string `json:"driver"`
}

type Script map[string]interface{}

type Skeleton struct {
	Package   string    `json:"package"`
	Version   string    `json:"version"`
	Author    string    `json:"author"`
	Email     string    `json:"email"`
	License   string    `json:"license"`
	Framework Framework `json:"framework"`
	Script    Script    `json:"script"`
}

func (s *Skeleton) ToJson() string {
	response, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(response)
}
func (s *Skeleton) Include() bool {
	f, err := os.ReadFile(gonsole)
	if err != nil {
		return false
	}
	err = json.Unmarshal(f, s)
	if err != nil {
		return false
	}
	return true

}
