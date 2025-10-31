package ies

import "rrc/utils"

// CSI-ReportConfig-timeRestrictionForChannelMeasurements ::= ENUMERATED
type CsiReportconfigTimerestrictionforchannelmeasurements struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigTimerestrictionforchannelmeasurementsEnumeratedNothing = iota
	CsiReportconfigTimerestrictionforchannelmeasurementsEnumeratedConfigured
	CsiReportconfigTimerestrictionforchannelmeasurementsEnumeratedNotconfigured
)
