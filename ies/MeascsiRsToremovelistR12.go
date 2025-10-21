package ies

import "rrc/utils"

// MeasCSI-RS-ToRemoveList-r12 ::= SEQUENCE OF MeasCSI-RS-Id-r12
// SIZE (1..maxCSI-RS-Meas-r12)
type MeascsiRsToremovelistR12 struct {
	Value utils.Sequence[MeascsiRsIdR12]
}
