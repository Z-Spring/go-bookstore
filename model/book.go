package model

type Book struct {
	Id        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Price     float64 `json:"price,omitempty"`
	ImagePath string  `json:"image_path,omitempty"`
}
