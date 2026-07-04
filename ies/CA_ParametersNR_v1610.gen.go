// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1610 (line 17305).

var cAParametersNRV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "parallelTxMsgA-SRS-PUCCH-PUSCH-r16", Optional: true},
		{Name: "msgA-SUL-r16", Optional: true},
		{Name: "jointSearchSpaceSwitchAcrossCells-r16", Optional: true},
		{Name: "half-DuplexTDD-CA-SameSCS-r16", Optional: true},
		{Name: "scellDormancyWithinActiveTime-r16", Optional: true},
		{Name: "scellDormancyOutsideActiveTime-r16", Optional: true},
		{Name: "crossCarrierA-CSI-trigDiffSCS-r16", Optional: true},
		{Name: "defaultQCL-CrossCarrierA-CSI-Trig-r16", Optional: true},
		{Name: "interCA-NonAlignedFrame-r16", Optional: true},
		{Name: "simul-SRS-Trans-BC-r16", Optional: true},
		{Name: "interFreqDAPS-r16", Optional: true},
		{Name: "codebookParametersPerBC-r16", Optional: true},
		{Name: "blindDetectFactor-r16", Optional: true},
		{Name: "pdcch-MonitoringCA-r16", Optional: true},
		{Name: "pdcch-BlindDetectionCA-Mixed-r16", Optional: true},
		{Name: "pdcch-BlindDetectionMCG-UE-r16", Optional: true},
		{Name: "pdcch-BlindDetectionSCG-UE-r16", Optional: true},
		{Name: "pdcch-BlindDetectionMCG-UE-Mixed-r16", Optional: true},
		{Name: "pdcch-BlindDetectionSCG-UE-Mixed-r16", Optional: true},
		{Name: "crossCarrierSchedulingDL-DiffSCS-r16", Optional: true},
		{Name: "crossCarrierSchedulingDefaultQCL-r16", Optional: true},
		{Name: "crossCarrierSchedulingUL-DiffSCS-r16", Optional: true},
		{Name: "simul-SRS-MIMO-Trans-BC-r16", Optional: true},
		{Name: "codebookParametersAdditionPerBC-r16", Optional: true},
		{Name: "codebookComboParametersAdditionPerBC-r16", Optional: true},
	},
}

const (
	CA_ParametersNR_v1610_ParallelTxMsgA_SRS_PUCCH_PUSCH_r16_Supported = 0
)

var cAParametersNRV1610ParallelTxMsgASRSPUCCHPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_ParallelTxMsgA_SRS_PUCCH_PUSCH_r16_Supported},
}

const (
	CA_ParametersNR_v1610_MsgA_SUL_r16_Supported = 0
)

var cAParametersNRV1610MsgASULR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_MsgA_SUL_r16_Supported},
}

const (
	CA_ParametersNR_v1610_JointSearchSpaceSwitchAcrossCells_r16_Supported = 0
)

var cAParametersNRV1610JointSearchSpaceSwitchAcrossCellsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_JointSearchSpaceSwitchAcrossCells_r16_Supported},
}

const (
	CA_ParametersNR_v1610_Half_DuplexTDD_CA_SameSCS_r16_Supported = 0
)

var cAParametersNRV1610HalfDuplexTDDCASameSCSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_Half_DuplexTDD_CA_SameSCS_r16_Supported},
}

const (
	CA_ParametersNR_v1610_ScellDormancyWithinActiveTime_r16_Supported = 0
)

var cAParametersNRV1610ScellDormancyWithinActiveTimeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_ScellDormancyWithinActiveTime_r16_Supported},
}

const (
	CA_ParametersNR_v1610_ScellDormancyOutsideActiveTime_r16_Supported = 0
)

var cAParametersNRV1610ScellDormancyOutsideActiveTimeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_ScellDormancyOutsideActiveTime_r16_Supported},
}

const (
	CA_ParametersNR_v1610_CrossCarrierA_CSI_TrigDiffSCS_r16_HigherA_CSI_SCS = 0
	CA_ParametersNR_v1610_CrossCarrierA_CSI_TrigDiffSCS_r16_LowerA_CSI_SCS  = 1
	CA_ParametersNR_v1610_CrossCarrierA_CSI_TrigDiffSCS_r16_Both            = 2
)

var cAParametersNRV1610CrossCarrierACSITrigDiffSCSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_CrossCarrierA_CSI_TrigDiffSCS_r16_HigherA_CSI_SCS, CA_ParametersNR_v1610_CrossCarrierA_CSI_TrigDiffSCS_r16_LowerA_CSI_SCS, CA_ParametersNR_v1610_CrossCarrierA_CSI_TrigDiffSCS_r16_Both},
}

const (
	CA_ParametersNR_v1610_DefaultQCL_CrossCarrierA_CSI_Trig_r16_DiffOnly = 0
	CA_ParametersNR_v1610_DefaultQCL_CrossCarrierA_CSI_Trig_r16_Both     = 1
)

var cAParametersNRV1610DefaultQCLCrossCarrierACSITrigR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_DefaultQCL_CrossCarrierA_CSI_Trig_r16_DiffOnly, CA_ParametersNR_v1610_DefaultQCL_CrossCarrierA_CSI_Trig_r16_Both},
}

const (
	CA_ParametersNR_v1610_InterCA_NonAlignedFrame_r16_Supported = 0
)

var cAParametersNRV1610InterCANonAlignedFrameR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_InterCA_NonAlignedFrame_r16_Supported},
}

const (
	CA_ParametersNR_v1610_Simul_SRS_Trans_BC_r16_N2 = 0
)

var cAParametersNRV1610SimulSRSTransBCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_Simul_SRS_Trans_BC_r16_N2},
}

var cAParametersNRV1610BlindDetectFactorR16Constraints = per.Constrained(1, 2)

var cAParametersNRV1610PdcchBlindDetectionMCGUER16Constraints = per.Constrained(1, 14)

var cAParametersNRV1610PdcchBlindDetectionSCGUER16Constraints = per.Constrained(1, 14)

const (
	CA_ParametersNR_v1610_CrossCarrierSchedulingDL_DiffSCS_r16_Low_To_High = 0
	CA_ParametersNR_v1610_CrossCarrierSchedulingDL_DiffSCS_r16_High_To_Low = 1
	CA_ParametersNR_v1610_CrossCarrierSchedulingDL_DiffSCS_r16_Both        = 2
)

var cAParametersNRV1610CrossCarrierSchedulingDLDiffSCSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_CrossCarrierSchedulingDL_DiffSCS_r16_Low_To_High, CA_ParametersNR_v1610_CrossCarrierSchedulingDL_DiffSCS_r16_High_To_Low, CA_ParametersNR_v1610_CrossCarrierSchedulingDL_DiffSCS_r16_Both},
}

const (
	CA_ParametersNR_v1610_CrossCarrierSchedulingDefaultQCL_r16_Diff_Only = 0
	CA_ParametersNR_v1610_CrossCarrierSchedulingDefaultQCL_r16_Both      = 1
)

var cAParametersNRV1610CrossCarrierSchedulingDefaultQCLR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_CrossCarrierSchedulingDefaultQCL_r16_Diff_Only, CA_ParametersNR_v1610_CrossCarrierSchedulingDefaultQCL_r16_Both},
}

const (
	CA_ParametersNR_v1610_CrossCarrierSchedulingUL_DiffSCS_r16_Low_To_High = 0
	CA_ParametersNR_v1610_CrossCarrierSchedulingUL_DiffSCS_r16_High_To_Low = 1
	CA_ParametersNR_v1610_CrossCarrierSchedulingUL_DiffSCS_r16_Both        = 2
)

var cAParametersNRV1610CrossCarrierSchedulingULDiffSCSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_CrossCarrierSchedulingUL_DiffSCS_r16_Low_To_High, CA_ParametersNR_v1610_CrossCarrierSchedulingUL_DiffSCS_r16_High_To_Low, CA_ParametersNR_v1610_CrossCarrierSchedulingUL_DiffSCS_r16_Both},
}

const (
	CA_ParametersNR_v1610_Simul_SRS_MIMO_Trans_BC_r16_N2 = 0
)

var cAParametersNRV1610SimulSRSMIMOTransBCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_Simul_SRS_MIMO_Trans_BC_r16_N2},
}

var cAParametersNRV1610InterFreqDAPSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "interFreqAsyncDAPS-r16", Optional: true},
		{Name: "interFreqDiffSCS-DAPS-r16", Optional: true},
		{Name: "interFreqMultiUL-TransmissionDAPS-r16", Optional: true},
		{Name: "interFreqSemiStaticPowerSharingDAPS-Mode1-r16", Optional: true},
		{Name: "interFreqSemiStaticPowerSharingDAPS-Mode2-r16", Optional: true},
		{Name: "interFreqDynamicPowerSharingDAPS-r16", Optional: true},
		{Name: "interFreqUL-TransCancellationDAPS-r16", Optional: true},
	},
}

const (
	CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqAsyncDAPS_r16_Supported = 0
)

var cAParametersNRV1610InterFreqDAPSR16InterFreqAsyncDAPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqAsyncDAPS_r16_Supported},
}

const (
	CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqDiffSCS_DAPS_r16_Supported = 0
)

var cAParametersNRV1610InterFreqDAPSR16InterFreqDiffSCSDAPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqDiffSCS_DAPS_r16_Supported},
}

const (
	CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqMultiUL_TransmissionDAPS_r16_Supported = 0
)

var cAParametersNRV1610InterFreqDAPSR16InterFreqMultiULTransmissionDAPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqMultiUL_TransmissionDAPS_r16_Supported},
}

const (
	CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqSemiStaticPowerSharingDAPS_Mode1_r16_Supported = 0
)

var cAParametersNRV1610InterFreqDAPSR16InterFreqSemiStaticPowerSharingDAPSMode1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqSemiStaticPowerSharingDAPS_Mode1_r16_Supported},
}

const (
	CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqSemiStaticPowerSharingDAPS_Mode2_r16_Supported = 0
)

var cAParametersNRV1610InterFreqDAPSR16InterFreqSemiStaticPowerSharingDAPSMode2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqSemiStaticPowerSharingDAPS_Mode2_r16_Supported},
}

const (
	CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqDynamicPowerSharingDAPS_r16_Short = 0
	CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqDynamicPowerSharingDAPS_r16_Long  = 1
)

var cAParametersNRV1610InterFreqDAPSR16InterFreqDynamicPowerSharingDAPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqDynamicPowerSharingDAPS_r16_Short, CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqDynamicPowerSharingDAPS_r16_Long},
}

const (
	CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqUL_TransCancellationDAPS_r16_Supported = 0
)

var cAParametersNRV1610InterFreqDAPSR16InterFreqULTransCancellationDAPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_InterFreqDAPS_r16_InterFreqUL_TransCancellationDAPS_r16_Supported},
}

const (
	CA_ParametersNR_v1610_Pdcch_MonitoringCA_r16_SupportedSpanArrangement_r16_AlignedOnly          = 0
	CA_ParametersNR_v1610_Pdcch_MonitoringCA_r16_SupportedSpanArrangement_r16_AlignedAndNonAligned = 1
)

var cAParametersNRV1610PdcchMonitoringCAR16SupportedSpanArrangementR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_Pdcch_MonitoringCA_r16_SupportedSpanArrangement_r16_AlignedOnly, CA_ParametersNR_v1610_Pdcch_MonitoringCA_r16_SupportedSpanArrangement_r16_AlignedAndNonAligned},
}

const (
	CA_ParametersNR_v1610_Pdcch_BlindDetectionCA_Mixed_r16_SupportedSpanArrangement_r16_AlignedOnly          = 0
	CA_ParametersNR_v1610_Pdcch_BlindDetectionCA_Mixed_r16_SupportedSpanArrangement_r16_AlignedAndNonAligned = 1
)

var cAParametersNRV1610PdcchBlindDetectionCAMixedR16SupportedSpanArrangementR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1610_Pdcch_BlindDetectionCA_Mixed_r16_SupportedSpanArrangement_r16_AlignedOnly, CA_ParametersNR_v1610_Pdcch_BlindDetectionCA_Mixed_r16_SupportedSpanArrangement_r16_AlignedAndNonAligned},
}

type CA_ParametersNR_v1610 struct {
	ParallelTxMsgA_SRS_PUCCH_PUSCH_r16    *int64
	MsgA_SUL_r16                          *int64
	JointSearchSpaceSwitchAcrossCells_r16 *int64
	Half_DuplexTDD_CA_SameSCS_r16         *int64
	ScellDormancyWithinActiveTime_r16     *int64
	ScellDormancyOutsideActiveTime_r16    *int64
	CrossCarrierA_CSI_TrigDiffSCS_r16     *int64
	DefaultQCL_CrossCarrierA_CSI_Trig_r16 *int64
	InterCA_NonAlignedFrame_r16           *int64
	Simul_SRS_Trans_BC_r16                *int64
	InterFreqDAPS_r16                     *struct {
		InterFreqAsyncDAPS_r16                        *int64
		InterFreqDiffSCS_DAPS_r16                     *int64
		InterFreqMultiUL_TransmissionDAPS_r16         *int64
		InterFreqSemiStaticPowerSharingDAPS_Mode1_r16 *int64
		InterFreqSemiStaticPowerSharingDAPS_Mode2_r16 *int64
		InterFreqDynamicPowerSharingDAPS_r16          *int64
		InterFreqUL_TransCancellationDAPS_r16         *int64
	}
	CodebookParametersPerBC_r16 *CodebookParameters_v1610
	BlindDetectFactor_r16       *int64
	Pdcch_MonitoringCA_r16      *struct {
		MaxNumberOfMonitoringCC_r16  int64
		SupportedSpanArrangement_r16 int64
	}
	Pdcch_BlindDetectionCA_Mixed_r16 *struct {
		Pdcch_BlindDetectionCA1_r16  int64
		Pdcch_BlindDetectionCA2_r16  int64
		SupportedSpanArrangement_r16 int64
	}
	Pdcch_BlindDetectionMCG_UE_r16       *int64
	Pdcch_BlindDetectionSCG_UE_r16       *int64
	Pdcch_BlindDetectionMCG_UE_Mixed_r16 *struct {
		Pdcch_BlindDetectionMCG_UE1_r16 int64
		Pdcch_BlindDetectionMCG_UE2_r16 int64
	}
	Pdcch_BlindDetectionSCG_UE_Mixed_r16 *struct {
		Pdcch_BlindDetectionSCG_UE1_r16 int64
		Pdcch_BlindDetectionSCG_UE2_r16 int64
	}
	CrossCarrierSchedulingDL_DiffSCS_r16     *int64
	CrossCarrierSchedulingDefaultQCL_r16     *int64
	CrossCarrierSchedulingUL_DiffSCS_r16     *int64
	Simul_SRS_MIMO_Trans_BC_r16              *int64
	CodebookParametersAdditionPerBC_r16      *CodebookParametersAdditionPerBC_r16
	CodebookComboParametersAdditionPerBC_r16 *CodebookComboParametersAdditionPerBC_r16
}

func (ie *CA_ParametersNR_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_r16 != nil, ie.MsgA_SUL_r16 != nil, ie.JointSearchSpaceSwitchAcrossCells_r16 != nil, ie.Half_DuplexTDD_CA_SameSCS_r16 != nil, ie.ScellDormancyWithinActiveTime_r16 != nil, ie.ScellDormancyOutsideActiveTime_r16 != nil, ie.CrossCarrierA_CSI_TrigDiffSCS_r16 != nil, ie.DefaultQCL_CrossCarrierA_CSI_Trig_r16 != nil, ie.InterCA_NonAlignedFrame_r16 != nil, ie.Simul_SRS_Trans_BC_r16 != nil, ie.InterFreqDAPS_r16 != nil, ie.CodebookParametersPerBC_r16 != nil, ie.BlindDetectFactor_r16 != nil, ie.Pdcch_MonitoringCA_r16 != nil, ie.Pdcch_BlindDetectionCA_Mixed_r16 != nil, ie.Pdcch_BlindDetectionMCG_UE_r16 != nil, ie.Pdcch_BlindDetectionSCG_UE_r16 != nil, ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16 != nil, ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16 != nil, ie.CrossCarrierSchedulingDL_DiffSCS_r16 != nil, ie.CrossCarrierSchedulingDefaultQCL_r16 != nil, ie.CrossCarrierSchedulingUL_DiffSCS_r16 != nil, ie.Simul_SRS_MIMO_Trans_BC_r16 != nil, ie.CodebookParametersAdditionPerBC_r16 != nil, ie.CodebookComboParametersAdditionPerBC_r16 != nil}); err != nil {
		return err
	}

	// 2. parallelTxMsgA-SRS-PUCCH-PUSCH-r16: enumerated
	{
		if ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_r16, cAParametersNRV1610ParallelTxMsgASRSPUCCHPUSCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. msgA-SUL-r16: enumerated
	{
		if ie.MsgA_SUL_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_SUL_r16, cAParametersNRV1610MsgASULR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. jointSearchSpaceSwitchAcrossCells-r16: enumerated
	{
		if ie.JointSearchSpaceSwitchAcrossCells_r16 != nil {
			if err := e.EncodeEnumerated(*ie.JointSearchSpaceSwitchAcrossCells_r16, cAParametersNRV1610JointSearchSpaceSwitchAcrossCellsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. half-DuplexTDD-CA-SameSCS-r16: enumerated
	{
		if ie.Half_DuplexTDD_CA_SameSCS_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Half_DuplexTDD_CA_SameSCS_r16, cAParametersNRV1610HalfDuplexTDDCASameSCSR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. scellDormancyWithinActiveTime-r16: enumerated
	{
		if ie.ScellDormancyWithinActiveTime_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ScellDormancyWithinActiveTime_r16, cAParametersNRV1610ScellDormancyWithinActiveTimeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. scellDormancyOutsideActiveTime-r16: enumerated
	{
		if ie.ScellDormancyOutsideActiveTime_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ScellDormancyOutsideActiveTime_r16, cAParametersNRV1610ScellDormancyOutsideActiveTimeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. crossCarrierA-CSI-trigDiffSCS-r16: enumerated
	{
		if ie.CrossCarrierA_CSI_TrigDiffSCS_r16 != nil {
			if err := e.EncodeEnumerated(*ie.CrossCarrierA_CSI_TrigDiffSCS_r16, cAParametersNRV1610CrossCarrierACSITrigDiffSCSR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. defaultQCL-CrossCarrierA-CSI-Trig-r16: enumerated
	{
		if ie.DefaultQCL_CrossCarrierA_CSI_Trig_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DefaultQCL_CrossCarrierA_CSI_Trig_r16, cAParametersNRV1610DefaultQCLCrossCarrierACSITrigR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. interCA-NonAlignedFrame-r16: enumerated
	{
		if ie.InterCA_NonAlignedFrame_r16 != nil {
			if err := e.EncodeEnumerated(*ie.InterCA_NonAlignedFrame_r16, cAParametersNRV1610InterCANonAlignedFrameR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. simul-SRS-Trans-BC-r16: enumerated
	{
		if ie.Simul_SRS_Trans_BC_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Simul_SRS_Trans_BC_r16, cAParametersNRV1610SimulSRSTransBCR16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. interFreqDAPS-r16: sequence
	{
		if ie.InterFreqDAPS_r16 != nil {
			c := ie.InterFreqDAPS_r16
			cAParametersNRV1610InterFreqDAPSR16Seq := e.NewSequenceEncoder(cAParametersNRV1610InterFreqDAPSR16Constraints)
			if err := cAParametersNRV1610InterFreqDAPSR16Seq.EncodePreamble([]bool{c.InterFreqAsyncDAPS_r16 != nil, c.InterFreqDiffSCS_DAPS_r16 != nil, c.InterFreqMultiUL_TransmissionDAPS_r16 != nil, c.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16 != nil, c.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16 != nil, c.InterFreqDynamicPowerSharingDAPS_r16 != nil, c.InterFreqUL_TransCancellationDAPS_r16 != nil}); err != nil {
				return err
			}
			if c.InterFreqAsyncDAPS_r16 != nil {
				if err := e.EncodeEnumerated((*c.InterFreqAsyncDAPS_r16), cAParametersNRV1610InterFreqDAPSR16InterFreqAsyncDAPSR16Constraints); err != nil {
					return err
				}
			}
			if c.InterFreqDiffSCS_DAPS_r16 != nil {
				if err := e.EncodeEnumerated((*c.InterFreqDiffSCS_DAPS_r16), cAParametersNRV1610InterFreqDAPSR16InterFreqDiffSCSDAPSR16Constraints); err != nil {
					return err
				}
			}
			if c.InterFreqMultiUL_TransmissionDAPS_r16 != nil {
				if err := e.EncodeEnumerated((*c.InterFreqMultiUL_TransmissionDAPS_r16), cAParametersNRV1610InterFreqDAPSR16InterFreqMultiULTransmissionDAPSR16Constraints); err != nil {
					return err
				}
			}
			if c.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16 != nil {
				if err := e.EncodeEnumerated((*c.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16), cAParametersNRV1610InterFreqDAPSR16InterFreqSemiStaticPowerSharingDAPSMode1R16Constraints); err != nil {
					return err
				}
			}
			if c.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16 != nil {
				if err := e.EncodeEnumerated((*c.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16), cAParametersNRV1610InterFreqDAPSR16InterFreqSemiStaticPowerSharingDAPSMode2R16Constraints); err != nil {
					return err
				}
			}
			if c.InterFreqDynamicPowerSharingDAPS_r16 != nil {
				if err := e.EncodeEnumerated((*c.InterFreqDynamicPowerSharingDAPS_r16), cAParametersNRV1610InterFreqDAPSR16InterFreqDynamicPowerSharingDAPSR16Constraints); err != nil {
					return err
				}
			}
			if c.InterFreqUL_TransCancellationDAPS_r16 != nil {
				if err := e.EncodeEnumerated((*c.InterFreqUL_TransCancellationDAPS_r16), cAParametersNRV1610InterFreqDAPSR16InterFreqULTransCancellationDAPSR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 13. codebookParametersPerBC-r16: ref
	{
		if ie.CodebookParametersPerBC_r16 != nil {
			if err := ie.CodebookParametersPerBC_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 14. blindDetectFactor-r16: integer
	{
		if ie.BlindDetectFactor_r16 != nil {
			if err := e.EncodeInteger(*ie.BlindDetectFactor_r16, cAParametersNRV1610BlindDetectFactorR16Constraints); err != nil {
				return err
			}
		}
	}

	// 15. pdcch-MonitoringCA-r16: sequence
	{
		if ie.Pdcch_MonitoringCA_r16 != nil {
			c := ie.Pdcch_MonitoringCA_r16
			if err := e.EncodeInteger(c.MaxNumberOfMonitoringCC_r16, per.Constrained(2, 16)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.SupportedSpanArrangement_r16, cAParametersNRV1610PdcchMonitoringCAR16SupportedSpanArrangementR16Constraints); err != nil {
				return err
			}
		}
	}

	// 16. pdcch-BlindDetectionCA-Mixed-r16: sequence
	{
		if ie.Pdcch_BlindDetectionCA_Mixed_r16 != nil {
			c := ie.Pdcch_BlindDetectionCA_Mixed_r16
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionCA1_r16, per.Constrained(1, 15)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionCA2_r16, per.Constrained(1, 15)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.SupportedSpanArrangement_r16, cAParametersNRV1610PdcchBlindDetectionCAMixedR16SupportedSpanArrangementR16Constraints); err != nil {
				return err
			}
		}
	}

	// 17. pdcch-BlindDetectionMCG-UE-r16: integer
	{
		if ie.Pdcch_BlindDetectionMCG_UE_r16 != nil {
			if err := e.EncodeInteger(*ie.Pdcch_BlindDetectionMCG_UE_r16, cAParametersNRV1610PdcchBlindDetectionMCGUER16Constraints); err != nil {
				return err
			}
		}
	}

	// 18. pdcch-BlindDetectionSCG-UE-r16: integer
	{
		if ie.Pdcch_BlindDetectionSCG_UE_r16 != nil {
			if err := e.EncodeInteger(*ie.Pdcch_BlindDetectionSCG_UE_r16, cAParametersNRV1610PdcchBlindDetectionSCGUER16Constraints); err != nil {
				return err
			}
		}
	}

	// 19. pdcch-BlindDetectionMCG-UE-Mixed-r16: sequence
	{
		if ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16 != nil {
			c := ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionMCG_UE1_r16, per.Constrained(0, 15)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionMCG_UE2_r16, per.Constrained(0, 15)); err != nil {
				return err
			}
		}
	}

	// 20. pdcch-BlindDetectionSCG-UE-Mixed-r16: sequence
	{
		if ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16 != nil {
			c := ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionSCG_UE1_r16, per.Constrained(0, 15)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionSCG_UE2_r16, per.Constrained(0, 15)); err != nil {
				return err
			}
		}
	}

	// 21. crossCarrierSchedulingDL-DiffSCS-r16: enumerated
	{
		if ie.CrossCarrierSchedulingDL_DiffSCS_r16 != nil {
			if err := e.EncodeEnumerated(*ie.CrossCarrierSchedulingDL_DiffSCS_r16, cAParametersNRV1610CrossCarrierSchedulingDLDiffSCSR16Constraints); err != nil {
				return err
			}
		}
	}

	// 22. crossCarrierSchedulingDefaultQCL-r16: enumerated
	{
		if ie.CrossCarrierSchedulingDefaultQCL_r16 != nil {
			if err := e.EncodeEnumerated(*ie.CrossCarrierSchedulingDefaultQCL_r16, cAParametersNRV1610CrossCarrierSchedulingDefaultQCLR16Constraints); err != nil {
				return err
			}
		}
	}

	// 23. crossCarrierSchedulingUL-DiffSCS-r16: enumerated
	{
		if ie.CrossCarrierSchedulingUL_DiffSCS_r16 != nil {
			if err := e.EncodeEnumerated(*ie.CrossCarrierSchedulingUL_DiffSCS_r16, cAParametersNRV1610CrossCarrierSchedulingULDiffSCSR16Constraints); err != nil {
				return err
			}
		}
	}

	// 24. simul-SRS-MIMO-Trans-BC-r16: enumerated
	{
		if ie.Simul_SRS_MIMO_Trans_BC_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Simul_SRS_MIMO_Trans_BC_r16, cAParametersNRV1610SimulSRSMIMOTransBCR16Constraints); err != nil {
				return err
			}
		}
	}

	// 25. codebookParametersAdditionPerBC-r16: ref
	{
		if ie.CodebookParametersAdditionPerBC_r16 != nil {
			if err := ie.CodebookParametersAdditionPerBC_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 26. codebookComboParametersAdditionPerBC-r16: ref
	{
		if ie.CodebookComboParametersAdditionPerBC_r16 != nil {
			if err := ie.CodebookComboParametersAdditionPerBC_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. parallelTxMsgA-SRS-PUCCH-PUSCH-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610ParallelTxMsgASRSPUCCHPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_r16 = &idx
		}
	}

	// 3. msgA-SUL-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610MsgASULR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_SUL_r16 = &idx
		}
	}

	// 4. jointSearchSpaceSwitchAcrossCells-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610JointSearchSpaceSwitchAcrossCellsR16Constraints)
			if err != nil {
				return err
			}
			ie.JointSearchSpaceSwitchAcrossCells_r16 = &idx
		}
	}

	// 5. half-DuplexTDD-CA-SameSCS-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610HalfDuplexTDDCASameSCSR16Constraints)
			if err != nil {
				return err
			}
			ie.Half_DuplexTDD_CA_SameSCS_r16 = &idx
		}
	}

	// 6. scellDormancyWithinActiveTime-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610ScellDormancyWithinActiveTimeR16Constraints)
			if err != nil {
				return err
			}
			ie.ScellDormancyWithinActiveTime_r16 = &idx
		}
	}

	// 7. scellDormancyOutsideActiveTime-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610ScellDormancyOutsideActiveTimeR16Constraints)
			if err != nil {
				return err
			}
			ie.ScellDormancyOutsideActiveTime_r16 = &idx
		}
	}

	// 8. crossCarrierA-CSI-trigDiffSCS-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610CrossCarrierACSITrigDiffSCSR16Constraints)
			if err != nil {
				return err
			}
			ie.CrossCarrierA_CSI_TrigDiffSCS_r16 = &idx
		}
	}

	// 9. defaultQCL-CrossCarrierA-CSI-Trig-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610DefaultQCLCrossCarrierACSITrigR16Constraints)
			if err != nil {
				return err
			}
			ie.DefaultQCL_CrossCarrierA_CSI_Trig_r16 = &idx
		}
	}

	// 10. interCA-NonAlignedFrame-r16: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610InterCANonAlignedFrameR16Constraints)
			if err != nil {
				return err
			}
			ie.InterCA_NonAlignedFrame_r16 = &idx
		}
	}

	// 11. simul-SRS-Trans-BC-r16: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610SimulSRSTransBCR16Constraints)
			if err != nil {
				return err
			}
			ie.Simul_SRS_Trans_BC_r16 = &idx
		}
	}

	// 12. interFreqDAPS-r16: sequence
	{
		if seq.IsComponentPresent(10) {
			ie.InterFreqDAPS_r16 = &struct {
				InterFreqAsyncDAPS_r16                        *int64
				InterFreqDiffSCS_DAPS_r16                     *int64
				InterFreqMultiUL_TransmissionDAPS_r16         *int64
				InterFreqSemiStaticPowerSharingDAPS_Mode1_r16 *int64
				InterFreqSemiStaticPowerSharingDAPS_Mode2_r16 *int64
				InterFreqDynamicPowerSharingDAPS_r16          *int64
				InterFreqUL_TransCancellationDAPS_r16         *int64
			}{}
			c := ie.InterFreqDAPS_r16
			cAParametersNRV1610InterFreqDAPSR16Seq := d.NewSequenceDecoder(cAParametersNRV1610InterFreqDAPSR16Constraints)
			if err := cAParametersNRV1610InterFreqDAPSR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if cAParametersNRV1610InterFreqDAPSR16Seq.IsComponentPresent(0) {
				c.InterFreqAsyncDAPS_r16 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1610InterFreqDAPSR16InterFreqAsyncDAPSR16Constraints)
				if err != nil {
					return err
				}
				(*c.InterFreqAsyncDAPS_r16) = v
			}
			if cAParametersNRV1610InterFreqDAPSR16Seq.IsComponentPresent(1) {
				c.InterFreqDiffSCS_DAPS_r16 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1610InterFreqDAPSR16InterFreqDiffSCSDAPSR16Constraints)
				if err != nil {
					return err
				}
				(*c.InterFreqDiffSCS_DAPS_r16) = v
			}
			if cAParametersNRV1610InterFreqDAPSR16Seq.IsComponentPresent(2) {
				c.InterFreqMultiUL_TransmissionDAPS_r16 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1610InterFreqDAPSR16InterFreqMultiULTransmissionDAPSR16Constraints)
				if err != nil {
					return err
				}
				(*c.InterFreqMultiUL_TransmissionDAPS_r16) = v
			}
			if cAParametersNRV1610InterFreqDAPSR16Seq.IsComponentPresent(3) {
				c.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1610InterFreqDAPSR16InterFreqSemiStaticPowerSharingDAPSMode1R16Constraints)
				if err != nil {
					return err
				}
				(*c.InterFreqSemiStaticPowerSharingDAPS_Mode1_r16) = v
			}
			if cAParametersNRV1610InterFreqDAPSR16Seq.IsComponentPresent(4) {
				c.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1610InterFreqDAPSR16InterFreqSemiStaticPowerSharingDAPSMode2R16Constraints)
				if err != nil {
					return err
				}
				(*c.InterFreqSemiStaticPowerSharingDAPS_Mode2_r16) = v
			}
			if cAParametersNRV1610InterFreqDAPSR16Seq.IsComponentPresent(5) {
				c.InterFreqDynamicPowerSharingDAPS_r16 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1610InterFreqDAPSR16InterFreqDynamicPowerSharingDAPSR16Constraints)
				if err != nil {
					return err
				}
				(*c.InterFreqDynamicPowerSharingDAPS_r16) = v
			}
			if cAParametersNRV1610InterFreqDAPSR16Seq.IsComponentPresent(6) {
				c.InterFreqUL_TransCancellationDAPS_r16 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1610InterFreqDAPSR16InterFreqULTransCancellationDAPSR16Constraints)
				if err != nil {
					return err
				}
				(*c.InterFreqUL_TransCancellationDAPS_r16) = v
			}
		}
	}

	// 13. codebookParametersPerBC-r16: ref
	{
		if seq.IsComponentPresent(11) {
			ie.CodebookParametersPerBC_r16 = new(CodebookParameters_v1610)
			if err := ie.CodebookParametersPerBC_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 14. blindDetectFactor-r16: integer
	{
		if seq.IsComponentPresent(12) {
			v, err := d.DecodeInteger(cAParametersNRV1610BlindDetectFactorR16Constraints)
			if err != nil {
				return err
			}
			ie.BlindDetectFactor_r16 = &v
		}
	}

	// 15. pdcch-MonitoringCA-r16: sequence
	{
		if seq.IsComponentPresent(13) {
			ie.Pdcch_MonitoringCA_r16 = &struct {
				MaxNumberOfMonitoringCC_r16  int64
				SupportedSpanArrangement_r16 int64
			}{}
			c := ie.Pdcch_MonitoringCA_r16
			{
				v, err := d.DecodeInteger(per.Constrained(2, 16))
				if err != nil {
					return err
				}
				c.MaxNumberOfMonitoringCC_r16 = v
			}
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1610PdcchMonitoringCAR16SupportedSpanArrangementR16Constraints)
				if err != nil {
					return err
				}
				c.SupportedSpanArrangement_r16 = v
			}
		}
	}

	// 16. pdcch-BlindDetectionCA-Mixed-r16: sequence
	{
		if seq.IsComponentPresent(14) {
			ie.Pdcch_BlindDetectionCA_Mixed_r16 = &struct {
				Pdcch_BlindDetectionCA1_r16  int64
				Pdcch_BlindDetectionCA2_r16  int64
				SupportedSpanArrangement_r16 int64
			}{}
			c := ie.Pdcch_BlindDetectionCA_Mixed_r16
			{
				v, err := d.DecodeInteger(per.Constrained(1, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionCA1_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionCA2_r16 = v
			}
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1610PdcchBlindDetectionCAMixedR16SupportedSpanArrangementR16Constraints)
				if err != nil {
					return err
				}
				c.SupportedSpanArrangement_r16 = v
			}
		}
	}

	// 17. pdcch-BlindDetectionMCG-UE-r16: integer
	{
		if seq.IsComponentPresent(15) {
			v, err := d.DecodeInteger(cAParametersNRV1610PdcchBlindDetectionMCGUER16Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionMCG_UE_r16 = &v
		}
	}

	// 18. pdcch-BlindDetectionSCG-UE-r16: integer
	{
		if seq.IsComponentPresent(16) {
			v, err := d.DecodeInteger(cAParametersNRV1610PdcchBlindDetectionSCGUER16Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionSCG_UE_r16 = &v
		}
	}

	// 19. pdcch-BlindDetectionMCG-UE-Mixed-r16: sequence
	{
		if seq.IsComponentPresent(17) {
			ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16 = &struct {
				Pdcch_BlindDetectionMCG_UE1_r16 int64
				Pdcch_BlindDetectionMCG_UE2_r16 int64
			}{}
			c := ie.Pdcch_BlindDetectionMCG_UE_Mixed_r16
			{
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionMCG_UE1_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionMCG_UE2_r16 = v
			}
		}
	}

	// 20. pdcch-BlindDetectionSCG-UE-Mixed-r16: sequence
	{
		if seq.IsComponentPresent(18) {
			ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16 = &struct {
				Pdcch_BlindDetectionSCG_UE1_r16 int64
				Pdcch_BlindDetectionSCG_UE2_r16 int64
			}{}
			c := ie.Pdcch_BlindDetectionSCG_UE_Mixed_r16
			{
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionSCG_UE1_r16 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionSCG_UE2_r16 = v
			}
		}
	}

	// 21. crossCarrierSchedulingDL-DiffSCS-r16: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610CrossCarrierSchedulingDLDiffSCSR16Constraints)
			if err != nil {
				return err
			}
			ie.CrossCarrierSchedulingDL_DiffSCS_r16 = &idx
		}
	}

	// 22. crossCarrierSchedulingDefaultQCL-r16: enumerated
	{
		if seq.IsComponentPresent(20) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610CrossCarrierSchedulingDefaultQCLR16Constraints)
			if err != nil {
				return err
			}
			ie.CrossCarrierSchedulingDefaultQCL_r16 = &idx
		}
	}

	// 23. crossCarrierSchedulingUL-DiffSCS-r16: enumerated
	{
		if seq.IsComponentPresent(21) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610CrossCarrierSchedulingULDiffSCSR16Constraints)
			if err != nil {
				return err
			}
			ie.CrossCarrierSchedulingUL_DiffSCS_r16 = &idx
		}
	}

	// 24. simul-SRS-MIMO-Trans-BC-r16: enumerated
	{
		if seq.IsComponentPresent(22) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1610SimulSRSMIMOTransBCR16Constraints)
			if err != nil {
				return err
			}
			ie.Simul_SRS_MIMO_Trans_BC_r16 = &idx
		}
	}

	// 25. codebookParametersAdditionPerBC-r16: ref
	{
		if seq.IsComponentPresent(23) {
			ie.CodebookParametersAdditionPerBC_r16 = new(CodebookParametersAdditionPerBC_r16)
			if err := ie.CodebookParametersAdditionPerBC_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 26. codebookComboParametersAdditionPerBC-r16: ref
	{
		if seq.IsComponentPresent(24) {
			ie.CodebookComboParametersAdditionPerBC_r16 = new(CodebookComboParametersAdditionPerBC_r16)
			if err := ie.CodebookComboParametersAdditionPerBC_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
