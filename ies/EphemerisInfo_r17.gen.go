// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EphemerisInfo-r17 (line 8262).

var ephemerisInfoR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "positionVelocity-r17"},
		{Name: "orbital-r17"},
	},
}

const (
	EphemerisInfo_r17_PositionVelocity_r17 = 0
	EphemerisInfo_r17_Orbital_r17          = 1
)

type EphemerisInfo_r17 struct {
	Choice               int
	PositionVelocity_r17 *PositionVelocity_r17
	Orbital_r17          *Orbital_r17
}

func (ie *EphemerisInfo_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(ephemerisInfoR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case EphemerisInfo_r17_PositionVelocity_r17:
		if err := ie.PositionVelocity_r17.Encode(e); err != nil {
			return err
		}
	case EphemerisInfo_r17_Orbital_r17:
		if err := ie.Orbital_r17.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "EphemerisInfo-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *EphemerisInfo_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(ephemerisInfoR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "EphemerisInfo-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case EphemerisInfo_r17_PositionVelocity_r17:
		ie.PositionVelocity_r17 = new(PositionVelocity_r17)
		if err := ie.PositionVelocity_r17.Decode(d); err != nil {
			return err
		}
	case EphemerisInfo_r17_Orbital_r17:
		ie.Orbital_r17 = new(Orbital_r17)
		if err := ie.Orbital_r17.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "EphemerisInfo-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
