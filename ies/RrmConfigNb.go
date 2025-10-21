package ies

import "rrc/utils"

// RRM-Config-NB ::= SEQUENCE
// Extensible
type RrmConfigNb struct {
	UeInactivetime *utils.ENUMERATED
}
