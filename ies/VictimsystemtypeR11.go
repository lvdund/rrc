package ies

import "rrc/utils"

// VictimSystemType-r11 ::= SEQUENCE
type VictimsystemtypeR11 struct {
	GpsR11       *utils.ENUMERATED
	GlonassR11   *utils.ENUMERATED
	BdsR11       *utils.ENUMERATED
	GalileoR11   *utils.ENUMERATED
	WlanR11      *utils.ENUMERATED
	BluetoothR11 *utils.ENUMERATED
}
