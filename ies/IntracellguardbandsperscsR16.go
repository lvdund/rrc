package ies

// IntraCellGuardBandsPerSCS-r16 ::= SEQUENCE
type IntracellguardbandsperscsR16 struct {
	GuardbandscsR16        Subcarrierspacing
	IntracellguardbandsR16 []GuardbandR16 `lb:1,ub:4`
}
