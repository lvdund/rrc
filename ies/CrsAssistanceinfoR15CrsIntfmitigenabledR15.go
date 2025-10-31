package ies

import "rrc/utils"

// CRS-AssistanceInfo-r15-crs-IntfMitigEnabled-r15 ::= ENUMERATED
type CrsAssistanceinfoR15CrsIntfmitigenabledR15 struct {
	Value utils.ENUMERATED
}

const (
	CrsAssistanceinfoR15CrsIntfmitigenabledR15EnumeratedNothing = iota
	CrsAssistanceinfoR15CrsIntfmitigenabledR15EnumeratedEnabled
)
