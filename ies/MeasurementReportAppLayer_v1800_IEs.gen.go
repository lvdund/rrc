// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasurementReportAppLayer-v1800-IEs (line 759).

var measurementReportAppLayerV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measurementReportAppLayerList-v1800", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type MeasurementReportAppLayer_v1800_IEs struct {
	MeasurementReportAppLayerList_v1800 *MeasurementReportAppLayerList_v1800
	NonCriticalExtension                *struct{}
}

func (ie *MeasurementReportAppLayer_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measurementReportAppLayerV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasurementReportAppLayerList_v1800 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. measurementReportAppLayerList-v1800: ref
	{
		if ie.MeasurementReportAppLayerList_v1800 != nil {
			if err := ie.MeasurementReportAppLayerList_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *MeasurementReportAppLayer_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measurementReportAppLayerV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measurementReportAppLayerList-v1800: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasurementReportAppLayerList_v1800 = new(MeasurementReportAppLayerList_v1800)
			if err := ie.MeasurementReportAppLayerList_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
