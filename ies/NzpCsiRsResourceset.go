package ies

import "rrc/utils"

// NZP-CSI-RS-ResourceSet ::= SEQUENCE
// Extensible
type NzpCsiRsResourceset struct {
	NzpCsiResourcesetid            NzpCsiRsResourcesetid
	NzpCsiRsResources              []NzpCsiRsResourceid `lb:1,ub:maxNrofNZPCsiRsResourcesperset`
	Repetition                     *NzpCsiRsResourcesetRepetition
	Aperiodictriggeringoffset      *utils.INTEGER `lb:0,ub:6`
	TrsInfo                        *utils.BOOLEAN
	AperiodictriggeringoffsetR16   *utils.INTEGER            `lb:0,ub:31,ext`
	PdcInfoR17                     *utils.BOOLEAN            `ext`
	CmrgroupingandpairingR17       *CmrgroupingandpairingR17 `ext`
	AperiodictriggeringoffsetR17   *utils.INTEGER            `lb:0,ub:124,ext`
	Aperiodictriggeringoffsetl2R17 *utils.INTEGER            `lb:0,ub:31,ext`
}
