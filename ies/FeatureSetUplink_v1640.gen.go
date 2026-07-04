// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplink-v1640 (line 20240).

var featureSetUplinkV1640Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "twoHARQ-ACK-Codebook-type1-r16", Optional: true},
		{Name: "twoHARQ-ACK-Codebook-type2-r16", Optional: true},
		{Name: "offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithSpanGap-fr1-r16", Optional: true},
	},
}

var featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
		{Name: "scs-60kHz-r16", Optional: true},
	},
}

const (
	FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set1 = 0
	FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set2 = 1
	FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set3 = 2
)

var featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs15kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set1, FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set2, FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_15kHz_r16_Set3},
}

const (
	FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set1 = 0
	FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set2 = 1
	FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set3 = 2
)

var featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set1, FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set2, FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_30kHz_r16_Set3},
}

const (
	FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set1 = 0
	FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set2 = 1
	FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set3 = 2
)

var featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set1, FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set2, FeatureSetUplink_v1640_OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16_Scs_60kHz_r16_Set3},
}

type FeatureSetUplink_v1640 struct {
	TwoHARQ_ACK_Codebook_Type1_r16                            *SubSlot_Config_r16
	TwoHARQ_ACK_Codebook_Type2_r16                            *SubSlot_Config_r16
	OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16 *struct {
		Scs_15kHz_r16 *int64
		Scs_30kHz_r16 *int64
		Scs_60kHz_r16 *int64
	}
}

func (ie *FeatureSetUplink_v1640) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV1640Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TwoHARQ_ACK_Codebook_Type1_r16 != nil, ie.TwoHARQ_ACK_Codebook_Type2_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16 != nil}); err != nil {
		return err
	}

	// 2. twoHARQ-ACK-Codebook-type1-r16: ref
	{
		if ie.TwoHARQ_ACK_Codebook_Type1_r16 != nil {
			if err := ie.TwoHARQ_ACK_Codebook_Type1_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. twoHARQ-ACK-Codebook-type2-r16: ref
	{
		if ie.TwoHARQ_ACK_Codebook_Type2_r16 != nil {
			if err := ie.TwoHARQ_ACK_Codebook_Type2_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithSpanGap-fr1-r16: sequence
	{
		if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16 != nil {
			c := ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16
			featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq := e.NewSequenceEncoder(featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Constraints)
			if err := featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil, c.Scs_60kHz_r16 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_r16), featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs15kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_r16), featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_r16), featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs60kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_v1640) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV1640Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. twoHARQ-ACK-Codebook-type1-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.TwoHARQ_ACK_Codebook_Type1_r16 = new(SubSlot_Config_r16)
			if err := ie.TwoHARQ_ACK_Codebook_Type1_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. twoHARQ-ACK-Codebook-type2-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.TwoHARQ_ACK_Codebook_Type2_r16 = new(SubSlot_Config_r16)
			if err := ie.TwoHARQ_ACK_Codebook_Type2_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. offsetSRS-CB-PUSCH-PDCCH-MonitorAnyOccWithSpanGap-fr1-r16: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16 = &struct {
				Scs_15kHz_r16 *int64
				Scs_30kHz_r16 *int64
				Scs_60kHz_r16 *int64
			}{}
			c := ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_Fr1_r16
			featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq := d.NewSequenceDecoder(featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Constraints)
			if err := featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs15kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r16) = v
			}
			if featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r16) = v
			}
			if featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1640OffsetSRSCBPUSCHPDCCHMonitorAnyOccWithSpanGapFr1R16Scs60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r16) = v
			}
		}
	}

	return nil
}
