package ies

// MeasConfig ::= SEQUENCE
// Extensible
type Measconfig struct {
	Measobjecttoremovelist   *Measobjecttoremovelist
	Measobjecttoaddmodlist   *Measobjecttoaddmodlist
	Reportconfigtoremovelist *Reportconfigtoremovelist
	Reportconfigtoaddmodlist *Reportconfigtoaddmodlist
	Measidtoremovelist       *Measidtoremovelist
	Measidtoaddmodlist       *Measidtoaddmodlist
	Quantityconfig           *Quantityconfig
	Measgapconfig            *Measgapconfig
	SMeasure                 *RsrpRange
	Preregistrationinfohrpd  *Preregistrationinfohrpd
	Speedstatepars           *MeasconfigSpeedstatepars
}
