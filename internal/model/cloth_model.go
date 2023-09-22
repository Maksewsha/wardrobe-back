package model

type Cloth struct {
	Id            uint64 `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	NumericalSize int    `json:"numericalSize"`
	UnifiedSize   string `json:"unifiedSize"`
	WardrobeId    uint64 `json:"wardrobe_id"`
	Poster        []byte `json:"poster"`
	Image         []byte `json:"image"`
}
