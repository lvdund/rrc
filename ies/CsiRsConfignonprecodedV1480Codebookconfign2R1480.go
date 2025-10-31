package ies

import "rrc/utils"

// CSI-RS-ConfigNonPrecoded-v1480-codebookConfigN2-r1480 ::= ENUMERATED
type CsiRsConfignonprecodedV1480Codebookconfign2R1480 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfignonprecodedV1480Codebookconfign2R1480EnumeratedNothing = iota
	CsiRsConfignonprecodedV1480Codebookconfign2R1480EnumeratedN5
	CsiRsConfignonprecodedV1480Codebookconfign2R1480EnumeratedN6
	CsiRsConfignonprecodedV1480Codebookconfign2R1480EnumeratedN7
)
