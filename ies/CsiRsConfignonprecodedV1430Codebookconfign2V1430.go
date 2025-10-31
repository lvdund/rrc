package ies

import "rrc/utils"

// CSI-RS-ConfigNonPrecoded-v1430-codebookConfigN2-v1430 ::= ENUMERATED
type CsiRsConfignonprecodedV1430Codebookconfign2V1430 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfignonprecodedV1430Codebookconfign2V1430EnumeratedNothing = iota
	CsiRsConfignonprecodedV1430Codebookconfign2V1430EnumeratedN5
	CsiRsConfignonprecodedV1430Codebookconfign2V1430EnumeratedN6
	CsiRsConfignonprecodedV1430Codebookconfign2V1430EnumeratedN7
)
