package ies

import "rrc/utils"

// VarMeasConfig ::= SEQUENCE
type Varmeasconfig struct {
	Measidlist            *Measidtoaddmodlist
	MeasidlistextR12      *MeasidtoaddmodlistextR12
	MeasidlistV1310       *MeasidtoaddmodlistV1310
	MeasidlistextV1310    *MeasidtoaddmodlistextV1310
	Measobjectlist        *Measobjecttoaddmodlist
	MeasobjectlistextR13  *MeasobjecttoaddmodlistextR13
	MeasobjectlistV9i0    *MeasobjecttoaddmodlistV9e0
	Reportconfiglist      *Reportconfigtoaddmodlist
	Quantityconfig        *Quantityconfig
	MeasscalefactorR12    *MeasscalefactorR12
	SMeasure              *utils.INTEGER `lb:0,ub:-44`
	Speedstatepars        *VarmeasconfigSpeedstatepars
	AllowinterruptionsR11 *utils.BOOLEAN
}
