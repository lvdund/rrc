package ies

import "rrc/utils"

// VarMeasReport ::= SEQUENCE
type Varmeasreport struct {
	Measid                      Measid
	Cellstriggeredlist          *Cellstriggeredlist
	Numberofreportssent         utils.INTEGER
	CliTriggeredlistR16         *CliTriggeredlistR16
	TxPoolmeastoaddmodlistnrR16 *TxPoolmeaslistR16
	RelaystriggeredlistR17      *RelaystriggeredlistR17
}
