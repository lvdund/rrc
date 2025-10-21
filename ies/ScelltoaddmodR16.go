package ies

import "rrc/utils"

// SCellToAddMod-r16 ::= SEQUENCE
// Extensible
type ScelltoaddmodR16 struct {
	ScellindexR16                        ScellindexR13
	CellidentificationR16                *interface{}
	RadioresourceconfigcommonscellR16    *RadioresourceconfigcommonscellR10
	RadioresourceconfigdedicatedscellR16 *RadioresourceconfigdedicatedscellR10
	AntennainfodedicatedscellR16         *AntennainfodedicatedV10i0
	SrsSwitchfromservcellindexR16        *utils.INTEGER
	ScellstateR16                        *utils.ENUMERATED
}
