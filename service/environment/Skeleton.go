package environment

import "encoding/json"

type Skeleton struct {
	Package   string                 `json:"package"`
	Version   string                 `json:"version"`
	Author    string                 `json:"author"`
	Email     string                 `json:"email"`
	License   string                 `json:"license"`
	Framework Framework              `json:"framework"`
	Script    map[string]interface{} `json:"script"`
}

func (s *Skeleton) ToJson() string {
	response, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(response)
}
