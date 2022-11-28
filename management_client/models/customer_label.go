package management_models

type CustomerLabel struct {
	Id        int64  `json:"id,omitempty"`
	LabelId   int    `json:"label_id,omitempty"`
	LabelName string `json:"label_name,omitempty"`
	Color     string `json:"color,omitempty"`
}
