package ies

// RAN-AreaConfig-r15 ::= SEQUENCE
type RanAreaconfigR15 struct {
	Trackingareacode5gcR15 Trackingareacode5gcR15
	RanAreacodelistR15     *[]RanAreacodeR15 `lb:1,ub:32`
}
