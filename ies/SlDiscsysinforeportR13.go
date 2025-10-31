package ies

import "rrc/utils"

// SL-DiscSysInfoReport-r13 ::= SEQUENCE
// Extensible
type SlDiscsysinforeportR13 struct {
	PlmnIdentitylistR13      *PlmnIdentitylist
	Cellidentity13           *Cellidentity
	Carrierfreqinfo13        *ArfcnValueeutraR9
	DiscrxresourcesR13       *SlDiscrxpoollistR12
	DisctxpoolcommonR13      *SlDisctxpoollistR12
	DisctxpowerinfoR13       *SlDisctxpowerinfolistR12
	DiscsyncconfigR13        *SlSyncconfignfreqR13
	DisccellselectioninfoR13 *SlDiscsysinforeportR13DisccellselectioninfoR13
	CellreselectioninfoR13   *SlDiscsysinforeportR13CellreselectioninfoR13
	TddConfigR13             *TddConfig
	FreqinfoR13              *SlDiscsysinforeportR13FreqinfoR13
	PMaxR13                  *PMax
	ReferencesignalpowerR13  *utils.INTEGER `lb:0,ub:50`
}
