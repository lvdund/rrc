package ies

import "rrc/utils"

// CGI-InfoNR ::= SEQUENCE
// Extensible
type CgiInfonr struct {
	PlmnIdentityinfolist       *PlmnIdentityinfolist
	Frequencybandlist          *Multifrequencybandlistnr
	Nosib1                     *CgiInfonrNosib1
	NpnIdentityinfolistR16     *NpnIdentityinfolistR16 `ext`
	CellreservedforotheruseR16 *utils.BOOLEAN          `ext`
}
