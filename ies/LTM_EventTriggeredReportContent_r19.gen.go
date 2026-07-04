// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-EventTriggeredReportContent-r19 (line 8876).

var lTMEventTriggeredReportContentR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberOfReportedBeams-r19"},
		{Name: "allowReportAnyBeam-r19", Optional: true},
		{Name: "reportCurrentBeam-r19", Optional: true},
	},
}

var lTMEventTriggeredReportContentR19MaxNumberOfReportedBeamsR19Constraints = per.Constrained(1, 16)

const (
	LTM_EventTriggeredReportContent_r19_AllowReportAnyBeam_r19_Enabled = 0
)

var lTMEventTriggeredReportContentR19AllowReportAnyBeamR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_EventTriggeredReportContent_r19_AllowReportAnyBeam_r19_Enabled},
}

const (
	LTM_EventTriggeredReportContent_r19_ReportCurrentBeam_r19_Enabled = 0
)

var lTMEventTriggeredReportContentR19ReportCurrentBeamR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_EventTriggeredReportContent_r19_ReportCurrentBeam_r19_Enabled},
}

type LTM_EventTriggeredReportContent_r19 struct {
	MaxNumberOfReportedBeams_r19 int64
	AllowReportAnyBeam_r19       *int64
	ReportCurrentBeam_r19        *int64
}

func (ie *LTM_EventTriggeredReportContent_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMEventTriggeredReportContentR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AllowReportAnyBeam_r19 != nil, ie.ReportCurrentBeam_r19 != nil}); err != nil {
		return err
	}

	// 3. maxNumberOfReportedBeams-r19: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberOfReportedBeams_r19, lTMEventTriggeredReportContentR19MaxNumberOfReportedBeamsR19Constraints); err != nil {
			return err
		}
	}

	// 4. allowReportAnyBeam-r19: enumerated
	{
		if ie.AllowReportAnyBeam_r19 != nil {
			if err := e.EncodeEnumerated(*ie.AllowReportAnyBeam_r19, lTMEventTriggeredReportContentR19AllowReportAnyBeamR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. reportCurrentBeam-r19: enumerated
	{
		if ie.ReportCurrentBeam_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ReportCurrentBeam_r19, lTMEventTriggeredReportContentR19ReportCurrentBeamR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_EventTriggeredReportContent_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMEventTriggeredReportContentR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. maxNumberOfReportedBeams-r19: integer
	{
		v0, err := d.DecodeInteger(lTMEventTriggeredReportContentR19MaxNumberOfReportedBeamsR19Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberOfReportedBeams_r19 = v0
	}

	// 4. allowReportAnyBeam-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(lTMEventTriggeredReportContentR19AllowReportAnyBeamR19Constraints)
			if err != nil {
				return err
			}
			ie.AllowReportAnyBeam_r19 = &idx
		}
	}

	// 5. reportCurrentBeam-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(lTMEventTriggeredReportContentR19ReportCurrentBeamR19Constraints)
			if err != nil {
				return err
			}
			ie.ReportCurrentBeam_r19 = &idx
		}
	}

	return nil
}
