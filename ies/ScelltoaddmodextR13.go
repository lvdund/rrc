package ies

import "rrc/utils"

// SCellToAddModExt-r13 ::= SEQUENCE
type ScelltoaddmodextR13 struct {
	ScellindexR13                        ScellindexR13
	CellidentificationR13                *interface{}
	RadioresourceconfigcommonscellR13    *RadioresourceconfigcommonscellR10
	RadioresourceconfigdedicatedscellR13 *RadioresourceconfigdedicatedscellR10
	AntennainfodedicatedscellR13         *AntennainfodedicatedV10i0
}
