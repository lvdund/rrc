package ies

// ExplicitListOfARFCNs ::= SEQUENCE OF ARFCN-ValueGERAN
// SIZE (0..31)
type Explicitlistofarfcns struct {
	Value []ArfcnValuegeran `lb:0,ub:31`
}
