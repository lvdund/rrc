package ies

// TrackingAreaCodeList-r10 ::= SEQUENCE OF TrackingAreaCode
// SIZE (1..8)
type TrackingareacodelistR10 struct {
	Value []Trackingareacode `lb:1,ub:8`
}
