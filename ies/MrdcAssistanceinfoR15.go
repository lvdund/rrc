package ies

// MRDC-AssistanceInfo-r15 ::= SEQUENCE
// Extensible
type MrdcAssistanceinfoR15 struct {
	AffectedcarrierfreqcombinfolistmrdcR15 []AffectedcarrierfreqcombinfomrdcR15 `lb:1,ub:maxCombIDCR11`
}
