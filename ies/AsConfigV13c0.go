package ies

import "rrc/utils"

// AS-Config-v13c0 ::= SEQUENCE
type AsConfigV13c0 struct {
	RadioresourceconfigdedicatedV13c01 *RadioresourceconfigdedicatedV1370
	RadioresourceconfigdedicatedV13c02 *RadioresourceconfigdedicatedV13c0
	ScelltoaddmodlistV13c0             *ScelltoaddmodlistV13c0
	ScelltoaddmodlistextV13c0          *ScelltoaddmodlistextV13c0
}
