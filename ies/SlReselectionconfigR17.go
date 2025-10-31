package ies

// SL-ReselectionConfig-r17 ::= SEQUENCE
type SlReselectionconfigR17 struct {
	SlRsrpThreshR17            *SlRsrpRangeR16
	SlFiltercoefficientrsrpR17 *Filtercoefficient
	SlHystminR17               *Hysteresis
}
