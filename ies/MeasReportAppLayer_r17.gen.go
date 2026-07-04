// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasReportAppLayer-r17 (line 768).

var measReportAppLayerR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measConfigAppLayerId-r17"},
		{Name: "measReportAppLayerContainer-r17", Optional: true},
		{Name: "appLayerSessionStatus-r17", Optional: true},
		{Name: "ran-VisibleMeasurements-r17", Optional: true},
	},
}

var measReportAppLayerR17MeasReportAppLayerContainerR17Constraints = per.SizeConstraints{}

const (
	MeasReportAppLayer_r17_AppLayerSessionStatus_r17_Start = 0
	MeasReportAppLayer_r17_AppLayerSessionStatus_r17_Stop  = 1
)

var measReportAppLayerR17AppLayerSessionStatusR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasReportAppLayer_r17_AppLayerSessionStatus_r17_Start, MeasReportAppLayer_r17_AppLayerSessionStatus_r17_Stop},
}

type MeasReportAppLayer_r17 struct {
	MeasConfigAppLayerId_r17        MeasConfigAppLayerId_r17
	MeasReportAppLayerContainer_r17 []byte
	AppLayerSessionStatus_r17       *int64
	Ran_VisibleMeasurements_r17     *RAN_VisibleMeasurements_r17
}

func (ie *MeasReportAppLayer_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measReportAppLayerR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasReportAppLayerContainer_r17 != nil, ie.AppLayerSessionStatus_r17 != nil, ie.Ran_VisibleMeasurements_r17 != nil}); err != nil {
		return err
	}

	// 2. measConfigAppLayerId-r17: ref
	{
		if err := ie.MeasConfigAppLayerId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. measReportAppLayerContainer-r17: octet-string
	{
		if ie.MeasReportAppLayerContainer_r17 != nil {
			if err := e.EncodeOctetString(ie.MeasReportAppLayerContainer_r17, measReportAppLayerR17MeasReportAppLayerContainerR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. appLayerSessionStatus-r17: enumerated
	{
		if ie.AppLayerSessionStatus_r17 != nil {
			if err := e.EncodeEnumerated(*ie.AppLayerSessionStatus_r17, measReportAppLayerR17AppLayerSessionStatusR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ran-VisibleMeasurements-r17: ref
	{
		if ie.Ran_VisibleMeasurements_r17 != nil {
			if err := ie.Ran_VisibleMeasurements_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasReportAppLayer_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measReportAppLayerR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measConfigAppLayerId-r17: ref
	{
		if err := ie.MeasConfigAppLayerId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. measReportAppLayerContainer-r17: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(measReportAppLayerR17MeasReportAppLayerContainerR17Constraints)
			if err != nil {
				return err
			}
			ie.MeasReportAppLayerContainer_r17 = v
		}
	}

	// 4. appLayerSessionStatus-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(measReportAppLayerR17AppLayerSessionStatusR17Constraints)
			if err != nil {
				return err
			}
			ie.AppLayerSessionStatus_r17 = &idx
		}
	}

	// 5. ran-VisibleMeasurements-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Ran_VisibleMeasurements_r17 = new(RAN_VisibleMeasurements_r17)
			if err := ie.Ran_VisibleMeasurements_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
