package ies

// MeasAndMobParametersXDD-Diff ::= SEQUENCE
// Extensible
type MeasandmobparametersxddDiff struct {
	IntraandinterfMeasandreport *MeasandmobparametersxddDiffIntraandinterfMeasandreport
	EventaMeasandreport         *MeasandmobparametersxddDiffEventaMeasandreport
	Handoverinterf              *MeasandmobparametersxddDiffHandoverinterf     `ext`
	HandoverlteEpc              *MeasandmobparametersxddDiffHandoverlteEpc     `ext`
	Handoverlte5gc              *MeasandmobparametersxddDiffHandoverlte5gc     `ext`
	SftdMeasnrNeigh             *MeasandmobparametersxddDiffSftdMeasnrNeigh    `ext`
	SftdMeasnrNeighDrx          *MeasandmobparametersxddDiffSftdMeasnrNeighDrx `ext`
	Dummy                       *MeasandmobparametersxddDiffDummy              `ext`
}
