package ies

// TrackingAreaCodeList-r16 ::= SEQUENCE OF TrackingAreaCode
// SIZE (1..8)
type TrackingareacodelistR16 struct {
	Value []Trackingareacode `lb:1,ub:8`
}
