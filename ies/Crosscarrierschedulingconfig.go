package ies

// CrossCarrierSchedulingConfig ::= SEQUENCE
// Extensible
type Crosscarrierschedulingconfig struct {
	Schedulingcellinfo         CrosscarrierschedulingconfigSchedulingcellinfo
	CarrierindicatorsizeR16    *CrosscarrierschedulingconfigCarrierindicatorsizeR16    `ext`
	EnabledefaultbeamforccsR16 *CrosscarrierschedulingconfigEnabledefaultbeamforccsR16 `ext`
	CcsBlinddetectionsplitR17  *CrosscarrierschedulingconfigCcsBlinddetectionsplitR17  `ext`
}
