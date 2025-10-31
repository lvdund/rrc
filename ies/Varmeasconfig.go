package ies

// VarMeasConfig ::= SEQUENCE
type Varmeasconfig struct {
	Measidlist       *Measidtoaddmodlist
	Measobjectlist   *Measobjecttoaddmodlist
	Reportconfiglist *Reportconfigtoaddmodlist
	Quantityconfig   *Quantityconfig
	SMeasureconfig   *VarmeasconfigSMeasureconfig
}
