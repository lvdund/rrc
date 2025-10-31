package ies

import "rrc/utils"

// CodebookConfig-r17-codebookType-type2-typeII-PortSelection-r17-valueOfN-r17 ::= ENUMERATED
type CodebookconfigR17CodebooktypeType2TypeiiPortselectionR17ValueofnR17 struct {
	Value utils.ENUMERATED
}

const (
	CodebookconfigR17CodebooktypeType2TypeiiPortselectionR17ValueofnR17EnumeratedNothing = iota
	CodebookconfigR17CodebooktypeType2TypeiiPortselectionR17ValueofnR17EnumeratedN2
	CodebookconfigR17CodebooktypeType2TypeiiPortselectionR17ValueofnR17EnumeratedN4
)
