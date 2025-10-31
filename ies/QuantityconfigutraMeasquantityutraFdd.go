package ies

import "rrc/utils"

// QuantityConfigUTRA-measQuantityUTRA-FDD ::= ENUMERATED
type QuantityconfigutraMeasquantityutraFdd struct {
	Value utils.ENUMERATED
}

const (
	QuantityconfigutraMeasquantityutraFddEnumeratedNothing = iota
	QuantityconfigutraMeasquantityutraFddEnumeratedCpich_Rscp
	QuantityconfigutraMeasquantityutraFddEnumeratedCpich_Ecn0
)
