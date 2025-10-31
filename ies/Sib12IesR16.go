package ies

import "rrc/utils"

// SIB12-IEs-r16 ::= SEQUENCE
// Extensible
type Sib12IesR16 struct {
	SlConfigcommonnrR16             SlConfigcommonnrR16
	Latenoncriticalextension        *utils.OCTETSTRING
	SlDrxConfigcommongcBcR17        *SlDrxConfiggcBcR17                  `ext`
	SlDiscconfigcommonR17           *SlDiscconfigcommonR17               `ext`
	SlL2u2nRelayR17                 *Sib12IesR16SlL2u2nRelayR17          `ext`
	SlNonrelaydiscoveryR17          *Sib12IesR16SlNonrelaydiscoveryR17   `ext`
	SlL3u2nRelaydiscoveryR17        *Sib12IesR16SlL3u2nRelaydiscoveryR17 `ext`
	SlTimersandconstantsremoteueR17 *UeTimersandconstantsremoteueR17     `ext`
}
