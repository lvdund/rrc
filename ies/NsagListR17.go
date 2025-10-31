package ies

// NSAG-List-r17 ::= SEQUENCE OF NSAG-ID-r17
// SIZE (1.. maxSliceInfo-r17)
type NsagListR17 struct {
	Value []NsagIdR17 `lb:1,ub:maxSliceInfoR17`
}
