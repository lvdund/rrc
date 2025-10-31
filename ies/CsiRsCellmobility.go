package ies

// CSI-RS-CellMobility ::= SEQUENCE
type CsiRsCellmobility struct {
	Cellid                    Physcellid
	CsiRsMeasurementbw        CsiRsCellmobilityCsiRsMeasurementbw
	Density                   *CsiRsCellmobilityDensity
	CsiRsResourcelistMobility []CsiRsResourceMobility `lb:1,ub:maxNrofCSIRsResourcesrrm`
}
