package ies

// IRAT-ParametersWLAN-r13 ::= SEQUENCE
type IratParameterswlanR13 struct {
	SupportedbandlistwlanR13 *[]WlanBandindicatorR13 `lb:1,ub:maxWLANBandsR13`
}
