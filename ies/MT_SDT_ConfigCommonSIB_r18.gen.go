// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MT-SDT-ConfigCommonSIB-r18 (line 2109).

var mTSDTConfigCommonSIBR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mt-SDT-RSRP-Threshold-r18", Optional: true},
		{Name: "sdt-LogicalChannelSR-DelayTimer-r18", Optional: true},
		{Name: "t319a-r18", Optional: true},
	},
}

const (
	MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf20   = 0
	MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf40   = 1
	MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf64   = 2
	MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf128  = 3
	MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf512  = 4
	MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf1024 = 5
	MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf2560 = 6
	MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Spare1 = 7
)

var mTSDTConfigCommonSIBR18SdtLogicalChannelSRDelayTimerR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf20, MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf40, MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf64, MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf128, MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf512, MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf1024, MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Sf2560, MT_SDT_ConfigCommonSIB_r18_Sdt_LogicalChannelSR_DelayTimer_r18_Spare1},
}

const (
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms100  = 0
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms200  = 1
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms300  = 2
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms400  = 3
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms600  = 4
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms1000 = 5
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms2000 = 6
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms3000 = 7
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms4000 = 8
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare7 = 9
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare6 = 10
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare5 = 11
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare4 = 12
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare3 = 13
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare2 = 14
	MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare1 = 15
)

var mTSDTConfigCommonSIBR18T319aR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms100, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms200, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms300, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms400, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms600, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms1000, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms2000, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms3000, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Ms4000, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare7, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare6, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare5, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare4, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare3, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare2, MT_SDT_ConfigCommonSIB_r18_T319a_r18_Spare1},
}

type MT_SDT_ConfigCommonSIB_r18 struct {
	Mt_SDT_RSRP_Threshold_r18           *RSRP_Range
	Sdt_LogicalChannelSR_DelayTimer_r18 *int64
	T319a_r18                           *int64
}

func (ie *MT_SDT_ConfigCommonSIB_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mTSDTConfigCommonSIBR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mt_SDT_RSRP_Threshold_r18 != nil, ie.Sdt_LogicalChannelSR_DelayTimer_r18 != nil, ie.T319a_r18 != nil}); err != nil {
		return err
	}

	// 2. mt-SDT-RSRP-Threshold-r18: ref
	{
		if ie.Mt_SDT_RSRP_Threshold_r18 != nil {
			if err := ie.Mt_SDT_RSRP_Threshold_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sdt-LogicalChannelSR-DelayTimer-r18: enumerated
	{
		if ie.Sdt_LogicalChannelSR_DelayTimer_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sdt_LogicalChannelSR_DelayTimer_r18, mTSDTConfigCommonSIBR18SdtLogicalChannelSRDelayTimerR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. t319a-r18: enumerated
	{
		if ie.T319a_r18 != nil {
			if err := e.EncodeEnumerated(*ie.T319a_r18, mTSDTConfigCommonSIBR18T319aR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MT_SDT_ConfigCommonSIB_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mTSDTConfigCommonSIBR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mt-SDT-RSRP-Threshold-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mt_SDT_RSRP_Threshold_r18 = new(RSRP_Range)
			if err := ie.Mt_SDT_RSRP_Threshold_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sdt-LogicalChannelSR-DelayTimer-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mTSDTConfigCommonSIBR18SdtLogicalChannelSRDelayTimerR18Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_LogicalChannelSR_DelayTimer_r18 = &idx
		}
	}

	// 4. t319a-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mTSDTConfigCommonSIBR18T319aR18Constraints)
			if err != nil {
				return err
			}
			ie.T319a_r18 = &idx
		}
	}

	return nil
}
