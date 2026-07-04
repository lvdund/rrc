// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v16k0 (line 19561).

var featureSetDownlinkV16k0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetSRS-CB-PUSCH-Ant-Switch-fr1-r16", Optional: true},
		{Name: "offsetSRS-CB-PUSCH-PDCCH-MonitorSingleOcc-fr1-r16", Optional: true},
		{Name: "offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithoutGap-fr1-r16", Optional: true},
		{Name: "offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithGap-fr1-r16", Optional: true},
		{Name: "offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithSpanGap-fr1-r16", Optional: true},
	},
}

const (
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16_Supported = 0
)

var featureSetDownlinkV16k0OffsetSRSCBPUSCHAntSwitchFr1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16_Supported},
}

const (
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16_Supported = 0
)

var featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorSingleOccFr1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16_Supported},
}

const (
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16_Supported = 0
)

var featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithoutGapFr1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16_Supported},
}

const (
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16_Supported = 0
)

var featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithGapFr1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16_Supported},
}

var featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
		{Name: "scs-60kHz-r16", Optional: true},
	},
}

const (
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set1 = 0
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set2 = 1
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set3 = 2
)

var featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs15kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set1, FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set2, FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set3},
}

const (
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set1 = 0
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set2 = 1
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set3 = 2
)

var featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set1, FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set2, FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set3},
}

const (
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set1 = 0
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set2 = 1
	FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set3 = 2
)

var featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set1, FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set2, FeatureSetDownlink_V16k0_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set3},
}

type FeatureSetDownlink_V16k0 struct {
	OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16                     *int64
	OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16         *int64
	OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16  *int64
	OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16     *int64
	OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16 *struct {
		Scs_15kHz_r16 *int64
		Scs_30kHz_r16 *int64
		Scs_60kHz_r16 *int64
	}
}

func (ie *FeatureSetDownlink_V16k0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV16k0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16 != nil}); err != nil {
		return err
	}

	// 2. offsetSRS-CB-PUSCH-Ant-Switch-fr1-r16: enumerated
	{
		if ie.OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16, featureSetDownlinkV16k0OffsetSRSCBPUSCHAntSwitchFr1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. offsetSRS-CB-PUSCH-PDCCH-MonitorSingleOcc-fr1-r16: enumerated
	{
		if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16, featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorSingleOccFr1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithoutGap-fr1-r16: enumerated
	{
		if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16, featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithoutGapFr1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithGap-fr1-r16: enumerated
	{
		if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16, featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithGapFr1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithSpanGap-fr1-r16: sequence
	{
		if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16 != nil {
			c := ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16
			featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq := e.NewSequenceEncoder(featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Constraints)
			if err := featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil, c.Scs_60kHz_r16 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_r16), featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs15kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_r16), featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_r16), featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs60kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_V16k0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV16k0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. offsetSRS-CB-PUSCH-Ant-Switch-fr1-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV16k0OffsetSRSCBPUSCHAntSwitchFr1R16Constraints)
			if err != nil {
				return err
			}
			ie.OffsetSRS_CB_PUSCH_Ant_Switch_Fr1_r16 = &idx
		}
	}

	// 3. offsetSRS-CB-PUSCH-PDCCH-MonitorSingleOcc-fr1-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorSingleOccFr1R16Constraints)
			if err != nil {
				return err
			}
			ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_Fr1_r16 = &idx
		}
	}

	// 4. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithoutGap-fr1-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithoutGapFr1R16Constraints)
			if err != nil {
				return err
			}
			ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_Fr1_r16 = &idx
		}
	}

	// 5. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithGap-fr1-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithGapFr1R16Constraints)
			if err != nil {
				return err
			}
			ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_Fr1_r16 = &idx
		}
	}

	// 6. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithSpanGap-fr1-r16: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16 = &struct {
				Scs_15kHz_r16 *int64
				Scs_30kHz_r16 *int64
				Scs_60kHz_r16 *int64
			}{}
			c := ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16
			featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq := d.NewSequenceDecoder(featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Constraints)
			if err := featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs15kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r16) = v
			}
			if featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r16) = v
			}
			if featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV16k0OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r16) = v
			}
		}
	}

	return nil
}
