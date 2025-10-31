package ies

import "rrc/utils"

// FailedLogicalChannelInfo-r15-failureType ::= ENUMERATED
type FailedlogicalchannelinfoR15Failuretype struct {
	Value utils.ENUMERATED
}

const (
	FailedlogicalchannelinfoR15FailuretypeEnumeratedNothing = iota
	FailedlogicalchannelinfoR15FailuretypeEnumeratedDuplication
	FailedlogicalchannelinfoR15FailuretypeEnumeratedSpare3
	FailedlogicalchannelinfoR15FailuretypeEnumeratedSpare2
	FailedlogicalchannelinfoR15FailuretypeEnumeratedSpare1
)
