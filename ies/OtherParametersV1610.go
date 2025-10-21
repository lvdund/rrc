package ies

import "rrc/utils"

// Other-Parameters-v1610 ::= SEQUENCE
type OtherParametersV1610 struct {
	ResumewithstoredmcgScellsR16 *utils.ENUMERATED
	ResumewithmcgScellconfigR16  *utils.ENUMERATED
	ResumewithstoredscgR16       *utils.ENUMERATED
	ResumewithscgConfigR16       *utils.ENUMERATED
	McgrlfRecoveryviascgR16      *utils.ENUMERATED
	OverheatingindforscgR16      *utils.ENUMERATED
}
