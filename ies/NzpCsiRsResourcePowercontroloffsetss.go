package ies

import "rrc/utils"

// NZP-CSI-RS-Resource-powerControlOffsetSS ::= ENUMERATED
type NzpCsiRsResourcePowercontroloffsetss struct {
	Value utils.ENUMERATED
}

const (
	NzpCsiRsResourcePowercontroloffsetssEnumeratedNothing = iota
	NzpCsiRsResourcePowercontroloffsetssEnumeratedDb_3
	NzpCsiRsResourcePowercontroloffsetssEnumeratedDb0
	NzpCsiRsResourcePowercontroloffsetssEnumeratedDb3
	NzpCsiRsResourcePowercontroloffsetssEnumeratedDb6
)
