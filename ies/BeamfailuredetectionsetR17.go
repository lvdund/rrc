package ies

// BeamFailureDetectionSet-r17 ::= SEQUENCE
// Extensible
type BeamfailuredetectionsetR17 struct {
	BfdresourcestoaddmodlistR17    *[]BeamlinkmonitoringrsR17   `lb:1,ub:maxNrofBFDResourcePerSetR17`
	BfdresourcestoreleaselistR17   *[]BeamlinkmonitoringrsIdR17 `lb:1,ub:maxNrofBFDResourcePerSetR17`
	BeamfailureinstancemaxcountR17 *BeamfailuredetectionsetR17BeamfailureinstancemaxcountR17
	BeamfailuredetectiontimerR17   *BeamfailuredetectionsetR17BeamfailuredetectiontimerR17
}
