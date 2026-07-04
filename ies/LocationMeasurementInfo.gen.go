// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LocationMeasurementInfo (line 8560).

var locationMeasurementInfoConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eutra-RSTD"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "eutra-FineTimingDetection"},
		{Name: "nr-PRS-Measurement-r16"},
	},
}

const (
	LocationMeasurementInfo_Eutra_RSTD = 0
)

type LocationMeasurementInfo struct {
	Choice     int
	Eutra_RSTD *EUTRA_RSTD_InfoList
}

func (ie *LocationMeasurementInfo) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(locationMeasurementInfoConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationMeasurementInfo_Eutra_RSTD:
		if err := ie.Eutra_RSTD.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "LocationMeasurementInfo",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *LocationMeasurementInfo) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(locationMeasurementInfoConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "LocationMeasurementInfo", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case LocationMeasurementInfo_Eutra_RSTD:
		ie.Eutra_RSTD = new(EUTRA_RSTD_InfoList)
		if err := ie.Eutra_RSTD.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "LocationMeasurementInfo", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
