package ies

import "rrc/utils"

// MAC-MainConfig ::= SEQUENCE
// Extensible
type MacMainconfig struct {
	UlSchConfig                 *interface{}
	DrxConfig                   *DrxConfig
	Timealignmenttimerdedicated Timealignmenttimer
	PhrConfig                   *interface{}
}
