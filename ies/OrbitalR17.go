package ies

import "rrc/utils"

// Orbital-r17 ::= SEQUENCE
type OrbitalR17 struct {
	SemimajoraxisR17 utils.INTEGER `lb:0,ub:8589934591`
	EccentricityR17  utils.INTEGER `lb:0,ub:1048575`
	PeriapsisR17     utils.INTEGER `lb:0,ub:268435455`
	LongitudeR17     utils.INTEGER `lb:0,ub:268435455`
	InclinationR17   utils.INTEGER `lb:0,ub:67108863`
	MeananomalyR17   utils.INTEGER `lb:0,ub:268435455`
}
