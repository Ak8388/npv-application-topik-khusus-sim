package helper

import (
	"math"
	"npv-application/entity"
)

// pv = fv/(1+%)2

func Pv(fv float64, inflasi []entity.Persen) ([]float64, []float64) {
	var sumgPv []float64
	var pv []float64
	var pv2 float64
	var pv3 float64

	for s := 0; s < len(inflasi); s++ {
		pv2 = 0
		for x := 1; x <= 11; x++ {
			pv3 = 0
			interst := (1 + ((inflasi[s].Inflasi / 100) / 12))
			pv3 = fv / math.Pow(interst, float64(x))
			pv2 += pv3
			pv = append(pv, pv3)

		}
		sumgPv = append(sumgPv, pv2)
	}

	return sumgPv, pv
}

func Irr(fv float64, modal float64, dp float64) float64 {
	var irr float64 = 0
	var pv2 float64
	var npv float64

	for {
		irr += 0.01
		for i := 1; i <= 11; i++ {
			interst := (1 + ((irr / 100) / 12))
			pv3 := fv / math.Pow(interst, float64(i))
			pv2 += pv3
		}

		npv = pv2 - (modal - dp)

		if npv <= 0 {
			break
		}
		pv2 = 0
		npv = 0
	}
	return irr
}
