// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IntendedServiceAreaInfo-areaCoordinates-r19 (line 4760).

var intendedServiceAreaInfoAreaCoordinatesR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "polygonArea-r19"},
		{Name: "circleArea-r19"},
	},
}

const (
	IntendedServiceAreaInfo_AreaCoordinates_r19_PolygonArea_r19 = 0
	IntendedServiceAreaInfo_AreaCoordinates_r19_CircleArea_r19  = 1
)

type IntendedServiceAreaInfo_AreaCoordinates_r19 struct {
	Choice          int
	PolygonArea_r19 *[]byte
	CircleArea_r19  *struct {
		Center_r19 ReferenceLocation_r17
		Radius_r19 int64
	}
}

func (ie *IntendedServiceAreaInfo_AreaCoordinates_r19) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(intendedServiceAreaInfoAreaCoordinatesR19Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case IntendedServiceAreaInfo_AreaCoordinates_r19_PolygonArea_r19:
		if err := e.EncodeOctetString((*ie.PolygonArea_r19), per.SizeConstraints{}); err != nil {
			return err
		}
	case IntendedServiceAreaInfo_AreaCoordinates_r19_CircleArea_r19:
		c := (*ie.CircleArea_r19)
		if err := c.Center_r19.Encode(e); err != nil {
			return err
		}
		if err := e.EncodeInteger(c.Radius_r19, per.Constrained(0, 65535)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "IntendedServiceAreaInfo-areaCoordinates-r19",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *IntendedServiceAreaInfo_AreaCoordinates_r19) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(intendedServiceAreaInfoAreaCoordinatesR19Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "IntendedServiceAreaInfo-areaCoordinates-r19", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case IntendedServiceAreaInfo_AreaCoordinates_r19_PolygonArea_r19:
		ie.PolygonArea_r19 = new([]byte)
		v, err := d.DecodeOctetString(per.SizeConstraints{})
		if err != nil {
			return err
		}
		(*ie.PolygonArea_r19) = v
	case IntendedServiceAreaInfo_AreaCoordinates_r19_CircleArea_r19:
		ie.CircleArea_r19 = &struct {
			Center_r19 ReferenceLocation_r17
			Radius_r19 int64
		}{}
		c := (*ie.CircleArea_r19)
		{
			if err := c.Center_r19.Decode(d); err != nil {
				return err
			}
		}
		{
			v, err := d.DecodeInteger(per.Constrained(0, 65535))
			if err != nil {
				return err
			}
			c.Radius_r19 = v
		}
	default:
		return &per.DecodeError{TypeName: "IntendedServiceAreaInfo-areaCoordinates-r19", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
