package ies

import "rrc/utils"

// AC-BarringConfig1XRTT-r9 ::= SEQUENCE
type AcBarringconfig1xrttR9 struct {
	AcBarring0to9R9 utils.INTEGER
	AcBarring10R9   utils.INTEGER
	AcBarring11R9   utils.INTEGER
	AcBarring12R9   utils.INTEGER
	AcBarring13R9   utils.INTEGER
	AcBarring14R9   utils.INTEGER
	AcBarring15R9   utils.INTEGER
	AcBarringmsgR9  utils.INTEGER
	AcBarringregR9  utils.INTEGER
	AcBarringemgR9  utils.INTEGER
}
