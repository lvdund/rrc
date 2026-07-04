// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SDT-ConfigCommonSIB-r17 (line 2082).

var sDTConfigCommonSIBR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sdt-RSRP-Threshold-r17", Optional: true},
		{Name: "sdt-LogicalChannelSR-DelayTimer-r17", Optional: true},
		{Name: "sdt-DataVolumeThreshold-r17"},
		{Name: "t319a-r17"},
	},
}

const (
	SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf20   = 0
	SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf40   = 1
	SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf64   = 2
	SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf128  = 3
	SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf512  = 4
	SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf1024 = 5
	SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf2560 = 6
	SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Spare1 = 7
)

var sDTConfigCommonSIBR17SdtLogicalChannelSRDelayTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf20, SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf40, SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf64, SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf128, SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf512, SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf1024, SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Sf2560, SDT_ConfigCommonSIB_r17_Sdt_LogicalChannelSR_DelayTimer_r17_Spare1},
}

const (
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte32    = 0
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte100   = 1
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte200   = 2
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte400   = 3
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte600   = 4
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte800   = 5
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte1000  = 6
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte2000  = 7
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte4000  = 8
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte8000  = 9
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte9000  = 10
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte10000 = 11
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte12000 = 12
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte24000 = 13
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte48000 = 14
	SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte96000 = 15
)

var sDTConfigCommonSIBR17SdtDataVolumeThresholdR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte32, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte100, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte200, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte400, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte600, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte800, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte1000, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte2000, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte4000, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte8000, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte9000, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte10000, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte12000, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte24000, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte48000, SDT_ConfigCommonSIB_r17_Sdt_DataVolumeThreshold_r17_Byte96000},
}

const (
	SDT_ConfigCommonSIB_r17_T319a_r17_Ms100  = 0
	SDT_ConfigCommonSIB_r17_T319a_r17_Ms200  = 1
	SDT_ConfigCommonSIB_r17_T319a_r17_Ms300  = 2
	SDT_ConfigCommonSIB_r17_T319a_r17_Ms400  = 3
	SDT_ConfigCommonSIB_r17_T319a_r17_Ms600  = 4
	SDT_ConfigCommonSIB_r17_T319a_r17_Ms1000 = 5
	SDT_ConfigCommonSIB_r17_T319a_r17_Ms2000 = 6
	SDT_ConfigCommonSIB_r17_T319a_r17_Ms3000 = 7
	SDT_ConfigCommonSIB_r17_T319a_r17_Ms4000 = 8
	SDT_ConfigCommonSIB_r17_T319a_r17_Spare7 = 9
	SDT_ConfigCommonSIB_r17_T319a_r17_Spare6 = 10
	SDT_ConfigCommonSIB_r17_T319a_r17_Spare5 = 11
	SDT_ConfigCommonSIB_r17_T319a_r17_Spare4 = 12
	SDT_ConfigCommonSIB_r17_T319a_r17_Spare3 = 13
	SDT_ConfigCommonSIB_r17_T319a_r17_Spare2 = 14
	SDT_ConfigCommonSIB_r17_T319a_r17_Spare1 = 15
)

var sDTConfigCommonSIBR17T319aR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDT_ConfigCommonSIB_r17_T319a_r17_Ms100, SDT_ConfigCommonSIB_r17_T319a_r17_Ms200, SDT_ConfigCommonSIB_r17_T319a_r17_Ms300, SDT_ConfigCommonSIB_r17_T319a_r17_Ms400, SDT_ConfigCommonSIB_r17_T319a_r17_Ms600, SDT_ConfigCommonSIB_r17_T319a_r17_Ms1000, SDT_ConfigCommonSIB_r17_T319a_r17_Ms2000, SDT_ConfigCommonSIB_r17_T319a_r17_Ms3000, SDT_ConfigCommonSIB_r17_T319a_r17_Ms4000, SDT_ConfigCommonSIB_r17_T319a_r17_Spare7, SDT_ConfigCommonSIB_r17_T319a_r17_Spare6, SDT_ConfigCommonSIB_r17_T319a_r17_Spare5, SDT_ConfigCommonSIB_r17_T319a_r17_Spare4, SDT_ConfigCommonSIB_r17_T319a_r17_Spare3, SDT_ConfigCommonSIB_r17_T319a_r17_Spare2, SDT_ConfigCommonSIB_r17_T319a_r17_Spare1},
}

type SDT_ConfigCommonSIB_r17 struct {
	Sdt_RSRP_Threshold_r17              *RSRP_Range
	Sdt_LogicalChannelSR_DelayTimer_r17 *int64
	Sdt_DataVolumeThreshold_r17         int64
	T319a_r17                           int64
}

func (ie *SDT_ConfigCommonSIB_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDTConfigCommonSIBR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sdt_RSRP_Threshold_r17 != nil, ie.Sdt_LogicalChannelSR_DelayTimer_r17 != nil}); err != nil {
		return err
	}

	// 2. sdt-RSRP-Threshold-r17: ref
	{
		if ie.Sdt_RSRP_Threshold_r17 != nil {
			if err := ie.Sdt_RSRP_Threshold_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sdt-LogicalChannelSR-DelayTimer-r17: enumerated
	{
		if ie.Sdt_LogicalChannelSR_DelayTimer_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sdt_LogicalChannelSR_DelayTimer_r17, sDTConfigCommonSIBR17SdtLogicalChannelSRDelayTimerR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sdt-DataVolumeThreshold-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sdt_DataVolumeThreshold_r17, sDTConfigCommonSIBR17SdtDataVolumeThresholdR17Constraints); err != nil {
			return err
		}
	}

	// 5. t319a-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.T319a_r17, sDTConfigCommonSIBR17T319aR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SDT_ConfigCommonSIB_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDTConfigCommonSIBR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sdt-RSRP-Threshold-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sdt_RSRP_Threshold_r17 = new(RSRP_Range)
			if err := ie.Sdt_RSRP_Threshold_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sdt-LogicalChannelSR-DelayTimer-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sDTConfigCommonSIBR17SdtLogicalChannelSRDelayTimerR17Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_LogicalChannelSR_DelayTimer_r17 = &idx
		}
	}

	// 4. sdt-DataVolumeThreshold-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(sDTConfigCommonSIBR17SdtDataVolumeThresholdR17Constraints)
		if err != nil {
			return err
		}
		ie.Sdt_DataVolumeThreshold_r17 = v2
	}

	// 5. t319a-r17: enumerated
	{
		v3, err := d.DecodeEnumerated(sDTConfigCommonSIBR17T319aR17Constraints)
		if err != nil {
			return err
		}
		ie.T319a_r17 = v3
	}

	return nil
}
