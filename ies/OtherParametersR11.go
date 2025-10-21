package ies

import "rrc/utils"

// Other-Parameters-r11 ::= SEQUENCE
type OtherParametersR11 struct {
	IndevicecoexindR11            *utils.ENUMERATED
	PowerprefindR11               *utils.ENUMERATED
	UeRxTxtimediffmeasurementsR11 *utils.ENUMERATED
}
