package ies

import "rrc/utils"

// BWP-DownlinkDedicatedSDT-r17 ::= SEQUENCE
// Extensible
type BwpDownlinkdedicatedsdtR17 struct {
	PdcchConfigR17 *utils.Setuprelease[PdcchConfig]
	PdschConfigR17 *utils.Setuprelease[PdschConfig]
}
