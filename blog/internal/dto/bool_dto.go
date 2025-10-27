package dto

type BoolDTO struct {
	Result bool `json:"result"`
}

var (
	True  = BoolDTO{Result: true}
	False = BoolDTO{Result: false}
)
