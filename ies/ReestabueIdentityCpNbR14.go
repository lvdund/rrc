package ies

import "rrc/utils"

// ReestabUE-Identity-CP-NB-r14 ::= SEQUENCE
type ReestabueIdentityCpNbR14 struct {
	STmsiR14      STmsi
	UlNasMacR14   utils.BITSTRING
	UlNasCountR14 utils.BITSTRING
}
