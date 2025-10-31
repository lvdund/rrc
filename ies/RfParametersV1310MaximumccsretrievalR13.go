package ies

import "rrc/utils"

// RF-Parameters-v1310-maximumCCsRetrieval-r13 ::= ENUMERATED
type RfParametersV1310MaximumccsretrievalR13 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV1310MaximumccsretrievalR13EnumeratedNothing = iota
	RfParametersV1310MaximumccsretrievalR13EnumeratedSupported
)
