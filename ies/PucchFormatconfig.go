package ies

import "rrc/utils"

// PUCCH-FormatConfig ::= SEQUENCE
type PucchFormatconfig struct {
	Interslotfrequencyhopping *PucchFormatconfigInterslotfrequencyhopping
	Additionaldmrs            *utils.BOOLEAN
	Maxcoderate               *PucchMaxcoderate
	Nrofslots                 *PucchFormatconfigNrofslots
	Pi2bpsk                   *PucchFormatconfigPi2bpsk
	SimultaneousharqAckCsi    *utils.BOOLEAN
}
