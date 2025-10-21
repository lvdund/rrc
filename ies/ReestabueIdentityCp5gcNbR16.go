package ies

import "rrc/utils"

// ReestabUE-Identity-CP-5GC-NB-r16 ::= SEQUENCE
type ReestabueIdentityCp5gcNbR16 struct {
	Truncated5gSTmsiR16 utils.BITSTRING
	UlNasMacR16         utils.BITSTRING
	UlNasCountR16       utils.BITSTRING
}
