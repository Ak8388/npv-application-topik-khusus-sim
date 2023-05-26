package entity

type Input struct {
	ModalUsaha float64  `json:"modal"`
	Dp         float64  `json:"dp"`
	Cicilan    float64  `json:"cicilan"`
	Persen     []Persen `json:"persen"`
}

type Persen struct {
	Inflasi float64 `json:"inflasi"`
}

type Irr struct {
	IRR float64 `json:"irr"`
}

type PV struct {
	Pv float64 `json:"pv"`
}

type NPV struct {
	Pv  []float64 `json:"pv"`
	Npv []float64 `json:"npv"`
	IRR float64   `json:"irr"`
}
