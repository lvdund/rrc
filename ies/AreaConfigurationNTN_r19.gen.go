// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AreaConfigurationNTN-r19 (line 601).

var areaConfigurationNTNR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "areaCoordinates-r19"},
	},
}

var areaConfigurationNTN_r19AreaCoordinatesR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "polygonArea-r19"},
		{Name: "circleArea-r19"},
	},
}

const (
	AreaConfigurationNTN_r19_AreaCoordinates_r19_PolygonArea_r19 = 0
	AreaConfigurationNTN_r19_AreaCoordinates_r19_CircleArea_r19  = 1
)

type AreaConfigurationNTN_r19 struct {
	AreaCoordinates_r19 struct {
		Choice          int
		PolygonArea_r19 *[]byte
		CircleArea_r19  *struct {
			CircularAreaReferenceLocation_r19 ReferenceLocation_r17
			CircularAreaRadius_r19            int64
		}
	}
}

func (ie *AreaConfigurationNTN_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(areaConfigurationNTNR19Constraints)
	_ = seq

	// 1. areaCoordinates-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(areaConfigurationNTN_r19AreaCoordinatesR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.AreaCoordinates_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.AreaCoordinates_r19.Choice {
		case AreaConfigurationNTN_r19_AreaCoordinates_r19_PolygonArea_r19:
			if err := e.EncodeOctetString((*ie.AreaCoordinates_r19.PolygonArea_r19), per.SizeConstraints{}); err != nil {
				return err
			}
		case AreaConfigurationNTN_r19_AreaCoordinates_r19_CircleArea_r19:
			c := (*ie.AreaCoordinates_r19.CircleArea_r19)
			if err := c.CircularAreaReferenceLocation_r19.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.CircularAreaRadius_r19, per.Constrained(1, 65535)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.AreaCoordinates_r19.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *AreaConfigurationNTN_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(areaConfigurationNTNR19Constraints)
	_ = seq

	// 1. areaCoordinates-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(areaConfigurationNTN_r19AreaCoordinatesR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.AreaCoordinates_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case AreaConfigurationNTN_r19_AreaCoordinates_r19_PolygonArea_r19:
			ie.AreaCoordinates_r19.PolygonArea_r19 = new([]byte)
			v, err := d.DecodeOctetString(per.SizeConstraints{})
			if err != nil {
				return err
			}
			(*ie.AreaCoordinates_r19.PolygonArea_r19) = v
		case AreaConfigurationNTN_r19_AreaCoordinates_r19_CircleArea_r19:
			ie.AreaCoordinates_r19.CircleArea_r19 = &struct {
				CircularAreaReferenceLocation_r19 ReferenceLocation_r17
				CircularAreaRadius_r19            int64
			}{}
			c := (*ie.AreaCoordinates_r19.CircleArea_r19)
			{
				if err := c.CircularAreaReferenceLocation_r19.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 65535))
				if err != nil {
					return err
				}
				c.CircularAreaRadius_r19 = v
			}
		}
	}

	return nil
}
