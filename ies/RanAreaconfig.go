package ies

// RAN-AreaConfig ::= SEQUENCE
type RanAreaconfig struct {
	Trackingareacode Trackingareacode
	RanAreacodelist  *[]RanAreacode `lb:1,ub:32`
}
