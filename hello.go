package modtest

import "encoding/json"

type Hello struct{
	Name string `json:"name"`
}


func NewHello() *Hello{
	return &Hello{Name: "foo"}
}

func (h *Hello) ToJsonString() string {
	return h.jsonDecodePrivate()
}

func FromJsonString(s string) *Hello{
	var h = NewHello()
	json.Unmarshal([]byte(s), h)
	return h
}

func (h *Hello) jsonDecodePrivate() string{
	r, _ := json.Marshal(h)
	return string(r)
}
