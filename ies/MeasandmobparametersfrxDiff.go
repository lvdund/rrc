package ies

// MeasAndMobParametersFRX-Diff ::= SEQUENCE
// Extensible
type MeasandmobparametersfrxDiff struct {
	SsSinrMeas                                  *MeasandmobparametersfrxDiffSsSinrMeas
	CsiRsrpAndrsrqMeaswithssb                   *MeasandmobparametersfrxDiffCsiRsrpAndrsrqMeaswithssb
	CsiRsrpAndrsrqMeaswithoutssb                *MeasandmobparametersfrxDiffCsiRsrpAndrsrqMeaswithoutssb
	CsiSinrMeas                                 *MeasandmobparametersfrxDiffCsiSinrMeas
	CsiRsRlm                                    *MeasandmobparametersfrxDiffCsiRsRlm
	Handoverinterf                              *MeasandmobparametersfrxDiffHandoverinterf                              `ext`
	HandoverlteEpc                              *MeasandmobparametersfrxDiffHandoverlteEpc                              `ext`
	Handoverlte5gc                              *MeasandmobparametersfrxDiffHandoverlte5gc                              `ext`
	MaxnumberresourceCsiRsRlm                   *MeasandmobparametersfrxDiffMaxnumberresourceCsiRsRlm                   `ext`
	SimultaneousrxdatassbDiffnumerology         *MeasandmobparametersfrxDiffSimultaneousrxdatassbDiffnumerology         `ext`
	NrAutonomousgapsR16                         *MeasandmobparametersfrxDiffNrAutonomousgapsR16                         `ext`
	NrAutonomousgapsEndcR16                     *MeasandmobparametersfrxDiffNrAutonomousgapsEndcR16                     `ext`
	NrAutonomousgapsNedcR16                     *MeasandmobparametersfrxDiffNrAutonomousgapsNedcR16                     `ext`
	NrAutonomousgapsNrdcR16                     *MeasandmobparametersfrxDiffNrAutonomousgapsNrdcR16                     `ext`
	Dummy                                       *MeasandmobparametersfrxDiffDummy                                       `ext`
	CliRssiMeasR16                              *MeasandmobparametersfrxDiffCliRssiMeasR16                              `ext`
	CliSrsRsrpMeasR16                           *MeasandmobparametersfrxDiffCliSrsRsrpMeasR16                           `ext`
	InterfrequencymeasNogapR16                  *MeasandmobparametersfrxDiffInterfrequencymeasNogapR16                  `ext`
	SimultaneousrxdatassbDiffnumerologyInterR16 *MeasandmobparametersfrxDiffSimultaneousrxdatassbDiffnumerologyInterR16 `ext`
	IdleinactivenrMeasreportR16                 *MeasandmobparametersfrxDiffIdleinactivenrMeasreportR16                 `ext`
	IdleinactivenrMeasbeamreportR16             *MeasandmobparametersfrxDiffIdleinactivenrMeasbeamreportR16             `ext`
	IncreasednumberofcsirspermoR16              *MeasandmobparametersfrxDiffIncreasednumberofcsirspermoR16              `ext`
}
