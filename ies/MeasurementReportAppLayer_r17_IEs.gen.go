// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasurementReportAppLayer-r17-IEs (line 753).

var measurementReportAppLayerR17IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measurementReportAppLayerList-r17"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var measurementReportAppLayerR17IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type MeasurementReportAppLayer_r17_IEs struct {
	MeasurementReportAppLayerList_r17 MeasurementReportAppLayerList_r17
	LateNonCriticalExtension          []byte
	NonCriticalExtension              *MeasurementReportAppLayer_v1800_IEs
}

func (ie *MeasurementReportAppLayer_r17_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measurementReportAppLayerR17IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. measurementReportAppLayerList-r17: ref
	{
		if err := ie.MeasurementReportAppLayerList_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, measurementReportAppLayerR17IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasurementReportAppLayer_r17_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measurementReportAppLayerR17IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measurementReportAppLayerList-r17: ref
	{
		if err := ie.MeasurementReportAppLayerList_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(measurementReportAppLayerR17IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(MeasurementReportAppLayer_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
