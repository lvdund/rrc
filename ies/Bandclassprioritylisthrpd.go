package ies

// BandClassPriorityListHRPD ::= SEQUENCE OF BandClassPriorityHRPD
// SIZE (1..maxCDMA-BandClass)
type Bandclassprioritylisthrpd struct {
	Value []Bandclasspriorityhrpd `lb:1,ub:maxCDMABandclass`
}
