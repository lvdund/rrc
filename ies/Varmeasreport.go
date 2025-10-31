package ies

import "rrc/utils"

// VarMeasReport ::= SEQUENCE
type Varmeasreport struct {
	Measid                Measid
	MeasidV1250           *MeasidV1250
	Cellstriggeredlist    *Cellstriggeredlist
	CsiRsTriggeredlistR12 *CsiRsTriggeredlistR12
	PoolstriggeredlistR14 *TxResourcepoolmeaslistR14
	Numberofreportssent   utils.INTEGER
}
