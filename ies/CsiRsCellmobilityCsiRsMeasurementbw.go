package ies

import "rrc/utils"

// CSI-RS-CellMobility-csi-rs-MeasurementBW ::= SEQUENCE
type CsiRsCellmobilityCsiRsMeasurementbw struct {
	Nrofprbs CsiRsCellmobilityCsiRsMeasurementbwNrofprbs
	Startprb utils.INTEGER `lb:0,ub:2169`
}
