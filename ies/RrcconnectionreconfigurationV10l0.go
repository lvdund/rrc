package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v10l0-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV10l0 struct {
	MobilitycontrolinfoV10l0 *MobilitycontrolinfoV10l0
	ScelltoaddmodlistV10l0   *ScelltoaddmodlistV10l0
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreconfigurationV12f0
}
