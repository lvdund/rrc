package ies

import "rrc/utils"

// ExplicitListOfARFCNs ::= SEQUENCE OF ARFCN-ValueGERAN
// SIZE (0..31)
type Explicitlistofarfcns struct {
	Value utils.Sequence[ArfcnValuegeran]
}
