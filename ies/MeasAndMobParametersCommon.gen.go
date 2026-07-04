// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersCommon (line 21112).

var measAndMobParametersCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedGapPattern", Optional: true},
		{Name: "ssb-RLM", Optional: true},
		{Name: "ssb-AndCSI-RS-RLM", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
		{Name: "extension-addition-6", Optional: true},
		{Name: "extension-addition-7", Optional: true},
		{Name: "extension-addition-8", Optional: true},
		{Name: "extension-addition-9", Optional: true},
		{Name: "extension-addition-10", Optional: true},
		{Name: "extension-addition-11", Optional: true},
		{Name: "extension-addition-12", Optional: true},
	},
}

var measAndMobParametersCommonSupportedGapPatternConstraints = per.FixedSize(22)

const (
	MeasAndMobParametersCommon_Ssb_RLM_Supported = 0
)

var measAndMobParametersCommonSsbRLMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ssb_RLM_Supported},
}

const (
	MeasAndMobParametersCommon_Ssb_AndCSI_RS_RLM_Supported = 0
)

var measAndMobParametersCommonSsbAndCSIRSRLMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ssb_AndCSI_RS_RLM_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_EventB_MeasAndReport_Supported = 0
)

var measAndMobParametersCommonExtEventBMeasAndReportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_EventB_MeasAndReport_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_HandoverFDD_TDD_Supported = 0
)

var measAndMobParametersCommonExtHandoverFDDTDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_HandoverFDD_TDD_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_CGI_Reporting_Supported = 0
)

var measAndMobParametersCommonExtEutraCGIReportingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_CGI_Reporting_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_Supported = 0
)

var measAndMobParametersCommonExtNrCGIReportingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_IndependentGapConfig_Supported = 0
)

var measAndMobParametersCommonExtIndependentGapConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_IndependentGapConfig_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_PeriodicEUTRA_MeasAndReport_Supported = 0
)

var measAndMobParametersCommonExtPeriodicEUTRAMeasAndReportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_PeriodicEUTRA_MeasAndReport_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_HandoverFR1_FR2_Supported = 0
)

var measAndMobParametersCommonExtHandoverFR1FR2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_HandoverFR1_FR2_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N4  = 0
	MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N8  = 1
	MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N16 = 2
	MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N32 = 3
	MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N64 = 4
	MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N96 = 5
)

var measAndMobParametersCommonExtMaxNumberCSIRSRRMRSSINRConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N4, MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N8, MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N16, MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N32, MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N64, MeasAndMobParametersCommon_Ext_MaxNumberCSI_RS_RRM_RS_SINR_N96},
}

const (
	MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_ENDC_Supported = 0
)

var measAndMobParametersCommonExtNrCGIReportingENDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_ENDC_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_CGI_Reporting_NEDC_Supported = 0
)

var measAndMobParametersCommonExtEutraCGIReportingNEDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_CGI_Reporting_NEDC_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_CGI_Reporting_NRDC_Supported = 0
)

var measAndMobParametersCommonExtEutraCGIReportingNRDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_CGI_Reporting_NRDC_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_NEDC_Supported = 0
)

var measAndMobParametersCommonExtNrCGIReportingNEDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_NEDC_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_NRDC_Supported = 0
)

var measAndMobParametersCommonExtNrCGIReportingNRDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_NRDC_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ReportAddNeighMeasForPeriodic_r16_Supported = 0
)

var measAndMobParametersCommonExtReportAddNeighMeasForPeriodicR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ReportAddNeighMeasForPeriodic_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Nr_NeedForGap_Reporting_r16_Supported = 0
)

var measAndMobParametersCommonExtNrNeedForGapReportingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Nr_NeedForGap_Reporting_r16_Supported},
}

var measAndMobParametersCommonSupportedGapPatternNRonlyR16Constraints = per.FixedSize(10)

const (
	MeasAndMobParametersCommon_Ext_SupportedGapPattern_NRonly_NEDC_r16_Supported = 0
)

var measAndMobParametersCommonExtSupportedGapPatternNRonlyNEDCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_SupportedGapPattern_NRonly_NEDC_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_MaxNumberCLI_RSSI_r16_N8  = 0
	MeasAndMobParametersCommon_Ext_MaxNumberCLI_RSSI_r16_N16 = 1
	MeasAndMobParametersCommon_Ext_MaxNumberCLI_RSSI_r16_N32 = 2
	MeasAndMobParametersCommon_Ext_MaxNumberCLI_RSSI_r16_N64 = 3
)

var measAndMobParametersCommonExtMaxNumberCLIRSSIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_MaxNumberCLI_RSSI_r16_N8, MeasAndMobParametersCommon_Ext_MaxNumberCLI_RSSI_r16_N16, MeasAndMobParametersCommon_Ext_MaxNumberCLI_RSSI_r16_N32, MeasAndMobParametersCommon_Ext_MaxNumberCLI_RSSI_r16_N64},
}

const (
	MeasAndMobParametersCommon_Ext_MaxNumberCLI_SRS_RSRP_r16_N4  = 0
	MeasAndMobParametersCommon_Ext_MaxNumberCLI_SRS_RSRP_r16_N8  = 1
	MeasAndMobParametersCommon_Ext_MaxNumberCLI_SRS_RSRP_r16_N16 = 2
	MeasAndMobParametersCommon_Ext_MaxNumberCLI_SRS_RSRP_r16_N32 = 3
)

var measAndMobParametersCommonExtMaxNumberCLISRSRSRPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_MaxNumberCLI_SRS_RSRP_r16_N4, MeasAndMobParametersCommon_Ext_MaxNumberCLI_SRS_RSRP_r16_N8, MeasAndMobParametersCommon_Ext_MaxNumberCLI_SRS_RSRP_r16_N16, MeasAndMobParametersCommon_Ext_MaxNumberCLI_SRS_RSRP_r16_N32},
}

const (
	MeasAndMobParametersCommon_Ext_MaxNumberPerSlotCLI_SRS_RSRP_r16_N2 = 0
	MeasAndMobParametersCommon_Ext_MaxNumberPerSlotCLI_SRS_RSRP_r16_N4 = 1
	MeasAndMobParametersCommon_Ext_MaxNumberPerSlotCLI_SRS_RSRP_r16_N8 = 2
)

var measAndMobParametersCommonExtMaxNumberPerSlotCLISRSRSRPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_MaxNumberPerSlotCLI_SRS_RSRP_r16_N2, MeasAndMobParametersCommon_Ext_MaxNumberPerSlotCLI_SRS_RSRP_r16_N4, MeasAndMobParametersCommon_Ext_MaxNumberPerSlotCLI_SRS_RSRP_r16_N8},
}

const (
	MeasAndMobParametersCommon_Ext_Mfbi_IAB_r16_Supported = 0
)

var measAndMobParametersCommonExtMfbiIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Mfbi_IAB_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Dummy_Supported = 0
)

var measAndMobParametersCommonExtDummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Dummy_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_NPN_r16_Supported = 0
)

var measAndMobParametersCommonExtNrCGIReportingNPNR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_NPN_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_IdleInactiveEUTRA_MeasReport_r16_Supported = 0
)

var measAndMobParametersCommonExtIdleInactiveEUTRAMeasReportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_IdleInactiveEUTRA_MeasReport_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_IdleInactive_ValidityArea_r16_Supported = 0
)

var measAndMobParametersCommonExtIdleInactiveValidityAreaR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_IdleInactive_ValidityArea_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_AutonomousGaps_r16_Supported = 0
)

var measAndMobParametersCommonExtEutraAutonomousGapsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_AutonomousGaps_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_AutonomousGaps_NEDC_r16_Supported = 0
)

var measAndMobParametersCommonExtEutraAutonomousGapsNEDCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_AutonomousGaps_NEDC_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_AutonomousGaps_NRDC_r16_Supported = 0
)

var measAndMobParametersCommonExtEutraAutonomousGapsNRDCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_AutonomousGaps_NRDC_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_PcellT312_r16_Supported = 0
)

var measAndMobParametersCommonExtPcellT312R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_PcellT312_r16_Supported},
}

var measAndMobParametersCommonSupportedGapPatternR16Constraints = per.FixedSize(2)

var measAndMobParametersCommonExtConcurrentMeasGapR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "concurrentPerUE-OnlyMeasGap-r17"},
		{Name: "concurrentPerUE-PerFRCombMeasGap-r17"},
	},
}

const (
	MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_OnlyMeasGap_r17      = 0
	MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_PerFRCombMeasGap_r17 = 1
)

const (
	MeasAndMobParametersCommon_Ext_Nr_NeedForGapNCSG_Reporting_r17_Supported = 0
)

var measAndMobParametersCommonExtNrNeedForGapNCSGReportingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Nr_NeedForGapNCSG_Reporting_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_NeedForGapNCSG_Reporting_r17_Supported = 0
)

var measAndMobParametersCommonExtEutraNeedForGapNCSGReportingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_NeedForGapNCSG_Reporting_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ncsg_MeasGapPerFR_r17_Supported = 0
)

var measAndMobParametersCommonExtNcsgMeasGapPerFRR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ncsg_MeasGapPerFR_r17_Supported},
}

var measAndMobParametersCommonNcsgMeasGapPatternsR17Constraints = per.FixedSize(24)

var measAndMobParametersCommonNcsgMeasGapNRPatternsR17Constraints = per.FixedSize(24)

const (
	MeasAndMobParametersCommon_Ext_PreconfiguredUE_AutonomousMeasGap_r17_Supported = 0
)

var measAndMobParametersCommonExtPreconfiguredUEAutonomousMeasGapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_PreconfiguredUE_AutonomousMeasGap_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_PreconfiguredNW_ControlledMeasGap_r17_Supported = 0
)

var measAndMobParametersCommonExtPreconfiguredNWControlledMeasGapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_PreconfiguredNW_ControlledMeasGap_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_HandoverFR1_FR2_2_r17_Supported = 0
)

var measAndMobParametersCommonExtHandoverFR1FR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_HandoverFR1_FR2_2_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_HandoverFR2_1_FR2_2_r17_Supported = 0
)

var measAndMobParametersCommonExtHandoverFR21FR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_HandoverFR2_1_FR2_2_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_IndependentGapConfigPRS_r17_Supported = 0
)

var measAndMobParametersCommonExtIndependentGapConfigPRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_IndependentGapConfigPRS_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Rrm_RelaxationRRC_ConnectedRedCap_r17_Supported = 0
)

var measAndMobParametersCommonExtRrmRelaxationRRCConnectedRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Rrm_RelaxationRRC_ConnectedRedCap_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ParallelMeasurementGap_r17_N2 = 0
)

var measAndMobParametersCommonExtParallelMeasurementGapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ParallelMeasurementGap_r17_N2},
}

const (
	MeasAndMobParametersCommon_Ext_CondHandoverWithSCG_NRDC_r17_Supported = 0
)

var measAndMobParametersCommonExtCondHandoverWithSCGNRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_CondHandoverWithSCG_NRDC_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_r17_Supported = 0
)

var measAndMobParametersCommonExtGNBIDLengthReportingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_ENDC_r17_Supported = 0
)

var measAndMobParametersCommonExtGNBIDLengthReportingENDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_ENDC_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_NEDC_r17_Supported = 0
)

var measAndMobParametersCommonExtGNBIDLengthReportingNEDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_NEDC_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_NRDC_r17_Supported = 0
)

var measAndMobParametersCommonExtGNBIDLengthReportingNRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_NRDC_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_NPN_r17_Supported = 0
)

var measAndMobParametersCommonExtGNBIDLengthReportingNPNR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_GNB_ID_LengthReporting_NPN_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ParallelSMTC_r17_N4 = 0
)

var measAndMobParametersCommonExtParallelSMTCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ParallelSMTC_r17_N4},
}

const (
	MeasAndMobParametersCommon_Ext_ConcurrentMeasGapEUTRA_r17_Supported = 0
)

var measAndMobParametersCommonExtConcurrentMeasGapEUTRAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ConcurrentMeasGapEUTRA_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ServiceLinkPropDelayDiffReporting_r17_Supported = 0
)

var measAndMobParametersCommonExtServiceLinkPropDelayDiffReportingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ServiceLinkPropDelayDiffReporting_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ncsg_SymbolLevelScheduleRestrictionInter_r17_Supported = 0
)

var measAndMobParametersCommonExtNcsgSymbolLevelScheduleRestrictionInterR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ncsg_SymbolLevelScheduleRestrictionInter_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_EventD1_MeasReportTrigger_r17_Supported = 0
)

var measAndMobParametersCommonExtEventD1MeasReportTriggerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_EventD1_MeasReportTrigger_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_InterSatMeas_r17_Supported = 0
)

var measAndMobParametersCommonExtInterSatMeasR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_InterSatMeas_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_DeriveSSB_IndexFromCellInterNon_NCSG_r17_Supported = 0
)

var measAndMobParametersCommonExtDeriveSSBIndexFromCellInterNonNCSGR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_DeriveSSB_IndexFromCellInterNon_NCSG_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_L3_MeasUnknownSCellActivation_r18_Supported = 0
)

var measAndMobParametersCommonExtL3MeasUnknownSCellActivationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_L3_MeasUnknownSCellActivation_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ShortMeasInterval_r18_Supported = 0
)

var measAndMobParametersCommonExtShortMeasIntervalR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ShortMeasInterval_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Nr_NeedForInterruptionReport_r18_Supported = 0
)

var measAndMobParametersCommonExtNrNeedForInterruptionReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Nr_NeedForInterruptionReport_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_MeasSequenceConfig_r18_Supported = 0
)

var measAndMobParametersCommonExtMeasSequenceConfigR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_MeasSequenceConfig_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_CellIndividualOffsetPerMeasEvent_r18_Supported = 0
)

var measAndMobParametersCommonExtCellIndividualOffsetPerMeasEventR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_CellIndividualOffsetPerMeasEvent_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_EventD2_MeasReportTrigger_r18_Supported = 0
)

var measAndMobParametersCommonExtEventD2MeasReportTriggerR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_EventD2_MeasReportTrigger_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ConcurrentMeasGapsPreMG_r18_Supported = 0
)

var measAndMobParametersCommonExtConcurrentMeasGapsPreMGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ConcurrentMeasGapsPreMG_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_DynamicCollision_r18_Supported = 0
)

var measAndMobParametersCommonExtDynamicCollisionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_DynamicCollision_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ConcurrentMeasGapsNCSG_r18_Supported = 0
)

var measAndMobParametersCommonExtConcurrentMeasGapsNCSGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ConcurrentMeasGapsNCSG_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_NoGapMeasurementOutsideBWP_r18_Supported = 0
)

var measAndMobParametersCommonExtEutraNoGapMeasurementOutsideBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_NoGapMeasurementOutsideBWP_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_NoGapMeasurementInsideBWP_r18_Supported = 0
)

var measAndMobParametersCommonExtEutraNoGapMeasurementInsideBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_NoGapMeasurementInsideBWP_r18_Supported},
}

var measAndMobParametersCommonEutraMeasEMWR18Constraints = per.FixedSize(6)

const (
	MeasAndMobParametersCommon_Ext_ConcurrentMeasCRS_InsideBWP_EUTRA_r18_Supported = 0
)

var measAndMobParametersCommonExtConcurrentMeasCRSInsideBWPEUTRAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ConcurrentMeasCRS_InsideBWP_EUTRA_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_InterFreqMeasGap_r18_Supported = 0
)

var measAndMobParametersCommonExtLtmInterFreqMeasGapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_InterFreqMeasGap_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Rach_LessHandoverInterFreq_r18_Supported = 0
)

var measAndMobParametersCommonExtRachLessHandoverInterFreqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Rach_LessHandoverInterFreq_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_EnterAndLeaveCellReport_r18_Supported = 0
)

var measAndMobParametersCommonExtEnterAndLeaveCellReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_EnterAndLeaveCellReport_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_BestCellChangeReport_r18_Supported = 0
)

var measAndMobParametersCommonExtBestCellChangeReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_BestCellChangeReport_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_SecondBestCellChangeReport_r18_Supported = 0
)

var measAndMobParametersCommonExtSecondBestCellChangeReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_SecondBestCellChangeReport_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_InterFreq_r18_Supported = 0
)

var measAndMobParametersCommonExtLtmInterFreqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_InterFreq_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_MCG_NRDC_r18_Supported = 0
)

var measAndMobParametersCommonExtLtmMCGNRDCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_MCG_NRDC_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_RACH_LessDG_r18_Supported = 0
)

var measAndMobParametersCommonExtLtmRACHLessDGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_RACH_LessDG_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_RACH_LessCG_r18_Supported = 0
)

var measAndMobParametersCommonExtLtmRACHLessCGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_RACH_LessCG_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_Recovery_r18_Supported = 0
)

var measAndMobParametersCommonExtLtmRecoveryR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_Recovery_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_ReferenceConfig_r18_Supported = 0
)

var measAndMobParametersCommonExtLtmReferenceConfigR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_ReferenceConfig_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_MCG_NRDC_Release_r18_Supported = 0
)

var measAndMobParametersCommonExtLtmMCGNRDCReleaseR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_MCG_NRDC_Release_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ntn_NeighbourCellInfoSupport_r18_Supported = 0
)

var measAndMobParametersCommonExtNtnNeighbourCellInfoSupportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ntn_NeighbourCellInfoSupport_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_InterFreqL1_OnlyInBC_r18_Supported = 0
)

var measAndMobParametersCommonExtLtmInterFreqL1OnlyInBCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_InterFreqL1_OnlyInBC_r18_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_HSDN_r19_Supported = 0
)

var measAndMobParametersCommonExtNrCGIReportingHSDNR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Nr_CGI_Reporting_HSDN_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Eutra_CGI_Reporting_HSDN_r19_Supported = 0
)

var measAndMobParametersCommonExtEutraCGIReportingHSDNR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Eutra_CGI_Reporting_HSDN_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_KeyUpdateMCG_r19_Supported = 0
)

var measAndMobParametersCommonExtLtmKeyUpdateMCGR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_KeyUpdateMCG_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_KeyUpdateSCG_r19_Supported = 0
)

var measAndMobParametersCommonExtLtmKeyUpdateSCGR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_KeyUpdateSCG_r19_Supported},
}

var measAndMobParametersCommonCltmEarlyTAIndicationR19Constraints = per.Constrained(1, 8)

const (
	MeasAndMobParametersCommon_Ext_Cltm_ExecutionConditionL1_r19_Supported = 0
)

var measAndMobParametersCommonExtCltmExecutionConditionL1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Cltm_ExecutionConditionL1_r19_Supported},
}

var measAndMobParametersCommonCltmExecutionConditionL3R19Constraints = per.Constrained(1, 2)

const (
	MeasAndMobParametersCommon_Ext_Ltm_EventMeasAndReport_r19_Supported = 0
)

var measAndMobParametersCommonExtLtmEventMeasAndReportR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_EventMeasAndReport_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_RecoveryWithKeyUpdate_r19_Supported = 0
)

var measAndMobParametersCommonExtLtmRecoveryWithKeyUpdateR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_RecoveryWithKeyUpdate_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_MCG_SCG_AdditionOrChange_r19_Supported = 0
)

var measAndMobParametersCommonExtLtmMCGSCGAdditionOrChangeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_MCG_SCG_AdditionOrChange_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_MultiCarrierSingleReportWithoutGap_r19_Supported = 0
)

var measAndMobParametersCommonExtMultiCarrierSingleReportWithoutGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_MultiCarrierSingleReportWithoutGap_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_SkipSSB_L1_RSRP_Meas_r19_Neighbour = 0
	MeasAndMobParametersCommon_Ext_SkipSSB_L1_RSRP_Meas_r19_Both      = 1
)

var measAndMobParametersCommonExtSkipSSBL1RSRPMeasR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_SkipSSB_L1_RSRP_Meas_r19_Neighbour, MeasAndMobParametersCommon_Ext_SkipSSB_L1_RSRP_Meas_r19_Both},
}

const (
	MeasAndMobParametersCommon_Ext_GapOccasionCancelRatioReport_r19_Supported = 0
)

var measAndMobParametersCommonExtGapOccasionCancelRatioReportR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_GapOccasionCancelRatioReport_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_TwoSMTC_PeriodicitiesIntraFreq_r19_Supported = 0
)

var measAndMobParametersCommonExtTwoSMTCPeriodicitiesIntraFreqR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_TwoSMTC_PeriodicitiesIntraFreq_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_TwoSMTC_PeriodicitiesInterFreq_r19_Supported = 0
)

var measAndMobParametersCommonExtTwoSMTCPeriodicitiesInterFreqR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_TwoSMTC_PeriodicitiesInterFreq_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ReportClosestReferenceLocations_r19_Supported = 0
)

var measAndMobParametersCommonExtReportClosestReferenceLocationsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ReportClosestReferenceLocations_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_SR_ConfIdInCellSwitchCommand_r19_Supported = 0
)

var measAndMobParametersCommonExtLtmSRConfIdInCellSwitchCommandR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_SR_ConfIdInCellSwitchCommand_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_SpeedStateDetection_r19_Supported = 0
)

var measAndMobParametersCommonExtSpeedStateDetectionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_SpeedStateDetection_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_EventD2_MultiReferenceLocations_r19_Supported = 0
)

var measAndMobParametersCommonExtEventD2MultiReferenceLocationsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_EventD2_MultiReferenceLocations_r19_Supported},
}

var measAndMobParametersCommonExtCondHandoverParametersCommonR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "condHandoverFDD-TDD-r16", Optional: true},
		{Name: "condHandoverFR1-FR2-r16", Optional: true},
	},
}

const (
	MeasAndMobParametersCommon_Ext_CondHandoverParametersCommon_r16_CondHandoverFDD_TDD_r16_Supported = 0
)

var measAndMobParametersCommonExtCondHandoverParametersCommonR16CondHandoverFDDTDDR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_CondHandoverParametersCommon_r16_CondHandoverFDD_TDD_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_CondHandoverParametersCommon_r16_CondHandoverFR1_FR2_r16_Supported = 0
)

var measAndMobParametersCommonExtCondHandoverParametersCommonR16CondHandoverFR1FR2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_CondHandoverParametersCommon_r16_CondHandoverFR1_FR2_r16_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_OnlyMeasGap_r17_Supported = 0
)

var measAndMobParametersCommonExtConcurrentMeasGapR17ConcurrentPerUEOnlyMeasGapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_OnlyMeasGap_r17_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_PerFRCombMeasGap_r17_Supported = 0
)

var measAndMobParametersCommonExtConcurrentMeasGapR17ConcurrentPerUEPerFRCombMeasGapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_PerFRCombMeasGap_r17_Supported},
}

var measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1-Only-r17", Optional: true},
		{Name: "fr2-Only-r17", Optional: true},
		{Name: "fr1-AndFR2-r17", Optional: true},
	},
}

const (
	MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr1_r18_Ms10 = 0
	MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr1_r18_Ms15 = 1
)

var measAndMobParametersCommonExtDummyLtmFastUEProcessingR18Fr1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr1_r18_Ms10, MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr1_r18_Ms15},
}

const (
	MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr2_r18_Ms10 = 0
	MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr2_r18_Ms15 = 1
)

var measAndMobParametersCommonExtDummyLtmFastUEProcessingR18Fr2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr2_r18_Ms10, MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr2_r18_Ms15},
}

const (
	MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr1_AndFR2_r18_Ms20 = 0
	MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr1_AndFR2_r18_Ms30 = 1
)

var measAndMobParametersCommonExtDummyLtmFastUEProcessingR18Fr1AndFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr1_AndFR2_r18_Ms20, MeasAndMobParametersCommon_Ext_Dummy_Ltm_FastUE_Processing_r18_Fr1_AndFR2_r18_Ms30},
}

var measAndMobParametersCommonExtLtmFastUEProcessingR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1-r18", Optional: true},
		{Name: "fr2-r18", Optional: true},
		{Name: "fr1-AndFR2-r18", Optional: true},
	},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr1_r18_Ms10 = 0
	MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr1_r18_Ms15 = 1
)

var measAndMobParametersCommonExtLtmFastUEProcessingR18Fr1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr1_r18_Ms10, MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr1_r18_Ms15},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr2_r18_Ms10 = 0
	MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr2_r18_Ms15 = 1
)

var measAndMobParametersCommonExtLtmFastUEProcessingR18Fr2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr2_r18_Ms10, MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr2_r18_Ms15},
}

const (
	MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr1_AndFR2_r18_Ms20 = 0
	MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr1_AndFR2_r18_Ms30 = 1
)

var measAndMobParametersCommonExtLtmFastUEProcessingR18Fr1AndFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr1_AndFR2_r18_Ms20, MeasAndMobParametersCommon_Ext_Ltm_FastUE_Processing_r18_Fr1_AndFR2_r18_Ms30},
}

var measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1-CA-NR-DC-r19", Optional: true},
		{Name: "fr1-FR2-CA-r19", Optional: true},
		{Name: "fr1-FR2-NR-DC-r19", Optional: true},
	},
}

const (
	MeasAndMobParametersCommon_Ext_ThreeCarrierMeasWithoutGap_r19_Fr1_CA_NR_DC_r19_Supported = 0
)

var measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Fr1CANRDCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ThreeCarrierMeasWithoutGap_r19_Fr1_CA_NR_DC_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ThreeCarrierMeasWithoutGap_r19_Fr1_FR2_CA_r19_Supported = 0
)

var measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Fr1FR2CAR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ThreeCarrierMeasWithoutGap_r19_Fr1_FR2_CA_r19_Supported},
}

const (
	MeasAndMobParametersCommon_Ext_ThreeCarrierMeasWithoutGap_r19_Fr1_FR2_NR_DC_r19_Supported = 0
)

var measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Fr1FR2NRDCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersCommon_Ext_ThreeCarrierMeasWithoutGap_r19_Fr1_FR2_NR_DC_r19_Supported},
}

type MeasAndMobParametersCommon struct {
	SupportedGapPattern               *per.BitString
	Ssb_RLM                           *int64
	Ssb_AndCSI_RS_RLM                 *int64
	EventB_MeasAndReport              *int64
	HandoverFDD_TDD                   *int64
	Eutra_CGI_Reporting               *int64
	Nr_CGI_Reporting                  *int64
	IndependentGapConfig              *int64
	PeriodicEUTRA_MeasAndReport       *int64
	HandoverFR1_FR2                   *int64
	MaxNumberCSI_RS_RRM_RS_SINR       *int64
	Nr_CGI_Reporting_ENDC             *int64
	Eutra_CGI_Reporting_NEDC          *int64
	Eutra_CGI_Reporting_NRDC          *int64
	Nr_CGI_Reporting_NEDC             *int64
	Nr_CGI_Reporting_NRDC             *int64
	ReportAddNeighMeasForPeriodic_r16 *int64
	CondHandoverParametersCommon_r16  *struct {
		CondHandoverFDD_TDD_r16 *int64
		CondHandoverFR1_FR2_r16 *int64
	}
	Nr_NeedForGap_Reporting_r16         *int64
	SupportedGapPattern_NRonly_r16      *per.BitString
	SupportedGapPattern_NRonly_NEDC_r16 *int64
	MaxNumberCLI_RSSI_r16               *int64
	MaxNumberCLI_SRS_RSRP_r16           *int64
	MaxNumberPerSlotCLI_SRS_RSRP_r16    *int64
	Mfbi_IAB_r16                        *int64
	Dummy                               *int64
	Nr_CGI_Reporting_NPN_r16            *int64
	IdleInactiveEUTRA_MeasReport_r16    *int64
	IdleInactive_ValidityArea_r16       *int64
	Eutra_AutonomousGaps_r16            *int64
	Eutra_AutonomousGaps_NEDC_r16       *int64
	Eutra_AutonomousGaps_NRDC_r16       *int64
	PcellT312_r16                       *int64
	SupportedGapPattern_r16             *per.BitString
	ConcurrentMeasGap_r17               *struct {
		Choice                               int
		ConcurrentPerUE_OnlyMeasGap_r17      *int64
		ConcurrentPerUE_PerFRCombMeasGap_r17 *int64
	}
	Nr_NeedForGapNCSG_Reporting_r17              *int64
	Eutra_NeedForGapNCSG_Reporting_r17           *int64
	Ncsg_MeasGapPerFR_r17                        *int64
	Ncsg_MeasGapPatterns_r17                     *per.BitString
	Ncsg_MeasGapNR_Patterns_r17                  *per.BitString
	PreconfiguredUE_AutonomousMeasGap_r17        *int64
	PreconfiguredNW_ControlledMeasGap_r17        *int64
	HandoverFR1_FR2_2_r17                        *int64
	HandoverFR2_1_FR2_2_r17                      *int64
	IndependentGapConfigPRS_r17                  *int64
	Rrm_RelaxationRRC_ConnectedRedCap_r17        *int64
	ParallelMeasurementGap_r17                   *int64
	CondHandoverWithSCG_NRDC_r17                 *int64
	GNB_ID_LengthReporting_r17                   *int64
	GNB_ID_LengthReporting_ENDC_r17              *int64
	GNB_ID_LengthReporting_NEDC_r17              *int64
	GNB_ID_LengthReporting_NRDC_r17              *int64
	GNB_ID_LengthReporting_NPN_r17               *int64
	ParallelSMTC_r17                             *int64
	ConcurrentMeasGapEUTRA_r17                   *int64
	ServiceLinkPropDelayDiffReporting_r17        *int64
	Ncsg_SymbolLevelScheduleRestrictionInter_r17 *int64
	EventD1_MeasReportTrigger_r17                *int64
	IndependentGapConfig_MaxCC_r17               *struct {
		Fr1_Only_r17   *int64
		Fr2_Only_r17   *int64
		Fr1_AndFR2_r17 *int64
	}
	InterSatMeas_r17                         *int64
	DeriveSSB_IndexFromCellInterNon_NCSG_r17 *int64
	L3_MeasUnknownSCellActivation_r18        *int64
	ShortMeasInterval_r18                    *int64
	Nr_NeedForInterruptionReport_r18         *int64
	MeasSequenceConfig_r18                   *int64
	CellIndividualOffsetPerMeasEvent_r18     *int64
	EventD2_MeasReportTrigger_r18            *int64
	ConcurrentMeasGapsPreMG_r18              *int64
	DynamicCollision_r18                     *int64
	ConcurrentMeasGapsNCSG_r18               *int64
	Eutra_NoGapMeasurementOutsideBWP_r18     *int64
	Eutra_NoGapMeasurementInsideBWP_r18      *int64
	Eutra_MeasEMW_r18                        *per.BitString
	ConcurrentMeasCRS_InsideBWP_EUTRA_r18    *int64
	Ltm_InterFreqMeasGap_r18                 *int64
	Dummy_Ltm_FastUE_Processing_r18          *struct {
		Fr1_r18        int64
		Fr2_r18        int64
		Fr1_AndFR2_r18 int64
	}
	Rach_LessHandoverInterFreq_r18 *int64
	EnterAndLeaveCellReport_r18    *int64
	BestCellChangeReport_r18       *int64
	SecondBestCellChangeReport_r18 *int64
	Ltm_InterFreq_r18              *int64
	Ltm_MCG_NRDC_r18               *int64
	Ltm_RACH_LessDG_r18            *int64
	Ltm_RACH_LessCG_r18            *int64
	Ltm_Recovery_r18               *int64
	Ltm_ReferenceConfig_r18        *int64
	Ltm_MCG_NRDC_Release_r18       *int64
	Ltm_FastUE_Processing_r18      *struct {
		Fr1_r18        *int64
		Fr2_r18        *int64
		Fr1_AndFR2_r18 *int64
	}
	Ntn_NeighbourCellInfoSupport_r18 *int64
	Ltm_InterFreqL1_OnlyInBC_r18     *int64
	Nr_CGI_Reporting_HSDN_r19        *int64
	Eutra_CGI_Reporting_HSDN_r19     *int64
	Ltm_KeyUpdateMCG_r19             *int64
	Ltm_KeyUpdateSCG_r19             *int64
	Cltm_EarlyTA_Indication_r19      *int64
	Cltm_ExecutionConditionL1_r19    *int64
	Cltm_ExecutionConditionL3_r19    *int64
	Ltm_EventMeasAndReport_r19       *int64
	Ltm_RecoveryWithKeyUpdate_r19    *int64
	Ltm_MCG_SCG_AdditionOrChange_r19 *int64
	ThreeCarrierMeasWithoutGap_r19   *struct {
		Fr1_CA_NR_DC_r19  *int64
		Fr1_FR2_CA_r19    *int64
		Fr1_FR2_NR_DC_r19 *int64
	}
	MultiCarrierSingleReportWithoutGap_r19 *int64
	SkipSSB_L1_RSRP_Meas_r19               *int64
	GapOccasionCancelRatioReport_r19       *int64
	TwoSMTC_PeriodicitiesIntraFreq_r19     *int64
	TwoSMTC_PeriodicitiesInterFreq_r19     *int64
	ReportClosestReferenceLocations_r19    *int64
	Ltm_SR_ConfIdInCellSwitchCommand_r19   *int64
	SpeedStateDetection_r19                *int64
	EventD2_MultiReferenceLocations_r19    *int64
}

func (ie *MeasAndMobParametersCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.EventB_MeasAndReport != nil || ie.HandoverFDD_TDD != nil || ie.Eutra_CGI_Reporting != nil || ie.Nr_CGI_Reporting != nil
	hasExtGroup1 := ie.IndependentGapConfig != nil || ie.PeriodicEUTRA_MeasAndReport != nil || ie.HandoverFR1_FR2 != nil || ie.MaxNumberCSI_RS_RRM_RS_SINR != nil
	hasExtGroup2 := ie.Nr_CGI_Reporting_ENDC != nil
	hasExtGroup3 := ie.Eutra_CGI_Reporting_NEDC != nil || ie.Eutra_CGI_Reporting_NRDC != nil || ie.Nr_CGI_Reporting_NEDC != nil || ie.Nr_CGI_Reporting_NRDC != nil
	hasExtGroup4 := ie.ReportAddNeighMeasForPeriodic_r16 != nil || ie.CondHandoverParametersCommon_r16 != nil || ie.Nr_NeedForGap_Reporting_r16 != nil || ie.SupportedGapPattern_NRonly_r16 != nil || ie.SupportedGapPattern_NRonly_NEDC_r16 != nil || ie.MaxNumberCLI_RSSI_r16 != nil || ie.MaxNumberCLI_SRS_RSRP_r16 != nil || ie.MaxNumberPerSlotCLI_SRS_RSRP_r16 != nil || ie.Mfbi_IAB_r16 != nil || ie.Dummy != nil || ie.Nr_CGI_Reporting_NPN_r16 != nil || ie.IdleInactiveEUTRA_MeasReport_r16 != nil || ie.IdleInactive_ValidityArea_r16 != nil || ie.Eutra_AutonomousGaps_r16 != nil || ie.Eutra_AutonomousGaps_NEDC_r16 != nil || ie.Eutra_AutonomousGaps_NRDC_r16 != nil || ie.PcellT312_r16 != nil || ie.SupportedGapPattern_r16 != nil
	hasExtGroup5 := ie.ConcurrentMeasGap_r17 != nil || ie.Nr_NeedForGapNCSG_Reporting_r17 != nil || ie.Eutra_NeedForGapNCSG_Reporting_r17 != nil || ie.Ncsg_MeasGapPerFR_r17 != nil || ie.Ncsg_MeasGapPatterns_r17 != nil || ie.Ncsg_MeasGapNR_Patterns_r17 != nil || ie.PreconfiguredUE_AutonomousMeasGap_r17 != nil || ie.PreconfiguredNW_ControlledMeasGap_r17 != nil || ie.HandoverFR1_FR2_2_r17 != nil || ie.HandoverFR2_1_FR2_2_r17 != nil || ie.IndependentGapConfigPRS_r17 != nil || ie.Rrm_RelaxationRRC_ConnectedRedCap_r17 != nil || ie.ParallelMeasurementGap_r17 != nil || ie.CondHandoverWithSCG_NRDC_r17 != nil || ie.GNB_ID_LengthReporting_r17 != nil || ie.GNB_ID_LengthReporting_ENDC_r17 != nil || ie.GNB_ID_LengthReporting_NEDC_r17 != nil || ie.GNB_ID_LengthReporting_NRDC_r17 != nil || ie.GNB_ID_LengthReporting_NPN_r17 != nil
	hasExtGroup6 := ie.ParallelSMTC_r17 != nil || ie.ConcurrentMeasGapEUTRA_r17 != nil || ie.ServiceLinkPropDelayDiffReporting_r17 != nil || ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil
	hasExtGroup7 := ie.EventD1_MeasReportTrigger_r17 != nil || ie.IndependentGapConfig_MaxCC_r17 != nil
	hasExtGroup8 := ie.InterSatMeas_r17 != nil || ie.DeriveSSB_IndexFromCellInterNon_NCSG_r17 != nil
	hasExtGroup9 := ie.L3_MeasUnknownSCellActivation_r18 != nil || ie.ShortMeasInterval_r18 != nil || ie.Nr_NeedForInterruptionReport_r18 != nil || ie.MeasSequenceConfig_r18 != nil || ie.CellIndividualOffsetPerMeasEvent_r18 != nil || ie.EventD2_MeasReportTrigger_r18 != nil || ie.ConcurrentMeasGapsPreMG_r18 != nil || ie.DynamicCollision_r18 != nil || ie.ConcurrentMeasGapsNCSG_r18 != nil || ie.Eutra_NoGapMeasurementOutsideBWP_r18 != nil || ie.Eutra_NoGapMeasurementInsideBWP_r18 != nil || ie.Eutra_MeasEMW_r18 != nil || ie.ConcurrentMeasCRS_InsideBWP_EUTRA_r18 != nil || ie.Ltm_InterFreqMeasGap_r18 != nil || ie.Dummy_Ltm_FastUE_Processing_r18 != nil || ie.Rach_LessHandoverInterFreq_r18 != nil || ie.EnterAndLeaveCellReport_r18 != nil || ie.BestCellChangeReport_r18 != nil || ie.SecondBestCellChangeReport_r18 != nil
	hasExtGroup10 := ie.Ltm_InterFreq_r18 != nil || ie.Ltm_MCG_NRDC_r18 != nil || ie.Ltm_RACH_LessDG_r18 != nil || ie.Ltm_RACH_LessCG_r18 != nil || ie.Ltm_Recovery_r18 != nil || ie.Ltm_ReferenceConfig_r18 != nil || ie.Ltm_MCG_NRDC_Release_r18 != nil || ie.Ltm_FastUE_Processing_r18 != nil || ie.Ntn_NeighbourCellInfoSupport_r18 != nil
	hasExtGroup11 := ie.Ltm_InterFreqL1_OnlyInBC_r18 != nil
	hasExtGroup12 := ie.Nr_CGI_Reporting_HSDN_r19 != nil || ie.Eutra_CGI_Reporting_HSDN_r19 != nil || ie.Ltm_KeyUpdateMCG_r19 != nil || ie.Ltm_KeyUpdateSCG_r19 != nil || ie.Cltm_EarlyTA_Indication_r19 != nil || ie.Cltm_ExecutionConditionL1_r19 != nil || ie.Cltm_ExecutionConditionL3_r19 != nil || ie.Ltm_EventMeasAndReport_r19 != nil || ie.Ltm_RecoveryWithKeyUpdate_r19 != nil || ie.Ltm_MCG_SCG_AdditionOrChange_r19 != nil || ie.ThreeCarrierMeasWithoutGap_r19 != nil || ie.MultiCarrierSingleReportWithoutGap_r19 != nil || ie.SkipSSB_L1_RSRP_Meas_r19 != nil || ie.GapOccasionCancelRatioReport_r19 != nil || ie.TwoSMTC_PeriodicitiesIntraFreq_r19 != nil || ie.TwoSMTC_PeriodicitiesInterFreq_r19 != nil || ie.ReportClosestReferenceLocations_r19 != nil || ie.Ltm_SR_ConfIdInCellSwitchCommand_r19 != nil || ie.SpeedStateDetection_r19 != nil || ie.EventD2_MultiReferenceLocations_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7 || hasExtGroup8 || hasExtGroup9 || hasExtGroup10 || hasExtGroup11 || hasExtGroup12

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedGapPattern != nil, ie.Ssb_RLM != nil, ie.Ssb_AndCSI_RS_RLM != nil}); err != nil {
		return err
	}

	// 3. supportedGapPattern: bit-string
	{
		if ie.SupportedGapPattern != nil {
			if err := e.EncodeBitString(*ie.SupportedGapPattern, measAndMobParametersCommonSupportedGapPatternConstraints); err != nil {
				return err
			}
		}
	}

	// 4. ssb-RLM: enumerated
	{
		if ie.Ssb_RLM != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_RLM, measAndMobParametersCommonSsbRLMConstraints); err != nil {
				return err
			}
		}
	}

	// 5. ssb-AndCSI-RS-RLM: enumerated
	{
		if ie.Ssb_AndCSI_RS_RLM != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_AndCSI_RS_RLM, measAndMobParametersCommonSsbAndCSIRSRLMConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7, hasExtGroup8, hasExtGroup9, hasExtGroup10, hasExtGroup11, hasExtGroup12}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "eventB-MeasAndReport", Optional: true},
					{Name: "handoverFDD-TDD", Optional: true},
					{Name: "eutra-CGI-Reporting", Optional: true},
					{Name: "nr-CGI-Reporting", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EventB_MeasAndReport != nil, ie.HandoverFDD_TDD != nil, ie.Eutra_CGI_Reporting != nil, ie.Nr_CGI_Reporting != nil}); err != nil {
				return err
			}

			if ie.EventB_MeasAndReport != nil {
				if err := ex.EncodeEnumerated(*ie.EventB_MeasAndReport, measAndMobParametersCommonExtEventBMeasAndReportConstraints); err != nil {
					return err
				}
			}

			if ie.HandoverFDD_TDD != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverFDD_TDD, measAndMobParametersCommonExtHandoverFDDTDDConstraints); err != nil {
					return err
				}
			}

			if ie.Eutra_CGI_Reporting != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_CGI_Reporting, measAndMobParametersCommonExtEutraCGIReportingConstraints); err != nil {
					return err
				}
			}

			if ie.Nr_CGI_Reporting != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_CGI_Reporting, measAndMobParametersCommonExtNrCGIReportingConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "independentGapConfig", Optional: true},
					{Name: "periodicEUTRA-MeasAndReport", Optional: true},
					{Name: "handoverFR1-FR2", Optional: true},
					{Name: "maxNumberCSI-RS-RRM-RS-SINR", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.IndependentGapConfig != nil, ie.PeriodicEUTRA_MeasAndReport != nil, ie.HandoverFR1_FR2 != nil, ie.MaxNumberCSI_RS_RRM_RS_SINR != nil}); err != nil {
				return err
			}

			if ie.IndependentGapConfig != nil {
				if err := ex.EncodeEnumerated(*ie.IndependentGapConfig, measAndMobParametersCommonExtIndependentGapConfigConstraints); err != nil {
					return err
				}
			}

			if ie.PeriodicEUTRA_MeasAndReport != nil {
				if err := ex.EncodeEnumerated(*ie.PeriodicEUTRA_MeasAndReport, measAndMobParametersCommonExtPeriodicEUTRAMeasAndReportConstraints); err != nil {
					return err
				}
			}

			if ie.HandoverFR1_FR2 != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverFR1_FR2, measAndMobParametersCommonExtHandoverFR1FR2Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberCSI_RS_RRM_RS_SINR != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberCSI_RS_RRM_RS_SINR, measAndMobParametersCommonExtMaxNumberCSIRSRRMRSSINRConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "nr-CGI-Reporting-ENDC", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Nr_CGI_Reporting_ENDC != nil}); err != nil {
				return err
			}

			if ie.Nr_CGI_Reporting_ENDC != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_CGI_Reporting_ENDC, measAndMobParametersCommonExtNrCGIReportingENDCConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "eutra-CGI-Reporting-NEDC", Optional: true},
					{Name: "eutra-CGI-Reporting-NRDC", Optional: true},
					{Name: "nr-CGI-Reporting-NEDC", Optional: true},
					{Name: "nr-CGI-Reporting-NRDC", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Eutra_CGI_Reporting_NEDC != nil, ie.Eutra_CGI_Reporting_NRDC != nil, ie.Nr_CGI_Reporting_NEDC != nil, ie.Nr_CGI_Reporting_NRDC != nil}); err != nil {
				return err
			}

			if ie.Eutra_CGI_Reporting_NEDC != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_CGI_Reporting_NEDC, measAndMobParametersCommonExtEutraCGIReportingNEDCConstraints); err != nil {
					return err
				}
			}

			if ie.Eutra_CGI_Reporting_NRDC != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_CGI_Reporting_NRDC, measAndMobParametersCommonExtEutraCGIReportingNRDCConstraints); err != nil {
					return err
				}
			}

			if ie.Nr_CGI_Reporting_NEDC != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_CGI_Reporting_NEDC, measAndMobParametersCommonExtNrCGIReportingNEDCConstraints); err != nil {
					return err
				}
			}

			if ie.Nr_CGI_Reporting_NRDC != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_CGI_Reporting_NRDC, measAndMobParametersCommonExtNrCGIReportingNRDCConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "reportAddNeighMeasForPeriodic-r16", Optional: true},
					{Name: "condHandoverParametersCommon-r16", Optional: true},
					{Name: "nr-NeedForGap-Reporting-r16", Optional: true},
					{Name: "supportedGapPattern-NRonly-r16", Optional: true},
					{Name: "supportedGapPattern-NRonly-NEDC-r16", Optional: true},
					{Name: "maxNumberCLI-RSSI-r16", Optional: true},
					{Name: "maxNumberCLI-SRS-RSRP-r16", Optional: true},
					{Name: "maxNumberPerSlotCLI-SRS-RSRP-r16", Optional: true},
					{Name: "mfbi-IAB-r16", Optional: true},
					{Name: "dummy", Optional: true},
					{Name: "nr-CGI-Reporting-NPN-r16", Optional: true},
					{Name: "idleInactiveEUTRA-MeasReport-r16", Optional: true},
					{Name: "idleInactive-ValidityArea-r16", Optional: true},
					{Name: "eutra-AutonomousGaps-r16", Optional: true},
					{Name: "eutra-AutonomousGaps-NEDC-r16", Optional: true},
					{Name: "eutra-AutonomousGaps-NRDC-r16", Optional: true},
					{Name: "pcellT312-r16", Optional: true},
					{Name: "supportedGapPattern-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportAddNeighMeasForPeriodic_r16 != nil, ie.CondHandoverParametersCommon_r16 != nil, ie.Nr_NeedForGap_Reporting_r16 != nil, ie.SupportedGapPattern_NRonly_r16 != nil, ie.SupportedGapPattern_NRonly_NEDC_r16 != nil, ie.MaxNumberCLI_RSSI_r16 != nil, ie.MaxNumberCLI_SRS_RSRP_r16 != nil, ie.MaxNumberPerSlotCLI_SRS_RSRP_r16 != nil, ie.Mfbi_IAB_r16 != nil, ie.Dummy != nil, ie.Nr_CGI_Reporting_NPN_r16 != nil, ie.IdleInactiveEUTRA_MeasReport_r16 != nil, ie.IdleInactive_ValidityArea_r16 != nil, ie.Eutra_AutonomousGaps_r16 != nil, ie.Eutra_AutonomousGaps_NEDC_r16 != nil, ie.Eutra_AutonomousGaps_NRDC_r16 != nil, ie.PcellT312_r16 != nil, ie.SupportedGapPattern_r16 != nil}); err != nil {
				return err
			}

			if ie.ReportAddNeighMeasForPeriodic_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ReportAddNeighMeasForPeriodic_r16, measAndMobParametersCommonExtReportAddNeighMeasForPeriodicR16Constraints); err != nil {
					return err
				}
			}

			if ie.CondHandoverParametersCommon_r16 != nil {
				c := ie.CondHandoverParametersCommon_r16
				measAndMobParametersCommonExtCondHandoverParametersCommonR16Seq := ex.NewSequenceEncoder(measAndMobParametersCommonExtCondHandoverParametersCommonR16Constraints)
				if err := measAndMobParametersCommonExtCondHandoverParametersCommonR16Seq.EncodePreamble([]bool{c.CondHandoverFDD_TDD_r16 != nil, c.CondHandoverFR1_FR2_r16 != nil}); err != nil {
					return err
				}
				if c.CondHandoverFDD_TDD_r16 != nil {
					if err := ex.EncodeEnumerated((*c.CondHandoverFDD_TDD_r16), measAndMobParametersCommonExtCondHandoverParametersCommonR16CondHandoverFDDTDDR16Constraints); err != nil {
						return err
					}
				}
				if c.CondHandoverFR1_FR2_r16 != nil {
					if err := ex.EncodeEnumerated((*c.CondHandoverFR1_FR2_r16), measAndMobParametersCommonExtCondHandoverParametersCommonR16CondHandoverFR1FR2R16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Nr_NeedForGap_Reporting_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_NeedForGap_Reporting_r16, measAndMobParametersCommonExtNrNeedForGapReportingR16Constraints); err != nil {
					return err
				}
			}

			if ie.SupportedGapPattern_NRonly_r16 != nil {
				if err := ex.EncodeBitString(*ie.SupportedGapPattern_NRonly_r16, measAndMobParametersCommonSupportedGapPatternNRonlyR16Constraints); err != nil {
					return err
				}
			}

			if ie.SupportedGapPattern_NRonly_NEDC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportedGapPattern_NRonly_NEDC_r16, measAndMobParametersCommonExtSupportedGapPatternNRonlyNEDCR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberCLI_RSSI_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberCLI_RSSI_r16, measAndMobParametersCommonExtMaxNumberCLIRSSIR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberCLI_SRS_RSRP_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberCLI_SRS_RSRP_r16, measAndMobParametersCommonExtMaxNumberCLISRSRSRPR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberPerSlotCLI_SRS_RSRP_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberPerSlotCLI_SRS_RSRP_r16, measAndMobParametersCommonExtMaxNumberPerSlotCLISRSRSRPR16Constraints); err != nil {
					return err
				}
			}

			if ie.Mfbi_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Mfbi_IAB_r16, measAndMobParametersCommonExtMfbiIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.Dummy != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy, measAndMobParametersCommonExtDummyConstraints); err != nil {
					return err
				}
			}

			if ie.Nr_CGI_Reporting_NPN_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_CGI_Reporting_NPN_r16, measAndMobParametersCommonExtNrCGIReportingNPNR16Constraints); err != nil {
					return err
				}
			}

			if ie.IdleInactiveEUTRA_MeasReport_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.IdleInactiveEUTRA_MeasReport_r16, measAndMobParametersCommonExtIdleInactiveEUTRAMeasReportR16Constraints); err != nil {
					return err
				}
			}

			if ie.IdleInactive_ValidityArea_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.IdleInactive_ValidityArea_r16, measAndMobParametersCommonExtIdleInactiveValidityAreaR16Constraints); err != nil {
					return err
				}
			}

			if ie.Eutra_AutonomousGaps_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_AutonomousGaps_r16, measAndMobParametersCommonExtEutraAutonomousGapsR16Constraints); err != nil {
					return err
				}
			}

			if ie.Eutra_AutonomousGaps_NEDC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_AutonomousGaps_NEDC_r16, measAndMobParametersCommonExtEutraAutonomousGapsNEDCR16Constraints); err != nil {
					return err
				}
			}

			if ie.Eutra_AutonomousGaps_NRDC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_AutonomousGaps_NRDC_r16, measAndMobParametersCommonExtEutraAutonomousGapsNRDCR16Constraints); err != nil {
					return err
				}
			}

			if ie.PcellT312_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.PcellT312_r16, measAndMobParametersCommonExtPcellT312R16Constraints); err != nil {
					return err
				}
			}

			if ie.SupportedGapPattern_r16 != nil {
				if err := ex.EncodeBitString(*ie.SupportedGapPattern_r16, measAndMobParametersCommonSupportedGapPatternR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 5:
		if hasExtGroup5 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "concurrentMeasGap-r17", Optional: true},
					{Name: "nr-NeedForGapNCSG-Reporting-r17", Optional: true},
					{Name: "eutra-NeedForGapNCSG-Reporting-r17", Optional: true},
					{Name: "ncsg-MeasGapPerFR-r17", Optional: true},
					{Name: "ncsg-MeasGapPatterns-r17", Optional: true},
					{Name: "ncsg-MeasGapNR-Patterns-r17", Optional: true},
					{Name: "preconfiguredUE-AutonomousMeasGap-r17", Optional: true},
					{Name: "preconfiguredNW-ControlledMeasGap-r17", Optional: true},
					{Name: "handoverFR1-FR2-2-r17", Optional: true},
					{Name: "handoverFR2-1-FR2-2-r17", Optional: true},
					{Name: "independentGapConfigPRS-r17", Optional: true},
					{Name: "rrm-RelaxationRRC-ConnectedRedCap-r17", Optional: true},
					{Name: "parallelMeasurementGap-r17", Optional: true},
					{Name: "condHandoverWithSCG-NRDC-r17", Optional: true},
					{Name: "gNB-ID-LengthReporting-r17", Optional: true},
					{Name: "gNB-ID-LengthReporting-ENDC-r17", Optional: true},
					{Name: "gNB-ID-LengthReporting-NEDC-r17", Optional: true},
					{Name: "gNB-ID-LengthReporting-NRDC-r17", Optional: true},
					{Name: "gNB-ID-LengthReporting-NPN-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ConcurrentMeasGap_r17 != nil, ie.Nr_NeedForGapNCSG_Reporting_r17 != nil, ie.Eutra_NeedForGapNCSG_Reporting_r17 != nil, ie.Ncsg_MeasGapPerFR_r17 != nil, ie.Ncsg_MeasGapPatterns_r17 != nil, ie.Ncsg_MeasGapNR_Patterns_r17 != nil, ie.PreconfiguredUE_AutonomousMeasGap_r17 != nil, ie.PreconfiguredNW_ControlledMeasGap_r17 != nil, ie.HandoverFR1_FR2_2_r17 != nil, ie.HandoverFR2_1_FR2_2_r17 != nil, ie.IndependentGapConfigPRS_r17 != nil, ie.Rrm_RelaxationRRC_ConnectedRedCap_r17 != nil, ie.ParallelMeasurementGap_r17 != nil, ie.CondHandoverWithSCG_NRDC_r17 != nil, ie.GNB_ID_LengthReporting_r17 != nil, ie.GNB_ID_LengthReporting_ENDC_r17 != nil, ie.GNB_ID_LengthReporting_NEDC_r17 != nil, ie.GNB_ID_LengthReporting_NRDC_r17 != nil, ie.GNB_ID_LengthReporting_NPN_r17 != nil}); err != nil {
				return err
			}

			if ie.ConcurrentMeasGap_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(measAndMobParametersCommonExtConcurrentMeasGapR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ConcurrentMeasGap_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ConcurrentMeasGap_r17).Choice {
				case MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_OnlyMeasGap_r17:
					if err := ex.EncodeEnumerated((*(*ie.ConcurrentMeasGap_r17).ConcurrentPerUE_OnlyMeasGap_r17), measAndMobParametersCommonExtConcurrentMeasGapR17ConcurrentPerUEOnlyMeasGapR17Constraints); err != nil {
						return err
					}
				case MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_PerFRCombMeasGap_r17:
					if err := ex.EncodeEnumerated((*(*ie.ConcurrentMeasGap_r17).ConcurrentPerUE_PerFRCombMeasGap_r17), measAndMobParametersCommonExtConcurrentMeasGapR17ConcurrentPerUEPerFRCombMeasGapR17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Nr_NeedForGapNCSG_Reporting_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_NeedForGapNCSG_Reporting_r17, measAndMobParametersCommonExtNrNeedForGapNCSGReportingR17Constraints); err != nil {
					return err
				}
			}

			if ie.Eutra_NeedForGapNCSG_Reporting_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_NeedForGapNCSG_Reporting_r17, measAndMobParametersCommonExtEutraNeedForGapNCSGReportingR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ncsg_MeasGapPerFR_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ncsg_MeasGapPerFR_r17, measAndMobParametersCommonExtNcsgMeasGapPerFRR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ncsg_MeasGapPatterns_r17 != nil {
				if err := ex.EncodeBitString(*ie.Ncsg_MeasGapPatterns_r17, measAndMobParametersCommonNcsgMeasGapPatternsR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ncsg_MeasGapNR_Patterns_r17 != nil {
				if err := ex.EncodeBitString(*ie.Ncsg_MeasGapNR_Patterns_r17, measAndMobParametersCommonNcsgMeasGapNRPatternsR17Constraints); err != nil {
					return err
				}
			}

			if ie.PreconfiguredUE_AutonomousMeasGap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PreconfiguredUE_AutonomousMeasGap_r17, measAndMobParametersCommonExtPreconfiguredUEAutonomousMeasGapR17Constraints); err != nil {
					return err
				}
			}

			if ie.PreconfiguredNW_ControlledMeasGap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PreconfiguredNW_ControlledMeasGap_r17, measAndMobParametersCommonExtPreconfiguredNWControlledMeasGapR17Constraints); err != nil {
					return err
				}
			}

			if ie.HandoverFR1_FR2_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverFR1_FR2_2_r17, measAndMobParametersCommonExtHandoverFR1FR22R17Constraints); err != nil {
					return err
				}
			}

			if ie.HandoverFR2_1_FR2_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverFR2_1_FR2_2_r17, measAndMobParametersCommonExtHandoverFR21FR22R17Constraints); err != nil {
					return err
				}
			}

			if ie.IndependentGapConfigPRS_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.IndependentGapConfigPRS_r17, measAndMobParametersCommonExtIndependentGapConfigPRSR17Constraints); err != nil {
					return err
				}
			}

			if ie.Rrm_RelaxationRRC_ConnectedRedCap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Rrm_RelaxationRRC_ConnectedRedCap_r17, measAndMobParametersCommonExtRrmRelaxationRRCConnectedRedCapR17Constraints); err != nil {
					return err
				}
			}

			if ie.ParallelMeasurementGap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ParallelMeasurementGap_r17, measAndMobParametersCommonExtParallelMeasurementGapR17Constraints); err != nil {
					return err
				}
			}

			if ie.CondHandoverWithSCG_NRDC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.CondHandoverWithSCG_NRDC_r17, measAndMobParametersCommonExtCondHandoverWithSCGNRDCR17Constraints); err != nil {
					return err
				}
			}

			if ie.GNB_ID_LengthReporting_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.GNB_ID_LengthReporting_r17, measAndMobParametersCommonExtGNBIDLengthReportingR17Constraints); err != nil {
					return err
				}
			}

			if ie.GNB_ID_LengthReporting_ENDC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.GNB_ID_LengthReporting_ENDC_r17, measAndMobParametersCommonExtGNBIDLengthReportingENDCR17Constraints); err != nil {
					return err
				}
			}

			if ie.GNB_ID_LengthReporting_NEDC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.GNB_ID_LengthReporting_NEDC_r17, measAndMobParametersCommonExtGNBIDLengthReportingNEDCR17Constraints); err != nil {
					return err
				}
			}

			if ie.GNB_ID_LengthReporting_NRDC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.GNB_ID_LengthReporting_NRDC_r17, measAndMobParametersCommonExtGNBIDLengthReportingNRDCR17Constraints); err != nil {
					return err
				}
			}

			if ie.GNB_ID_LengthReporting_NPN_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.GNB_ID_LengthReporting_NPN_r17, measAndMobParametersCommonExtGNBIDLengthReportingNPNR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 6:
		if hasExtGroup6 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "parallelSMTC-r17", Optional: true},
					{Name: "concurrentMeasGapEUTRA-r17", Optional: true},
					{Name: "serviceLinkPropDelayDiffReporting-r17", Optional: true},
					{Name: "ncsg-SymbolLevelScheduleRestrictionInter-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ParallelSMTC_r17 != nil, ie.ConcurrentMeasGapEUTRA_r17 != nil, ie.ServiceLinkPropDelayDiffReporting_r17 != nil, ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil}); err != nil {
				return err
			}

			if ie.ParallelSMTC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ParallelSMTC_r17, measAndMobParametersCommonExtParallelSMTCR17Constraints); err != nil {
					return err
				}
			}

			if ie.ConcurrentMeasGapEUTRA_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ConcurrentMeasGapEUTRA_r17, measAndMobParametersCommonExtConcurrentMeasGapEUTRAR17Constraints); err != nil {
					return err
				}
			}

			if ie.ServiceLinkPropDelayDiffReporting_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ServiceLinkPropDelayDiffReporting_r17, measAndMobParametersCommonExtServiceLinkPropDelayDiffReportingR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17, measAndMobParametersCommonExtNcsgSymbolLevelScheduleRestrictionInterR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 7:
		if hasExtGroup7 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "eventD1-MeasReportTrigger-r17", Optional: true},
					{Name: "independentGapConfig-maxCC-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EventD1_MeasReportTrigger_r17 != nil, ie.IndependentGapConfig_MaxCC_r17 != nil}); err != nil {
				return err
			}

			if ie.EventD1_MeasReportTrigger_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.EventD1_MeasReportTrigger_r17, measAndMobParametersCommonExtEventD1MeasReportTriggerR17Constraints); err != nil {
					return err
				}
			}

			if ie.IndependentGapConfig_MaxCC_r17 != nil {
				c := ie.IndependentGapConfig_MaxCC_r17
				measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Seq := ex.NewSequenceEncoder(measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Constraints)
				if err := measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Seq.EncodePreamble([]bool{c.Fr1_Only_r17 != nil, c.Fr2_Only_r17 != nil, c.Fr1_AndFR2_r17 != nil}); err != nil {
					return err
				}
				if c.Fr1_Only_r17 != nil {
					if err := ex.EncodeInteger((*c.Fr1_Only_r17), per.Constrained(1, 32)); err != nil {
						return err
					}
				}
				if c.Fr2_Only_r17 != nil {
					if err := ex.EncodeInteger((*c.Fr2_Only_r17), per.Constrained(1, 32)); err != nil {
						return err
					}
				}
				if c.Fr1_AndFR2_r17 != nil {
					if err := ex.EncodeInteger((*c.Fr1_AndFR2_r17), per.Constrained(1, 32)); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 8:
		if hasExtGroup8 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "interSatMeas-r17", Optional: true},
					{Name: "deriveSSB-IndexFromCellInterNon-NCSG-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InterSatMeas_r17 != nil, ie.DeriveSSB_IndexFromCellInterNon_NCSG_r17 != nil}); err != nil {
				return err
			}

			if ie.InterSatMeas_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.InterSatMeas_r17, measAndMobParametersCommonExtInterSatMeasR17Constraints); err != nil {
					return err
				}
			}

			if ie.DeriveSSB_IndexFromCellInterNon_NCSG_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.DeriveSSB_IndexFromCellInterNon_NCSG_r17, measAndMobParametersCommonExtDeriveSSBIndexFromCellInterNonNCSGR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 9:
		if hasExtGroup9 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "l3-MeasUnknownSCellActivation-r18", Optional: true},
					{Name: "shortMeasInterval-r18", Optional: true},
					{Name: "nr-NeedForInterruptionReport-r18", Optional: true},
					{Name: "measSequenceConfig-r18", Optional: true},
					{Name: "cellIndividualOffsetPerMeasEvent-r18", Optional: true},
					{Name: "eventD2-MeasReportTrigger-r18", Optional: true},
					{Name: "concurrentMeasGapsPreMG-r18", Optional: true},
					{Name: "dynamicCollision-r18", Optional: true},
					{Name: "concurrentMeasGapsNCSG-r18", Optional: true},
					{Name: "eutra-NoGapMeasurementOutsideBWP-r18", Optional: true},
					{Name: "eutra-NoGapMeasurementInsideBWP-r18", Optional: true},
					{Name: "eutra-MeasEMW-r18", Optional: true},
					{Name: "concurrentMeasCRS-InsideBWP-EUTRA-r18", Optional: true},
					{Name: "ltm-InterFreqMeasGap-r18", Optional: true},
					{Name: "dummy-ltm-FastUE-Processing-r18", Optional: true},
					{Name: "rach-LessHandoverInterFreq-r18", Optional: true},
					{Name: "enterAndLeaveCellReport-r18", Optional: true},
					{Name: "bestCellChangeReport-r18", Optional: true},
					{Name: "secondBestCellChangeReport-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.L3_MeasUnknownSCellActivation_r18 != nil, ie.ShortMeasInterval_r18 != nil, ie.Nr_NeedForInterruptionReport_r18 != nil, ie.MeasSequenceConfig_r18 != nil, ie.CellIndividualOffsetPerMeasEvent_r18 != nil, ie.EventD2_MeasReportTrigger_r18 != nil, ie.ConcurrentMeasGapsPreMG_r18 != nil, ie.DynamicCollision_r18 != nil, ie.ConcurrentMeasGapsNCSG_r18 != nil, ie.Eutra_NoGapMeasurementOutsideBWP_r18 != nil, ie.Eutra_NoGapMeasurementInsideBWP_r18 != nil, ie.Eutra_MeasEMW_r18 != nil, ie.ConcurrentMeasCRS_InsideBWP_EUTRA_r18 != nil, ie.Ltm_InterFreqMeasGap_r18 != nil, ie.Dummy_Ltm_FastUE_Processing_r18 != nil, ie.Rach_LessHandoverInterFreq_r18 != nil, ie.EnterAndLeaveCellReport_r18 != nil, ie.BestCellChangeReport_r18 != nil, ie.SecondBestCellChangeReport_r18 != nil}); err != nil {
				return err
			}

			if ie.L3_MeasUnknownSCellActivation_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.L3_MeasUnknownSCellActivation_r18, measAndMobParametersCommonExtL3MeasUnknownSCellActivationR18Constraints); err != nil {
					return err
				}
			}

			if ie.ShortMeasInterval_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ShortMeasInterval_r18, measAndMobParametersCommonExtShortMeasIntervalR18Constraints); err != nil {
					return err
				}
			}

			if ie.Nr_NeedForInterruptionReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_NeedForInterruptionReport_r18, measAndMobParametersCommonExtNrNeedForInterruptionReportR18Constraints); err != nil {
					return err
				}
			}

			if ie.MeasSequenceConfig_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MeasSequenceConfig_r18, measAndMobParametersCommonExtMeasSequenceConfigR18Constraints); err != nil {
					return err
				}
			}

			if ie.CellIndividualOffsetPerMeasEvent_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.CellIndividualOffsetPerMeasEvent_r18, measAndMobParametersCommonExtCellIndividualOffsetPerMeasEventR18Constraints); err != nil {
					return err
				}
			}

			if ie.EventD2_MeasReportTrigger_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EventD2_MeasReportTrigger_r18, measAndMobParametersCommonExtEventD2MeasReportTriggerR18Constraints); err != nil {
					return err
				}
			}

			if ie.ConcurrentMeasGapsPreMG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ConcurrentMeasGapsPreMG_r18, measAndMobParametersCommonExtConcurrentMeasGapsPreMGR18Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicCollision_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicCollision_r18, measAndMobParametersCommonExtDynamicCollisionR18Constraints); err != nil {
					return err
				}
			}

			if ie.ConcurrentMeasGapsNCSG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ConcurrentMeasGapsNCSG_r18, measAndMobParametersCommonExtConcurrentMeasGapsNCSGR18Constraints); err != nil {
					return err
				}
			}

			if ie.Eutra_NoGapMeasurementOutsideBWP_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_NoGapMeasurementOutsideBWP_r18, measAndMobParametersCommonExtEutraNoGapMeasurementOutsideBWPR18Constraints); err != nil {
					return err
				}
			}

			if ie.Eutra_NoGapMeasurementInsideBWP_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_NoGapMeasurementInsideBWP_r18, measAndMobParametersCommonExtEutraNoGapMeasurementInsideBWPR18Constraints); err != nil {
					return err
				}
			}

			if ie.Eutra_MeasEMW_r18 != nil {
				if err := ex.EncodeBitString(*ie.Eutra_MeasEMW_r18, measAndMobParametersCommonEutraMeasEMWR18Constraints); err != nil {
					return err
				}
			}

			if ie.ConcurrentMeasCRS_InsideBWP_EUTRA_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ConcurrentMeasCRS_InsideBWP_EUTRA_r18, measAndMobParametersCommonExtConcurrentMeasCRSInsideBWPEUTRAR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_InterFreqMeasGap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_InterFreqMeasGap_r18, measAndMobParametersCommonExtLtmInterFreqMeasGapR18Constraints); err != nil {
					return err
				}
			}

			if ie.Dummy_Ltm_FastUE_Processing_r18 != nil {
				c := ie.Dummy_Ltm_FastUE_Processing_r18
				if err := ex.EncodeEnumerated(c.Fr1_r18, measAndMobParametersCommonExtDummyLtmFastUEProcessingR18Fr1R18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.Fr2_r18, measAndMobParametersCommonExtDummyLtmFastUEProcessingR18Fr2R18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.Fr1_AndFR2_r18, measAndMobParametersCommonExtDummyLtmFastUEProcessingR18Fr1AndFR2R18Constraints); err != nil {
					return err
				}
			}

			if ie.Rach_LessHandoverInterFreq_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Rach_LessHandoverInterFreq_r18, measAndMobParametersCommonExtRachLessHandoverInterFreqR18Constraints); err != nil {
					return err
				}
			}

			if ie.EnterAndLeaveCellReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EnterAndLeaveCellReport_r18, measAndMobParametersCommonExtEnterAndLeaveCellReportR18Constraints); err != nil {
					return err
				}
			}

			if ie.BestCellChangeReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.BestCellChangeReport_r18, measAndMobParametersCommonExtBestCellChangeReportR18Constraints); err != nil {
					return err
				}
			}

			if ie.SecondBestCellChangeReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SecondBestCellChangeReport_r18, measAndMobParametersCommonExtSecondBestCellChangeReportR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 10:
		if hasExtGroup10 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ltm-InterFreq-r18", Optional: true},
					{Name: "ltm-MCG-NRDC-r18", Optional: true},
					{Name: "ltm-RACH-LessDG-r18", Optional: true},
					{Name: "ltm-RACH-LessCG-r18", Optional: true},
					{Name: "ltm-Recovery-r18", Optional: true},
					{Name: "ltm-ReferenceConfig-r18", Optional: true},
					{Name: "ltm-MCG-NRDC-Release-r18", Optional: true},
					{Name: "ltm-FastUE-Processing-r18", Optional: true},
					{Name: "ntn-NeighbourCellInfoSupport-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_InterFreq_r18 != nil, ie.Ltm_MCG_NRDC_r18 != nil, ie.Ltm_RACH_LessDG_r18 != nil, ie.Ltm_RACH_LessCG_r18 != nil, ie.Ltm_Recovery_r18 != nil, ie.Ltm_ReferenceConfig_r18 != nil, ie.Ltm_MCG_NRDC_Release_r18 != nil, ie.Ltm_FastUE_Processing_r18 != nil, ie.Ntn_NeighbourCellInfoSupport_r18 != nil}); err != nil {
				return err
			}

			if ie.Ltm_InterFreq_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_InterFreq_r18, measAndMobParametersCommonExtLtmInterFreqR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_MCG_NRDC_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_MCG_NRDC_r18, measAndMobParametersCommonExtLtmMCGNRDCR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_RACH_LessDG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_RACH_LessDG_r18, measAndMobParametersCommonExtLtmRACHLessDGR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_RACH_LessCG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_RACH_LessCG_r18, measAndMobParametersCommonExtLtmRACHLessCGR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_Recovery_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_Recovery_r18, measAndMobParametersCommonExtLtmRecoveryR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_ReferenceConfig_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_ReferenceConfig_r18, measAndMobParametersCommonExtLtmReferenceConfigR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_MCG_NRDC_Release_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_MCG_NRDC_Release_r18, measAndMobParametersCommonExtLtmMCGNRDCReleaseR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_FastUE_Processing_r18 != nil {
				c := ie.Ltm_FastUE_Processing_r18
				measAndMobParametersCommonExtLtmFastUEProcessingR18Seq := ex.NewSequenceEncoder(measAndMobParametersCommonExtLtmFastUEProcessingR18Constraints)
				if err := measAndMobParametersCommonExtLtmFastUEProcessingR18Seq.EncodePreamble([]bool{c.Fr1_r18 != nil, c.Fr2_r18 != nil, c.Fr1_AndFR2_r18 != nil}); err != nil {
					return err
				}
				if c.Fr1_r18 != nil {
					if err := ex.EncodeEnumerated((*c.Fr1_r18), measAndMobParametersCommonExtLtmFastUEProcessingR18Fr1R18Constraints); err != nil {
						return err
					}
				}
				if c.Fr2_r18 != nil {
					if err := ex.EncodeEnumerated((*c.Fr2_r18), measAndMobParametersCommonExtLtmFastUEProcessingR18Fr2R18Constraints); err != nil {
						return err
					}
				}
				if c.Fr1_AndFR2_r18 != nil {
					if err := ex.EncodeEnumerated((*c.Fr1_AndFR2_r18), measAndMobParametersCommonExtLtmFastUEProcessingR18Fr1AndFR2R18Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Ntn_NeighbourCellInfoSupport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ntn_NeighbourCellInfoSupport_r18, measAndMobParametersCommonExtNtnNeighbourCellInfoSupportR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 11:
		if hasExtGroup11 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ltm-interFreqL1-OnlyInBC-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_InterFreqL1_OnlyInBC_r18 != nil}); err != nil {
				return err
			}

			if ie.Ltm_InterFreqL1_OnlyInBC_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_InterFreqL1_OnlyInBC_r18, measAndMobParametersCommonExtLtmInterFreqL1OnlyInBCR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 12:
		if hasExtGroup12 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "nr-CGI-Reporting-HSDN-r19", Optional: true},
					{Name: "eutra-CGI-Reporting-HSDN-r19", Optional: true},
					{Name: "ltm-KeyUpdateMCG-r19", Optional: true},
					{Name: "ltm-KeyUpdateSCG-r19", Optional: true},
					{Name: "cltm-EarlyTA-Indication-r19", Optional: true},
					{Name: "cltm-ExecutionConditionL1-r19", Optional: true},
					{Name: "cltm-ExecutionConditionL3-r19", Optional: true},
					{Name: "ltm-EventMeasAndReport-r19", Optional: true},
					{Name: "ltm-RecoveryWithKeyUpdate-r19", Optional: true},
					{Name: "ltm-MCG-SCG-AdditionOrChange-r19", Optional: true},
					{Name: "threeCarrierMeasWithoutGap-r19", Optional: true},
					{Name: "multiCarrierSingleReportWithoutGap-r19", Optional: true},
					{Name: "skipSSB-L1-RSRP-Meas-r19", Optional: true},
					{Name: "gapOccasionCancelRatioReport-r19", Optional: true},
					{Name: "twoSMTC-PeriodicitiesIntraFreq-r19", Optional: true},
					{Name: "twoSMTC-PeriodicitiesInterFreq-r19", Optional: true},
					{Name: "reportClosestReferenceLocations-r19", Optional: true},
					{Name: "ltm-SR-ConfIdInCellSwitchCommand-r19", Optional: true},
					{Name: "speedStateDetection-r19", Optional: true},
					{Name: "eventD2-MultiReferenceLocations-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Nr_CGI_Reporting_HSDN_r19 != nil, ie.Eutra_CGI_Reporting_HSDN_r19 != nil, ie.Ltm_KeyUpdateMCG_r19 != nil, ie.Ltm_KeyUpdateSCG_r19 != nil, ie.Cltm_EarlyTA_Indication_r19 != nil, ie.Cltm_ExecutionConditionL1_r19 != nil, ie.Cltm_ExecutionConditionL3_r19 != nil, ie.Ltm_EventMeasAndReport_r19 != nil, ie.Ltm_RecoveryWithKeyUpdate_r19 != nil, ie.Ltm_MCG_SCG_AdditionOrChange_r19 != nil, ie.ThreeCarrierMeasWithoutGap_r19 != nil, ie.MultiCarrierSingleReportWithoutGap_r19 != nil, ie.SkipSSB_L1_RSRP_Meas_r19 != nil, ie.GapOccasionCancelRatioReport_r19 != nil, ie.TwoSMTC_PeriodicitiesIntraFreq_r19 != nil, ie.TwoSMTC_PeriodicitiesInterFreq_r19 != nil, ie.ReportClosestReferenceLocations_r19 != nil, ie.Ltm_SR_ConfIdInCellSwitchCommand_r19 != nil, ie.SpeedStateDetection_r19 != nil, ie.EventD2_MultiReferenceLocations_r19 != nil}); err != nil {
				return err
			}

			if ie.Nr_CGI_Reporting_HSDN_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_CGI_Reporting_HSDN_r19, measAndMobParametersCommonExtNrCGIReportingHSDNR19Constraints); err != nil {
					return err
				}
			}

			if ie.Eutra_CGI_Reporting_HSDN_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Eutra_CGI_Reporting_HSDN_r19, measAndMobParametersCommonExtEutraCGIReportingHSDNR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_KeyUpdateMCG_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_KeyUpdateMCG_r19, measAndMobParametersCommonExtLtmKeyUpdateMCGR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_KeyUpdateSCG_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_KeyUpdateSCG_r19, measAndMobParametersCommonExtLtmKeyUpdateSCGR19Constraints); err != nil {
					return err
				}
			}

			if ie.Cltm_EarlyTA_Indication_r19 != nil {
				if err := ex.EncodeInteger(*ie.Cltm_EarlyTA_Indication_r19, measAndMobParametersCommonCltmEarlyTAIndicationR19Constraints); err != nil {
					return err
				}
			}

			if ie.Cltm_ExecutionConditionL1_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Cltm_ExecutionConditionL1_r19, measAndMobParametersCommonExtCltmExecutionConditionL1R19Constraints); err != nil {
					return err
				}
			}

			if ie.Cltm_ExecutionConditionL3_r19 != nil {
				if err := ex.EncodeInteger(*ie.Cltm_ExecutionConditionL3_r19, measAndMobParametersCommonCltmExecutionConditionL3R19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_EventMeasAndReport_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_EventMeasAndReport_r19, measAndMobParametersCommonExtLtmEventMeasAndReportR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_RecoveryWithKeyUpdate_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_RecoveryWithKeyUpdate_r19, measAndMobParametersCommonExtLtmRecoveryWithKeyUpdateR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_MCG_SCG_AdditionOrChange_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_MCG_SCG_AdditionOrChange_r19, measAndMobParametersCommonExtLtmMCGSCGAdditionOrChangeR19Constraints); err != nil {
					return err
				}
			}

			if ie.ThreeCarrierMeasWithoutGap_r19 != nil {
				c := ie.ThreeCarrierMeasWithoutGap_r19
				measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Seq := ex.NewSequenceEncoder(measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Constraints)
				if err := measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Seq.EncodePreamble([]bool{c.Fr1_CA_NR_DC_r19 != nil, c.Fr1_FR2_CA_r19 != nil, c.Fr1_FR2_NR_DC_r19 != nil}); err != nil {
					return err
				}
				if c.Fr1_CA_NR_DC_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Fr1_CA_NR_DC_r19), measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Fr1CANRDCR19Constraints); err != nil {
						return err
					}
				}
				if c.Fr1_FR2_CA_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Fr1_FR2_CA_r19), measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Fr1FR2CAR19Constraints); err != nil {
						return err
					}
				}
				if c.Fr1_FR2_NR_DC_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Fr1_FR2_NR_DC_r19), measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Fr1FR2NRDCR19Constraints); err != nil {
						return err
					}
				}
			}

			if ie.MultiCarrierSingleReportWithoutGap_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiCarrierSingleReportWithoutGap_r19, measAndMobParametersCommonExtMultiCarrierSingleReportWithoutGapR19Constraints); err != nil {
					return err
				}
			}

			if ie.SkipSSB_L1_RSRP_Meas_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.SkipSSB_L1_RSRP_Meas_r19, measAndMobParametersCommonExtSkipSSBL1RSRPMeasR19Constraints); err != nil {
					return err
				}
			}

			if ie.GapOccasionCancelRatioReport_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.GapOccasionCancelRatioReport_r19, measAndMobParametersCommonExtGapOccasionCancelRatioReportR19Constraints); err != nil {
					return err
				}
			}

			if ie.TwoSMTC_PeriodicitiesIntraFreq_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoSMTC_PeriodicitiesIntraFreq_r19, measAndMobParametersCommonExtTwoSMTCPeriodicitiesIntraFreqR19Constraints); err != nil {
					return err
				}
			}

			if ie.TwoSMTC_PeriodicitiesInterFreq_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoSMTC_PeriodicitiesInterFreq_r19, measAndMobParametersCommonExtTwoSMTCPeriodicitiesInterFreqR19Constraints); err != nil {
					return err
				}
			}

			if ie.ReportClosestReferenceLocations_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.ReportClosestReferenceLocations_r19, measAndMobParametersCommonExtReportClosestReferenceLocationsR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_SR_ConfIdInCellSwitchCommand_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_SR_ConfIdInCellSwitchCommand_r19, measAndMobParametersCommonExtLtmSRConfIdInCellSwitchCommandR19Constraints); err != nil {
					return err
				}
			}

			if ie.SpeedStateDetection_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.SpeedStateDetection_r19, measAndMobParametersCommonExtSpeedStateDetectionR19Constraints); err != nil {
					return err
				}
			}

			if ie.EventD2_MultiReferenceLocations_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.EventD2_MultiReferenceLocations_r19, measAndMobParametersCommonExtEventD2MultiReferenceLocationsR19Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasAndMobParametersCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. supportedGapPattern: bit-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeBitString(measAndMobParametersCommonSupportedGapPatternConstraints)
			if err != nil {
				return err
			}
			ie.SupportedGapPattern = &v
		}
	}

	// 4. ssb-RLM: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(measAndMobParametersCommonSsbRLMConstraints)
			if err != nil {
				return err
			}
			ie.Ssb_RLM = &idx
		}
	}

	// 5. ssb-AndCSI-RS-RLM: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(measAndMobParametersCommonSsbAndCSIRSRLMConstraints)
			if err != nil {
				return err
			}
			ie.Ssb_AndCSI_RS_RLM = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "eventB-MeasAndReport", Optional: true},
				{Name: "handoverFDD-TDD", Optional: true},
				{Name: "eutra-CGI-Reporting", Optional: true},
				{Name: "nr-CGI-Reporting", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEventBMeasAndReportConstraints)
			if err != nil {
				return err
			}
			ie.EventB_MeasAndReport = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtHandoverFDDTDDConstraints)
			if err != nil {
				return err
			}
			ie.HandoverFDD_TDD = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraCGIReportingConstraints)
			if err != nil {
				return err
			}
			ie.Eutra_CGI_Reporting = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNrCGIReportingConstraints)
			if err != nil {
				return err
			}
			ie.Nr_CGI_Reporting = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "independentGapConfig", Optional: true},
				{Name: "periodicEUTRA-MeasAndReport", Optional: true},
				{Name: "handoverFR1-FR2", Optional: true},
				{Name: "maxNumberCSI-RS-RRM-RS-SINR", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtIndependentGapConfigConstraints)
			if err != nil {
				return err
			}
			ie.IndependentGapConfig = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtPeriodicEUTRAMeasAndReportConstraints)
			if err != nil {
				return err
			}
			ie.PeriodicEUTRA_MeasAndReport = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtHandoverFR1FR2Constraints)
			if err != nil {
				return err
			}
			ie.HandoverFR1_FR2 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtMaxNumberCSIRSRRMRSSINRConstraints)
			if err != nil {
				return err
			}
			ie.MaxNumberCSI_RS_RRM_RS_SINR = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "nr-CGI-Reporting-ENDC", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNrCGIReportingENDCConstraints)
			if err != nil {
				return err
			}
			ie.Nr_CGI_Reporting_ENDC = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "eutra-CGI-Reporting-NEDC", Optional: true},
				{Name: "eutra-CGI-Reporting-NRDC", Optional: true},
				{Name: "nr-CGI-Reporting-NEDC", Optional: true},
				{Name: "nr-CGI-Reporting-NRDC", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraCGIReportingNEDCConstraints)
			if err != nil {
				return err
			}
			ie.Eutra_CGI_Reporting_NEDC = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraCGIReportingNRDCConstraints)
			if err != nil {
				return err
			}
			ie.Eutra_CGI_Reporting_NRDC = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNrCGIReportingNEDCConstraints)
			if err != nil {
				return err
			}
			ie.Nr_CGI_Reporting_NEDC = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNrCGIReportingNRDCConstraints)
			if err != nil {
				return err
			}
			ie.Nr_CGI_Reporting_NRDC = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "reportAddNeighMeasForPeriodic-r16", Optional: true},
				{Name: "condHandoverParametersCommon-r16", Optional: true},
				{Name: "nr-NeedForGap-Reporting-r16", Optional: true},
				{Name: "supportedGapPattern-NRonly-r16", Optional: true},
				{Name: "supportedGapPattern-NRonly-NEDC-r16", Optional: true},
				{Name: "maxNumberCLI-RSSI-r16", Optional: true},
				{Name: "maxNumberCLI-SRS-RSRP-r16", Optional: true},
				{Name: "maxNumberPerSlotCLI-SRS-RSRP-r16", Optional: true},
				{Name: "mfbi-IAB-r16", Optional: true},
				{Name: "dummy", Optional: true},
				{Name: "nr-CGI-Reporting-NPN-r16", Optional: true},
				{Name: "idleInactiveEUTRA-MeasReport-r16", Optional: true},
				{Name: "idleInactive-ValidityArea-r16", Optional: true},
				{Name: "eutra-AutonomousGaps-r16", Optional: true},
				{Name: "eutra-AutonomousGaps-NEDC-r16", Optional: true},
				{Name: "eutra-AutonomousGaps-NRDC-r16", Optional: true},
				{Name: "pcellT312-r16", Optional: true},
				{Name: "supportedGapPattern-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtReportAddNeighMeasForPeriodicR16Constraints)
			if err != nil {
				return err
			}
			ie.ReportAddNeighMeasForPeriodic_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CondHandoverParametersCommon_r16 = &struct {
				CondHandoverFDD_TDD_r16 *int64
				CondHandoverFR1_FR2_r16 *int64
			}{}
			c := ie.CondHandoverParametersCommon_r16
			measAndMobParametersCommonExtCondHandoverParametersCommonR16Seq := dx.NewSequenceDecoder(measAndMobParametersCommonExtCondHandoverParametersCommonR16Constraints)
			if err := measAndMobParametersCommonExtCondHandoverParametersCommonR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if measAndMobParametersCommonExtCondHandoverParametersCommonR16Seq.IsComponentPresent(0) {
				c.CondHandoverFDD_TDD_r16 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtCondHandoverParametersCommonR16CondHandoverFDDTDDR16Constraints)
				if err != nil {
					return err
				}
				(*c.CondHandoverFDD_TDD_r16) = v
			}
			if measAndMobParametersCommonExtCondHandoverParametersCommonR16Seq.IsComponentPresent(1) {
				c.CondHandoverFR1_FR2_r16 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtCondHandoverParametersCommonR16CondHandoverFR1FR2R16Constraints)
				if err != nil {
					return err
				}
				(*c.CondHandoverFR1_FR2_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNrNeedForGapReportingR16Constraints)
			if err != nil {
				return err
			}
			ie.Nr_NeedForGap_Reporting_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeBitString(measAndMobParametersCommonSupportedGapPatternNRonlyR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportedGapPattern_NRonly_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtSupportedGapPatternNRonlyNEDCR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportedGapPattern_NRonly_NEDC_r16 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtMaxNumberCLIRSSIR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberCLI_RSSI_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtMaxNumberCLISRSRSRPR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberCLI_SRS_RSRP_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtMaxNumberPerSlotCLISRSRSRPR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberPerSlotCLI_SRS_RSRP_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtMfbiIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Mfbi_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtDummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNrCGIReportingNPNR16Constraints)
			if err != nil {
				return err
			}
			ie.Nr_CGI_Reporting_NPN_r16 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtIdleInactiveEUTRAMeasReportR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleInactiveEUTRA_MeasReport_r16 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtIdleInactiveValidityAreaR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleInactive_ValidityArea_r16 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraAutonomousGapsR16Constraints)
			if err != nil {
				return err
			}
			ie.Eutra_AutonomousGaps_r16 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraAutonomousGapsNEDCR16Constraints)
			if err != nil {
				return err
			}
			ie.Eutra_AutonomousGaps_NEDC_r16 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraAutonomousGapsNRDCR16Constraints)
			if err != nil {
				return err
			}
			ie.Eutra_AutonomousGaps_NRDC_r16 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtPcellT312R16Constraints)
			if err != nil {
				return err
			}
			ie.PcellT312_r16 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeBitString(measAndMobParametersCommonSupportedGapPatternR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportedGapPattern_r16 = &v
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "concurrentMeasGap-r17", Optional: true},
				{Name: "nr-NeedForGapNCSG-Reporting-r17", Optional: true},
				{Name: "eutra-NeedForGapNCSG-Reporting-r17", Optional: true},
				{Name: "ncsg-MeasGapPerFR-r17", Optional: true},
				{Name: "ncsg-MeasGapPatterns-r17", Optional: true},
				{Name: "ncsg-MeasGapNR-Patterns-r17", Optional: true},
				{Name: "preconfiguredUE-AutonomousMeasGap-r17", Optional: true},
				{Name: "preconfiguredNW-ControlledMeasGap-r17", Optional: true},
				{Name: "handoverFR1-FR2-2-r17", Optional: true},
				{Name: "handoverFR2-1-FR2-2-r17", Optional: true},
				{Name: "independentGapConfigPRS-r17", Optional: true},
				{Name: "rrm-RelaxationRRC-ConnectedRedCap-r17", Optional: true},
				{Name: "parallelMeasurementGap-r17", Optional: true},
				{Name: "condHandoverWithSCG-NRDC-r17", Optional: true},
				{Name: "gNB-ID-LengthReporting-r17", Optional: true},
				{Name: "gNB-ID-LengthReporting-ENDC-r17", Optional: true},
				{Name: "gNB-ID-LengthReporting-NEDC-r17", Optional: true},
				{Name: "gNB-ID-LengthReporting-NRDC-r17", Optional: true},
				{Name: "gNB-ID-LengthReporting-NPN-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ConcurrentMeasGap_r17 = &struct {
				Choice                               int
				ConcurrentPerUE_OnlyMeasGap_r17      *int64
				ConcurrentPerUE_PerFRCombMeasGap_r17 *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(measAndMobParametersCommonExtConcurrentMeasGapR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ConcurrentMeasGap_r17).Choice = int(index)
			switch index {
			case MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_OnlyMeasGap_r17:
				(*ie.ConcurrentMeasGap_r17).ConcurrentPerUE_OnlyMeasGap_r17 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtConcurrentMeasGapR17ConcurrentPerUEOnlyMeasGapR17Constraints)
				if err != nil {
					return err
				}
				(*(*ie.ConcurrentMeasGap_r17).ConcurrentPerUE_OnlyMeasGap_r17) = v
			case MeasAndMobParametersCommon_Ext_ConcurrentMeasGap_r17_ConcurrentPerUE_PerFRCombMeasGap_r17:
				(*ie.ConcurrentMeasGap_r17).ConcurrentPerUE_PerFRCombMeasGap_r17 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtConcurrentMeasGapR17ConcurrentPerUEPerFRCombMeasGapR17Constraints)
				if err != nil {
					return err
				}
				(*(*ie.ConcurrentMeasGap_r17).ConcurrentPerUE_PerFRCombMeasGap_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNrNeedForGapNCSGReportingR17Constraints)
			if err != nil {
				return err
			}
			ie.Nr_NeedForGapNCSG_Reporting_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraNeedForGapNCSGReportingR17Constraints)
			if err != nil {
				return err
			}
			ie.Eutra_NeedForGapNCSG_Reporting_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNcsgMeasGapPerFRR17Constraints)
			if err != nil {
				return err
			}
			ie.Ncsg_MeasGapPerFR_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeBitString(measAndMobParametersCommonNcsgMeasGapPatternsR17Constraints)
			if err != nil {
				return err
			}
			ie.Ncsg_MeasGapPatterns_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeBitString(measAndMobParametersCommonNcsgMeasGapNRPatternsR17Constraints)
			if err != nil {
				return err
			}
			ie.Ncsg_MeasGapNR_Patterns_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtPreconfiguredUEAutonomousMeasGapR17Constraints)
			if err != nil {
				return err
			}
			ie.PreconfiguredUE_AutonomousMeasGap_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtPreconfiguredNWControlledMeasGapR17Constraints)
			if err != nil {
				return err
			}
			ie.PreconfiguredNW_ControlledMeasGap_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtHandoverFR1FR22R17Constraints)
			if err != nil {
				return err
			}
			ie.HandoverFR1_FR2_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtHandoverFR21FR22R17Constraints)
			if err != nil {
				return err
			}
			ie.HandoverFR2_1_FR2_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtIndependentGapConfigPRSR17Constraints)
			if err != nil {
				return err
			}
			ie.IndependentGapConfigPRS_r17 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtRrmRelaxationRRCConnectedRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.Rrm_RelaxationRRC_ConnectedRedCap_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtParallelMeasurementGapR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelMeasurementGap_r17 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtCondHandoverWithSCGNRDCR17Constraints)
			if err != nil {
				return err
			}
			ie.CondHandoverWithSCG_NRDC_r17 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtGNBIDLengthReportingR17Constraints)
			if err != nil {
				return err
			}
			ie.GNB_ID_LengthReporting_r17 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtGNBIDLengthReportingENDCR17Constraints)
			if err != nil {
				return err
			}
			ie.GNB_ID_LengthReporting_ENDC_r17 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtGNBIDLengthReportingNEDCR17Constraints)
			if err != nil {
				return err
			}
			ie.GNB_ID_LengthReporting_NEDC_r17 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtGNBIDLengthReportingNRDCR17Constraints)
			if err != nil {
				return err
			}
			ie.GNB_ID_LengthReporting_NRDC_r17 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtGNBIDLengthReportingNPNR17Constraints)
			if err != nil {
				return err
			}
			ie.GNB_ID_LengthReporting_NPN_r17 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "parallelSMTC-r17", Optional: true},
				{Name: "concurrentMeasGapEUTRA-r17", Optional: true},
				{Name: "serviceLinkPropDelayDiffReporting-r17", Optional: true},
				{Name: "ncsg-SymbolLevelScheduleRestrictionInter-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtParallelSMTCR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelSMTC_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtConcurrentMeasGapEUTRAR17Constraints)
			if err != nil {
				return err
			}
			ie.ConcurrentMeasGapEUTRA_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtServiceLinkPropDelayDiffReportingR17Constraints)
			if err != nil {
				return err
			}
			ie.ServiceLinkPropDelayDiffReporting_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNcsgSymbolLevelScheduleRestrictionInterR17Constraints)
			if err != nil {
				return err
			}
			ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17 = &v
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "eventD1-MeasReportTrigger-r17", Optional: true},
				{Name: "independentGapConfig-maxCC-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEventD1MeasReportTriggerR17Constraints)
			if err != nil {
				return err
			}
			ie.EventD1_MeasReportTrigger_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.IndependentGapConfig_MaxCC_r17 = &struct {
				Fr1_Only_r17   *int64
				Fr2_Only_r17   *int64
				Fr1_AndFR2_r17 *int64
			}{}
			c := ie.IndependentGapConfig_MaxCC_r17
			measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Seq := dx.NewSequenceDecoder(measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Constraints)
			if err := measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Seq.IsComponentPresent(0) {
				c.Fr1_Only_r17 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.Fr1_Only_r17) = v
			}
			if measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Seq.IsComponentPresent(1) {
				c.Fr2_Only_r17 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.Fr2_Only_r17) = v
			}
			if measAndMobParametersCommonExtIndependentGapConfigMaxCCR17Seq.IsComponentPresent(2) {
				c.Fr1_AndFR2_r17 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.Fr1_AndFR2_r17) = v
			}
		}
	}

	// Extension group 8:
	if seq.IsExtensionPresent(8) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "interSatMeas-r17", Optional: true},
				{Name: "deriveSSB-IndexFromCellInterNon-NCSG-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtInterSatMeasR17Constraints)
			if err != nil {
				return err
			}
			ie.InterSatMeas_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtDeriveSSBIndexFromCellInterNonNCSGR17Constraints)
			if err != nil {
				return err
			}
			ie.DeriveSSB_IndexFromCellInterNon_NCSG_r17 = &v
		}
	}

	// Extension group 9:
	if seq.IsExtensionPresent(9) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "l3-MeasUnknownSCellActivation-r18", Optional: true},
				{Name: "shortMeasInterval-r18", Optional: true},
				{Name: "nr-NeedForInterruptionReport-r18", Optional: true},
				{Name: "measSequenceConfig-r18", Optional: true},
				{Name: "cellIndividualOffsetPerMeasEvent-r18", Optional: true},
				{Name: "eventD2-MeasReportTrigger-r18", Optional: true},
				{Name: "concurrentMeasGapsPreMG-r18", Optional: true},
				{Name: "dynamicCollision-r18", Optional: true},
				{Name: "concurrentMeasGapsNCSG-r18", Optional: true},
				{Name: "eutra-NoGapMeasurementOutsideBWP-r18", Optional: true},
				{Name: "eutra-NoGapMeasurementInsideBWP-r18", Optional: true},
				{Name: "eutra-MeasEMW-r18", Optional: true},
				{Name: "concurrentMeasCRS-InsideBWP-EUTRA-r18", Optional: true},
				{Name: "ltm-InterFreqMeasGap-r18", Optional: true},
				{Name: "dummy-ltm-FastUE-Processing-r18", Optional: true},
				{Name: "rach-LessHandoverInterFreq-r18", Optional: true},
				{Name: "enterAndLeaveCellReport-r18", Optional: true},
				{Name: "bestCellChangeReport-r18", Optional: true},
				{Name: "secondBestCellChangeReport-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtL3MeasUnknownSCellActivationR18Constraints)
			if err != nil {
				return err
			}
			ie.L3_MeasUnknownSCellActivation_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtShortMeasIntervalR18Constraints)
			if err != nil {
				return err
			}
			ie.ShortMeasInterval_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNrNeedForInterruptionReportR18Constraints)
			if err != nil {
				return err
			}
			ie.Nr_NeedForInterruptionReport_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtMeasSequenceConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.MeasSequenceConfig_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtCellIndividualOffsetPerMeasEventR18Constraints)
			if err != nil {
				return err
			}
			ie.CellIndividualOffsetPerMeasEvent_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEventD2MeasReportTriggerR18Constraints)
			if err != nil {
				return err
			}
			ie.EventD2_MeasReportTrigger_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtConcurrentMeasGapsPreMGR18Constraints)
			if err != nil {
				return err
			}
			ie.ConcurrentMeasGapsPreMG_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtDynamicCollisionR18Constraints)
			if err != nil {
				return err
			}
			ie.DynamicCollision_r18 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtConcurrentMeasGapsNCSGR18Constraints)
			if err != nil {
				return err
			}
			ie.ConcurrentMeasGapsNCSG_r18 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraNoGapMeasurementOutsideBWPR18Constraints)
			if err != nil {
				return err
			}
			ie.Eutra_NoGapMeasurementOutsideBWP_r18 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraNoGapMeasurementInsideBWPR18Constraints)
			if err != nil {
				return err
			}
			ie.Eutra_NoGapMeasurementInsideBWP_r18 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeBitString(measAndMobParametersCommonEutraMeasEMWR18Constraints)
			if err != nil {
				return err
			}
			ie.Eutra_MeasEMW_r18 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtConcurrentMeasCRSInsideBWPEUTRAR18Constraints)
			if err != nil {
				return err
			}
			ie.ConcurrentMeasCRS_InsideBWP_EUTRA_r18 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmInterFreqMeasGapR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_InterFreqMeasGap_r18 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			ie.Dummy_Ltm_FastUE_Processing_r18 = &struct {
				Fr1_r18        int64
				Fr2_r18        int64
				Fr1_AndFR2_r18 int64
			}{}
			c := ie.Dummy_Ltm_FastUE_Processing_r18
			{
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtDummyLtmFastUEProcessingR18Fr1R18Constraints)
				if err != nil {
					return err
				}
				c.Fr1_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtDummyLtmFastUEProcessingR18Fr2R18Constraints)
				if err != nil {
					return err
				}
				c.Fr2_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtDummyLtmFastUEProcessingR18Fr1AndFR2R18Constraints)
				if err != nil {
					return err
				}
				c.Fr1_AndFR2_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtRachLessHandoverInterFreqR18Constraints)
			if err != nil {
				return err
			}
			ie.Rach_LessHandoverInterFreq_r18 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEnterAndLeaveCellReportR18Constraints)
			if err != nil {
				return err
			}
			ie.EnterAndLeaveCellReport_r18 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtBestCellChangeReportR18Constraints)
			if err != nil {
				return err
			}
			ie.BestCellChangeReport_r18 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtSecondBestCellChangeReportR18Constraints)
			if err != nil {
				return err
			}
			ie.SecondBestCellChangeReport_r18 = &v
		}
	}

	// Extension group 10:
	if seq.IsExtensionPresent(10) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ltm-InterFreq-r18", Optional: true},
				{Name: "ltm-MCG-NRDC-r18", Optional: true},
				{Name: "ltm-RACH-LessDG-r18", Optional: true},
				{Name: "ltm-RACH-LessCG-r18", Optional: true},
				{Name: "ltm-Recovery-r18", Optional: true},
				{Name: "ltm-ReferenceConfig-r18", Optional: true},
				{Name: "ltm-MCG-NRDC-Release-r18", Optional: true},
				{Name: "ltm-FastUE-Processing-r18", Optional: true},
				{Name: "ntn-NeighbourCellInfoSupport-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmInterFreqR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_InterFreq_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmMCGNRDCR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_MCG_NRDC_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmRACHLessDGR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_RACH_LessDG_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmRACHLessCGR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_RACH_LessCG_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmRecoveryR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_Recovery_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmReferenceConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_ReferenceConfig_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmMCGNRDCReleaseR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_MCG_NRDC_Release_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			ie.Ltm_FastUE_Processing_r18 = &struct {
				Fr1_r18        *int64
				Fr2_r18        *int64
				Fr1_AndFR2_r18 *int64
			}{}
			c := ie.Ltm_FastUE_Processing_r18
			measAndMobParametersCommonExtLtmFastUEProcessingR18Seq := dx.NewSequenceDecoder(measAndMobParametersCommonExtLtmFastUEProcessingR18Constraints)
			if err := measAndMobParametersCommonExtLtmFastUEProcessingR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if measAndMobParametersCommonExtLtmFastUEProcessingR18Seq.IsComponentPresent(0) {
				c.Fr1_r18 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmFastUEProcessingR18Fr1R18Constraints)
				if err != nil {
					return err
				}
				(*c.Fr1_r18) = v
			}
			if measAndMobParametersCommonExtLtmFastUEProcessingR18Seq.IsComponentPresent(1) {
				c.Fr2_r18 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmFastUEProcessingR18Fr2R18Constraints)
				if err != nil {
					return err
				}
				(*c.Fr2_r18) = v
			}
			if measAndMobParametersCommonExtLtmFastUEProcessingR18Seq.IsComponentPresent(2) {
				c.Fr1_AndFR2_r18 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmFastUEProcessingR18Fr1AndFR2R18Constraints)
				if err != nil {
					return err
				}
				(*c.Fr1_AndFR2_r18) = v
			}
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNtnNeighbourCellInfoSupportR18Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_NeighbourCellInfoSupport_r18 = &v
		}
	}

	// Extension group 11:
	if seq.IsExtensionPresent(11) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ltm-interFreqL1-OnlyInBC-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmInterFreqL1OnlyInBCR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_InterFreqL1_OnlyInBC_r18 = &v
		}
	}

	// Extension group 12:
	if seq.IsExtensionPresent(12) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "nr-CGI-Reporting-HSDN-r19", Optional: true},
				{Name: "eutra-CGI-Reporting-HSDN-r19", Optional: true},
				{Name: "ltm-KeyUpdateMCG-r19", Optional: true},
				{Name: "ltm-KeyUpdateSCG-r19", Optional: true},
				{Name: "cltm-EarlyTA-Indication-r19", Optional: true},
				{Name: "cltm-ExecutionConditionL1-r19", Optional: true},
				{Name: "cltm-ExecutionConditionL3-r19", Optional: true},
				{Name: "ltm-EventMeasAndReport-r19", Optional: true},
				{Name: "ltm-RecoveryWithKeyUpdate-r19", Optional: true},
				{Name: "ltm-MCG-SCG-AdditionOrChange-r19", Optional: true},
				{Name: "threeCarrierMeasWithoutGap-r19", Optional: true},
				{Name: "multiCarrierSingleReportWithoutGap-r19", Optional: true},
				{Name: "skipSSB-L1-RSRP-Meas-r19", Optional: true},
				{Name: "gapOccasionCancelRatioReport-r19", Optional: true},
				{Name: "twoSMTC-PeriodicitiesIntraFreq-r19", Optional: true},
				{Name: "twoSMTC-PeriodicitiesInterFreq-r19", Optional: true},
				{Name: "reportClosestReferenceLocations-r19", Optional: true},
				{Name: "ltm-SR-ConfIdInCellSwitchCommand-r19", Optional: true},
				{Name: "speedStateDetection-r19", Optional: true},
				{Name: "eventD2-MultiReferenceLocations-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtNrCGIReportingHSDNR19Constraints)
			if err != nil {
				return err
			}
			ie.Nr_CGI_Reporting_HSDN_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEutraCGIReportingHSDNR19Constraints)
			if err != nil {
				return err
			}
			ie.Eutra_CGI_Reporting_HSDN_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmKeyUpdateMCGR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_KeyUpdateMCG_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmKeyUpdateSCGR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_KeyUpdateSCG_r19 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeInteger(measAndMobParametersCommonCltmEarlyTAIndicationR19Constraints)
			if err != nil {
				return err
			}
			ie.Cltm_EarlyTA_Indication_r19 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtCltmExecutionConditionL1R19Constraints)
			if err != nil {
				return err
			}
			ie.Cltm_ExecutionConditionL1_r19 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeInteger(measAndMobParametersCommonCltmExecutionConditionL3R19Constraints)
			if err != nil {
				return err
			}
			ie.Cltm_ExecutionConditionL3_r19 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmEventMeasAndReportR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_EventMeasAndReport_r19 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmRecoveryWithKeyUpdateR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_RecoveryWithKeyUpdate_r19 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmMCGSCGAdditionOrChangeR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_MCG_SCG_AdditionOrChange_r19 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			ie.ThreeCarrierMeasWithoutGap_r19 = &struct {
				Fr1_CA_NR_DC_r19  *int64
				Fr1_FR2_CA_r19    *int64
				Fr1_FR2_NR_DC_r19 *int64
			}{}
			c := ie.ThreeCarrierMeasWithoutGap_r19
			measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Seq := dx.NewSequenceDecoder(measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Constraints)
			if err := measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Seq.IsComponentPresent(0) {
				c.Fr1_CA_NR_DC_r19 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Fr1CANRDCR19Constraints)
				if err != nil {
					return err
				}
				(*c.Fr1_CA_NR_DC_r19) = v
			}
			if measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Seq.IsComponentPresent(1) {
				c.Fr1_FR2_CA_r19 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Fr1FR2CAR19Constraints)
				if err != nil {
					return err
				}
				(*c.Fr1_FR2_CA_r19) = v
			}
			if measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Seq.IsComponentPresent(2) {
				c.Fr1_FR2_NR_DC_r19 = new(int64)
				v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtThreeCarrierMeasWithoutGapR19Fr1FR2NRDCR19Constraints)
				if err != nil {
					return err
				}
				(*c.Fr1_FR2_NR_DC_r19) = v
			}
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtMultiCarrierSingleReportWithoutGapR19Constraints)
			if err != nil {
				return err
			}
			ie.MultiCarrierSingleReportWithoutGap_r19 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtSkipSSBL1RSRPMeasR19Constraints)
			if err != nil {
				return err
			}
			ie.SkipSSB_L1_RSRP_Meas_r19 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtGapOccasionCancelRatioReportR19Constraints)
			if err != nil {
				return err
			}
			ie.GapOccasionCancelRatioReport_r19 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtTwoSMTCPeriodicitiesIntraFreqR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoSMTC_PeriodicitiesIntraFreq_r19 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtTwoSMTCPeriodicitiesInterFreqR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoSMTC_PeriodicitiesInterFreq_r19 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtReportClosestReferenceLocationsR19Constraints)
			if err != nil {
				return err
			}
			ie.ReportClosestReferenceLocations_r19 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtLtmSRConfIdInCellSwitchCommandR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_SR_ConfIdInCellSwitchCommand_r19 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtSpeedStateDetectionR19Constraints)
			if err != nil {
				return err
			}
			ie.SpeedStateDetection_r19 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			v, err := dx.DecodeEnumerated(measAndMobParametersCommonExtEventD2MultiReferenceLocationsR19Constraints)
			if err != nil {
				return err
			}
			ie.EventD2_MultiReferenceLocations_r19 = &v
		}
	}

	return nil
}
