package ies

// TCI-State ::= SEQUENCE
// Extensible
type TciState struct {
	TciStateid               TciStateid
	QclType1                 QclInfo
	QclType2                 *QclInfo
	AdditionalpciR17         *AdditionalpciindexR17    `ext`
	PathlossreferencersIdR17 *PathlossreferencersIdR17 `ext`
	UlPowercontrolR17        *UplinkPowercontrolidR17  `ext`
}
