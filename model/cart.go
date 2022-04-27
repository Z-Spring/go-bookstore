package model

type Cart struct {
	Id       int     `json:"id"`
	Pid      int     `json:"pid"`
	Uid      string  `json:"uid"`
	Count    int     `json:"count"`
	SumPrice float64 `json:"sum_price"`
}
