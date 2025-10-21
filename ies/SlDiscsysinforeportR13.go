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
	DisccellselectioninfoR13 *interface{}
	CellreselectioninfoR13   *interface{}
	TddConfigR13             *TddConfig
	FreqinfoR13              *interface{}
	PMaxR13                  *PMax
	ReferencesignalpowerR13  *utils.INTEGER
}
