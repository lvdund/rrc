package ies

import "rrc/utils"

// Other-Parameters-r11-ue-Rx-TxTimeDiffMeasurements-r11 ::= ENUMERATED
type OtherParametersR11UeRxTxtimediffmeasurementsR11 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersR11UeRxTxtimediffmeasurementsR11EnumeratedNothing = iota
	OtherParametersR11UeRxTxtimediffmeasurementsR11EnumeratedSupported
)
