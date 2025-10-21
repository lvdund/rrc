package ies

import "rrc/utils"

// DRB-ToReleaseList ::= SEQUENCE OF DRB-Identity
// SIZE (1..maxDRB)
type DrbToreleaselist struct {
	Value utils.Sequence[DrbIdentity]
}
