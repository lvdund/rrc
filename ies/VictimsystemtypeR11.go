package ies

// VictimSystemType-r11 ::= SEQUENCE
type VictimsystemtypeR11 struct {
	GpsR11       *bool
	GlonassR11   *bool
	BdsR11       *bool
	GalileoR11   *bool
	WlanR11      *bool
	BluetoothR11 *bool
}
