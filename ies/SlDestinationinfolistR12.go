package ies

import "rrc/utils"

// SL-DestinationInfoList-r12 ::= SEQUENCE OF SL-DestinationIdentity-r12
// SIZE (1..maxSL-Dest-r12)
type SlDestinationinfolistR12 struct {
	Value []SlDestinationidentityR12
}
