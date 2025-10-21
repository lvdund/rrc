package ies

import "rrc/utils"

// GWUS-Config-NB-r16 ::= SEQUENCE
// Extensible
type GwusConfigNbR16 struct {
	GroupalternationR16        *utils.ENUMERATED
	CommonsequenceR16          *utils.ENUMERATED
	TimeparametersR16          *WusConfigNbR15
	ResourceconfigdrxR16       GwusResourceconfigNbR16
	ResourceconfigEdrxShortR16 *GwusResourceconfigNbR16
	ResourceconfigEdrxLongR16  *GwusResourceconfigNbR16
	ProbthreshlistR16          *GwusProbthreshlistNbR16
}
