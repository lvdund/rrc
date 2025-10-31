package ies

import "rrc/utils"

// DummyH ::= SEQUENCE
type Dummyh struct {
	Burstlength                      utils.INTEGER `lb:0,ub:2`
	Maxsimultaneousresourcesetspercc utils.INTEGER `lb:0,ub:8`
	Maxconfiguredresourcesetspercc   utils.INTEGER `lb:0,ub:64`
	Maxconfiguredresourcesetsallcc   utils.INTEGER `lb:0,ub:128`
}
