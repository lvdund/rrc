package ies

// CSI-RS-ResourceConfigMobility ::= SEQUENCE
// Extensible
type CsiRsResourceconfigmobility struct {
	Subcarrierspacing     Subcarrierspacing
	CsiRsCelllistMobility []CsiRsCellmobility `lb:1,ub:maxNrofCSIRsCellsrrm`
	Refservcellindex      *Servcellindex      `ext`
}
