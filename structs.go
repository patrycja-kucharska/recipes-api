package main

type Ingredient struct {
	Id   string `db:"id" json:"id"`
	Type string `db:"type" json:"type"`
	Name string `db:"name" json:"name"`
	Unit string `db:"unit_name" json:"unitName"`
}

type Response struct {
	Code   string      `json:"code"`
	Reason string      `json:"reason,omitempty"`
	Body   interface{} `json:"body,omitempty"`
}
