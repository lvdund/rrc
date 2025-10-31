package ies

// PDCCH-MonitoringOccasions-r16 ::= SEQUENCE
type PdcchMonitoringoccasionsR16 struct {
	Period7span3R16 *PdcchMonitoringoccasionsR16Period7span3R16
	Period4span3R16 *PdcchMonitoringoccasionsR16Period4span3R16
	Period2span2R16 *PdcchMonitoringoccasionsR16Period2span2R16
}
