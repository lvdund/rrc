// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplink-v1540 (line 20108).

var featureSetUplinkV1540Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "zeroSlotOffsetAperiodicSRS", Optional: true},
		{Name: "pa-PhaseDiscontinuityImpacts", Optional: true},
		{Name: "pusch-SeparationWithGap", Optional: true},
		{Name: "pusch-ProcessingType2", Optional: true},
		{Name: "ul-MCS-TableAlt-DynamicIndication", Optional: true},
	},
}

const (
	FeatureSetUplink_v1540_ZeroSlotOffsetAperiodicSRS_Supported = 0
)

var featureSetUplinkV1540ZeroSlotOffsetAperiodicSRSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1540_ZeroSlotOffsetAperiodicSRS_Supported},
}

const (
	FeatureSetUplink_v1540_Pa_PhaseDiscontinuityImpacts_Supported = 0
)

var featureSetUplinkV1540PaPhaseDiscontinuityImpactsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1540_Pa_PhaseDiscontinuityImpacts_Supported},
}

const (
	FeatureSetUplink_v1540_Pusch_SeparationWithGap_Supported = 0
)

var featureSetUplinkV1540PuschSeparationWithGapConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1540_Pusch_SeparationWithGap_Supported},
}

const (
	FeatureSetUplink_v1540_Ul_MCS_TableAlt_DynamicIndication_Supported = 0
)

var featureSetUplinkV1540UlMCSTableAltDynamicIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1540_Ul_MCS_TableAlt_DynamicIndication_Supported},
}

var featureSetUplinkV1540PuschProcessingType2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

type FeatureSetUplink_v1540 struct {
	ZeroSlotOffsetAperiodicSRS   *int64
	Pa_PhaseDiscontinuityImpacts *int64
	Pusch_SeparationWithGap      *int64
	Pusch_ProcessingType2        *struct {
		Scs_15kHz *ProcessingParameters
		Scs_30kHz *ProcessingParameters
		Scs_60kHz *ProcessingParameters
	}
	Ul_MCS_TableAlt_DynamicIndication *int64
}

func (ie *FeatureSetUplink_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV1540Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ZeroSlotOffsetAperiodicSRS != nil, ie.Pa_PhaseDiscontinuityImpacts != nil, ie.Pusch_SeparationWithGap != nil, ie.Pusch_ProcessingType2 != nil, ie.Ul_MCS_TableAlt_DynamicIndication != nil}); err != nil {
		return err
	}

	// 2. zeroSlotOffsetAperiodicSRS: enumerated
	{
		if ie.ZeroSlotOffsetAperiodicSRS != nil {
			if err := e.EncodeEnumerated(*ie.ZeroSlotOffsetAperiodicSRS, featureSetUplinkV1540ZeroSlotOffsetAperiodicSRSConstraints); err != nil {
				return err
			}
		}
	}

	// 3. pa-PhaseDiscontinuityImpacts: enumerated
	{
		if ie.Pa_PhaseDiscontinuityImpacts != nil {
			if err := e.EncodeEnumerated(*ie.Pa_PhaseDiscontinuityImpacts, featureSetUplinkV1540PaPhaseDiscontinuityImpactsConstraints); err != nil {
				return err
			}
		}
	}

	// 4. pusch-SeparationWithGap: enumerated
	{
		if ie.Pusch_SeparationWithGap != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_SeparationWithGap, featureSetUplinkV1540PuschSeparationWithGapConstraints); err != nil {
				return err
			}
		}
	}

	// 5. pusch-ProcessingType2: sequence
	{
		if ie.Pusch_ProcessingType2 != nil {
			c := ie.Pusch_ProcessingType2
			featureSetUplinkV1540PuschProcessingType2Seq := e.NewSequenceEncoder(featureSetUplinkV1540PuschProcessingType2Constraints)
			if err := featureSetUplinkV1540PuschProcessingType2Seq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
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

	// 6. ul-MCS-TableAlt-DynamicIndication: enumerated
	{
		if ie.Ul_MCS_TableAlt_DynamicIndication != nil {
			if err := e.EncodeEnumerated(*ie.Ul_MCS_TableAlt_DynamicIndication, featureSetUplinkV1540UlMCSTableAltDynamicIndicationConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV1540Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. zeroSlotOffsetAperiodicSRS: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1540ZeroSlotOffsetAperiodicSRSConstraints)
			if err != nil {
				return err
			}
			ie.ZeroSlotOffsetAperiodicSRS = &idx
		}
	}

	// 3. pa-PhaseDiscontinuityImpacts: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1540PaPhaseDiscontinuityImpactsConstraints)
			if err != nil {
				return err
			}
			ie.Pa_PhaseDiscontinuityImpacts = &idx
		}
	}

	// 4. pusch-SeparationWithGap: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1540PuschSeparationWithGapConstraints)
			if err != nil {
				return err
			}
			ie.Pusch_SeparationWithGap = &idx
		}
	}

	// 5. pusch-ProcessingType2: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.Pusch_ProcessingType2 = &struct {
				Scs_15kHz *ProcessingParameters
				Scs_30kHz *ProcessingParameters
				Scs_60kHz *ProcessingParameters
			}{}
			c := ie.Pusch_ProcessingType2
			featureSetUplinkV1540PuschProcessingType2Seq := d.NewSequenceDecoder(featureSetUplinkV1540PuschProcessingType2Constraints)
			if err := featureSetUplinkV1540PuschProcessingType2Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1540PuschProcessingType2Seq.IsComponentPresent(0) {
				c.Scs_15kHz = new(ProcessingParameters)
				if err := (*c.Scs_15kHz).Decode(d); err != nil {
					return err
				}
			}
			if featureSetUplinkV1540PuschProcessingType2Seq.IsComponentPresent(1) {
				c.Scs_30kHz = new(ProcessingParameters)
				if err := (*c.Scs_30kHz).Decode(d); err != nil {
					return err
				}
			}
			if featureSetUplinkV1540PuschProcessingType2Seq.IsComponentPresent(2) {
				c.Scs_60kHz = new(ProcessingParameters)
				if err := (*c.Scs_60kHz).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. ul-MCS-TableAlt-DynamicIndication: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1540UlMCSTableAltDynamicIndicationConstraints)
			if err != nil {
				return err
			}
			ie.Ul_MCS_TableAlt_DynamicIndication = &idx
		}
	}

	return nil
}
