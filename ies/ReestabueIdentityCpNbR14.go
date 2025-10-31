package ies

import "rrc/utils"

// ReestabUE-Identity-CP-NB-r14 ::= SEQUENCE
type ReestabueIdentityCpNbR14 struct {
	STmsiR14      STmsi
	UlNasMacR14   utils.BITSTRING `lb:16,ub:16`
	UlNasCountR14 utils.BITSTRING `lb:5,ub:5`
}
