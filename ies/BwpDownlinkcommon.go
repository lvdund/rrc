package ies

import "rrc/utils"

// BWP-DownlinkCommon ::= SEQUENCE
// Extensible
type BwpDownlinkcommon struct {
	Genericparameters Bwp
	PdcchConfigcommon *utils.Setuprelease[PdcchConfigcommon]
	PdschConfigcommon *utils.Setuprelease[PdschConfigcommon]
}
