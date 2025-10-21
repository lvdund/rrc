package ies

import "rrc/utils"

// CSI-RS-TriggeredList-r12 ::= SEQUENCE OF MeasCSI-RS-Id-r12
// SIZE (1..maxCSI-RS-Meas-r12)
type CsiRsTriggeredlistR12 struct {
	Value utils.Sequence[MeascsiRsIdR12]
}
