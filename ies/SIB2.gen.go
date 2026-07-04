// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB2 (line 3702).

var sIB2Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cellReselectionInfoCommon"},
		{Name: "cellReselectionServingFreqInfo"},
		{Name: "intraFreqCellReselectionInfo"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var sIB2CellEquivalentSizeR17Constraints = per.Constrained(2, 16)

const (
	SIB2_Ext_Uav_PrioritizedFrequency_r19_True = 0
)

var sIB2ExtUavPrioritizedFrequencyR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_Ext_Uav_PrioritizedFrequency_r19_True},
}

var sIB2CellReselectionInfoCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofSS-BlocksToAverage", Optional: true},
		{Name: "absThreshSS-BlocksConsolidation", Optional: true},
		{Name: "rangeToBestCell", Optional: true},
		{Name: "q-Hyst"},
		{Name: "speedStateReselectionPars", Optional: true},
	},
}

const (
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB0  = 0
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB1  = 1
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB2  = 2
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB3  = 3
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB4  = 4
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB5  = 5
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB6  = 6
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB8  = 7
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB10 = 8
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB12 = 9
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB14 = 10
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB16 = 11
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB18 = 12
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB20 = 13
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB22 = 14
	SIB2_CellReselectionInfoCommon_Q_Hyst_DB24 = 15
)

var sIB2CellReselectionInfoCommonQHystConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_CellReselectionInfoCommon_Q_Hyst_DB0, SIB2_CellReselectionInfoCommon_Q_Hyst_DB1, SIB2_CellReselectionInfoCommon_Q_Hyst_DB2, SIB2_CellReselectionInfoCommon_Q_Hyst_DB3, SIB2_CellReselectionInfoCommon_Q_Hyst_DB4, SIB2_CellReselectionInfoCommon_Q_Hyst_DB5, SIB2_CellReselectionInfoCommon_Q_Hyst_DB6, SIB2_CellReselectionInfoCommon_Q_Hyst_DB8, SIB2_CellReselectionInfoCommon_Q_Hyst_DB10, SIB2_CellReselectionInfoCommon_Q_Hyst_DB12, SIB2_CellReselectionInfoCommon_Q_Hyst_DB14, SIB2_CellReselectionInfoCommon_Q_Hyst_DB16, SIB2_CellReselectionInfoCommon_Q_Hyst_DB18, SIB2_CellReselectionInfoCommon_Q_Hyst_DB20, SIB2_CellReselectionInfoCommon_Q_Hyst_DB22, SIB2_CellReselectionInfoCommon_Q_Hyst_DB24},
}

const (
	SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_Medium_DB_6 = 0
	SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_Medium_DB_4 = 1
	SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_Medium_DB_2 = 2
	SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_Medium_DB0  = 3
)

var sIB2CellReselectionInfoCommonSpeedStateReselectionParsQHystSFSfMediumConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_Medium_DB_6, SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_Medium_DB_4, SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_Medium_DB_2, SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_Medium_DB0},
}

const (
	SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_High_DB_6 = 0
	SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_High_DB_4 = 1
	SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_High_DB_2 = 2
	SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_High_DB0  = 3
)

var sIB2CellReselectionInfoCommonSpeedStateReselectionParsQHystSFSfHighConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_High_DB_6, SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_High_DB_4, SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_High_DB_2, SIB2_CellReselectionInfoCommon_SpeedStateReselectionPars_Q_HystSF_Sf_High_DB0},
}

var sIB2CellReselectionServingFreqInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "s-NonIntraSearchP", Optional: true},
		{Name: "s-NonIntraSearchQ", Optional: true},
		{Name: "threshServingLowP"},
		{Name: "threshServingLowQ", Optional: true},
		{Name: "cellReselectionPriority"},
		{Name: "cellReselectionSubPriority", Optional: true},
	},
}

var sIB2IntraFreqCellReselectionInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "q-RxLevMin"},
		{Name: "q-RxLevMinSUL", Optional: true},
		{Name: "q-QualMin", Optional: true},
		{Name: "s-IntraSearchP"},
		{Name: "s-IntraSearchQ", Optional: true},
		{Name: "t-ReselectionNR"},
		{Name: "frequencyBandList", Optional: true},
		{Name: "frequencyBandListSUL", Optional: true},
		{Name: "p-Max", Optional: true},
		{Name: "smtc", Optional: true},
		{Name: "ss-RSSI-Measurement", Optional: true},
		{Name: "ssb-ToMeasure", Optional: true},
		{Name: "deriveSSB-IndexFromCell"},
	},
}

var sIB2ExtRelaxedMeasurementR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lowMobilityEvaluation-r16", Optional: true},
		{Name: "cellEdgeEvaluation-r16", Optional: true},
		{Name: "combineRelaxedMeasCondition-r16", Optional: true},
		{Name: "highPriorityMeasRelax-r16", Optional: true},
	},
}

const (
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB3    = 0
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB6    = 1
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB9    = 2
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB12   = 3
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB15   = 4
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_Spare3 = 5
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_Spare2 = 6
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_Spare1 = 7
)

var sIB2ExtRelaxedMeasurementR16LowMobilityEvaluationR16SSearchDeltaPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB3, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB6, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB9, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB12, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_DB15, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_Spare3, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_Spare2, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_S_SearchDeltaP_r16_Spare1},
}

const (
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S5     = 0
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S10    = 1
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S20    = 2
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S30    = 3
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S60    = 4
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S120   = 5
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S180   = 6
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S240   = 7
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S300   = 8
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare7 = 9
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare6 = 10
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare5 = 11
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare4 = 12
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare3 = 13
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare2 = 14
	SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare1 = 15
)

var sIB2ExtRelaxedMeasurementR16LowMobilityEvaluationR16TSearchDeltaPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S5, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S10, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S20, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S30, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S60, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S120, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S180, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S240, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_S300, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare7, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare6, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare5, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare4, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare3, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare2, SIB2_Ext_RelaxedMeasurement_r16_LowMobilityEvaluation_r16_T_SearchDeltaP_r16_Spare1},
}

var sIB2ExtRelaxedMeasurementR16CellEdgeEvaluationR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "s-SearchThresholdP-r16"},
		{Name: "s-SearchThresholdQ-r16", Optional: true},
	},
}

const (
	SIB2_Ext_RelaxedMeasurement_r16_CombineRelaxedMeasCondition_r16_True = 0
)

var sIB2ExtRelaxedMeasurementR16CombineRelaxedMeasConditionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_Ext_RelaxedMeasurement_r16_CombineRelaxedMeasCondition_r16_True},
}

const (
	SIB2_Ext_RelaxedMeasurement_r16_HighPriorityMeasRelax_r16_True = 0
)

var sIB2ExtRelaxedMeasurementR16HighPriorityMeasRelaxR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_Ext_RelaxedMeasurement_r16_HighPriorityMeasRelax_r16_True},
}

var sIB2ExtRelaxedMeasurementR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "stationaryMobilityEvaluation-r17"},
		{Name: "cellEdgeEvaluationWhileStationary-r17", Optional: true},
		{Name: "combineRelaxedMeasCondition2-r17", Optional: true},
	},
}

const (
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB2    = 0
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB3    = 1
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB6    = 2
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB9    = 3
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB12   = 4
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB15   = 5
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_Spare2 = 6
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_Spare1 = 7
)

var sIB2ExtRelaxedMeasurementR17StationaryMobilityEvaluationR17SSearchDeltaPStationaryR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB2, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB3, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB6, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB9, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB12, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_DB15, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_Spare2, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_S_SearchDeltaP_Stationary_r17_Spare1},
}

const (
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S5     = 0
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S10    = 1
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S20    = 2
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S30    = 3
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S60    = 4
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S120   = 5
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S180   = 6
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S240   = 7
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S300   = 8
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare7 = 9
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare6 = 10
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare5 = 11
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare4 = 12
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare3 = 13
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare2 = 14
	SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare1 = 15
)

var sIB2ExtRelaxedMeasurementR17StationaryMobilityEvaluationR17TSearchDeltaPStationaryR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S5, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S10, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S20, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S30, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S60, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S120, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S180, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S240, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_S300, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare7, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare6, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare5, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare4, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare3, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare2, SIB2_Ext_RelaxedMeasurement_r17_StationaryMobilityEvaluation_r17_T_SearchDeltaP_Stationary_r17_Spare1},
}

var sIB2ExtRelaxedMeasurementR17CellEdgeEvaluationWhileStationaryR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "s-SearchThresholdP2-r17"},
		{Name: "s-SearchThresholdQ2-r17", Optional: true},
	},
}

const (
	SIB2_Ext_RelaxedMeasurement_r17_CombineRelaxedMeasCondition2_r17_True = 0
)

var sIB2ExtRelaxedMeasurementR17CombineRelaxedMeasCondition2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB2_Ext_RelaxedMeasurement_r17_CombineRelaxedMeasCondition2_r17_True},
}

var sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "altitudeMin-r19", Optional: true},
		{Name: "altitudeMax-r19", Optional: true},
		{Name: "altitudeHyst-r19", Optional: true},
	},
}

var sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellEdgeEvaluationOnMR-ForLR-OnSSB-r19", Optional: true},
		{Name: "cellEdgeEvaluationOnMR-ForLR-OnLPSS-r19", Optional: true},
		{Name: "cellEdgeEvaluationOnLR-ForLR-OnLPSS-r19", Optional: true},
		{Name: "cellEdgeEvaluationOnLR-ForLR-OnSSB-r19", Optional: true},
	},
}

var sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnSSBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "s-SearchThresholdP3-r19"},
		{Name: "s-SearchThresholdQ3-r19", Optional: true},
	},
}

var sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "s-SearchThresholdP4-r19"},
		{Name: "s-SearchThresholdQ4-r19", Optional: true},
	},
}

var sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrpThresholdLR-r19"},
		{Name: "rsrqThresholdLR-r19", Optional: true},
	},
}

var sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnSSBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrpThresholdLR2-r19"},
		{Name: "rsrqThresholdLR2-r19", Optional: true},
	},
}

var sIB2ExtOffloadMeasurementForServingCellR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellEdgeEvaluationOnMR-ForLR-OnSSB-r19", Optional: true},
		{Name: "cellEdgeEvaluationOnMR-ForLR-OnLPSS-r19", Optional: true},
		{Name: "cellEdgeEvaluationOnLR-ForLR-OnLPSS-r19", Optional: true},
		{Name: "cellEdgeEvaluationOnLR-ForLR-OnSSB-r19", Optional: true},
		{Name: "cellEdgeEvaluationOnLR-ForLR-OnLPSS-Exit-r19", Optional: true},
		{Name: "cellEdgeEvaluationOnLR-ForLR-OnSSB-Exit-r19", Optional: true},
	},
}

var sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnSSBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "s-SearchThresholdP5-r19"},
		{Name: "s-SearchThresholdQ5-r19", Optional: true},
	},
}

var sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "s-SearchThresholdP6-r19"},
		{Name: "s-SearchThresholdQ6-r19", Optional: true},
	},
}

var sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrpThresholdLR3-r19"},
		{Name: "rsrqThresholdLR3-r19", Optional: true},
	},
}

var sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrpThresholdLR4-r19"},
		{Name: "rsrqThresholdLR4-r19", Optional: true},
	},
}

var sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSExitR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrpThresholdLR5-r19"},
		{Name: "rsrqThresholdLR5-r19", Optional: true},
	},
}

var sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBExitR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrpThresholdLR6-r19"},
		{Name: "rsrqThresholdLR6-r19", Optional: true},
	},
}

type SIB2 struct {
	CellReselectionInfoCommon struct {
		NrofSS_BlocksToAverage          *int64
		AbsThreshSS_BlocksConsolidation *ThresholdNR
		RangeToBestCell                 *RangeToBestCell
		Q_Hyst                          int64
		SpeedStateReselectionPars       *struct {
			MobilityStateParameters MobilityStateParameters
			Q_HystSF                struct {
				Sf_Medium int64
				Sf_High   int64
			}
		}
	}
	CellReselectionServingFreqInfo struct {
		S_NonIntraSearchP          *ReselectionThreshold
		S_NonIntraSearchQ          *ReselectionThresholdQ
		ThreshServingLowP          ReselectionThreshold
		ThreshServingLowQ          *ReselectionThresholdQ
		CellReselectionPriority    CellReselectionPriority
		CellReselectionSubPriority *CellReselectionSubPriority
	}
	IntraFreqCellReselectionInfo struct {
		Q_RxLevMin              Q_RxLevMin
		Q_RxLevMinSUL           *Q_RxLevMin
		Q_QualMin               *Q_QualMin
		S_IntraSearchP          ReselectionThreshold
		S_IntraSearchQ          *ReselectionThresholdQ
		T_ReselectionNR         T_Reselection
		FrequencyBandList       *MultiFrequencyBandListNR_SIB
		FrequencyBandListSUL    *MultiFrequencyBandListNR_SIB
		P_Max                   *P_Max
		Smtc                    *SSB_MTC
		Ss_RSSI_Measurement     *SS_RSSI_Measurement
		Ssb_ToMeasure           *SSB_ToMeasure
		DeriveSSB_IndexFromCell bool
	}
	RelaxedMeasurement_r16 *struct {
		LowMobilityEvaluation_r16 *struct {
			S_SearchDeltaP_r16 int64
			T_SearchDeltaP_r16 int64
		}
		CellEdgeEvaluation_r16 *struct {
			S_SearchThresholdP_r16 ReselectionThreshold
			S_SearchThresholdQ_r16 *ReselectionThresholdQ
		}
		CombineRelaxedMeasCondition_r16 *int64
		HighPriorityMeasRelax_r16       *int64
	}
	CellEquivalentSize_r17 *int64
	RelaxedMeasurement_r17 *struct {
		StationaryMobilityEvaluation_r17 struct {
			S_SearchDeltaP_Stationary_r17 int64
			T_SearchDeltaP_Stationary_r17 int64
		}
		CellEdgeEvaluationWhileStationary_r17 *struct {
			S_SearchThresholdP2_r17 ReselectionThreshold
			S_SearchThresholdQ2_r17 *ReselectionThresholdQ
		}
		CombineRelaxedMeasCondition2_r17 *int64
	}
	Uav_PrioritizedFrequency_r19              *int64
	Uav_PrioritizedFrequencyAltitudeRange_r19 *struct {
		AltitudeMin_r19  *Altitude_r18
		AltitudeMax_r19  *Altitude_r18
		AltitudeHyst_r19 *HysteresisAltitude_r18
	}
	Ssb_ToMeasureAltitudeBasedList_r19                 *SSB_ToMeasureAltitudeBasedList_r18
	RelaxedMeasurementForServingAndNeighboringCell_r19 *struct {
		CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 *struct {
			S_SearchThresholdP3_r19 ReselectionThreshold
			S_SearchThresholdQ3_r19 *ReselectionThresholdQ
		}
		CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 *struct {
			S_SearchThresholdP4_r19 ReselectionThreshold
			S_SearchThresholdQ4_r19 *ReselectionThresholdQ
		}
		CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 *struct {
			RsrpThresholdLR_r19 ThresholdP_LR_r19
			RsrqThresholdLR_r19 *ThresholdQ_LR_r19
		}
		CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 *struct {
			RsrpThresholdLR2_r19 ThresholdP_LR_r19
			RsrqThresholdLR2_r19 *ThresholdQ_LR_r19
		}
	}
	OffloadMeasurementForServingCell_r19 *struct {
		CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 *struct {
			S_SearchThresholdP5_r19 ReselectionThreshold
			S_SearchThresholdQ5_r19 *ReselectionThresholdQ
		}
		CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 *struct {
			S_SearchThresholdP6_r19 ReselectionThreshold
			S_SearchThresholdQ6_r19 *ReselectionThresholdQ
		}
		CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 *struct {
			RsrpThresholdLR3_r19 ThresholdP_LR_r19
			RsrqThresholdLR3_r19 *ThresholdQ_LR_r19
		}
		CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 *struct {
			RsrpThresholdLR4_r19 ThresholdP_LR_r19
			RsrqThresholdLR4_r19 *ThresholdQ_LR_r19
		}
		CellEdgeEvaluationOnLR_ForLR_OnLPSS_Exit_r19 *struct {
			RsrpThresholdLR5_r19 ThresholdP_LR_r19
			RsrqThresholdLR5_r19 *ThresholdQ_LR_r19
		}
		CellEdgeEvaluationOnLR_ForLR_OnSSB_Exit_r19 *struct {
			RsrpThresholdLR6_r19 ThresholdP_LR_r19
			RsrqThresholdLR6_r19 *ThresholdQ_LR_r19
		}
	}
}

func (ie *SIB2) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB2Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.RelaxedMeasurement_r16 != nil
	hasExtGroup1 := ie.CellEquivalentSize_r17 != nil || ie.RelaxedMeasurement_r17 != nil
	hasExtGroup2 := ie.Uav_PrioritizedFrequency_r19 != nil || ie.Uav_PrioritizedFrequencyAltitudeRange_r19 != nil || ie.Ssb_ToMeasureAltitudeBasedList_r19 != nil || ie.RelaxedMeasurementForServingAndNeighboringCell_r19 != nil || ie.OffloadMeasurementForServingCell_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. cellReselectionInfoCommon: sequence
	{
		{
			c := &ie.CellReselectionInfoCommon
			sIB2CellReselectionInfoCommonSeq := e.NewSequenceEncoder(sIB2CellReselectionInfoCommonConstraints)
			if err := sIB2CellReselectionInfoCommonSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := sIB2CellReselectionInfoCommonSeq.EncodePreamble([]bool{c.NrofSS_BlocksToAverage != nil, c.AbsThreshSS_BlocksConsolidation != nil, c.RangeToBestCell != nil, c.SpeedStateReselectionPars != nil}); err != nil {
				return err
			}
			if c.NrofSS_BlocksToAverage != nil {
				if err := e.EncodeInteger((*c.NrofSS_BlocksToAverage), per.Constrained(2, common.MaxNrofSS_BlocksToAverage)); err != nil {
					return err
				}
			}
			if c.AbsThreshSS_BlocksConsolidation != nil {
				if err := c.AbsThreshSS_BlocksConsolidation.Encode(e); err != nil {
					return err
				}
			}
			if c.RangeToBestCell != nil {
				if err := c.RangeToBestCell.Encode(e); err != nil {
					return err
				}
			}
			if err := e.EncodeEnumerated(c.Q_Hyst, sIB2CellReselectionInfoCommonQHystConstraints); err != nil {
				return err
			}
			if c.SpeedStateReselectionPars != nil {
				c := (*c.SpeedStateReselectionPars)
				if err := c.MobilityStateParameters.Encode(e); err != nil {
					return err
				}
				{
					c := &c.Q_HystSF
					if err := e.EncodeEnumerated(c.Sf_Medium, sIB2CellReselectionInfoCommonSpeedStateReselectionParsQHystSFSfMediumConstraints); err != nil {
						return err
					}
					if err := e.EncodeEnumerated(c.Sf_High, sIB2CellReselectionInfoCommonSpeedStateReselectionParsQHystSFSfHighConstraints); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. cellReselectionServingFreqInfo: sequence
	{
		{
			c := &ie.CellReselectionServingFreqInfo
			sIB2CellReselectionServingFreqInfoSeq := e.NewSequenceEncoder(sIB2CellReselectionServingFreqInfoConstraints)
			if err := sIB2CellReselectionServingFreqInfoSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := sIB2CellReselectionServingFreqInfoSeq.EncodePreamble([]bool{c.S_NonIntraSearchP != nil, c.S_NonIntraSearchQ != nil, c.ThreshServingLowQ != nil, c.CellReselectionSubPriority != nil}); err != nil {
				return err
			}
			if c.S_NonIntraSearchP != nil {
				if err := c.S_NonIntraSearchP.Encode(e); err != nil {
					return err
				}
			}
			if c.S_NonIntraSearchQ != nil {
				if err := c.S_NonIntraSearchQ.Encode(e); err != nil {
					return err
				}
			}
			if err := c.ThreshServingLowP.Encode(e); err != nil {
				return err
			}
			if c.ThreshServingLowQ != nil {
				if err := c.ThreshServingLowQ.Encode(e); err != nil {
					return err
				}
			}
			if err := c.CellReselectionPriority.Encode(e); err != nil {
				return err
			}
			if c.CellReselectionSubPriority != nil {
				if err := c.CellReselectionSubPriority.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. intraFreqCellReselectionInfo: sequence
	{
		{
			c := &ie.IntraFreqCellReselectionInfo
			sIB2IntraFreqCellReselectionInfoSeq := e.NewSequenceEncoder(sIB2IntraFreqCellReselectionInfoConstraints)
			if err := sIB2IntraFreqCellReselectionInfoSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := sIB2IntraFreqCellReselectionInfoSeq.EncodePreamble([]bool{c.Q_RxLevMinSUL != nil, c.Q_QualMin != nil, c.S_IntraSearchQ != nil, c.FrequencyBandList != nil, c.FrequencyBandListSUL != nil, c.P_Max != nil, c.Smtc != nil, c.Ss_RSSI_Measurement != nil, c.Ssb_ToMeasure != nil}); err != nil {
				return err
			}
			if err := c.Q_RxLevMin.Encode(e); err != nil {
				return err
			}
			if c.Q_RxLevMinSUL != nil {
				if err := c.Q_RxLevMinSUL.Encode(e); err != nil {
					return err
				}
			}
			if c.Q_QualMin != nil {
				if err := c.Q_QualMin.Encode(e); err != nil {
					return err
				}
			}
			if err := c.S_IntraSearchP.Encode(e); err != nil {
				return err
			}
			if c.S_IntraSearchQ != nil {
				if err := c.S_IntraSearchQ.Encode(e); err != nil {
					return err
				}
			}
			if err := c.T_ReselectionNR.Encode(e); err != nil {
				return err
			}
			if c.FrequencyBandList != nil {
				if err := c.FrequencyBandList.Encode(e); err != nil {
					return err
				}
			}
			if c.FrequencyBandListSUL != nil {
				if err := c.FrequencyBandListSUL.Encode(e); err != nil {
					return err
				}
			}
			if c.P_Max != nil {
				if err := c.P_Max.Encode(e); err != nil {
					return err
				}
			}
			if c.Smtc != nil {
				if err := c.Smtc.Encode(e); err != nil {
					return err
				}
			}
			if c.Ss_RSSI_Measurement != nil {
				if err := c.Ss_RSSI_Measurement.Encode(e); err != nil {
					return err
				}
			}
			if c.Ssb_ToMeasure != nil {
				if err := c.Ssb_ToMeasure.Encode(e); err != nil {
					return err
				}
			}
			if err := e.EncodeBoolean(c.DeriveSSB_IndexFromCell); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "relaxedMeasurement-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RelaxedMeasurement_r16 != nil}); err != nil {
				return err
			}

			if ie.RelaxedMeasurement_r16 != nil {
				c := ie.RelaxedMeasurement_r16
				sIB2ExtRelaxedMeasurementR16Seq := ex.NewSequenceEncoder(sIB2ExtRelaxedMeasurementR16Constraints)
				if err := sIB2ExtRelaxedMeasurementR16Seq.EncodePreamble([]bool{c.LowMobilityEvaluation_r16 != nil, c.CellEdgeEvaluation_r16 != nil, c.CombineRelaxedMeasCondition_r16 != nil, c.HighPriorityMeasRelax_r16 != nil}); err != nil {
					return err
				}
				if c.LowMobilityEvaluation_r16 != nil {
					c := (*c.LowMobilityEvaluation_r16)
					if err := ex.EncodeEnumerated(c.S_SearchDeltaP_r16, sIB2ExtRelaxedMeasurementR16LowMobilityEvaluationR16SSearchDeltaPR16Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.T_SearchDeltaP_r16, sIB2ExtRelaxedMeasurementR16LowMobilityEvaluationR16TSearchDeltaPR16Constraints); err != nil {
						return err
					}
				}
				if c.CellEdgeEvaluation_r16 != nil {
					c := (*c.CellEdgeEvaluation_r16)
					sIB2ExtRelaxedMeasurementR16CellEdgeEvaluationR16Seq := ex.NewSequenceEncoder(sIB2ExtRelaxedMeasurementR16CellEdgeEvaluationR16Constraints)
					if err := sIB2ExtRelaxedMeasurementR16CellEdgeEvaluationR16Seq.EncodePreamble([]bool{c.S_SearchThresholdQ_r16 != nil}); err != nil {
						return err
					}
					if err := c.S_SearchThresholdP_r16.Encode(ex); err != nil {
						return err
					}
					if c.S_SearchThresholdQ_r16 != nil {
						if err := c.S_SearchThresholdQ_r16.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CombineRelaxedMeasCondition_r16 != nil {
					if err := ex.EncodeEnumerated((*c.CombineRelaxedMeasCondition_r16), sIB2ExtRelaxedMeasurementR16CombineRelaxedMeasConditionR16Constraints); err != nil {
						return err
					}
				}
				if c.HighPriorityMeasRelax_r16 != nil {
					if err := ex.EncodeEnumerated((*c.HighPriorityMeasRelax_r16), sIB2ExtRelaxedMeasurementR16HighPriorityMeasRelaxR16Constraints); err != nil {
						return err
					}
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
					{Name: "cellEquivalentSize-r17", Optional: true},
					{Name: "relaxedMeasurement-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CellEquivalentSize_r17 != nil, ie.RelaxedMeasurement_r17 != nil}); err != nil {
				return err
			}

			if ie.CellEquivalentSize_r17 != nil {
				if err := ex.EncodeInteger(*ie.CellEquivalentSize_r17, sIB2CellEquivalentSizeR17Constraints); err != nil {
					return err
				}
			}

			if ie.RelaxedMeasurement_r17 != nil {
				c := ie.RelaxedMeasurement_r17
				sIB2ExtRelaxedMeasurementR17Seq := ex.NewSequenceEncoder(sIB2ExtRelaxedMeasurementR17Constraints)
				if err := sIB2ExtRelaxedMeasurementR17Seq.EncodePreamble([]bool{c.CellEdgeEvaluationWhileStationary_r17 != nil, c.CombineRelaxedMeasCondition2_r17 != nil}); err != nil {
					return err
				}
				{
					c := &c.StationaryMobilityEvaluation_r17
					if err := ex.EncodeEnumerated(c.S_SearchDeltaP_Stationary_r17, sIB2ExtRelaxedMeasurementR17StationaryMobilityEvaluationR17SSearchDeltaPStationaryR17Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.T_SearchDeltaP_Stationary_r17, sIB2ExtRelaxedMeasurementR17StationaryMobilityEvaluationR17TSearchDeltaPStationaryR17Constraints); err != nil {
						return err
					}
				}
				if c.CellEdgeEvaluationWhileStationary_r17 != nil {
					c := (*c.CellEdgeEvaluationWhileStationary_r17)
					sIB2ExtRelaxedMeasurementR17CellEdgeEvaluationWhileStationaryR17Seq := ex.NewSequenceEncoder(sIB2ExtRelaxedMeasurementR17CellEdgeEvaluationWhileStationaryR17Constraints)
					if err := sIB2ExtRelaxedMeasurementR17CellEdgeEvaluationWhileStationaryR17Seq.EncodePreamble([]bool{c.S_SearchThresholdQ2_r17 != nil}); err != nil {
						return err
					}
					if err := c.S_SearchThresholdP2_r17.Encode(ex); err != nil {
						return err
					}
					if c.S_SearchThresholdQ2_r17 != nil {
						if err := c.S_SearchThresholdQ2_r17.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CombineRelaxedMeasCondition2_r17 != nil {
					if err := ex.EncodeEnumerated((*c.CombineRelaxedMeasCondition2_r17), sIB2ExtRelaxedMeasurementR17CombineRelaxedMeasCondition2R17Constraints); err != nil {
						return err
					}
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
					{Name: "uav-PrioritizedFrequency-r19", Optional: true},
					{Name: "uav-PrioritizedFrequencyAltitudeRange-r19", Optional: true},
					{Name: "ssb-ToMeasureAltitudeBasedList-r19", Optional: true},
					{Name: "relaxedMeasurementForServingAndNeighboringCell-r19", Optional: true},
					{Name: "offloadMeasurementForServingCell-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Uav_PrioritizedFrequency_r19 != nil, ie.Uav_PrioritizedFrequencyAltitudeRange_r19 != nil, ie.Ssb_ToMeasureAltitudeBasedList_r19 != nil, ie.RelaxedMeasurementForServingAndNeighboringCell_r19 != nil, ie.OffloadMeasurementForServingCell_r19 != nil}); err != nil {
				return err
			}

			if ie.Uav_PrioritizedFrequency_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Uav_PrioritizedFrequency_r19, sIB2ExtUavPrioritizedFrequencyR19Constraints); err != nil {
					return err
				}
			}

			if ie.Uav_PrioritizedFrequencyAltitudeRange_r19 != nil {
				c := ie.Uav_PrioritizedFrequencyAltitudeRange_r19
				sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Seq := ex.NewSequenceEncoder(sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Constraints)
				if err := sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Seq.EncodePreamble([]bool{c.AltitudeMin_r19 != nil, c.AltitudeMax_r19 != nil, c.AltitudeHyst_r19 != nil}); err != nil {
					return err
				}
				if c.AltitudeMin_r19 != nil {
					if err := c.AltitudeMin_r19.Encode(ex); err != nil {
						return err
					}
				}
				if c.AltitudeMax_r19 != nil {
					if err := c.AltitudeMax_r19.Encode(ex); err != nil {
						return err
					}
				}
				if c.AltitudeHyst_r19 != nil {
					if err := c.AltitudeHyst_r19.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ssb_ToMeasureAltitudeBasedList_r19 != nil {
				if err := ie.Ssb_ToMeasureAltitudeBasedList_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.RelaxedMeasurementForServingAndNeighboringCell_r19 != nil {
				c := ie.RelaxedMeasurementForServingAndNeighboringCell_r19
				sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Seq := ex.NewSequenceEncoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Constraints)
				if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Seq.EncodePreamble([]bool{c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 != nil, c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 != nil, c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 != nil, c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 != nil}); err != nil {
					return err
				}
				if c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 != nil {
					c := (*c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19)
					sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq := ex.NewSequenceEncoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnSSBR19Constraints)
					if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq.EncodePreamble([]bool{c.S_SearchThresholdQ3_r19 != nil}); err != nil {
						return err
					}
					if err := c.S_SearchThresholdP3_r19.Encode(ex); err != nil {
						return err
					}
					if c.S_SearchThresholdQ3_r19 != nil {
						if err := c.S_SearchThresholdQ3_r19.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 != nil {
					c := (*c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19)
					sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq := ex.NewSequenceEncoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Constraints)
					if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq.EncodePreamble([]bool{c.S_SearchThresholdQ4_r19 != nil}); err != nil {
						return err
					}
					if err := c.S_SearchThresholdP4_r19.Encode(ex); err != nil {
						return err
					}
					if c.S_SearchThresholdQ4_r19 != nil {
						if err := c.S_SearchThresholdQ4_r19.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 != nil {
					c := (*c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19)
					sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq := ex.NewSequenceEncoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Constraints)
					if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq.EncodePreamble([]bool{c.RsrqThresholdLR_r19 != nil}); err != nil {
						return err
					}
					if err := c.RsrpThresholdLR_r19.Encode(ex); err != nil {
						return err
					}
					if c.RsrqThresholdLR_r19 != nil {
						if err := c.RsrqThresholdLR_r19.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 != nil {
					c := (*c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19)
					sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq := ex.NewSequenceEncoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnSSBR19Constraints)
					if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq.EncodePreamble([]bool{c.RsrqThresholdLR2_r19 != nil}); err != nil {
						return err
					}
					if err := c.RsrpThresholdLR2_r19.Encode(ex); err != nil {
						return err
					}
					if c.RsrqThresholdLR2_r19 != nil {
						if err := c.RsrqThresholdLR2_r19.Encode(ex); err != nil {
							return err
						}
					}
				}
			}

			if ie.OffloadMeasurementForServingCell_r19 != nil {
				c := ie.OffloadMeasurementForServingCell_r19
				sIB2ExtOffloadMeasurementForServingCellR19Seq := ex.NewSequenceEncoder(sIB2ExtOffloadMeasurementForServingCellR19Constraints)
				if err := sIB2ExtOffloadMeasurementForServingCellR19Seq.EncodePreamble([]bool{c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 != nil, c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 != nil, c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 != nil, c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 != nil, c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_Exit_r19 != nil, c.CellEdgeEvaluationOnLR_ForLR_OnSSB_Exit_r19 != nil}); err != nil {
					return err
				}
				if c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 != nil {
					c := (*c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19)
					sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq := ex.NewSequenceEncoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnSSBR19Constraints)
					if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq.EncodePreamble([]bool{c.S_SearchThresholdQ5_r19 != nil}); err != nil {
						return err
					}
					if err := c.S_SearchThresholdP5_r19.Encode(ex); err != nil {
						return err
					}
					if c.S_SearchThresholdQ5_r19 != nil {
						if err := c.S_SearchThresholdQ5_r19.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 != nil {
					c := (*c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19)
					sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq := ex.NewSequenceEncoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Constraints)
					if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq.EncodePreamble([]bool{c.S_SearchThresholdQ6_r19 != nil}); err != nil {
						return err
					}
					if err := c.S_SearchThresholdP6_r19.Encode(ex); err != nil {
						return err
					}
					if c.S_SearchThresholdQ6_r19 != nil {
						if err := c.S_SearchThresholdQ6_r19.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 != nil {
					c := (*c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19)
					sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq := ex.NewSequenceEncoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Constraints)
					if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq.EncodePreamble([]bool{c.RsrqThresholdLR3_r19 != nil}); err != nil {
						return err
					}
					if err := c.RsrpThresholdLR3_r19.Encode(ex); err != nil {
						return err
					}
					if c.RsrqThresholdLR3_r19 != nil {
						if err := c.RsrqThresholdLR3_r19.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 != nil {
					c := (*c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19)
					sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq := ex.NewSequenceEncoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBR19Constraints)
					if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq.EncodePreamble([]bool{c.RsrqThresholdLR4_r19 != nil}); err != nil {
						return err
					}
					if err := c.RsrpThresholdLR4_r19.Encode(ex); err != nil {
						return err
					}
					if c.RsrqThresholdLR4_r19 != nil {
						if err := c.RsrqThresholdLR4_r19.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_Exit_r19 != nil {
					c := (*c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_Exit_r19)
					sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSExitR19Seq := ex.NewSequenceEncoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSExitR19Constraints)
					if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSExitR19Seq.EncodePreamble([]bool{c.RsrqThresholdLR5_r19 != nil}); err != nil {
						return err
					}
					if err := c.RsrpThresholdLR5_r19.Encode(ex); err != nil {
						return err
					}
					if c.RsrqThresholdLR5_r19 != nil {
						if err := c.RsrqThresholdLR5_r19.Encode(ex); err != nil {
							return err
						}
					}
				}
				if c.CellEdgeEvaluationOnLR_ForLR_OnSSB_Exit_r19 != nil {
					c := (*c.CellEdgeEvaluationOnLR_ForLR_OnSSB_Exit_r19)
					sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBExitR19Seq := ex.NewSequenceEncoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBExitR19Constraints)
					if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBExitR19Seq.EncodePreamble([]bool{c.RsrqThresholdLR6_r19 != nil}); err != nil {
						return err
					}
					if err := c.RsrpThresholdLR6_r19.Encode(ex); err != nil {
						return err
					}
					if c.RsrqThresholdLR6_r19 != nil {
						if err := c.RsrqThresholdLR6_r19.Encode(ex); err != nil {
							return err
						}
					}
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

func (ie *SIB2) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB2Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. cellReselectionInfoCommon: sequence
	{
		{
			c := &ie.CellReselectionInfoCommon
			sIB2CellReselectionInfoCommonSeq := d.NewSequenceDecoder(sIB2CellReselectionInfoCommonConstraints)
			if err := sIB2CellReselectionInfoCommonSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := sIB2CellReselectionInfoCommonSeq.DecodePreamble(); err != nil {
				return err
			}
			if sIB2CellReselectionInfoCommonSeq.IsComponentPresent(0) {
				c.NrofSS_BlocksToAverage = new(int64)
				v, err := d.DecodeInteger(per.Constrained(2, common.MaxNrofSS_BlocksToAverage))
				if err != nil {
					return err
				}
				(*c.NrofSS_BlocksToAverage) = v
			}
			if sIB2CellReselectionInfoCommonSeq.IsComponentPresent(1) {
				c.AbsThreshSS_BlocksConsolidation = new(ThresholdNR)
				if err := (*c.AbsThreshSS_BlocksConsolidation).Decode(d); err != nil {
					return err
				}
			}
			if sIB2CellReselectionInfoCommonSeq.IsComponentPresent(2) {
				c.RangeToBestCell = new(RangeToBestCell)
				if err := (*c.RangeToBestCell).Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeEnumerated(sIB2CellReselectionInfoCommonQHystConstraints)
				if err != nil {
					return err
				}
				c.Q_Hyst = v
			}
			if sIB2CellReselectionInfoCommonSeq.IsComponentPresent(4) {
				c.SpeedStateReselectionPars = &struct {
					MobilityStateParameters MobilityStateParameters
					Q_HystSF                struct {
						Sf_Medium int64
						Sf_High   int64
					}
				}{}
				c := (*c.SpeedStateReselectionPars)
				{
					if err := c.MobilityStateParameters.Decode(d); err != nil {
						return err
					}
				}
				{
					c := &c.Q_HystSF
					{
						v, err := d.DecodeEnumerated(sIB2CellReselectionInfoCommonSpeedStateReselectionParsQHystSFSfMediumConstraints)
						if err != nil {
							return err
						}
						c.Sf_Medium = v
					}
					{
						v, err := d.DecodeEnumerated(sIB2CellReselectionInfoCommonSpeedStateReselectionParsQHystSFSfHighConstraints)
						if err != nil {
							return err
						}
						c.Sf_High = v
					}
				}
			}
		}
	}

	// 3. cellReselectionServingFreqInfo: sequence
	{
		{
			c := &ie.CellReselectionServingFreqInfo
			sIB2CellReselectionServingFreqInfoSeq := d.NewSequenceDecoder(sIB2CellReselectionServingFreqInfoConstraints)
			if err := sIB2CellReselectionServingFreqInfoSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := sIB2CellReselectionServingFreqInfoSeq.DecodePreamble(); err != nil {
				return err
			}
			if sIB2CellReselectionServingFreqInfoSeq.IsComponentPresent(0) {
				c.S_NonIntraSearchP = new(ReselectionThreshold)
				if err := (*c.S_NonIntraSearchP).Decode(d); err != nil {
					return err
				}
			}
			if sIB2CellReselectionServingFreqInfoSeq.IsComponentPresent(1) {
				c.S_NonIntraSearchQ = new(ReselectionThresholdQ)
				if err := (*c.S_NonIntraSearchQ).Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.ThreshServingLowP.Decode(d); err != nil {
					return err
				}
			}
			if sIB2CellReselectionServingFreqInfoSeq.IsComponentPresent(3) {
				c.ThreshServingLowQ = new(ReselectionThresholdQ)
				if err := (*c.ThreshServingLowQ).Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.CellReselectionPriority.Decode(d); err != nil {
					return err
				}
			}
			if sIB2CellReselectionServingFreqInfoSeq.IsComponentPresent(5) {
				c.CellReselectionSubPriority = new(CellReselectionSubPriority)
				if err := (*c.CellReselectionSubPriority).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. intraFreqCellReselectionInfo: sequence
	{
		{
			c := &ie.IntraFreqCellReselectionInfo
			sIB2IntraFreqCellReselectionInfoSeq := d.NewSequenceDecoder(sIB2IntraFreqCellReselectionInfoConstraints)
			if err := sIB2IntraFreqCellReselectionInfoSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := sIB2IntraFreqCellReselectionInfoSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.Q_RxLevMin.Decode(d); err != nil {
					return err
				}
			}
			if sIB2IntraFreqCellReselectionInfoSeq.IsComponentPresent(1) {
				c.Q_RxLevMinSUL = new(Q_RxLevMin)
				if err := (*c.Q_RxLevMinSUL).Decode(d); err != nil {
					return err
				}
			}
			if sIB2IntraFreqCellReselectionInfoSeq.IsComponentPresent(2) {
				c.Q_QualMin = new(Q_QualMin)
				if err := (*c.Q_QualMin).Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.S_IntraSearchP.Decode(d); err != nil {
					return err
				}
			}
			if sIB2IntraFreqCellReselectionInfoSeq.IsComponentPresent(4) {
				c.S_IntraSearchQ = new(ReselectionThresholdQ)
				if err := (*c.S_IntraSearchQ).Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.T_ReselectionNR.Decode(d); err != nil {
					return err
				}
			}
			if sIB2IntraFreqCellReselectionInfoSeq.IsComponentPresent(6) {
				c.FrequencyBandList = new(MultiFrequencyBandListNR_SIB)
				if err := (*c.FrequencyBandList).Decode(d); err != nil {
					return err
				}
			}
			if sIB2IntraFreqCellReselectionInfoSeq.IsComponentPresent(7) {
				c.FrequencyBandListSUL = new(MultiFrequencyBandListNR_SIB)
				if err := (*c.FrequencyBandListSUL).Decode(d); err != nil {
					return err
				}
			}
			if sIB2IntraFreqCellReselectionInfoSeq.IsComponentPresent(8) {
				c.P_Max = new(P_Max)
				if err := (*c.P_Max).Decode(d); err != nil {
					return err
				}
			}
			if sIB2IntraFreqCellReselectionInfoSeq.IsComponentPresent(9) {
				c.Smtc = new(SSB_MTC)
				if err := (*c.Smtc).Decode(d); err != nil {
					return err
				}
			}
			if sIB2IntraFreqCellReselectionInfoSeq.IsComponentPresent(10) {
				c.Ss_RSSI_Measurement = new(SS_RSSI_Measurement)
				if err := (*c.Ss_RSSI_Measurement).Decode(d); err != nil {
					return err
				}
			}
			if sIB2IntraFreqCellReselectionInfoSeq.IsComponentPresent(11) {
				c.Ssb_ToMeasure = new(SSB_ToMeasure)
				if err := (*c.Ssb_ToMeasure).Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.DeriveSSB_IndexFromCell = v
			}
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
				{Name: "relaxedMeasurement-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.RelaxedMeasurement_r16 = &struct {
				LowMobilityEvaluation_r16 *struct {
					S_SearchDeltaP_r16 int64
					T_SearchDeltaP_r16 int64
				}
				CellEdgeEvaluation_r16 *struct {
					S_SearchThresholdP_r16 ReselectionThreshold
					S_SearchThresholdQ_r16 *ReselectionThresholdQ
				}
				CombineRelaxedMeasCondition_r16 *int64
				HighPriorityMeasRelax_r16       *int64
			}{}
			c := ie.RelaxedMeasurement_r16
			sIB2ExtRelaxedMeasurementR16Seq := dx.NewSequenceDecoder(sIB2ExtRelaxedMeasurementR16Constraints)
			if err := sIB2ExtRelaxedMeasurementR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if sIB2ExtRelaxedMeasurementR16Seq.IsComponentPresent(0) {
				c.LowMobilityEvaluation_r16 = &struct {
					S_SearchDeltaP_r16 int64
					T_SearchDeltaP_r16 int64
				}{}
				c := (*c.LowMobilityEvaluation_r16)
				{
					v, err := dx.DecodeEnumerated(sIB2ExtRelaxedMeasurementR16LowMobilityEvaluationR16SSearchDeltaPR16Constraints)
					if err != nil {
						return err
					}
					c.S_SearchDeltaP_r16 = v
				}
				{
					v, err := dx.DecodeEnumerated(sIB2ExtRelaxedMeasurementR16LowMobilityEvaluationR16TSearchDeltaPR16Constraints)
					if err != nil {
						return err
					}
					c.T_SearchDeltaP_r16 = v
				}
			}
			if sIB2ExtRelaxedMeasurementR16Seq.IsComponentPresent(1) {
				c.CellEdgeEvaluation_r16 = &struct {
					S_SearchThresholdP_r16 ReselectionThreshold
					S_SearchThresholdQ_r16 *ReselectionThresholdQ
				}{}
				c := (*c.CellEdgeEvaluation_r16)
				sIB2ExtRelaxedMeasurementR16CellEdgeEvaluationR16Seq := dx.NewSequenceDecoder(sIB2ExtRelaxedMeasurementR16CellEdgeEvaluationR16Constraints)
				if err := sIB2ExtRelaxedMeasurementR16CellEdgeEvaluationR16Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.S_SearchThresholdP_r16.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtRelaxedMeasurementR16CellEdgeEvaluationR16Seq.IsComponentPresent(1) {
					c.S_SearchThresholdQ_r16 = new(ReselectionThresholdQ)
					if err := (*c.S_SearchThresholdQ_r16).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtRelaxedMeasurementR16Seq.IsComponentPresent(2) {
				c.CombineRelaxedMeasCondition_r16 = new(int64)
				v, err := dx.DecodeEnumerated(sIB2ExtRelaxedMeasurementR16CombineRelaxedMeasConditionR16Constraints)
				if err != nil {
					return err
				}
				(*c.CombineRelaxedMeasCondition_r16) = v
			}
			if sIB2ExtRelaxedMeasurementR16Seq.IsComponentPresent(3) {
				c.HighPriorityMeasRelax_r16 = new(int64)
				v, err := dx.DecodeEnumerated(sIB2ExtRelaxedMeasurementR16HighPriorityMeasRelaxR16Constraints)
				if err != nil {
					return err
				}
				(*c.HighPriorityMeasRelax_r16) = v
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cellEquivalentSize-r17", Optional: true},
				{Name: "relaxedMeasurement-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sIB2CellEquivalentSizeR17Constraints)
			if err != nil {
				return err
			}
			ie.CellEquivalentSize_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.RelaxedMeasurement_r17 = &struct {
				StationaryMobilityEvaluation_r17 struct {
					S_SearchDeltaP_Stationary_r17 int64
					T_SearchDeltaP_Stationary_r17 int64
				}
				CellEdgeEvaluationWhileStationary_r17 *struct {
					S_SearchThresholdP2_r17 ReselectionThreshold
					S_SearchThresholdQ2_r17 *ReselectionThresholdQ
				}
				CombineRelaxedMeasCondition2_r17 *int64
			}{}
			c := ie.RelaxedMeasurement_r17
			sIB2ExtRelaxedMeasurementR17Seq := dx.NewSequenceDecoder(sIB2ExtRelaxedMeasurementR17Constraints)
			if err := sIB2ExtRelaxedMeasurementR17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				c := &c.StationaryMobilityEvaluation_r17
				{
					v, err := dx.DecodeEnumerated(sIB2ExtRelaxedMeasurementR17StationaryMobilityEvaluationR17SSearchDeltaPStationaryR17Constraints)
					if err != nil {
						return err
					}
					c.S_SearchDeltaP_Stationary_r17 = v
				}
				{
					v, err := dx.DecodeEnumerated(sIB2ExtRelaxedMeasurementR17StationaryMobilityEvaluationR17TSearchDeltaPStationaryR17Constraints)
					if err != nil {
						return err
					}
					c.T_SearchDeltaP_Stationary_r17 = v
				}
			}
			if sIB2ExtRelaxedMeasurementR17Seq.IsComponentPresent(1) {
				c.CellEdgeEvaluationWhileStationary_r17 = &struct {
					S_SearchThresholdP2_r17 ReselectionThreshold
					S_SearchThresholdQ2_r17 *ReselectionThresholdQ
				}{}
				c := (*c.CellEdgeEvaluationWhileStationary_r17)
				sIB2ExtRelaxedMeasurementR17CellEdgeEvaluationWhileStationaryR17Seq := dx.NewSequenceDecoder(sIB2ExtRelaxedMeasurementR17CellEdgeEvaluationWhileStationaryR17Constraints)
				if err := sIB2ExtRelaxedMeasurementR17CellEdgeEvaluationWhileStationaryR17Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.S_SearchThresholdP2_r17.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtRelaxedMeasurementR17CellEdgeEvaluationWhileStationaryR17Seq.IsComponentPresent(1) {
					c.S_SearchThresholdQ2_r17 = new(ReselectionThresholdQ)
					if err := (*c.S_SearchThresholdQ2_r17).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtRelaxedMeasurementR17Seq.IsComponentPresent(2) {
				c.CombineRelaxedMeasCondition2_r17 = new(int64)
				v, err := dx.DecodeEnumerated(sIB2ExtRelaxedMeasurementR17CombineRelaxedMeasCondition2R17Constraints)
				if err != nil {
					return err
				}
				(*c.CombineRelaxedMeasCondition2_r17) = v
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "uav-PrioritizedFrequency-r19", Optional: true},
				{Name: "uav-PrioritizedFrequencyAltitudeRange-r19", Optional: true},
				{Name: "ssb-ToMeasureAltitudeBasedList-r19", Optional: true},
				{Name: "relaxedMeasurementForServingAndNeighboringCell-r19", Optional: true},
				{Name: "offloadMeasurementForServingCell-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sIB2ExtUavPrioritizedFrequencyR19Constraints)
			if err != nil {
				return err
			}
			ie.Uav_PrioritizedFrequency_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Uav_PrioritizedFrequencyAltitudeRange_r19 = &struct {
				AltitudeMin_r19  *Altitude_r18
				AltitudeMax_r19  *Altitude_r18
				AltitudeHyst_r19 *HysteresisAltitude_r18
			}{}
			c := ie.Uav_PrioritizedFrequencyAltitudeRange_r19
			sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Seq := dx.NewSequenceDecoder(sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Constraints)
			if err := sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Seq.IsComponentPresent(0) {
				c.AltitudeMin_r19 = new(Altitude_r18)
				if err := (*c.AltitudeMin_r19).Decode(dx); err != nil {
					return err
				}
			}
			if sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Seq.IsComponentPresent(1) {
				c.AltitudeMax_r19 = new(Altitude_r18)
				if err := (*c.AltitudeMax_r19).Decode(dx); err != nil {
					return err
				}
			}
			if sIB2ExtUavPrioritizedFrequencyAltitudeRangeR19Seq.IsComponentPresent(2) {
				c.AltitudeHyst_r19 = new(HysteresisAltitude_r18)
				if err := (*c.AltitudeHyst_r19).Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ssb_ToMeasureAltitudeBasedList_r19 = new(SSB_ToMeasureAltitudeBasedList_r18)
			if err := ie.Ssb_ToMeasureAltitudeBasedList_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.RelaxedMeasurementForServingAndNeighboringCell_r19 = &struct {
				CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 *struct {
					S_SearchThresholdP3_r19 ReselectionThreshold
					S_SearchThresholdQ3_r19 *ReselectionThresholdQ
				}
				CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 *struct {
					S_SearchThresholdP4_r19 ReselectionThreshold
					S_SearchThresholdQ4_r19 *ReselectionThresholdQ
				}
				CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 *struct {
					RsrpThresholdLR_r19 ThresholdP_LR_r19
					RsrqThresholdLR_r19 *ThresholdQ_LR_r19
				}
				CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 *struct {
					RsrpThresholdLR2_r19 ThresholdP_LR_r19
					RsrqThresholdLR2_r19 *ThresholdQ_LR_r19
				}
			}{}
			c := ie.RelaxedMeasurementForServingAndNeighboringCell_r19
			sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Seq := dx.NewSequenceDecoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Constraints)
			if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Seq.IsComponentPresent(0) {
				c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 = &struct {
					S_SearchThresholdP3_r19 ReselectionThreshold
					S_SearchThresholdQ3_r19 *ReselectionThresholdQ
				}{}
				c := (*c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19)
				sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq := dx.NewSequenceDecoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnSSBR19Constraints)
				if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.S_SearchThresholdP3_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq.IsComponentPresent(1) {
					c.S_SearchThresholdQ3_r19 = new(ReselectionThresholdQ)
					if err := (*c.S_SearchThresholdQ3_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Seq.IsComponentPresent(1) {
				c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 = &struct {
					S_SearchThresholdP4_r19 ReselectionThreshold
					S_SearchThresholdQ4_r19 *ReselectionThresholdQ
				}{}
				c := (*c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19)
				sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq := dx.NewSequenceDecoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Constraints)
				if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.S_SearchThresholdP4_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq.IsComponentPresent(1) {
					c.S_SearchThresholdQ4_r19 = new(ReselectionThresholdQ)
					if err := (*c.S_SearchThresholdQ4_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Seq.IsComponentPresent(2) {
				c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 = &struct {
					RsrpThresholdLR_r19 ThresholdP_LR_r19
					RsrqThresholdLR_r19 *ThresholdQ_LR_r19
				}{}
				c := (*c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19)
				sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq := dx.NewSequenceDecoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Constraints)
				if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.RsrpThresholdLR_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq.IsComponentPresent(1) {
					c.RsrqThresholdLR_r19 = new(ThresholdQ_LR_r19)
					if err := (*c.RsrqThresholdLR_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19Seq.IsComponentPresent(3) {
				c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 = &struct {
					RsrpThresholdLR2_r19 ThresholdP_LR_r19
					RsrqThresholdLR2_r19 *ThresholdQ_LR_r19
				}{}
				c := (*c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19)
				sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq := dx.NewSequenceDecoder(sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnSSBR19Constraints)
				if err := sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.RsrpThresholdLR2_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtRelaxedMeasurementForServingAndNeighboringCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq.IsComponentPresent(1) {
					c.RsrqThresholdLR2_r19 = new(ThresholdQ_LR_r19)
					if err := (*c.RsrqThresholdLR2_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.OffloadMeasurementForServingCell_r19 = &struct {
				CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 *struct {
					S_SearchThresholdP5_r19 ReselectionThreshold
					S_SearchThresholdQ5_r19 *ReselectionThresholdQ
				}
				CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 *struct {
					S_SearchThresholdP6_r19 ReselectionThreshold
					S_SearchThresholdQ6_r19 *ReselectionThresholdQ
				}
				CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 *struct {
					RsrpThresholdLR3_r19 ThresholdP_LR_r19
					RsrqThresholdLR3_r19 *ThresholdQ_LR_r19
				}
				CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 *struct {
					RsrpThresholdLR4_r19 ThresholdP_LR_r19
					RsrqThresholdLR4_r19 *ThresholdQ_LR_r19
				}
				CellEdgeEvaluationOnLR_ForLR_OnLPSS_Exit_r19 *struct {
					RsrpThresholdLR5_r19 ThresholdP_LR_r19
					RsrqThresholdLR5_r19 *ThresholdQ_LR_r19
				}
				CellEdgeEvaluationOnLR_ForLR_OnSSB_Exit_r19 *struct {
					RsrpThresholdLR6_r19 ThresholdP_LR_r19
					RsrqThresholdLR6_r19 *ThresholdQ_LR_r19
				}
			}{}
			c := ie.OffloadMeasurementForServingCell_r19
			sIB2ExtOffloadMeasurementForServingCellR19Seq := dx.NewSequenceDecoder(sIB2ExtOffloadMeasurementForServingCellR19Constraints)
			if err := sIB2ExtOffloadMeasurementForServingCellR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if sIB2ExtOffloadMeasurementForServingCellR19Seq.IsComponentPresent(0) {
				c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19 = &struct {
					S_SearchThresholdP5_r19 ReselectionThreshold
					S_SearchThresholdQ5_r19 *ReselectionThresholdQ
				}{}
				c := (*c.CellEdgeEvaluationOnMR_ForLR_OnSSB_r19)
				sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq := dx.NewSequenceDecoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnSSBR19Constraints)
				if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.S_SearchThresholdP5_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnSSBR19Seq.IsComponentPresent(1) {
					c.S_SearchThresholdQ5_r19 = new(ReselectionThresholdQ)
					if err := (*c.S_SearchThresholdQ5_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtOffloadMeasurementForServingCellR19Seq.IsComponentPresent(1) {
				c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19 = &struct {
					S_SearchThresholdP6_r19 ReselectionThreshold
					S_SearchThresholdQ6_r19 *ReselectionThresholdQ
				}{}
				c := (*c.CellEdgeEvaluationOnMR_ForLR_OnLPSS_r19)
				sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq := dx.NewSequenceDecoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Constraints)
				if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.S_SearchThresholdP6_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnMRForLROnLPSSR19Seq.IsComponentPresent(1) {
					c.S_SearchThresholdQ6_r19 = new(ReselectionThresholdQ)
					if err := (*c.S_SearchThresholdQ6_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtOffloadMeasurementForServingCellR19Seq.IsComponentPresent(2) {
				c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19 = &struct {
					RsrpThresholdLR3_r19 ThresholdP_LR_r19
					RsrqThresholdLR3_r19 *ThresholdQ_LR_r19
				}{}
				c := (*c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_r19)
				sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq := dx.NewSequenceDecoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Constraints)
				if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.RsrpThresholdLR3_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSR19Seq.IsComponentPresent(1) {
					c.RsrqThresholdLR3_r19 = new(ThresholdQ_LR_r19)
					if err := (*c.RsrqThresholdLR3_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtOffloadMeasurementForServingCellR19Seq.IsComponentPresent(3) {
				c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19 = &struct {
					RsrpThresholdLR4_r19 ThresholdP_LR_r19
					RsrqThresholdLR4_r19 *ThresholdQ_LR_r19
				}{}
				c := (*c.CellEdgeEvaluationOnLR_ForLR_OnSSB_r19)
				sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq := dx.NewSequenceDecoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBR19Constraints)
				if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.RsrpThresholdLR4_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBR19Seq.IsComponentPresent(1) {
					c.RsrqThresholdLR4_r19 = new(ThresholdQ_LR_r19)
					if err := (*c.RsrqThresholdLR4_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtOffloadMeasurementForServingCellR19Seq.IsComponentPresent(4) {
				c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_Exit_r19 = &struct {
					RsrpThresholdLR5_r19 ThresholdP_LR_r19
					RsrqThresholdLR5_r19 *ThresholdQ_LR_r19
				}{}
				c := (*c.CellEdgeEvaluationOnLR_ForLR_OnLPSS_Exit_r19)
				sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSExitR19Seq := dx.NewSequenceDecoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSExitR19Constraints)
				if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSExitR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.RsrpThresholdLR5_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnLPSSExitR19Seq.IsComponentPresent(1) {
					c.RsrqThresholdLR5_r19 = new(ThresholdQ_LR_r19)
					if err := (*c.RsrqThresholdLR5_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
			if sIB2ExtOffloadMeasurementForServingCellR19Seq.IsComponentPresent(5) {
				c.CellEdgeEvaluationOnLR_ForLR_OnSSB_Exit_r19 = &struct {
					RsrpThresholdLR6_r19 ThresholdP_LR_r19
					RsrqThresholdLR6_r19 *ThresholdQ_LR_r19
				}{}
				c := (*c.CellEdgeEvaluationOnLR_ForLR_OnSSB_Exit_r19)
				sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBExitR19Seq := dx.NewSequenceDecoder(sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBExitR19Constraints)
				if err := sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBExitR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.RsrpThresholdLR6_r19.Decode(dx); err != nil {
						return err
					}
				}
				if sIB2ExtOffloadMeasurementForServingCellR19CellEdgeEvaluationOnLRForLROnSSBExitR19Seq.IsComponentPresent(1) {
					c.RsrqThresholdLR6_r19 = new(ThresholdQ_LR_r19)
					if err := (*c.RsrqThresholdLR6_r19).Decode(dx); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
