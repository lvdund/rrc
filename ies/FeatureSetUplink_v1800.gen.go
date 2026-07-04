// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetUplink-v1800 (line 20309).

var featureSetUplinkV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxDelayValueBeyondD-Basic-r18", Optional: true},
		{Name: "tdcp-NumberDelayValue-r18", Optional: true},
		{Name: "phaseReportMoreThanOne-r18", Optional: true},
		{Name: "maxNumberTRS-ResourceSet-r18", Optional: true},
		{Name: "maxNumberTDCP-PerBWP-r18", Optional: true},
		{Name: "pusch-DMRS-TypeEnh-r18", Optional: true},
		{Name: "ul-DMRS-SingleDCI-M-TRP-r18", Optional: true},
		{Name: "ul-DMRS-M-DCI-M-TRP-r18", Optional: true},
		{Name: "srs-AntennaSwitching8T8R2SP-1Periodic-r18", Optional: true},
		{Name: "pucch-SingleDCI-STx2P-SFN-r18", Optional: true},
		{Name: "posSRS-BWA-RRC-Connected-r18", Optional: true},
		{Name: "posSRS-BWA-IndependentCA-RRC-Connected-r18", Optional: true},
		{Name: "posSRS-BWA-AffectedBandList-r18", Optional: true},
		{Name: "rach-EarlyTA-BandList-r18", Optional: true},
		{Name: "simultaneous-2-1-HARQ-ACK-CB-r18", Optional: true},
		{Name: "simultaneous-2-2-HARQ-ACK-CB-r18", Optional: true},
		{Name: "ul-IntraUE-MuxEnh-r18", Optional: true},
		{Name: "txDiversity4Tx-r18", Optional: true},
		{Name: "powerBoosting-pi2BPSK-QPSK-r18", Optional: true},
		{Name: "powerBoosting-pi2BPSK-QPSK-Modified-r18", Optional: true},
		{Name: "txDiversity2Tx-r18", Optional: true},
		{Name: "ue-PowerClassPerBandPerBC-v1820", Optional: true},
	},
}

const (
	FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl2  = 0
	FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl3  = 1
	FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl4  = 2
	FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl5  = 3
	FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl6  = 4
	FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl10 = 5
)

var featureSetUplinkV1800MaxDelayValueBeyondDBasicR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl2, FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl3, FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl4, FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl5, FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl6, FeatureSetUplink_v1800_MaxDelayValueBeyondD_Basic_r18_Sl10},
}

var featureSetUplinkV1800TdcpNumberDelayValueR18Constraints = per.Constrained(2, 4)

const (
	FeatureSetUplink_v1800_PhaseReportMoreThanOne_r18_Supported = 0
)

var featureSetUplinkV1800PhaseReportMoreThanOneR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_PhaseReportMoreThanOne_r18_Supported},
}

var featureSetUplinkV1800MaxNumberTRSResourceSetR18Constraints = per.Constrained(2, 3)

var featureSetUplinkV1800MaxNumberTDCPPerBWPR18Constraints = per.Constrained(1, 4)

const (
	FeatureSetUplink_v1800_Ul_DMRS_SingleDCI_M_TRP_r18_Supported = 0
)

var featureSetUplinkV1800UlDMRSSingleDCIMTRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Ul_DMRS_SingleDCI_M_TRP_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Ul_DMRS_M_DCI_M_TRP_r18_Supported = 0
)

var featureSetUplinkV1800UlDMRSMDCIMTRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Ul_DMRS_M_DCI_M_TRP_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Srs_AntennaSwitching8T8R2SP_1Periodic_r18_Supported = 0
)

var featureSetUplinkV1800SrsAntennaSwitching8T8R2SP1PeriodicR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Srs_AntennaSwitching8T8R2SP_1Periodic_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pucch_SingleDCI_STx2P_SFN_r18_Pf0_2   = 0
	FeatureSetUplink_v1800_Pucch_SingleDCI_STx2P_SFN_r18_Pf1_3_4 = 1
	FeatureSetUplink_v1800_Pucch_SingleDCI_STx2P_SFN_r18_Pf0_4   = 2
)

var featureSetUplinkV1800PucchSingleDCISTx2PSFNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pucch_SingleDCI_STx2P_SFN_r18_Pf0_2, FeatureSetUplink_v1800_Pucch_SingleDCI_STx2P_SFN_r18_Pf1_3_4, FeatureSetUplink_v1800_Pucch_SingleDCI_STx2P_SFN_r18_Pf0_4},
}

var featureSetUplinkV1800PosSRSBWAAffectedBandListR18Constraints = per.SizeRange(1, common.MaxBands)

var featureSetUplinkV1800RachEarlyTABandListR18Constraints = per.SizeRange(1, common.MaxBandsMRDC)

const (
	FeatureSetUplink_v1800_TxDiversity4Tx_r18_Supported = 0
)

var featureSetUplinkV1800TxDiversity4TxR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_TxDiversity4Tx_r18_Supported},
}

const (
	FeatureSetUplink_v1800_PowerBoosting_Pi2BPSK_QPSK_r18_Supported = 0
)

var featureSetUplinkV1800PowerBoostingPi2BPSKQPSKR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_PowerBoosting_Pi2BPSK_QPSK_r18_Supported},
}

const (
	FeatureSetUplink_v1800_PowerBoosting_Pi2BPSK_QPSK_Modified_r18_Supported = 0
)

var featureSetUplinkV1800PowerBoostingPi2BPSKQPSKModifiedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_PowerBoosting_Pi2BPSK_QPSK_Modified_r18_Supported},
}

const (
	FeatureSetUplink_v1800_TxDiversity2Tx_r18_Supported = 0
)

var featureSetUplinkV1800TxDiversity2TxR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_TxDiversity2Tx_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Ue_PowerClassPerBandPerBC_v1820_Pc5 = 0
)

var featureSetUplinkV1800UePowerClassPerBandPerBCV1820Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Ue_PowerClassPerBandPerBC_v1820_Pc5},
}

var featureSetUplinkV1800PuschDMRSTypeEnhR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dmrs-Type-r18"},
		{Name: "pusch-TypeA-DMRS-r18", Optional: true},
		{Name: "dummy", Optional: true},
		{Name: "pusch-TypeB-DMRS-r18", Optional: true},
		{Name: "pusch-rank-1-4-1Port-r18", Optional: true},
		{Name: "pusch-rank-5-8-1Port-r18", Optional: true},
		{Name: "pusch-rank-1-4-2Port-r18", Optional: true},
		{Name: "pusch-rank-5-8-2Port-r18", Optional: true},
	},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Dmrs_Type_r18_Etype1 = 0
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Dmrs_Type_r18_Both   = 1
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18DmrsTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Dmrs_Type_r18_Etype1, FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Dmrs_Type_r18_Both},
}

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dmrs-TypeA-r18"},
		{Name: "pusch-2SymbolFL-DMRS-r18", Optional: true},
		{Name: "pusch-2SymbolFL-DMRS-Addition2Symbol-r18", Optional: true},
		{Name: "pusch-1SymbolFL-DMRS-Addition3Symbol-r18", Optional: true},
		{Name: "pusch-1SymbolFL-DMRS-BeyondOnePort-r18", Optional: true},
	},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Dmrs_TypeA_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18DmrsTypeAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Dmrs_TypeA_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Pusch_2SymbolFL_DMRS_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch2SymbolFLDMRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Pusch_2SymbolFL_DMRS_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Pusch_2SymbolFL_DMRS_Addition2Symbol_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch2SymbolFLDMRSAddition2SymbolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Pusch_2SymbolFL_DMRS_Addition2Symbol_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Pusch_1SymbolFL_DMRS_Addition3Symbol_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch1SymbolFLDMRSAddition3SymbolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Pusch_1SymbolFL_DMRS_Addition3Symbol_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Pusch_1SymbolFL_DMRS_BeyondOnePort_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch1SymbolFLDMRSBeyondOnePortR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeA_DMRS_r18_Pusch_1SymbolFL_DMRS_BeyondOnePort_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Dummy_Rel15 = 0
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Dummy_Both  = 1
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18DummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Dummy_Rel15, FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Dummy_Both},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeB_DMRS_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeBDMRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_TypeB_DMRS_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_Rank_1_4_1Port_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank141PortR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_Rank_1_4_1Port_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_Rank_5_8_1Port_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank581PortR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_Rank_5_8_1Port_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_Rank_1_4_2Port_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank142PortR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_Rank_1_4_2Port_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_Rank_5_8_2Port_r18_Supported = 0
)

var featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank582PortR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Pusch_DMRS_TypeEnh_r18_Pusch_Rank_5_8_2Port_r18_Supported},
}

const (
	FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationLowPriority_r18_Sym0 = 0
	FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationLowPriority_r18_Sym1 = 1
	FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationLowPriority_r18_Sym2 = 2
)

var featureSetUplinkV1800UlIntraUEMuxEnhR18PuschPreparationLowPriorityR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationLowPriority_r18_Sym0, FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationLowPriority_r18_Sym1, FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationLowPriority_r18_Sym2},
}

const (
	FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationHighPriority_r18_Sym0 = 0
	FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationHighPriority_r18_Sym1 = 1
	FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationHighPriority_r18_Sym2 = 2
)

var featureSetUplinkV1800UlIntraUEMuxEnhR18PuschPreparationHighPriorityR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationHighPriority_r18_Sym0, FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationHighPriority_r18_Sym1, FeatureSetUplink_v1800_Ul_IntraUE_MuxEnh_r18_Pusch_PreparationHighPriority_r18_Sym2},
}

type FeatureSetUplink_v1800 struct {
	MaxDelayValueBeyondD_Basic_r18 *int64
	Tdcp_NumberDelayValue_r18      *int64
	PhaseReportMoreThanOne_r18     *int64
	MaxNumberTRS_ResourceSet_r18   *int64
	MaxNumberTDCP_PerBWP_r18       *int64
	Pusch_DMRS_TypeEnh_r18         *struct {
		Dmrs_Type_r18        int64
		Pusch_TypeA_DMRS_r18 *struct {
			Dmrs_TypeA_r18                           int64
			Pusch_2SymbolFL_DMRS_r18                 *int64
			Pusch_2SymbolFL_DMRS_Addition2Symbol_r18 *int64
			Pusch_1SymbolFL_DMRS_Addition3Symbol_r18 *int64
			Pusch_1SymbolFL_DMRS_BeyondOnePort_r18   *int64
		}
		Dummy                    *int64
		Pusch_TypeB_DMRS_r18     *int64
		Pusch_Rank_1_4_1Port_r18 *int64
		Pusch_Rank_5_8_1Port_r18 *int64
		Pusch_Rank_1_4_2Port_r18 *int64
		Pusch_Rank_5_8_2Port_r18 *int64
	}
	Ul_DMRS_SingleDCI_M_TRP_r18                *int64
	Ul_DMRS_M_DCI_M_TRP_r18                    *int64
	Srs_AntennaSwitching8T8R2SP_1Periodic_r18  *int64
	Pucch_SingleDCI_STx2P_SFN_r18              *int64
	PosSRS_BWA_RRC_Connected_r18               *PosSRS_BWA_RRC_Connected_r18
	PosSRS_BWA_IndependentCA_RRC_Connected_r18 *PosSRS_BWA_IndependentCA_RRC_Connected_r18
	PosSRS_BWA_AffectedBandList_r18            []FreqBandIndicatorNR
	Rach_EarlyTA_BandList_r18                  []bool
	Simultaneous_2_1_HARQ_ACK_CB_r18           *SubSlot_Config_r16
	Simultaneous_2_2_HARQ_ACK_CB_r18           *SubSlot_Config_r16
	Ul_IntraUE_MuxEnh_r18                      *struct {
		Pusch_PreparationLowPriority_r18  int64
		Pusch_PreparationHighPriority_r18 int64
	}
	TxDiversity4Tx_r18                      *int64
	PowerBoosting_Pi2BPSK_QPSK_r18          *int64
	PowerBoosting_Pi2BPSK_QPSK_Modified_r18 *int64
	TxDiversity2Tx_r18                      *int64
	Ue_PowerClassPerBandPerBC_v1820         *int64
}

func (ie *FeatureSetUplink_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxDelayValueBeyondD_Basic_r18 != nil, ie.Tdcp_NumberDelayValue_r18 != nil, ie.PhaseReportMoreThanOne_r18 != nil, ie.MaxNumberTRS_ResourceSet_r18 != nil, ie.MaxNumberTDCP_PerBWP_r18 != nil, ie.Pusch_DMRS_TypeEnh_r18 != nil, ie.Ul_DMRS_SingleDCI_M_TRP_r18 != nil, ie.Ul_DMRS_M_DCI_M_TRP_r18 != nil, ie.Srs_AntennaSwitching8T8R2SP_1Periodic_r18 != nil, ie.Pucch_SingleDCI_STx2P_SFN_r18 != nil, ie.PosSRS_BWA_RRC_Connected_r18 != nil, ie.PosSRS_BWA_IndependentCA_RRC_Connected_r18 != nil, ie.PosSRS_BWA_AffectedBandList_r18 != nil, ie.Rach_EarlyTA_BandList_r18 != nil, ie.Simultaneous_2_1_HARQ_ACK_CB_r18 != nil, ie.Simultaneous_2_2_HARQ_ACK_CB_r18 != nil, ie.Ul_IntraUE_MuxEnh_r18 != nil, ie.TxDiversity4Tx_r18 != nil, ie.PowerBoosting_Pi2BPSK_QPSK_r18 != nil, ie.PowerBoosting_Pi2BPSK_QPSK_Modified_r18 != nil, ie.TxDiversity2Tx_r18 != nil, ie.Ue_PowerClassPerBandPerBC_v1820 != nil}); err != nil {
		return err
	}

	// 2. maxDelayValueBeyondD-Basic-r18: enumerated
	{
		if ie.MaxDelayValueBeyondD_Basic_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaxDelayValueBeyondD_Basic_r18, featureSetUplinkV1800MaxDelayValueBeyondDBasicR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. tdcp-NumberDelayValue-r18: integer
	{
		if ie.Tdcp_NumberDelayValue_r18 != nil {
			if err := e.EncodeInteger(*ie.Tdcp_NumberDelayValue_r18, featureSetUplinkV1800TdcpNumberDelayValueR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. phaseReportMoreThanOne-r18: enumerated
	{
		if ie.PhaseReportMoreThanOne_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PhaseReportMoreThanOne_r18, featureSetUplinkV1800PhaseReportMoreThanOneR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maxNumberTRS-ResourceSet-r18: integer
	{
		if ie.MaxNumberTRS_ResourceSet_r18 != nil {
			if err := e.EncodeInteger(*ie.MaxNumberTRS_ResourceSet_r18, featureSetUplinkV1800MaxNumberTRSResourceSetR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. maxNumberTDCP-PerBWP-r18: integer
	{
		if ie.MaxNumberTDCP_PerBWP_r18 != nil {
			if err := e.EncodeInteger(*ie.MaxNumberTDCP_PerBWP_r18, featureSetUplinkV1800MaxNumberTDCPPerBWPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. pusch-DMRS-TypeEnh-r18: sequence
	{
		if ie.Pusch_DMRS_TypeEnh_r18 != nil {
			c := ie.Pusch_DMRS_TypeEnh_r18
			featureSetUplinkV1800PuschDMRSTypeEnhR18Seq := e.NewSequenceEncoder(featureSetUplinkV1800PuschDMRSTypeEnhR18Constraints)
			if err := featureSetUplinkV1800PuschDMRSTypeEnhR18Seq.EncodePreamble([]bool{c.Pusch_TypeA_DMRS_r18 != nil, c.Dummy != nil, c.Pusch_TypeB_DMRS_r18 != nil, c.Pusch_Rank_1_4_1Port_r18 != nil, c.Pusch_Rank_5_8_1Port_r18 != nil, c.Pusch_Rank_1_4_2Port_r18 != nil, c.Pusch_Rank_5_8_2Port_r18 != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Dmrs_Type_r18, featureSetUplinkV1800PuschDMRSTypeEnhR18DmrsTypeR18Constraints); err != nil {
				return err
			}
			if c.Pusch_TypeA_DMRS_r18 != nil {
				c := (*c.Pusch_TypeA_DMRS_r18)
				featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Seq := e.NewSequenceEncoder(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Constraints)
				if err := featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Seq.EncodePreamble([]bool{c.Pusch_2SymbolFL_DMRS_r18 != nil, c.Pusch_2SymbolFL_DMRS_Addition2Symbol_r18 != nil, c.Pusch_1SymbolFL_DMRS_Addition3Symbol_r18 != nil, c.Pusch_1SymbolFL_DMRS_BeyondOnePort_r18 != nil}); err != nil {
					return err
				}
				if err := e.EncodeEnumerated(c.Dmrs_TypeA_r18, featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18DmrsTypeAR18Constraints); err != nil {
					return err
				}
				if c.Pusch_2SymbolFL_DMRS_r18 != nil {
					if err := e.EncodeEnumerated((*c.Pusch_2SymbolFL_DMRS_r18), featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch2SymbolFLDMRSR18Constraints); err != nil {
						return err
					}
				}
				if c.Pusch_2SymbolFL_DMRS_Addition2Symbol_r18 != nil {
					if err := e.EncodeEnumerated((*c.Pusch_2SymbolFL_DMRS_Addition2Symbol_r18), featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch2SymbolFLDMRSAddition2SymbolR18Constraints); err != nil {
						return err
					}
				}
				if c.Pusch_1SymbolFL_DMRS_Addition3Symbol_r18 != nil {
					if err := e.EncodeEnumerated((*c.Pusch_1SymbolFL_DMRS_Addition3Symbol_r18), featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch1SymbolFLDMRSAddition3SymbolR18Constraints); err != nil {
						return err
					}
				}
				if c.Pusch_1SymbolFL_DMRS_BeyondOnePort_r18 != nil {
					if err := e.EncodeEnumerated((*c.Pusch_1SymbolFL_DMRS_BeyondOnePort_r18), featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch1SymbolFLDMRSBeyondOnePortR18Constraints); err != nil {
						return err
					}
				}
			}
			if c.Dummy != nil {
				if err := e.EncodeEnumerated((*c.Dummy), featureSetUplinkV1800PuschDMRSTypeEnhR18DummyConstraints); err != nil {
					return err
				}
			}
			if c.Pusch_TypeB_DMRS_r18 != nil {
				if err := e.EncodeEnumerated((*c.Pusch_TypeB_DMRS_r18), featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeBDMRSR18Constraints); err != nil {
					return err
				}
			}
			if c.Pusch_Rank_1_4_1Port_r18 != nil {
				if err := e.EncodeEnumerated((*c.Pusch_Rank_1_4_1Port_r18), featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank141PortR18Constraints); err != nil {
					return err
				}
			}
			if c.Pusch_Rank_5_8_1Port_r18 != nil {
				if err := e.EncodeEnumerated((*c.Pusch_Rank_5_8_1Port_r18), featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank581PortR18Constraints); err != nil {
					return err
				}
			}
			if c.Pusch_Rank_1_4_2Port_r18 != nil {
				if err := e.EncodeEnumerated((*c.Pusch_Rank_1_4_2Port_r18), featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank142PortR18Constraints); err != nil {
					return err
				}
			}
			if c.Pusch_Rank_5_8_2Port_r18 != nil {
				if err := e.EncodeEnumerated((*c.Pusch_Rank_5_8_2Port_r18), featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank582PortR18Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 8. ul-DMRS-SingleDCI-M-TRP-r18: enumerated
	{
		if ie.Ul_DMRS_SingleDCI_M_TRP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_DMRS_SingleDCI_M_TRP_r18, featureSetUplinkV1800UlDMRSSingleDCIMTRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. ul-DMRS-M-DCI-M-TRP-r18: enumerated
	{
		if ie.Ul_DMRS_M_DCI_M_TRP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_DMRS_M_DCI_M_TRP_r18, featureSetUplinkV1800UlDMRSMDCIMTRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. srs-AntennaSwitching8T8R2SP-1Periodic-r18: enumerated
	{
		if ie.Srs_AntennaSwitching8T8R2SP_1Periodic_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_AntennaSwitching8T8R2SP_1Periodic_r18, featureSetUplinkV1800SrsAntennaSwitching8T8R2SP1PeriodicR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. pucch-SingleDCI-STx2P-SFN-r18: enumerated
	{
		if ie.Pucch_SingleDCI_STx2P_SFN_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_SingleDCI_STx2P_SFN_r18, featureSetUplinkV1800PucchSingleDCISTx2PSFNR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. posSRS-BWA-RRC-Connected-r18: ref
	{
		if ie.PosSRS_BWA_RRC_Connected_r18 != nil {
			if err := ie.PosSRS_BWA_RRC_Connected_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 13. posSRS-BWA-IndependentCA-RRC-Connected-r18: ref
	{
		if ie.PosSRS_BWA_IndependentCA_RRC_Connected_r18 != nil {
			if err := ie.PosSRS_BWA_IndependentCA_RRC_Connected_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 14. posSRS-BWA-AffectedBandList-r18: sequence-of
	{
		if ie.PosSRS_BWA_AffectedBandList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetUplinkV1800PosSRSBWAAffectedBandListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PosSRS_BWA_AffectedBandList_r18))); err != nil {
				return err
			}
			for i := range ie.PosSRS_BWA_AffectedBandList_r18 {
				if err := ie.PosSRS_BWA_AffectedBandList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 15. rach-EarlyTA-BandList-r18: sequence-of
	{
		if ie.Rach_EarlyTA_BandList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(featureSetUplinkV1800RachEarlyTABandListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Rach_EarlyTA_BandList_r18))); err != nil {
				return err
			}
			for i := range ie.Rach_EarlyTA_BandList_r18 {
				if err := e.EncodeBoolean(ie.Rach_EarlyTA_BandList_r18[i]); err != nil {
					return err
				}
			}
		}
	}

	// 16. simultaneous-2-1-HARQ-ACK-CB-r18: ref
	{
		if ie.Simultaneous_2_1_HARQ_ACK_CB_r18 != nil {
			if err := ie.Simultaneous_2_1_HARQ_ACK_CB_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 17. simultaneous-2-2-HARQ-ACK-CB-r18: ref
	{
		if ie.Simultaneous_2_2_HARQ_ACK_CB_r18 != nil {
			if err := ie.Simultaneous_2_2_HARQ_ACK_CB_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 18. ul-IntraUE-MuxEnh-r18: sequence
	{
		if ie.Ul_IntraUE_MuxEnh_r18 != nil {
			c := ie.Ul_IntraUE_MuxEnh_r18
			if err := e.EncodeEnumerated(c.Pusch_PreparationLowPriority_r18, featureSetUplinkV1800UlIntraUEMuxEnhR18PuschPreparationLowPriorityR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Pusch_PreparationHighPriority_r18, featureSetUplinkV1800UlIntraUEMuxEnhR18PuschPreparationHighPriorityR18Constraints); err != nil {
				return err
			}
		}
	}

	// 19. txDiversity4Tx-r18: enumerated
	{
		if ie.TxDiversity4Tx_r18 != nil {
			if err := e.EncodeEnumerated(*ie.TxDiversity4Tx_r18, featureSetUplinkV1800TxDiversity4TxR18Constraints); err != nil {
				return err
			}
		}
	}

	// 20. powerBoosting-pi2BPSK-QPSK-r18: enumerated
	{
		if ie.PowerBoosting_Pi2BPSK_QPSK_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PowerBoosting_Pi2BPSK_QPSK_r18, featureSetUplinkV1800PowerBoostingPi2BPSKQPSKR18Constraints); err != nil {
				return err
			}
		}
	}

	// 21. powerBoosting-pi2BPSK-QPSK-Modified-r18: enumerated
	{
		if ie.PowerBoosting_Pi2BPSK_QPSK_Modified_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PowerBoosting_Pi2BPSK_QPSK_Modified_r18, featureSetUplinkV1800PowerBoostingPi2BPSKQPSKModifiedR18Constraints); err != nil {
				return err
			}
		}
	}

	// 22. txDiversity2Tx-r18: enumerated
	{
		if ie.TxDiversity2Tx_r18 != nil {
			if err := e.EncodeEnumerated(*ie.TxDiversity2Tx_r18, featureSetUplinkV1800TxDiversity2TxR18Constraints); err != nil {
				return err
			}
		}
	}

	// 23. ue-PowerClassPerBandPerBC-v1820: enumerated
	{
		if ie.Ue_PowerClassPerBandPerBC_v1820 != nil {
			if err := e.EncodeEnumerated(*ie.Ue_PowerClassPerBandPerBC_v1820, featureSetUplinkV1800UePowerClassPerBandPerBCV1820Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxDelayValueBeyondD-Basic-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800MaxDelayValueBeyondDBasicR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxDelayValueBeyondD_Basic_r18 = &idx
		}
	}

	// 3. tdcp-NumberDelayValue-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(featureSetUplinkV1800TdcpNumberDelayValueR18Constraints)
			if err != nil {
				return err
			}
			ie.Tdcp_NumberDelayValue_r18 = &v
		}
	}

	// 4. phaseReportMoreThanOne-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800PhaseReportMoreThanOneR18Constraints)
			if err != nil {
				return err
			}
			ie.PhaseReportMoreThanOne_r18 = &idx
		}
	}

	// 5. maxNumberTRS-ResourceSet-r18: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(featureSetUplinkV1800MaxNumberTRSResourceSetR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberTRS_ResourceSet_r18 = &v
		}
	}

	// 6. maxNumberTDCP-PerBWP-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(featureSetUplinkV1800MaxNumberTDCPPerBWPR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberTDCP_PerBWP_r18 = &v
		}
	}

	// 7. pusch-DMRS-TypeEnh-r18: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.Pusch_DMRS_TypeEnh_r18 = &struct {
				Dmrs_Type_r18        int64
				Pusch_TypeA_DMRS_r18 *struct {
					Dmrs_TypeA_r18                           int64
					Pusch_2SymbolFL_DMRS_r18                 *int64
					Pusch_2SymbolFL_DMRS_Addition2Symbol_r18 *int64
					Pusch_1SymbolFL_DMRS_Addition3Symbol_r18 *int64
					Pusch_1SymbolFL_DMRS_BeyondOnePort_r18   *int64
				}
				Dummy                    *int64
				Pusch_TypeB_DMRS_r18     *int64
				Pusch_Rank_1_4_1Port_r18 *int64
				Pusch_Rank_5_8_1Port_r18 *int64
				Pusch_Rank_1_4_2Port_r18 *int64
				Pusch_Rank_5_8_2Port_r18 *int64
			}{}
			c := ie.Pusch_DMRS_TypeEnh_r18
			featureSetUplinkV1800PuschDMRSTypeEnhR18Seq := d.NewSequenceDecoder(featureSetUplinkV1800PuschDMRSTypeEnhR18Constraints)
			if err := featureSetUplinkV1800PuschDMRSTypeEnhR18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18DmrsTypeR18Constraints)
				if err != nil {
					return err
				}
				c.Dmrs_Type_r18 = v
			}
			if featureSetUplinkV1800PuschDMRSTypeEnhR18Seq.IsComponentPresent(1) {
				c.Pusch_TypeA_DMRS_r18 = &struct {
					Dmrs_TypeA_r18                           int64
					Pusch_2SymbolFL_DMRS_r18                 *int64
					Pusch_2SymbolFL_DMRS_Addition2Symbol_r18 *int64
					Pusch_1SymbolFL_DMRS_Addition3Symbol_r18 *int64
					Pusch_1SymbolFL_DMRS_BeyondOnePort_r18   *int64
				}{}
				c := (*c.Pusch_TypeA_DMRS_r18)
				featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Seq := d.NewSequenceDecoder(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Constraints)
				if err := featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18DmrsTypeAR18Constraints)
					if err != nil {
						return err
					}
					c.Dmrs_TypeA_r18 = v
				}
				if featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Seq.IsComponentPresent(1) {
					c.Pusch_2SymbolFL_DMRS_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch2SymbolFLDMRSR18Constraints)
					if err != nil {
						return err
					}
					(*c.Pusch_2SymbolFL_DMRS_r18) = v
				}
				if featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Seq.IsComponentPresent(2) {
					c.Pusch_2SymbolFL_DMRS_Addition2Symbol_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch2SymbolFLDMRSAddition2SymbolR18Constraints)
					if err != nil {
						return err
					}
					(*c.Pusch_2SymbolFL_DMRS_Addition2Symbol_r18) = v
				}
				if featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Seq.IsComponentPresent(3) {
					c.Pusch_1SymbolFL_DMRS_Addition3Symbol_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch1SymbolFLDMRSAddition3SymbolR18Constraints)
					if err != nil {
						return err
					}
					(*c.Pusch_1SymbolFL_DMRS_Addition3Symbol_r18) = v
				}
				if featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Seq.IsComponentPresent(4) {
					c.Pusch_1SymbolFL_DMRS_BeyondOnePort_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeADMRSR18Pusch1SymbolFLDMRSBeyondOnePortR18Constraints)
					if err != nil {
						return err
					}
					(*c.Pusch_1SymbolFL_DMRS_BeyondOnePort_r18) = v
				}
			}
			if featureSetUplinkV1800PuschDMRSTypeEnhR18Seq.IsComponentPresent(2) {
				c.Dummy = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18DummyConstraints)
				if err != nil {
					return err
				}
				(*c.Dummy) = v
			}
			if featureSetUplinkV1800PuschDMRSTypeEnhR18Seq.IsComponentPresent(3) {
				c.Pusch_TypeB_DMRS_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschTypeBDMRSR18Constraints)
				if err != nil {
					return err
				}
				(*c.Pusch_TypeB_DMRS_r18) = v
			}
			if featureSetUplinkV1800PuschDMRSTypeEnhR18Seq.IsComponentPresent(4) {
				c.Pusch_Rank_1_4_1Port_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank141PortR18Constraints)
				if err != nil {
					return err
				}
				(*c.Pusch_Rank_1_4_1Port_r18) = v
			}
			if featureSetUplinkV1800PuschDMRSTypeEnhR18Seq.IsComponentPresent(5) {
				c.Pusch_Rank_5_8_1Port_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank581PortR18Constraints)
				if err != nil {
					return err
				}
				(*c.Pusch_Rank_5_8_1Port_r18) = v
			}
			if featureSetUplinkV1800PuschDMRSTypeEnhR18Seq.IsComponentPresent(6) {
				c.Pusch_Rank_1_4_2Port_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank142PortR18Constraints)
				if err != nil {
					return err
				}
				(*c.Pusch_Rank_1_4_2Port_r18) = v
			}
			if featureSetUplinkV1800PuschDMRSTypeEnhR18Seq.IsComponentPresent(7) {
				c.Pusch_Rank_5_8_2Port_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1800PuschDMRSTypeEnhR18PuschRank582PortR18Constraints)
				if err != nil {
					return err
				}
				(*c.Pusch_Rank_5_8_2Port_r18) = v
			}
		}
	}

	// 8. ul-DMRS-SingleDCI-M-TRP-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800UlDMRSSingleDCIMTRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Ul_DMRS_SingleDCI_M_TRP_r18 = &idx
		}
	}

	// 9. ul-DMRS-M-DCI-M-TRP-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800UlDMRSMDCIMTRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Ul_DMRS_M_DCI_M_TRP_r18 = &idx
		}
	}

	// 10. srs-AntennaSwitching8T8R2SP-1Periodic-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800SrsAntennaSwitching8T8R2SP1PeriodicR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_AntennaSwitching8T8R2SP_1Periodic_r18 = &idx
		}
	}

	// 11. pucch-SingleDCI-STx2P-SFN-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800PucchSingleDCISTx2PSFNR18Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_SingleDCI_STx2P_SFN_r18 = &idx
		}
	}

	// 12. posSRS-BWA-RRC-Connected-r18: ref
	{
		if seq.IsComponentPresent(10) {
			ie.PosSRS_BWA_RRC_Connected_r18 = new(PosSRS_BWA_RRC_Connected_r18)
			if err := ie.PosSRS_BWA_RRC_Connected_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 13. posSRS-BWA-IndependentCA-RRC-Connected-r18: ref
	{
		if seq.IsComponentPresent(11) {
			ie.PosSRS_BWA_IndependentCA_RRC_Connected_r18 = new(PosSRS_BWA_IndependentCA_RRC_Connected_r18)
			if err := ie.PosSRS_BWA_IndependentCA_RRC_Connected_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 14. posSRS-BWA-AffectedBandList-r18: sequence-of
	{
		if seq.IsComponentPresent(12) {
			seqOf := d.NewSequenceOfDecoder(featureSetUplinkV1800PosSRSBWAAffectedBandListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PosSRS_BWA_AffectedBandList_r18 = make([]FreqBandIndicatorNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PosSRS_BWA_AffectedBandList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 15. rach-EarlyTA-BandList-r18: sequence-of
	{
		if seq.IsComponentPresent(13) {
			seqOf := d.NewSequenceOfDecoder(featureSetUplinkV1800RachEarlyTABandListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Rach_EarlyTA_BandList_r18 = make([]bool, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				ie.Rach_EarlyTA_BandList_r18[i] = v
			}
		}
	}

	// 16. simultaneous-2-1-HARQ-ACK-CB-r18: ref
	{
		if seq.IsComponentPresent(14) {
			ie.Simultaneous_2_1_HARQ_ACK_CB_r18 = new(SubSlot_Config_r16)
			if err := ie.Simultaneous_2_1_HARQ_ACK_CB_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 17. simultaneous-2-2-HARQ-ACK-CB-r18: ref
	{
		if seq.IsComponentPresent(15) {
			ie.Simultaneous_2_2_HARQ_ACK_CB_r18 = new(SubSlot_Config_r16)
			if err := ie.Simultaneous_2_2_HARQ_ACK_CB_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 18. ul-IntraUE-MuxEnh-r18: sequence
	{
		if seq.IsComponentPresent(16) {
			ie.Ul_IntraUE_MuxEnh_r18 = &struct {
				Pusch_PreparationLowPriority_r18  int64
				Pusch_PreparationHighPriority_r18 int64
			}{}
			c := ie.Ul_IntraUE_MuxEnh_r18
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1800UlIntraUEMuxEnhR18PuschPreparationLowPriorityR18Constraints)
				if err != nil {
					return err
				}
				c.Pusch_PreparationLowPriority_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1800UlIntraUEMuxEnhR18PuschPreparationHighPriorityR18Constraints)
				if err != nil {
					return err
				}
				c.Pusch_PreparationHighPriority_r18 = v
			}
		}
	}

	// 19. txDiversity4Tx-r18: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800TxDiversity4TxR18Constraints)
			if err != nil {
				return err
			}
			ie.TxDiversity4Tx_r18 = &idx
		}
	}

	// 20. powerBoosting-pi2BPSK-QPSK-r18: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800PowerBoostingPi2BPSKQPSKR18Constraints)
			if err != nil {
				return err
			}
			ie.PowerBoosting_Pi2BPSK_QPSK_r18 = &idx
		}
	}

	// 21. powerBoosting-pi2BPSK-QPSK-Modified-r18: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800PowerBoostingPi2BPSKQPSKModifiedR18Constraints)
			if err != nil {
				return err
			}
			ie.PowerBoosting_Pi2BPSK_QPSK_Modified_r18 = &idx
		}
	}

	// 22. txDiversity2Tx-r18: enumerated
	{
		if seq.IsComponentPresent(20) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800TxDiversity2TxR18Constraints)
			if err != nil {
				return err
			}
			ie.TxDiversity2Tx_r18 = &idx
		}
	}

	// 23. ue-PowerClassPerBandPerBC-v1820: enumerated
	{
		if seq.IsComponentPresent(21) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1800UePowerClassPerBandPerBCV1820Constraints)
			if err != nil {
				return err
			}
			ie.Ue_PowerClassPerBandPerBC_v1820 = &idx
		}
	}

	return nil
}
