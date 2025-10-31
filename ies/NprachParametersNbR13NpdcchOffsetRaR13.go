package ies

import "rrc/utils"

// NPRACH-Parameters-NB-r13-npdcch-Offset-RA-r13 ::= ENUMERATED
type NprachParametersNbR13NpdcchOffsetRaR13 struct {
	Value utils.ENUMERATED
}

const (
	NprachParametersNbR13NpdcchOffsetRaR13EnumeratedNothing = iota
	NprachParametersNbR13NpdcchOffsetRaR13EnumeratedZero
	NprachParametersNbR13NpdcchOffsetRaR13EnumeratedOneeighth
	NprachParametersNbR13NpdcchOffsetRaR13EnumeratedOnefourth
	NprachParametersNbR13NpdcchOffsetRaR13EnumeratedThreeeighth
)
