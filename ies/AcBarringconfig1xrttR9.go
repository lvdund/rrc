package ies

import "rrc/utils"

// AC-BarringConfig1XRTT-r9 ::= SEQUENCE
type AcBarringconfig1xrttR9 struct {
	AcBarring0to9R9 utils.INTEGER `lb:0,ub:63`
	AcBarring10R9   utils.INTEGER `lb:0,ub:7`
	AcBarring11R9   utils.INTEGER `lb:0,ub:7`
	AcBarring12R9   utils.INTEGER `lb:0,ub:7`
	AcBarring13R9   utils.INTEGER `lb:0,ub:7`
	AcBarring14R9   utils.INTEGER `lb:0,ub:7`
	AcBarring15R9   utils.INTEGER `lb:0,ub:7`
	AcBarringmsgR9  utils.INTEGER `lb:0,ub:7`
	AcBarringregR9  utils.INTEGER `lb:0,ub:7`
	AcBarringemgR9  utils.INTEGER `lb:0,ub:7`
}
