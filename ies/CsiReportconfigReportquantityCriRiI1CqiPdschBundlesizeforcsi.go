package ies

import "rrc/utils"

// CSI-ReportConfig-reportQuantity-cri-RI-i1-CQI-pdsch-BundleSizeForCSI ::= ENUMERATED
type CsiReportconfigReportquantityCriRiI1CqiPdschBundlesizeforcsi struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigReportquantityCriRiI1CqiPdschBundlesizeforcsiEnumeratedNothing = iota
	CsiReportconfigReportquantityCriRiI1CqiPdschBundlesizeforcsiEnumeratedN2
	CsiReportconfigReportquantityCriRiI1CqiPdschBundlesizeforcsiEnumeratedN4
)
