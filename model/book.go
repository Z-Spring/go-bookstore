package model

type Book struct {
	Id        int     `sql:"not null" json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Price     float64 `json:"price,omitempty"`
	ImagePath string  `sql:"imagepath" json:"image_path,omitempty"`
}
