package ies

import "rrc/utils"

// QuantityConfigCDMA2000-measQuantityCDMA2000 ::= ENUMERATED
type Quantityconfigcdma2000Measquantitycdma2000 struct {
	Value utils.ENUMERATED
}

const (
	Quantityconfigcdma2000Measquantitycdma2000EnumeratedNothing = iota
	Quantityconfigcdma2000Measquantitycdma2000EnumeratedPilotstrength
	Quantityconfigcdma2000Measquantitycdma2000EnumeratedPilotpnphaseandpilotstrength
)
