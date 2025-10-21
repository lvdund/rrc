package ies

import "rrc/utils"

// BCCH-Config ::= SEQUENCE
type BcchConfig struct {
	Modificationperiodcoeff utils.ENUMERATED
}
