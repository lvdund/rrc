package ies

import "rrc/utils"

// MAC-ParametersXDD-Diff-secondaryDRX-Group-r16 ::= ENUMERATED
type MacParametersxddDiffSecondarydrxGroupR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersxddDiffSecondarydrxGroupR16EnumeratedNothing = iota
	MacParametersxddDiffSecondarydrxGroupR16EnumeratedSupported
)
