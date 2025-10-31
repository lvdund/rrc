package ies

// GWUS-Config-NB-r16 ::= SEQUENCE
// Extensible
type GwusConfigNbR16 struct {
	GroupalternationR16        *bool
	CommonsequenceR16          *GwusConfigNbR16CommonsequenceR16
	TimeparametersR16          *WusConfigNbR15
	ResourceconfigdrxR16       GwusResourceconfigNbR16
	ResourceconfigEdrxShortR16 *GwusResourceconfigNbR16
	ResourceconfigEdrxLongR16  *GwusResourceconfigNbR16
	ProbthreshlistR16          *GwusProbthreshlistNbR16
}
