package ies

import "rrc/utils"

// RF-Parameters-v1530-powerClass-14dBm-r15 ::= ENUMERATED
type RfParametersV1530Powerclass14dbmR15 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV1530Powerclass14dbmR15EnumeratedNothing = iota
	RfParametersV1530Powerclass14dbmR15EnumeratedSupported
)
