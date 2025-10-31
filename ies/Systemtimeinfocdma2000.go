package ies

import "rrc/utils"

// SystemTimeInfoCDMA2000 ::= SEQUENCE
type Systemtimeinfocdma2000 struct {
	CdmaEutraSynchronisation utils.BOOLEAN
	CdmaSystemtime           Systemtimeinfocdma2000CdmaSystemtime
}
