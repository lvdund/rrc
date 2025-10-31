package ies

// DefaultDC-Location-r17 ::= CHOICE
const (
	DefaultdcLocationR17ChoiceNothing = iota
	DefaultdcLocationR17ChoiceUl
	DefaultdcLocationR17ChoiceDl
	DefaultdcLocationR17ChoiceUlanddl
)

type DefaultdcLocationR17 struct {
	Choice  uint64
	Ul      *FrequencycomponentR17
	Dl      *FrequencycomponentR17
	Ulanddl *FrequencycomponentR17
}
