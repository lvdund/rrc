package ies

import "rrc/utils"

// VictimSystemType-r16 ::= SEQUENCE
// Extensible
type VictimsystemtypeR16 struct {
	GpsR16       *utils.BOOLEAN
	GlonassR16   *utils.BOOLEAN
	BdsR16       *utils.BOOLEAN
	GalileoR16   *utils.BOOLEAN
	NavicR16     *utils.BOOLEAN
	WlanR16      *utils.BOOLEAN
	BluetoothR16 *utils.BOOLEAN
}
