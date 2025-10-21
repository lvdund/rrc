package ies

import "rrc/utils"

// GWUS-Config-r16 ::= SEQUENCE
type GwusConfigR16 struct {
	GroupalternationR16        *utils.ENUMERATED
	CommonsequenceR16          *utils.ENUMERATED
	TimeparametersR16          *GwusTimeparametersR16
	ResourceconfigdrxR16       GwusResourceconfigR16
	ResourceconfigEdrxShortR16 *GwusResourceconfigR16
	ResourceconfigEdrxLongR16  *GwusResourceconfigR16
	ProbthreshlistR16          *GwusProbthreshlistR16
	GroupnarrowbandlistR16     *GwusGroupnarrowbandlistR16
}
