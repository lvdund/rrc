package ies

import "rrc/utils"

// SL-PQI-r16-sl-Non-StandardizedPQI-r16 ::= SEQUENCE
// Extensible
type SlPqiR16SlNonStandardizedpqiR16 struct {
	SlResourcetypeR16       *SlPqiR16SlNonStandardizedpqiR16SlResourcetypeR16
	SlPrioritylevelR16      *utils.INTEGER `lb:0,ub:8`
	SlPacketdelaybudgetR16  *utils.INTEGER `lb:0,ub:1023`
	SlPacketerrorrateR16    *utils.INTEGER `lb:0,ub:9`
	SlAveragingwindowR16    *utils.INTEGER `lb:0,ub:4095`
	SlMaxdataburstvolumeR16 *utils.INTEGER `lb:0,ub:4095`
}
