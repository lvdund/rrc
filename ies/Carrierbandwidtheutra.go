package ies

import "rrc/utils"

// CarrierBandwidthEUTRA ::= SEQUENCE
type Carrierbandwidtheutra struct {
	DlBandwidth utils.ENUMERATED
	UlBandwidth *utils.ENUMERATED
}
