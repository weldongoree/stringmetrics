package sorenson

import (
	"math/big"
	"stringmetrics.io/internal/setutils"
	"stringmetrics.io/pkg/constants"
)

func Coefficient(s1 string, s2 string, alph constants.Alphabet) {
	p1 := make(map[string]int)
	p2 := make(map[string]int)

}

func PmfCoefficient(p1 setutils.Pmfset, p2 setutils.Pmfset) big.Rat {
	union := p1.Union(p2)
	numerator := 2 * union.Cardinality()

}

func RuneCoefficient(s1 string, s2 string) {

}
