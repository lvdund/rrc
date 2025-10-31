package ies

// CG-Config-v1540-IEs-reportCGI-RequestNR-requestedCellInfo ::= SEQUENCE
type CgConfigV1540IesReportcgiRequestnrRequestedcellinfo struct {
	Ssbfrequency            ArfcnValuenr
	Cellforwhichtoreportcgi Physcellid
}
