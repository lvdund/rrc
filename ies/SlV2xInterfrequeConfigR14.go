package ies

import "rrc/utils"

// SL-V2X-InterFreqUE-Config-r14 ::= SEQUENCE
// Extensible
type SlV2xInterfrequeConfigR14 struct {
	PhyscellidlistR14             *PhyscellidlistR13
	TypetxsyncR14                 *SlTypetxsyncR14
	V2xSyncconfigR14              *SlSyncconfiglistnfreqv2xR14
	V2xCommrxpoolR14              *SlCommrxpoollistv2xR14
	V2xCommtxpoolnormalR14        *SlCommtxpoollistv2xR14
	P2xCommtxpoolnormalR14        *SlCommtxpoollistv2xR14
	V2xCommtxpoolexceptionalR14   *SlCommresourcepoolv2xR14
	V2xResourceselectionconfigR14 *SlCommtxpoolsensingconfigR14
	ZoneconfigR14                 *SlZoneconfigR14
	OffsetdfnR14                  *utils.INTEGER `lb:0,ub:1000`
}
