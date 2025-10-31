package ies

// CSI-SemiPersistentOnPUSCH-TriggerState ::= SEQUENCE
// Extensible
type CsiSemipersistentonpuschTriggerstate struct {
	Associatedreportconfiginfo CsiReportconfigid
	SpCsiMultiplexingmodeR17   *CsiSemipersistentonpuschTriggerstateSpCsiMultiplexingmodeR17 `ext`
}
