package ies

import "rrc/utils"

// MeasObjectGERAN ::= SEQUENCE
// Extensible
type Measobjectgeran struct {
	Carrierfreqs            Carrierfreqsgeran
	Offsetfreq              QOffsetrangeinterrat
	NccPermitted            utils.BITSTRING `lb:8,ub:8`
	Cellforwhichtoreportcgi *Physcellidgeran
}
