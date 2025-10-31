package ies

import "rrc/utils"

// SRS-TPC-PDCCH-Config-r14-setup ::= SEQUENCE
type SrsTpcPdcchConfigR14Setup struct {
	SrsTpcRntiR14            utils.BITSTRING     `lb:16,ub:16`
	Startingbitofformat3bR14 utils.INTEGER       `lb:0,ub:31`
	Fieldtypeformat3bR14     utils.INTEGER       `lb:0,ub:4`
	SrsCcSetindexlistR14     *[]SrsCcSetindexR14 `lb:1,ub:4`
}
