// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1860 (line 17899).

var cAParametersNRV1860Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxSSB-PerFreqLayerL1-Meas-r18", Optional: true},
	},
}

var cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedMaxSSB-PerFreqLayerWithoutGaps-r18", Optional: true},
		{Name: "supportedMaxSSB-PerFreqLayerWithGaps-r18", Optional: true},
	},
}

const (
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N1  = 0
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N2  = 1
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N3  = 2
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N4  = 3
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N5  = 4
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N6  = 5
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N7  = 6
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N8  = 7
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N12 = 8
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N16 = 9
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N20 = 10
	CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N24 = 11
)

var cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18SupportedMaxSSBPerFreqLayerWithoutGapsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N1, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N2, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N3, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N4, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N5, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N6, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N7, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N8, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N12, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N16, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N20, CA_ParametersNR_v1860_MaxSSB_PerFreqLayerL1_Meas_r18_SupportedMaxSSB_PerFreqLayerWithoutGaps_r18_N24},
}

type CA_ParametersNR_v1860 struct {
	MaxSSB_PerFreqLayerL1_Meas_r18 *struct {
		SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 *int64
		SupportedMaxSSB_PerFreqLayerWithGaps_r18    *int64
	}
}

func (ie *CA_ParametersNR_v1860) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1860Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxSSB_PerFreqLayerL1_Meas_r18 != nil}); err != nil {
		return err
	}

	// 2. maxSSB-PerFreqLayerL1-Meas-r18: sequence
	{
		if ie.MaxSSB_PerFreqLayerL1_Meas_r18 != nil {
			c := ie.MaxSSB_PerFreqLayerL1_Meas_r18
			cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18Seq := e.NewSequenceEncoder(cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18Constraints)
			if err := cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18Seq.EncodePreamble([]bool{c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 != nil, c.SupportedMaxSSB_PerFreqLayerWithGaps_r18 != nil}); err != nil {
				return err
			}
			if c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 != nil {
				if err := e.EncodeEnumerated((*c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18), cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18SupportedMaxSSBPerFreqLayerWithoutGapsR18Constraints); err != nil {
					return err
				}
			}
			if c.SupportedMaxSSB_PerFreqLayerWithGaps_r18 != nil {
				if err := e.EncodeInteger((*c.SupportedMaxSSB_PerFreqLayerWithGaps_r18), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1860) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1860Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxSSB-PerFreqLayerL1-Meas-r18: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.MaxSSB_PerFreqLayerL1_Meas_r18 = &struct {
				SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 *int64
				SupportedMaxSSB_PerFreqLayerWithGaps_r18    *int64
			}{}
			c := ie.MaxSSB_PerFreqLayerL1_Meas_r18
			cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18Seq := d.NewSequenceDecoder(cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18Constraints)
			if err := cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18Seq.IsComponentPresent(0) {
				c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18SupportedMaxSSBPerFreqLayerWithoutGapsR18Constraints)
				if err != nil {
					return err
				}
				(*c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18) = v
			}
			if cAParametersNRV1860MaxSSBPerFreqLayerL1MeasR18Seq.IsComponentPresent(1) {
				c.SupportedMaxSSB_PerFreqLayerWithGaps_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.SupportedMaxSSB_PerFreqLayerWithGaps_r18) = v
			}
		}
	}

	return nil
}
