package ies

// RadioLinkMonitoringConfig ::= SEQUENCE
// Extensible
type Radiolinkmonitoringconfig struct {
	Failuredetectionresourcestoaddmodlist  *[]Radiolinkmonitoringrs   `lb:1,ub:maxNrofFailureDetectionResources`
	Failuredetectionresourcestoreleaselist *[]RadiolinkmonitoringrsId `lb:1,ub:maxNrofFailureDetectionResources`
	Beamfailureinstancemaxcount            *RadiolinkmonitoringconfigBeamfailureinstancemaxcount
	Beamfailuredetectiontimer              *RadiolinkmonitoringconfigBeamfailuredetectiontimer
	BeamfailureR17                         *BeamfailuredetectionR17 `ext`
}
