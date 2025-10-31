package ies

import "rrc/utils"

// MeasAndMobParametersCommon ::= SEQUENCE
// Extensible
type Measandmobparameterscommon struct {
	Supportedgappattern                        *utils.BITSTRING `lb:22,ub:22`
	SsbRlm                                     *MeasandmobparameterscommonSsbRlm
	SsbAndcsiRsRlm                             *MeasandmobparameterscommonSsbAndcsiRsRlm
	EventbMeasandreport                        *MeasandmobparameterscommonEventbMeasandreport                        `ext`
	HandoverfddTdd                             *MeasandmobparameterscommonHandoverfddTdd                             `ext`
	EutraCgiReporting                          *MeasandmobparameterscommonEutraCgiReporting                          `ext`
	NrCgiReporting                             *MeasandmobparameterscommonNrCgiReporting                             `ext`
	Independentgapconfig                       *MeasandmobparameterscommonIndependentgapconfig                       `ext`
	PeriodiceutraMeasandreport                 *MeasandmobparameterscommonPeriodiceutraMeasandreport                 `ext`
	Handoverfr1Fr2                             *MeasandmobparameterscommonHandoverfr1Fr2                             `ext`
	MaxnumbercsiRsRrmRsSinr                    *MeasandmobparameterscommonMaxnumbercsiRsRrmRsSinr                    `ext`
	NrCgiReportingEndc                         *MeasandmobparameterscommonNrCgiReportingEndc                         `ext`
	EutraCgiReportingNedc                      *MeasandmobparameterscommonEutraCgiReportingNedc                      `ext`
	EutraCgiReportingNrdc                      *MeasandmobparameterscommonEutraCgiReportingNrdc                      `ext`
	NrCgiReportingNedc                         *MeasandmobparameterscommonNrCgiReportingNedc                         `ext`
	NrCgiReportingNrdc                         *MeasandmobparameterscommonNrCgiReportingNrdc                         `ext`
	ReportaddneighmeasforperiodicR16           *MeasandmobparameterscommonReportaddneighmeasforperiodicR16           `ext`
	CondhandoverparameterscommonR16            *MeasandmobparameterscommonCondhandoverparameterscommonR16            `ext`
	NrNeedforgapReportingR16                   *MeasandmobparameterscommonNrNeedforgapReportingR16                   `ext`
	SupportedgappatternNronlyR16               *utils.BITSTRING                                                      `lb:10,ub:10,ext`
	SupportedgappatternNronlyNedcR16           *MeasandmobparameterscommonSupportedgappatternNronlyNedcR16           `ext`
	MaxnumbercliRssiR16                        *MeasandmobparameterscommonMaxnumbercliRssiR16                        `ext`
	MaxnumbercliSrsRsrpR16                     *MeasandmobparameterscommonMaxnumbercliSrsRsrpR16                     `ext`
	MaxnumberperslotcliSrsRsrpR16              *MeasandmobparameterscommonMaxnumberperslotcliSrsRsrpR16              `ext`
	MfbiIabR16                                 *MeasandmobparameterscommonMfbiIabR16                                 `ext`
	Dummy                                      *MeasandmobparameterscommonDummy                                      `ext`
	NrCgiReportingNpnR16                       *MeasandmobparameterscommonNrCgiReportingNpnR16                       `ext`
	IdleinactiveeutraMeasreportR16             *MeasandmobparameterscommonIdleinactiveeutraMeasreportR16             `ext`
	IdleinactiveValidityareaR16                *MeasandmobparameterscommonIdleinactiveValidityareaR16                `ext`
	EutraAutonomousgapsR16                     *MeasandmobparameterscommonEutraAutonomousgapsR16                     `ext`
	EutraAutonomousgapsNedcR16                 *MeasandmobparameterscommonEutraAutonomousgapsNedcR16                 `ext`
	EutraAutonomousgapsNrdcR16                 *MeasandmobparameterscommonEutraAutonomousgapsNrdcR16                 `ext`
	Pcellt312R16                               *MeasandmobparameterscommonPcellt312R16                               `ext`
	SupportedgappatternR16                     *utils.BITSTRING                                                      `lb:2,ub:2,ext`
	ConcurrentmeasgapR17                       *MeasandmobparameterscommonConcurrentmeasgapR17                       `ext`
	NrNeedforgapncsgReportingR17               *MeasandmobparameterscommonNrNeedforgapncsgReportingR17               `ext`
	EutraNeedforgapncsgReportingR17            *MeasandmobparameterscommonEutraNeedforgapncsgReportingR17            `ext`
	NcsgMeasgapperfrR17                        *MeasandmobparameterscommonNcsgMeasgapperfrR17                        `ext`
	NcsgMeasgappatternsR17                     *utils.BITSTRING                                                      `lb:24,ub:24,ext`
	NcsgMeasgapnrPatternsR17                   *utils.BITSTRING                                                      `lb:24,ub:24,ext`
	PreconfiguredueAutonomousmeasgapR17        *MeasandmobparameterscommonPreconfiguredueAutonomousmeasgapR17        `ext`
	PreconfigurednwControlledmeasgapR17        *MeasandmobparameterscommonPreconfigurednwControlledmeasgapR17        `ext`
	Handoverfr1Fr22R17                         *MeasandmobparameterscommonHandoverfr1Fr22R17                         `ext`
	Handoverfr21Fr22R17                        *MeasandmobparameterscommonHandoverfr21Fr22R17                        `ext`
	IndependentgapconfigprsR17                 *MeasandmobparameterscommonIndependentgapconfigprsR17                 `ext`
	RrmRelaxationrrcConnectedredcapR17         *MeasandmobparameterscommonRrmRelaxationrrcConnectedredcapR17         `ext`
	ParallelmeasurementgapR17                  *MeasandmobparameterscommonParallelmeasurementgapR17                  `ext`
	CondhandoverwithscgNrdcR17                 *MeasandmobparameterscommonCondhandoverwithscgNrdcR17                 `ext`
	GnbIdLengthreportingR17                    *MeasandmobparameterscommonGnbIdLengthreportingR17                    `ext`
	GnbIdLengthreportingEndcR17                *MeasandmobparameterscommonGnbIdLengthreportingEndcR17                `ext`
	GnbIdLengthreportingNedcR17                *MeasandmobparameterscommonGnbIdLengthreportingNedcR17                `ext`
	GnbIdLengthreportingNrdcR17                *MeasandmobparameterscommonGnbIdLengthreportingNrdcR17                `ext`
	GnbIdLengthreportingNpnR17                 *MeasandmobparameterscommonGnbIdLengthreportingNpnR17                 `ext`
	ParallelsmtcR17                            *MeasandmobparameterscommonParallelsmtcR17                            `ext`
	ConcurrentmeasgapeutraR17                  *MeasandmobparameterscommonConcurrentmeasgapeutraR17                  `ext`
	ServicelinkpropdelaydiffreportingR17       *MeasandmobparameterscommonServicelinkpropdelaydiffreportingR17       `ext`
	NcsgSymbollevelschedulerestrictioninterR17 *MeasandmobparameterscommonNcsgSymbollevelschedulerestrictioninterR17 `ext`
	Eventd1MeasreporttriggerR17                *MeasandmobparameterscommonEventd1MeasreporttriggerR17                `ext`
	IndependentgapconfigMaxccR17               *MeasandmobparameterscommonIndependentgapconfigMaxccR17               `ext`
}
