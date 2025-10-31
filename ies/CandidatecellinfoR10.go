package ies

// CandidateCellInfo-r10 ::= SEQUENCE
// Extensible
type CandidatecellinfoR10 struct {
	PhyscellidR10    Physcellid
	DlCarrierfreqR10 ArfcnValueeutra
	RsrpresultR10    *RsrpRange
	RsrqresultR10    *RsrqRange
}
