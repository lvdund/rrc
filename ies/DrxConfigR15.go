package ies

import "rrc/utils"

// DRX-Config-r15 ::= SEQUENCE
type DrxConfigR15 struct {
	DrxRetransmissiontimershortttiR15   *utils.ENUMERATED
	DrxUlRetransmissiontimershortttiR15 *utils.ENUMERATED
}
