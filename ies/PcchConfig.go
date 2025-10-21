package ies

import "rrc/utils"

// PCCH-Config ::= SEQUENCE
type PcchConfig struct {
	Defaultpagingcycle utils.ENUMERATED
	Nb                 utils.ENUMERATED
}
