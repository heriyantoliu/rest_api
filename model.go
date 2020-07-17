package main

type Material struct {
	Code    string `json:"code"`
	Desc    string `json:"desc"`
	Uom     string `json:"uom"`
	Release bool   `json:"release"`
	Group   string `json:"group"`
}
