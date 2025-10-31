package ies

import "rrc/utils"

// IAB-IP-Usage-r16 ::= ENUMERATED
type IabIpUsageR16 struct {
	Value utils.ENUMERATED
}

const (
	IabIpUsageR16EnumeratedNothing = iota
	IabIpUsageR16EnumeratedF1_C
	IabIpUsageR16EnumeratedF1_U
	IabIpUsageR16EnumeratedNon_F1
	IabIpUsageR16EnumeratedSpare
)
