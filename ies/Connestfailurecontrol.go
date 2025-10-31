package ies

import "rrc/utils"

// ConnEstFailureControl ::= SEQUENCE
type Connestfailurecontrol struct {
	Connestfailcount          ConnestfailurecontrolConnestfailcount
	Connestfailoffsetvalidity ConnestfailurecontrolConnestfailoffsetvalidity
	Connestfailoffset         *utils.INTEGER `lb:0,ub:15`
}
