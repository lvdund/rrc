package ies

import "rrc/utils"

// TPC-PDCCH-Config-setup ::= SEQUENCE
type TpcPdcchConfigSetup struct {
	TpcRnti  utils.BITSTRING `lb:16,ub:16`
	TpcIndex TpcIndex
}
