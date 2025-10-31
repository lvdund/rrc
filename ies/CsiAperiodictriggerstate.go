package ies

// CSI-AperiodicTriggerState ::= SEQUENCE
// Extensible
type CsiAperiodictriggerstate struct {
	Associatedreportconfiginfolist []CsiAssociatedreportconfiginfo                   `lb:1,ub:maxNrofReportConfigPerAperiodicTrigger`
	ApCsiMultiplexingmodeR17       *CsiAperiodictriggerstateApCsiMultiplexingmodeR17 `ext`
}
