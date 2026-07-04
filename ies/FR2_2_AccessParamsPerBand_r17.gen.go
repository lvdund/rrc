// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FR2-2-AccessParamsPerBand-r17 (line 20771).

var fR22AccessParamsPerBandR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-FR2-2-SCS-120kHz-r17", Optional: true},
		{Name: "ul-FR2-2-SCS-120kHz-r17", Optional: true},
		{Name: "initialAccessSSB-120kHz-r17", Optional: true},
		{Name: "widebandPRACH-SCS-120kHz-r17", Optional: true},
		{Name: "multiRB-PUCCH-SCS-120kHz-r17", Optional: true},
		{Name: "multiPDSCH-SingleDCI-FR2-2-SCS-120kHz-r17", Optional: true},
		{Name: "multiPUSCH-SingleDCI-FR2-2-SCS-120kHz-r17", Optional: true},
		{Name: "dl-FR2-2-SCS-480kHz-r17", Optional: true},
		{Name: "ul-FR2-2-SCS-480kHz-r17", Optional: true},
		{Name: "initialAccessSSB-480kHz-r17", Optional: true},
		{Name: "widebandPRACH-SCS-480kHz-r17", Optional: true},
		{Name: "multiRB-PUCCH-SCS-480kHz-r17", Optional: true},
		{Name: "enhancedPDCCH-monitoringSCS-480kHz-r17", Optional: true},
		{Name: "dl-FR2-2-SCS-960kHz-r17", Optional: true},
		{Name: "ul-FR2-2-SCS-960kHz-r17", Optional: true},
		{Name: "multiRB-PUCCH-SCS-960kHz-r17", Optional: true},
		{Name: "enhancedPDCCH-monitoringSCS-960kHz-r17", Optional: true},
		{Name: "type1-ChannelAccess-FR2-2-r17", Optional: true},
		{Name: "type2-ChannelAccess-FR2-2-r17", Optional: true},
		{Name: "reduced-BeamSwitchTiming-FR2-2-r17", Optional: true},
		{Name: "support32-DL-HARQ-ProcessPerSCS-r17", Optional: true},
		{Name: "support32-UL-HARQ-ProcessPerSCS-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	FR2_2_AccessParamsPerBand_r17_Dl_FR2_2_SCS_120kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17DlFR22SCS120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Dl_FR2_2_SCS_120kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Ul_FR2_2_SCS_120kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17UlFR22SCS120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Ul_FR2_2_SCS_120kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_InitialAccessSSB_120kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17InitialAccessSSB120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_InitialAccessSSB_120kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_WidebandPRACH_SCS_120kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17WidebandPRACHSCS120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_WidebandPRACH_SCS_120kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_MultiRB_PUCCH_SCS_120kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17MultiRBPUCCHSCS120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_MultiRB_PUCCH_SCS_120kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17MultiPDSCHSingleDCIFR22SCS120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17MultiPUSCHSingleDCIFR22SCS120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Dl_FR2_2_SCS_480kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17DlFR22SCS480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Dl_FR2_2_SCS_480kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Ul_FR2_2_SCS_480kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17UlFR22SCS480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Ul_FR2_2_SCS_480kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_InitialAccessSSB_480kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17InitialAccessSSB480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_InitialAccessSSB_480kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_WidebandPRACH_SCS_480kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17WidebandPRACHSCS480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_WidebandPRACH_SCS_480kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_MultiRB_PUCCH_SCS_480kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17MultiRBPUCCHSCS480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_MultiRB_PUCCH_SCS_480kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_EnhancedPDCCH_MonitoringSCS_480kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_EnhancedPDCCH_MonitoringSCS_480kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Dl_FR2_2_SCS_960kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17DlFR22SCS960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Dl_FR2_2_SCS_960kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Ul_FR2_2_SCS_960kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17UlFR22SCS960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Ul_FR2_2_SCS_960kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_MultiRB_PUCCH_SCS_960kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17MultiRBPUCCHSCS960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_MultiRB_PUCCH_SCS_960kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Type1_ChannelAccess_FR2_2_r17_Supported = 0
)

var fR22AccessParamsPerBandR17Type1ChannelAccessFR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Type1_ChannelAccess_FR2_2_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Type2_ChannelAccess_FR2_2_r17_Supported = 0
)

var fR22AccessParamsPerBandR17Type2ChannelAccessFR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Type2_ChannelAccess_FR2_2_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Reduced_BeamSwitchTiming_FR2_2_r17_Supported = 0
)

var fR22AccessParamsPerBandR17ReducedBeamSwitchTimingFR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Reduced_BeamSwitchTiming_FR2_2_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Ext_Modulation64_QAM_PUSCH_FR2_2_r17_Supported = 0
)

var fR22AccessParamsPerBandR17ExtModulation64QAMPUSCHFR22R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Ext_Modulation64_QAM_PUSCH_FR2_2_r17_Supported},
}

var fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-monitoring4-1-r17", Optional: true},
		{Name: "pdcch-monitoring4-2-r17", Optional: true},
		{Name: "pdcch-monitoring8-4-r17", Optional: true},
	},
}

const (
	FR2_2_AccessParamsPerBand_r17_EnhancedPDCCH_MonitoringSCS_960kHz_r17_Pdcch_Monitoring4_1_r17_Supported = 0
)

var fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17PdcchMonitoring41R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_EnhancedPDCCH_MonitoringSCS_960kHz_r17_Pdcch_Monitoring4_1_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_EnhancedPDCCH_MonitoringSCS_960kHz_r17_Pdcch_Monitoring4_2_r17_Supported = 0
)

var fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17PdcchMonitoring42R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_EnhancedPDCCH_MonitoringSCS_960kHz_r17_Pdcch_Monitoring4_2_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_EnhancedPDCCH_MonitoringSCS_960kHz_r17_Pdcch_Monitoring8_4_r17_Supported = 0
)

var fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17PdcchMonitoring84R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_EnhancedPDCCH_MonitoringSCS_960kHz_r17_Pdcch_Monitoring8_4_r17_Supported},
}

var fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-120kHz-r17", Optional: true},
		{Name: "scs-480kHz-r17", Optional: true},
		{Name: "scs-960kHz-r17", Optional: true},
	},
}

const (
	FR2_2_AccessParamsPerBand_r17_Support32_DL_HARQ_ProcessPerSCS_r17_Scs_120kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Scs120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Support32_DL_HARQ_ProcessPerSCS_r17_Scs_120kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Support32_DL_HARQ_ProcessPerSCS_r17_Scs_480kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Scs480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Support32_DL_HARQ_ProcessPerSCS_r17_Scs_480kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Support32_DL_HARQ_ProcessPerSCS_r17_Scs_960kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Scs960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Support32_DL_HARQ_ProcessPerSCS_r17_Scs_960kHz_r17_Supported},
}

var fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-120kHz-r17", Optional: true},
		{Name: "scs-480kHz-r17", Optional: true},
		{Name: "scs-960kHz-r17", Optional: true},
	},
}

const (
	FR2_2_AccessParamsPerBand_r17_Support32_UL_HARQ_ProcessPerSCS_r17_Scs_120kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Scs120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Support32_UL_HARQ_ProcessPerSCS_r17_Scs_120kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Support32_UL_HARQ_ProcessPerSCS_r17_Scs_480kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Scs480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Support32_UL_HARQ_ProcessPerSCS_r17_Scs_480kHz_r17_Supported},
}

const (
	FR2_2_AccessParamsPerBand_r17_Support32_UL_HARQ_ProcessPerSCS_r17_Scs_960kHz_r17_Supported = 0
)

var fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Scs960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FR2_2_AccessParamsPerBand_r17_Support32_UL_HARQ_ProcessPerSCS_r17_Scs_960kHz_r17_Supported},
}

type FR2_2_AccessParamsPerBand_r17 struct {
	Dl_FR2_2_SCS_120kHz_r17                   *int64
	Ul_FR2_2_SCS_120kHz_r17                   *int64
	InitialAccessSSB_120kHz_r17               *int64
	WidebandPRACH_SCS_120kHz_r17              *int64
	MultiRB_PUCCH_SCS_120kHz_r17              *int64
	MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 *int64
	MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 *int64
	Dl_FR2_2_SCS_480kHz_r17                   *int64
	Ul_FR2_2_SCS_480kHz_r17                   *int64
	InitialAccessSSB_480kHz_r17               *int64
	WidebandPRACH_SCS_480kHz_r17              *int64
	MultiRB_PUCCH_SCS_480kHz_r17              *int64
	EnhancedPDCCH_MonitoringSCS_480kHz_r17    *int64
	Dl_FR2_2_SCS_960kHz_r17                   *int64
	Ul_FR2_2_SCS_960kHz_r17                   *int64
	MultiRB_PUCCH_SCS_960kHz_r17              *int64
	EnhancedPDCCH_MonitoringSCS_960kHz_r17    *struct {
		Pdcch_Monitoring4_1_r17 *int64
		Pdcch_Monitoring4_2_r17 *int64
		Pdcch_Monitoring8_4_r17 *int64
	}
	Type1_ChannelAccess_FR2_2_r17       *int64
	Type2_ChannelAccess_FR2_2_r17       *int64
	Reduced_BeamSwitchTiming_FR2_2_r17  *int64
	Support32_DL_HARQ_ProcessPerSCS_r17 *struct {
		Scs_120kHz_r17 *int64
		Scs_480kHz_r17 *int64
		Scs_960kHz_r17 *int64
	}
	Support32_UL_HARQ_ProcessPerSCS_r17 *struct {
		Scs_120kHz_r17 *int64
		Scs_480kHz_r17 *int64
		Scs_960kHz_r17 *int64
	}
	Modulation64_QAM_PUSCH_FR2_2_r17 *int64
}

func (ie *FR2_2_AccessParamsPerBand_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(fR22AccessParamsPerBandR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Modulation64_QAM_PUSCH_FR2_2_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_FR2_2_SCS_120kHz_r17 != nil, ie.Ul_FR2_2_SCS_120kHz_r17 != nil, ie.InitialAccessSSB_120kHz_r17 != nil, ie.WidebandPRACH_SCS_120kHz_r17 != nil, ie.MultiRB_PUCCH_SCS_120kHz_r17 != nil, ie.MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil, ie.MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil, ie.Dl_FR2_2_SCS_480kHz_r17 != nil, ie.Ul_FR2_2_SCS_480kHz_r17 != nil, ie.InitialAccessSSB_480kHz_r17 != nil, ie.WidebandPRACH_SCS_480kHz_r17 != nil, ie.MultiRB_PUCCH_SCS_480kHz_r17 != nil, ie.EnhancedPDCCH_MonitoringSCS_480kHz_r17 != nil, ie.Dl_FR2_2_SCS_960kHz_r17 != nil, ie.Ul_FR2_2_SCS_960kHz_r17 != nil, ie.MultiRB_PUCCH_SCS_960kHz_r17 != nil, ie.EnhancedPDCCH_MonitoringSCS_960kHz_r17 != nil, ie.Type1_ChannelAccess_FR2_2_r17 != nil, ie.Type2_ChannelAccess_FR2_2_r17 != nil, ie.Reduced_BeamSwitchTiming_FR2_2_r17 != nil, ie.Support32_DL_HARQ_ProcessPerSCS_r17 != nil, ie.Support32_UL_HARQ_ProcessPerSCS_r17 != nil}); err != nil {
		return err
	}

	// 3. dl-FR2-2-SCS-120kHz-r17: enumerated
	{
		if ie.Dl_FR2_2_SCS_120kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_FR2_2_SCS_120kHz_r17, fR22AccessParamsPerBandR17DlFR22SCS120kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ul-FR2-2-SCS-120kHz-r17: enumerated
	{
		if ie.Ul_FR2_2_SCS_120kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_FR2_2_SCS_120kHz_r17, fR22AccessParamsPerBandR17UlFR22SCS120kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. initialAccessSSB-120kHz-r17: enumerated
	{
		if ie.InitialAccessSSB_120kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.InitialAccessSSB_120kHz_r17, fR22AccessParamsPerBandR17InitialAccessSSB120kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. widebandPRACH-SCS-120kHz-r17: enumerated
	{
		if ie.WidebandPRACH_SCS_120kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.WidebandPRACH_SCS_120kHz_r17, fR22AccessParamsPerBandR17WidebandPRACHSCS120kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. multiRB-PUCCH-SCS-120kHz-r17: enumerated
	{
		if ie.MultiRB_PUCCH_SCS_120kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MultiRB_PUCCH_SCS_120kHz_r17, fR22AccessParamsPerBandR17MultiRBPUCCHSCS120kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. multiPDSCH-SingleDCI-FR2-2-SCS-120kHz-r17: enumerated
	{
		if ie.MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17, fR22AccessParamsPerBandR17MultiPDSCHSingleDCIFR22SCS120kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. multiPUSCH-SingleDCI-FR2-2-SCS-120kHz-r17: enumerated
	{
		if ie.MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17, fR22AccessParamsPerBandR17MultiPUSCHSingleDCIFR22SCS120kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. dl-FR2-2-SCS-480kHz-r17: enumerated
	{
		if ie.Dl_FR2_2_SCS_480kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_FR2_2_SCS_480kHz_r17, fR22AccessParamsPerBandR17DlFR22SCS480kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. ul-FR2-2-SCS-480kHz-r17: enumerated
	{
		if ie.Ul_FR2_2_SCS_480kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_FR2_2_SCS_480kHz_r17, fR22AccessParamsPerBandR17UlFR22SCS480kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. initialAccessSSB-480kHz-r17: enumerated
	{
		if ie.InitialAccessSSB_480kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.InitialAccessSSB_480kHz_r17, fR22AccessParamsPerBandR17InitialAccessSSB480kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. widebandPRACH-SCS-480kHz-r17: enumerated
	{
		if ie.WidebandPRACH_SCS_480kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.WidebandPRACH_SCS_480kHz_r17, fR22AccessParamsPerBandR17WidebandPRACHSCS480kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 14. multiRB-PUCCH-SCS-480kHz-r17: enumerated
	{
		if ie.MultiRB_PUCCH_SCS_480kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MultiRB_PUCCH_SCS_480kHz_r17, fR22AccessParamsPerBandR17MultiRBPUCCHSCS480kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 15. enhancedPDCCH-monitoringSCS-480kHz-r17: enumerated
	{
		if ie.EnhancedPDCCH_MonitoringSCS_480kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.EnhancedPDCCH_MonitoringSCS_480kHz_r17, fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS480kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 16. dl-FR2-2-SCS-960kHz-r17: enumerated
	{
		if ie.Dl_FR2_2_SCS_960kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_FR2_2_SCS_960kHz_r17, fR22AccessParamsPerBandR17DlFR22SCS960kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 17. ul-FR2-2-SCS-960kHz-r17: enumerated
	{
		if ie.Ul_FR2_2_SCS_960kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_FR2_2_SCS_960kHz_r17, fR22AccessParamsPerBandR17UlFR22SCS960kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 18. multiRB-PUCCH-SCS-960kHz-r17: enumerated
	{
		if ie.MultiRB_PUCCH_SCS_960kHz_r17 != nil {
			if err := e.EncodeEnumerated(*ie.MultiRB_PUCCH_SCS_960kHz_r17, fR22AccessParamsPerBandR17MultiRBPUCCHSCS960kHzR17Constraints); err != nil {
				return err
			}
		}
	}

	// 19. enhancedPDCCH-monitoringSCS-960kHz-r17: sequence
	{
		if ie.EnhancedPDCCH_MonitoringSCS_960kHz_r17 != nil {
			c := ie.EnhancedPDCCH_MonitoringSCS_960kHz_r17
			fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Seq := e.NewSequenceEncoder(fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Constraints)
			if err := fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Seq.EncodePreamble([]bool{c.Pdcch_Monitoring4_1_r17 != nil, c.Pdcch_Monitoring4_2_r17 != nil, c.Pdcch_Monitoring8_4_r17 != nil}); err != nil {
				return err
			}
			if c.Pdcch_Monitoring4_1_r17 != nil {
				if err := e.EncodeEnumerated((*c.Pdcch_Monitoring4_1_r17), fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17PdcchMonitoring41R17Constraints); err != nil {
					return err
				}
			}
			if c.Pdcch_Monitoring4_2_r17 != nil {
				if err := e.EncodeEnumerated((*c.Pdcch_Monitoring4_2_r17), fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17PdcchMonitoring42R17Constraints); err != nil {
					return err
				}
			}
			if c.Pdcch_Monitoring8_4_r17 != nil {
				if err := e.EncodeEnumerated((*c.Pdcch_Monitoring8_4_r17), fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17PdcchMonitoring84R17Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 20. type1-ChannelAccess-FR2-2-r17: enumerated
	{
		if ie.Type1_ChannelAccess_FR2_2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Type1_ChannelAccess_FR2_2_r17, fR22AccessParamsPerBandR17Type1ChannelAccessFR22R17Constraints); err != nil {
				return err
			}
		}
	}

	// 21. type2-ChannelAccess-FR2-2-r17: enumerated
	{
		if ie.Type2_ChannelAccess_FR2_2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Type2_ChannelAccess_FR2_2_r17, fR22AccessParamsPerBandR17Type2ChannelAccessFR22R17Constraints); err != nil {
				return err
			}
		}
	}

	// 22. reduced-BeamSwitchTiming-FR2-2-r17: enumerated
	{
		if ie.Reduced_BeamSwitchTiming_FR2_2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Reduced_BeamSwitchTiming_FR2_2_r17, fR22AccessParamsPerBandR17ReducedBeamSwitchTimingFR22R17Constraints); err != nil {
				return err
			}
		}
	}

	// 23. support32-DL-HARQ-ProcessPerSCS-r17: sequence
	{
		if ie.Support32_DL_HARQ_ProcessPerSCS_r17 != nil {
			c := ie.Support32_DL_HARQ_ProcessPerSCS_r17
			fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Seq := e.NewSequenceEncoder(fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Constraints)
			if err := fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Seq.EncodePreamble([]bool{c.Scs_120kHz_r17 != nil, c.Scs_480kHz_r17 != nil, c.Scs_960kHz_r17 != nil}); err != nil {
				return err
			}
			if c.Scs_120kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz_r17), fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Scs120kHzR17Constraints); err != nil {
					return err
				}
			}
			if c.Scs_480kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.Scs_480kHz_r17), fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Scs480kHzR17Constraints); err != nil {
					return err
				}
			}
			if c.Scs_960kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.Scs_960kHz_r17), fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Scs960kHzR17Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 24. support32-UL-HARQ-ProcessPerSCS-r17: sequence
	{
		if ie.Support32_UL_HARQ_ProcessPerSCS_r17 != nil {
			c := ie.Support32_UL_HARQ_ProcessPerSCS_r17
			fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Seq := e.NewSequenceEncoder(fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Constraints)
			if err := fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Seq.EncodePreamble([]bool{c.Scs_120kHz_r17 != nil, c.Scs_480kHz_r17 != nil, c.Scs_960kHz_r17 != nil}); err != nil {
				return err
			}
			if c.Scs_120kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz_r17), fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Scs120kHzR17Constraints); err != nil {
					return err
				}
			}
			if c.Scs_480kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.Scs_480kHz_r17), fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Scs480kHzR17Constraints); err != nil {
					return err
				}
			}
			if c.Scs_960kHz_r17 != nil {
				if err := e.EncodeEnumerated((*c.Scs_960kHz_r17), fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Scs960kHzR17Constraints); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "modulation64-QAM-PUSCH-FR2-2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Modulation64_QAM_PUSCH_FR2_2_r17 != nil}); err != nil {
				return err
			}

			if ie.Modulation64_QAM_PUSCH_FR2_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Modulation64_QAM_PUSCH_FR2_2_r17, fR22AccessParamsPerBandR17ExtModulation64QAMPUSCHFR22R17Constraints); err != nil {
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

func (ie *FR2_2_AccessParamsPerBand_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(fR22AccessParamsPerBandR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dl-FR2-2-SCS-120kHz-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17DlFR22SCS120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.Dl_FR2_2_SCS_120kHz_r17 = &idx
		}
	}

	// 4. ul-FR2-2-SCS-120kHz-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17UlFR22SCS120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FR2_2_SCS_120kHz_r17 = &idx
		}
	}

	// 5. initialAccessSSB-120kHz-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17InitialAccessSSB120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.InitialAccessSSB_120kHz_r17 = &idx
		}
	}

	// 6. widebandPRACH-SCS-120kHz-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17WidebandPRACHSCS120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.WidebandPRACH_SCS_120kHz_r17 = &idx
		}
	}

	// 7. multiRB-PUCCH-SCS-120kHz-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17MultiRBPUCCHSCS120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiRB_PUCCH_SCS_120kHz_r17 = &idx
		}
	}

	// 8. multiPDSCH-SingleDCI-FR2-2-SCS-120kHz-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17MultiPDSCHSingleDCIFR22SCS120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 = &idx
		}
	}

	// 9. multiPUSCH-SingleDCI-FR2-2-SCS-120kHz-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17MultiPUSCHSingleDCIFR22SCS120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 = &idx
		}
	}

	// 10. dl-FR2-2-SCS-480kHz-r17: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17DlFR22SCS480kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.Dl_FR2_2_SCS_480kHz_r17 = &idx
		}
	}

	// 11. ul-FR2-2-SCS-480kHz-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17UlFR22SCS480kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FR2_2_SCS_480kHz_r17 = &idx
		}
	}

	// 12. initialAccessSSB-480kHz-r17: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17InitialAccessSSB480kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.InitialAccessSSB_480kHz_r17 = &idx
		}
	}

	// 13. widebandPRACH-SCS-480kHz-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17WidebandPRACHSCS480kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.WidebandPRACH_SCS_480kHz_r17 = &idx
		}
	}

	// 14. multiRB-PUCCH-SCS-480kHz-r17: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17MultiRBPUCCHSCS480kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiRB_PUCCH_SCS_480kHz_r17 = &idx
		}
	}

	// 15. enhancedPDCCH-monitoringSCS-480kHz-r17: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS480kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedPDCCH_MonitoringSCS_480kHz_r17 = &idx
		}
	}

	// 16. dl-FR2-2-SCS-960kHz-r17: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17DlFR22SCS960kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.Dl_FR2_2_SCS_960kHz_r17 = &idx
		}
	}

	// 17. ul-FR2-2-SCS-960kHz-r17: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17UlFR22SCS960kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FR2_2_SCS_960kHz_r17 = &idx
		}
	}

	// 18. multiRB-PUCCH-SCS-960kHz-r17: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17MultiRBPUCCHSCS960kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiRB_PUCCH_SCS_960kHz_r17 = &idx
		}
	}

	// 19. enhancedPDCCH-monitoringSCS-960kHz-r17: sequence
	{
		if seq.IsComponentPresent(16) {
			ie.EnhancedPDCCH_MonitoringSCS_960kHz_r17 = &struct {
				Pdcch_Monitoring4_1_r17 *int64
				Pdcch_Monitoring4_2_r17 *int64
				Pdcch_Monitoring8_4_r17 *int64
			}{}
			c := ie.EnhancedPDCCH_MonitoringSCS_960kHz_r17
			fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Seq := d.NewSequenceDecoder(fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Constraints)
			if err := fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Seq.IsComponentPresent(0) {
				c.Pdcch_Monitoring4_1_r17 = new(int64)
				v, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17PdcchMonitoring41R17Constraints)
				if err != nil {
					return err
				}
				(*c.Pdcch_Monitoring4_1_r17) = v
			}
			if fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Seq.IsComponentPresent(1) {
				c.Pdcch_Monitoring4_2_r17 = new(int64)
				v, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17PdcchMonitoring42R17Constraints)
				if err != nil {
					return err
				}
				(*c.Pdcch_Monitoring4_2_r17) = v
			}
			if fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17Seq.IsComponentPresent(2) {
				c.Pdcch_Monitoring8_4_r17 = new(int64)
				v, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17EnhancedPDCCHMonitoringSCS960kHzR17PdcchMonitoring84R17Constraints)
				if err != nil {
					return err
				}
				(*c.Pdcch_Monitoring8_4_r17) = v
			}
		}
	}

	// 20. type1-ChannelAccess-FR2-2-r17: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17Type1ChannelAccessFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.Type1_ChannelAccess_FR2_2_r17 = &idx
		}
	}

	// 21. type2-ChannelAccess-FR2-2-r17: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17Type2ChannelAccessFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.Type2_ChannelAccess_FR2_2_r17 = &idx
		}
	}

	// 22. reduced-BeamSwitchTiming-FR2-2-r17: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17ReducedBeamSwitchTimingFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.Reduced_BeamSwitchTiming_FR2_2_r17 = &idx
		}
	}

	// 23. support32-DL-HARQ-ProcessPerSCS-r17: sequence
	{
		if seq.IsComponentPresent(20) {
			ie.Support32_DL_HARQ_ProcessPerSCS_r17 = &struct {
				Scs_120kHz_r17 *int64
				Scs_480kHz_r17 *int64
				Scs_960kHz_r17 *int64
			}{}
			c := ie.Support32_DL_HARQ_ProcessPerSCS_r17
			fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Seq := d.NewSequenceDecoder(fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Constraints)
			if err := fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Seq.IsComponentPresent(0) {
				c.Scs_120kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Scs120kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r17) = v
			}
			if fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Seq.IsComponentPresent(1) {
				c.Scs_480kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Scs480kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_480kHz_r17) = v
			}
			if fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Seq.IsComponentPresent(2) {
				c.Scs_960kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17Support32DLHARQProcessPerSCSR17Scs960kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_960kHz_r17) = v
			}
		}
	}

	// 24. support32-UL-HARQ-ProcessPerSCS-r17: sequence
	{
		if seq.IsComponentPresent(21) {
			ie.Support32_UL_HARQ_ProcessPerSCS_r17 = &struct {
				Scs_120kHz_r17 *int64
				Scs_480kHz_r17 *int64
				Scs_960kHz_r17 *int64
			}{}
			c := ie.Support32_UL_HARQ_ProcessPerSCS_r17
			fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Seq := d.NewSequenceDecoder(fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Constraints)
			if err := fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Seq.IsComponentPresent(0) {
				c.Scs_120kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Scs120kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r17) = v
			}
			if fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Seq.IsComponentPresent(1) {
				c.Scs_480kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Scs480kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_480kHz_r17) = v
			}
			if fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Seq.IsComponentPresent(2) {
				c.Scs_960kHz_r17 = new(int64)
				v, err := d.DecodeEnumerated(fR22AccessParamsPerBandR17Support32ULHARQProcessPerSCSR17Scs960kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_960kHz_r17) = v
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
				{Name: "modulation64-QAM-PUSCH-FR2-2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(fR22AccessParamsPerBandR17ExtModulation64QAMPUSCHFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.Modulation64_QAM_PUSCH_FR2_2_r17 = &v
		}
	}

	return nil
}
