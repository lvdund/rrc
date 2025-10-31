package ies

import "rrc/utils"

// MeasConfig ::= SEQUENCE
// Extensible
type Measconfig struct {
	Measobjecttoremovelist       *Measobjecttoremovelist
	Measobjecttoaddmodlist       *Measobjecttoaddmodlist
	Reportconfigtoremovelist     *Reportconfigtoremovelist
	Reportconfigtoaddmodlist     *Reportconfigtoaddmodlist
	Measidtoremovelist           *Measidtoremovelist
	Measidtoaddmodlist           *Measidtoaddmodlist
	SMeasureconfig               *MeasconfigSMeasureconfig
	Quantityconfig               *Quantityconfig
	Measgapconfig                *Measgapconfig
	Measgapsharingconfig         *Measgapsharingconfig
	InterfrequencyconfigNogapR16 *utils.BOOLEAN `ext`
}
