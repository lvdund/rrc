package ies

import "rrc/utils"

// MasterKeyUpdate ::= SEQUENCE
// Extensible
type Masterkeyupdate struct {
	Keysetchangeindicator utils.BOOLEAN
	Nexthopchainingcount  Nexthopchainingcount
	NasContainer          *utils.OCTETSTRING
}
