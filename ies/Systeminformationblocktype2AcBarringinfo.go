package ies

import "rrc/utils"

// SystemInformationBlockType2-ac-BarringInfo ::= SEQUENCE
type Systeminformationblocktype2AcBarringinfo struct {
	AcBarringforemergency    utils.BOOLEAN
	AcBarringformoSignalling *AcBarringconfig
	AcBarringformoData       *AcBarringconfig
}
