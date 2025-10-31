package ies

import "rrc/utils"

// BandSidelink-r16-congestionControlSidelink-r16-cbr-ReportSidelink-r16 ::= ENUMERATED
type BandsidelinkR16CongestioncontrolsidelinkR16CbrReportsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16CongestioncontrolsidelinkR16CbrReportsidelinkR16EnumeratedNothing = iota
	BandsidelinkR16CongestioncontrolsidelinkR16CbrReportsidelinkR16EnumeratedSupported
)
