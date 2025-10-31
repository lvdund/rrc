package ies

// SlotOrSubslotPDSCH-Config-r15-setup ::= SEQUENCE
// Extensible
type SlotorsubslotpdschConfigR15Setup struct {
	AltcqiTablesttiR15        *SlotorsubslotpdschConfigR15SetupAltcqiTablesttiR15
	AltcqiTable1024qamSttiR15 *SlotorsubslotpdschConfigR15SetupAltcqiTable1024qamSttiR15
	ResourceallocationR15     *SlotorsubslotpdschConfigR15SetupResourceallocationR15
	TbsindexaltSttiR15        *SlotorsubslotpdschConfigR15SetupTbsindexaltSttiR15
	Tbsindexalt2SttiR15       *SlotorsubslotpdschConfigR15SetupTbsindexalt2SttiR15
	Tbsindexalt3SttiR15       *SlotorsubslotpdschConfigR15SetupTbsindexalt3SttiR15
}
