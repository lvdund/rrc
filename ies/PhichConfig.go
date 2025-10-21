package ies

import "rrc/utils"

// PHICH-Config ::= SEQUENCE
type PhichConfig struct {
	PhichDuration utils.ENUMERATED
	PhichResource utils.ENUMERATED
}
