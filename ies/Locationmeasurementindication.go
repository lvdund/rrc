package ies

import "rrc/utils"

// LocationMeasurementIndication-IEs ::= SEQUENCE
type Locationmeasurementindication struct {
	Measurementindication    utils.Setuprelease[Locationmeasurementinfo]
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *LocationmeasurementindicationIesNoncriticalextension
}
