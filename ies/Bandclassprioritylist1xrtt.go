package ies

// BandClassPriorityList1XRTT ::= SEQUENCE OF BandClassPriority1XRTT
// SIZE (1..maxCDMA-BandClass)
type Bandclassprioritylist1xrtt struct {
	Value []Bandclasspriority1xrtt `lb:1,ub:maxCDMABandclass`
}
