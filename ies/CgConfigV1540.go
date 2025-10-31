package ies

// CG-Config-v1540-IEs ::= SEQUENCE
type CgConfigV1540 struct {
	Pscellfrequency      *ArfcnValuenr
	ReportcgiRequestnr   *CgConfigV1540IesReportcgiRequestnr
	PhInfoscg            *PhTypelistscg
	Noncriticalextension *CgConfigV1560
}
