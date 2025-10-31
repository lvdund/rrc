package ies

// PLMN-IdentityInfoNR-r15 ::= SEQUENCE
type PlmnIdentityinfonrR15 struct {
	PlmnIdentitylistR15 PlmnIdentitylistnrR15
	TrackingareacodeR15 *TrackingareacodenrR15
	RanAreacodeR15      *RanAreacodeR15
	CellidentityR15     CellidentitynrR15
}
