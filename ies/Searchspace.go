package ies

import "rrc/utils"

// SearchSpace ::= SEQUENCE
// Extensible
type Searchspace struct {
	Searchspaceid                      Searchspaceid
	Controlresourcesetid               *Controlresourcesetid
	Monitoringslotperiodicityandoffset *SearchspaceMonitoringslotperiodicityandoffset
	Duration                           *utils.INTEGER   `lb:0,ub:2559`
	Monitoringsymbolswithinslot        *utils.BITSTRING `lb:14,ub:14`
	Nrofcandidates                     *SearchspaceNrofcandidates
	Searchspacetype                    *SearchspaceSearchspacetype `ext`
}
