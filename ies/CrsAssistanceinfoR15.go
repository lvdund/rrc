package ies

// CRS-AssistanceInfo-r15 ::= SEQUENCE
type CrsAssistanceinfoR15 struct {
	PhyscellidR15          Physcellid
	CrsIntfmitigenabledR15 *CrsAssistanceinfoR15CrsIntfmitigenabledR15
}
