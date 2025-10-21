package ies

import "rrc/utils"

// SL-DiscTxConfigScheduled-r13 ::= SEQUENCE
// Extensible
type SlDisctxconfigscheduledR13 struct {
	DisctxconfigR13      *SlDiscresourcepoolR12
	DisctfIndexlistR13   *SlTfIndexpairlistR12b
	DischoppingconfigR13 *SlHoppingconfigdiscR12
}
