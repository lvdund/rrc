package ies

import "rrc/utils"

// LWA-Parameters-v1430-wlan-ReportAnyWLAN-r14 ::= ENUMERATED
type LwaParametersV1430WlanReportanywlanR14 struct {
	Value utils.ENUMERATED
}

const (
	LwaParametersV1430WlanReportanywlanR14EnumeratedNothing = iota
	LwaParametersV1430WlanReportanywlanR14EnumeratedSupported
)
