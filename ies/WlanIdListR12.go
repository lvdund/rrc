package ies

import "rrc/utils"

// WLAN-Id-List-r12 ::= SEQUENCE OF WLAN-Identifiers-r12
// SIZE (1..maxWLAN-Id-r12)
type WlanIdListR12 struct {
	Value utils.Sequence[WlanIdentifiersR12]
}
