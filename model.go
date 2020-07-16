package main

type Material struct {
	MaterialCode string `json:"materialCode"`
	MaterialDesc string `json:"materialDesc"`
	Uom          string `json:"uom"`
	Release      bool   `json:"release"`
}
