package model

type ClothDTO struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	UnifiedSize   string `json:"unifiedSize"`
	NumericalSize int    `json:"numericalSize"`
}

type ClothDAO struct {
	Id            int
	Name          string
	Type          string
	UnifiedSize   string
	NumericalSize int
}
