// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v1540 (line 19478).

var featureSetDownlinkV1540Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "oneFL-DMRS-TwoAdditionalDMRS-DL", Optional: true},
		{Name: "additionalDMRS-DL-Alt", Optional: true},
		{Name: "twoFL-DMRS-TwoAdditionalDMRS-DL", Optional: true},
		{Name: "oneFL-DMRS-ThreeAdditionalDMRS-DL", Optional: true},
		{Name: "pdcch-MonitoringAnyOccasionsWithSpanGap", Optional: true},
		{Name: "pdsch-SeparationWithGap", Optional: true},
		{Name: "pdsch-ProcessingType2", Optional: true},
		{Name: "pdsch-ProcessingType2-Limited", Optional: true},
		{Name: "dl-MCS-TableAlt-DynamicIndication", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1540_OneFL_DMRS_TwoAdditionalDMRS_DL_Supported = 0
)

var featureSetDownlinkV1540OneFLDMRSTwoAdditionalDMRSDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_OneFL_DMRS_TwoAdditionalDMRS_DL_Supported},
}

const (
	FeatureSetDownlink_v1540_AdditionalDMRS_DL_Alt_Supported = 0
)

var featureSetDownlinkV1540AdditionalDMRSDLAltConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_AdditionalDMRS_DL_Alt_Supported},
}

const (
	FeatureSetDownlink_v1540_TwoFL_DMRS_TwoAdditionalDMRS_DL_Supported = 0
)

var featureSetDownlinkV1540TwoFLDMRSTwoAdditionalDMRSDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_TwoFL_DMRS_TwoAdditionalDMRS_DL_Supported},
}

const (
	FeatureSetDownlink_v1540_OneFL_DMRS_ThreeAdditionalDMRS_DL_Supported = 0
)

var featureSetDownlinkV1540OneFLDMRSThreeAdditionalDMRSDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_OneFL_DMRS_ThreeAdditionalDMRS_DL_Supported},
}

const (
	FeatureSetDownlink_v1540_Pdsch_SeparationWithGap_Supported = 0
)

var featureSetDownlinkV1540PdschSeparationWithGapConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_Pdsch_SeparationWithGap_Supported},
}

const (
	FeatureSetDownlink_v1540_Dl_MCS_TableAlt_DynamicIndication_Supported = 0
)

var featureSetDownlinkV1540DlMCSTableAltDynamicIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_Dl_MCS_TableAlt_DynamicIndication_Supported},
}

var featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_15kHz_Set1 = 0
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_15kHz_Set2 = 1
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_15kHz_Set3 = 2
)

var featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_15kHz_Set1, FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_15kHz_Set2, FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_15kHz_Set3},
}

const (
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_30kHz_Set1 = 0
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_30kHz_Set2 = 1
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_30kHz_Set3 = 2
)

var featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_30kHz_Set1, FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_30kHz_Set2, FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_30kHz_Set3},
}

const (
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_60kHz_Set1 = 0
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_60kHz_Set2 = 1
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_60kHz_Set3 = 2
)

var featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_60kHz_Set1, FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_60kHz_Set2, FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_60kHz_Set3},
}

const (
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_120kHz_Set1 = 0
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_120kHz_Set2 = 1
	FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_120kHz_Set3 = 2
)

var featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_120kHz_Set1, FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_120kHz_Set2, FeatureSetDownlink_v1540_Pdcch_MonitoringAnyOccasionsWithSpanGap_Scs_120kHz_Set3},
}

var featureSetDownlinkV1540PdschProcessingType2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1540_Pdsch_ProcessingType2_Limited_DifferentTB_PerSlot_SCS_30kHz_Upto1 = 0
	FeatureSetDownlink_v1540_Pdsch_ProcessingType2_Limited_DifferentTB_PerSlot_SCS_30kHz_Upto2 = 1
	FeatureSetDownlink_v1540_Pdsch_ProcessingType2_Limited_DifferentTB_PerSlot_SCS_30kHz_Upto4 = 2
	FeatureSetDownlink_v1540_Pdsch_ProcessingType2_Limited_DifferentTB_PerSlot_SCS_30kHz_Upto7 = 3
)

var featureSetDownlinkV1540PdschProcessingType2LimitedDifferentTBPerSlotSCS30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1540_Pdsch_ProcessingType2_Limited_DifferentTB_PerSlot_SCS_30kHz_Upto1, FeatureSetDownlink_v1540_Pdsch_ProcessingType2_Limited_DifferentTB_PerSlot_SCS_30kHz_Upto2, FeatureSetDownlink_v1540_Pdsch_ProcessingType2_Limited_DifferentTB_PerSlot_SCS_30kHz_Upto4, FeatureSetDownlink_v1540_Pdsch_ProcessingType2_Limited_DifferentTB_PerSlot_SCS_30kHz_Upto7},
}

type FeatureSetDownlink_v1540 struct {
	OneFL_DMRS_TwoAdditionalDMRS_DL         *int64
	AdditionalDMRS_DL_Alt                   *int64
	TwoFL_DMRS_TwoAdditionalDMRS_DL         *int64
	OneFL_DMRS_ThreeAdditionalDMRS_DL       *int64
	Pdcch_MonitoringAnyOccasionsWithSpanGap *struct {
		Scs_15kHz  *int64
		Scs_30kHz  *int64
		Scs_60kHz  *int64
		Scs_120kHz *int64
	}
	Pdsch_SeparationWithGap *int64
	Pdsch_ProcessingType2   *struct {
		Scs_15kHz *ProcessingParameters
		Scs_30kHz *ProcessingParameters
		Scs_60kHz *ProcessingParameters
	}
	Pdsch_ProcessingType2_Limited     *struct{ DifferentTB_PerSlot_SCS_30kHz int64 }
	Dl_MCS_TableAlt_DynamicIndication *int64
}

func (ie *FeatureSetDownlink_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV1540Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OneFL_DMRS_TwoAdditionalDMRS_DL != nil, ie.AdditionalDMRS_DL_Alt != nil, ie.TwoFL_DMRS_TwoAdditionalDMRS_DL != nil, ie.OneFL_DMRS_ThreeAdditionalDMRS_DL != nil, ie.Pdcch_MonitoringAnyOccasionsWithSpanGap != nil, ie.Pdsch_SeparationWithGap != nil, ie.Pdsch_ProcessingType2 != nil, ie.Pdsch_ProcessingType2_Limited != nil, ie.Dl_MCS_TableAlt_DynamicIndication != nil}); err != nil {
		return err
	}

	// 2. oneFL-DMRS-TwoAdditionalDMRS-DL: enumerated
	{
		if ie.OneFL_DMRS_TwoAdditionalDMRS_DL != nil {
			if err := e.EncodeEnumerated(*ie.OneFL_DMRS_TwoAdditionalDMRS_DL, featureSetDownlinkV1540OneFLDMRSTwoAdditionalDMRSDLConstraints); err != nil {
				return err
			}
		}
	}

	// 3. additionalDMRS-DL-Alt: enumerated
	{
		if ie.AdditionalDMRS_DL_Alt != nil {
			if err := e.EncodeEnumerated(*ie.AdditionalDMRS_DL_Alt, featureSetDownlinkV1540AdditionalDMRSDLAltConstraints); err != nil {
				return err
			}
		}
	}

	// 4. twoFL-DMRS-TwoAdditionalDMRS-DL: enumerated
	{
		if ie.TwoFL_DMRS_TwoAdditionalDMRS_DL != nil {
			if err := e.EncodeEnumerated(*ie.TwoFL_DMRS_TwoAdditionalDMRS_DL, featureSetDownlinkV1540TwoFLDMRSTwoAdditionalDMRSDLConstraints); err != nil {
				return err
			}
		}
	}

	// 5. oneFL-DMRS-ThreeAdditionalDMRS-DL: enumerated
	{
		if ie.OneFL_DMRS_ThreeAdditionalDMRS_DL != nil {
			if err := e.EncodeEnumerated(*ie.OneFL_DMRS_ThreeAdditionalDMRS_DL, featureSetDownlinkV1540OneFLDMRSThreeAdditionalDMRSDLConstraints); err != nil {
				return err
			}
		}
	}

	// 6. pdcch-MonitoringAnyOccasionsWithSpanGap: sequence
	{
		if ie.Pdcch_MonitoringAnyOccasionsWithSpanGap != nil {
			c := ie.Pdcch_MonitoringAnyOccasionsWithSpanGap
			featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapSeq := e.NewSequenceEncoder(featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapConstraints)
			if err := featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz), featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs15kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz), featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs30kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz), featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs60kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz), featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs120kHzConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 7. pdsch-SeparationWithGap: enumerated
	{
		if ie.Pdsch_SeparationWithGap != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_SeparationWithGap, featureSetDownlinkV1540PdschSeparationWithGapConstraints); err != nil {
				return err
			}
		}
	}

	// 8. pdsch-ProcessingType2: sequence
	{
		if ie.Pdsch_ProcessingType2 != nil {
			c := ie.Pdsch_ProcessingType2
			featureSetDownlinkV1540PdschProcessingType2Seq := e.NewSequenceEncoder(featureSetDownlinkV1540PdschProcessingType2Constraints)
			if err := featureSetDownlinkV1540PdschProcessingType2Seq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz != nil {
				if err := c.Scs_15kHz.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_30kHz != nil {
				if err := c.Scs_30kHz.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_60kHz != nil {
				if err := c.Scs_60kHz.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. pdsch-ProcessingType2-Limited: sequence
	{
		if ie.Pdsch_ProcessingType2_Limited != nil {
			c := ie.Pdsch_ProcessingType2_Limited
			if err := e.EncodeEnumerated(c.DifferentTB_PerSlot_SCS_30kHz, featureSetDownlinkV1540PdschProcessingType2LimitedDifferentTBPerSlotSCS30kHzConstraints); err != nil {
				return err
			}
		}
	}

	// 10. dl-MCS-TableAlt-DynamicIndication: enumerated
	{
		if ie.Dl_MCS_TableAlt_DynamicIndication != nil {
			if err := e.EncodeEnumerated(*ie.Dl_MCS_TableAlt_DynamicIndication, featureSetDownlinkV1540DlMCSTableAltDynamicIndicationConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV1540Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. oneFL-DMRS-TwoAdditionalDMRS-DL: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1540OneFLDMRSTwoAdditionalDMRSDLConstraints)
			if err != nil {
				return err
			}
			ie.OneFL_DMRS_TwoAdditionalDMRS_DL = &idx
		}
	}

	// 3. additionalDMRS-DL-Alt: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1540AdditionalDMRSDLAltConstraints)
			if err != nil {
				return err
			}
			ie.AdditionalDMRS_DL_Alt = &idx
		}
	}

	// 4. twoFL-DMRS-TwoAdditionalDMRS-DL: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1540TwoFLDMRSTwoAdditionalDMRSDLConstraints)
			if err != nil {
				return err
			}
			ie.TwoFL_DMRS_TwoAdditionalDMRS_DL = &idx
		}
	}

	// 5. oneFL-DMRS-ThreeAdditionalDMRS-DL: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1540OneFLDMRSThreeAdditionalDMRSDLConstraints)
			if err != nil {
				return err
			}
			ie.OneFL_DMRS_ThreeAdditionalDMRS_DL = &idx
		}
	}

	// 6. pdcch-MonitoringAnyOccasionsWithSpanGap: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.Pdcch_MonitoringAnyOccasionsWithSpanGap = &struct {
				Scs_15kHz  *int64
				Scs_30kHz  *int64
				Scs_60kHz  *int64
				Scs_120kHz *int64
			}{}
			c := ie.Pdcch_MonitoringAnyOccasionsWithSpanGap
			featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapSeq := d.NewSequenceDecoder(featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapConstraints)
			if err := featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapSeq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapSeq.IsComponentPresent(0) {
				c.Scs_15kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs15kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz) = v
			}
			if featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapSeq.IsComponentPresent(1) {
				c.Scs_30kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs30kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz) = v
			}
			if featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapSeq.IsComponentPresent(2) {
				c.Scs_60kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs60kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz) = v
			}
			if featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapSeq.IsComponentPresent(3) {
				c.Scs_120kHz = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1540PdcchMonitoringAnyOccasionsWithSpanGapScs120kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz) = v
			}
		}
	}

	// 7. pdsch-SeparationWithGap: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1540PdschSeparationWithGapConstraints)
			if err != nil {
				return err
			}
			ie.Pdsch_SeparationWithGap = &idx
		}
	}

	// 8. pdsch-ProcessingType2: sequence
	{
		if seq.IsComponentPresent(6) {
			ie.Pdsch_ProcessingType2 = &struct {
				Scs_15kHz *ProcessingParameters
				Scs_30kHz *ProcessingParameters
				Scs_60kHz *ProcessingParameters
			}{}
			c := ie.Pdsch_ProcessingType2
			featureSetDownlinkV1540PdschProcessingType2Seq := d.NewSequenceDecoder(featureSetDownlinkV1540PdschProcessingType2Constraints)
			if err := featureSetDownlinkV1540PdschProcessingType2Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1540PdschProcessingType2Seq.IsComponentPresent(0) {
				c.Scs_15kHz = new(ProcessingParameters)
				if err := (*c.Scs_15kHz).Decode(d); err != nil {
					return err
				}
			}
			if featureSetDownlinkV1540PdschProcessingType2Seq.IsComponentPresent(1) {
				c.Scs_30kHz = new(ProcessingParameters)
				if err := (*c.Scs_30kHz).Decode(d); err != nil {
					return err
				}
			}
			if featureSetDownlinkV1540PdschProcessingType2Seq.IsComponentPresent(2) {
				c.Scs_60kHz = new(ProcessingParameters)
				if err := (*c.Scs_60kHz).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. pdsch-ProcessingType2-Limited: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.Pdsch_ProcessingType2_Limited = &struct{ DifferentTB_PerSlot_SCS_30kHz int64 }{}
			c := ie.Pdsch_ProcessingType2_Limited
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1540PdschProcessingType2LimitedDifferentTBPerSlotSCS30kHzConstraints)
				if err != nil {
					return err
				}
				c.DifferentTB_PerSlot_SCS_30kHz = v
			}
		}
	}

	// 10. dl-MCS-TableAlt-DynamicIndication: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1540DlMCSTableAltDynamicIndicationConstraints)
			if err != nil {
				return err
			}
			ie.Dl_MCS_TableAlt_DynamicIndication = &idx
		}
	}

	return nil
}
