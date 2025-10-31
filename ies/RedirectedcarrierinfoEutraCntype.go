package ies

import "rrc/utils"

// RedirectedCarrierInfo-EUTRA-cnType ::= ENUMERATED
type RedirectedcarrierinfoEutraCntype struct {
	Value utils.ENUMERATED
}

const (
	RedirectedcarrierinfoEutraCntypeEnumeratedNothing = iota
	RedirectedcarrierinfoEutraCntypeEnumeratedEpc
	RedirectedcarrierinfoEutraCntypeEnumeratedFivegc
)
