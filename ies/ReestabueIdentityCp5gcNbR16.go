package ies

import "rrc/utils"

// ReestabUE-Identity-CP-5GC-NB-r16 ::= SEQUENCE
type ReestabueIdentityCp5gcNbR16 struct {
	Truncated5gSTmsiR16 utils.BITSTRING `lb:40,ub:40`
	UlNasMacR16         utils.BITSTRING `lb:16,ub:16`
	UlNasCountR16       utils.BITSTRING `lb:5,ub:5`
}
