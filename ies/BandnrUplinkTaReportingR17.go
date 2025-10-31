package ies

import "rrc/utils"

// BandNR-uplink-TA-Reporting-r17 ::= ENUMERATED
type BandnrUplinkTaReportingR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrUplinkTaReportingR17EnumeratedNothing = iota
	BandnrUplinkTaReportingR17EnumeratedSupported
)
