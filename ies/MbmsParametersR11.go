package ies

import "rrc/utils"

// MBMS-Parameters-r11 ::= SEQUENCE
type MbmsParametersR11 struct {
	MbmsScellR11          *utils.ENUMERATED
	MbmsNonservingcellR11 *utils.ENUMERATED
}
