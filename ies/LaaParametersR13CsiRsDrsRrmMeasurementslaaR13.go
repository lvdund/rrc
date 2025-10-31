package ies

import "rrc/utils"

// LAA-Parameters-r13-csi-RS-DRS-RRM-MeasurementsLAA-r13 ::= ENUMERATED
type LaaParametersR13CsiRsDrsRrmMeasurementslaaR13 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersR13CsiRsDrsRrmMeasurementslaaR13EnumeratedNothing = iota
	LaaParametersR13CsiRsDrsRrmMeasurementslaaR13EnumeratedSupported
)
