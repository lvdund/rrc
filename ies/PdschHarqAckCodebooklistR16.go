package ies

import "rrc/utils"

// PDSCH-HARQ-ACK-CodebookList-r16 ::= SEQUENCE OF utils.ENUMERATED // SIZE (1..2)
type PdschHarqAckCodebooklistR16 struct {
	Value []utils.ENUMERATED `lb:1,ub:2`
}
