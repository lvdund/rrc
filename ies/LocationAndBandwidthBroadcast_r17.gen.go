// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LocationAndBandwidthBroadcast-r17 (line 28398).

var locationAndBandwidthBroadcastR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sameAsSib1ConfiguredLocationAndBW"},
		{Name: "locationAndBandwidth"},
	},
}

const (
	LocationAndBandwidthBroadcast_r17_SameAsSib1ConfiguredLocationAndBW = 0
	LocationAndBandwidthBroadcast_r17_LocationAndBandwidth              = 1
)

type LocationAndBandwidthBroadcast_r17 struct {
	Choice                            int
	SameAsSib1ConfiguredLocationAndBW *struct{}
	LocationAndBandwidth              *int64
}

func (ie *LocationAndBandwidthBroadcast_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(locationAndBandwidthBroadcastR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationAndBandwidthBroadcast_r17_SameAsSib1ConfiguredLocationAndBW:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case LocationAndBandwidthBroadcast_r17_LocationAndBandwidth:
		if err := e.EncodeInteger((*ie.LocationAndBandwidth), per.Constrained(0, 37949)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "LocationAndBandwidthBroadcast-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *LocationAndBandwidthBroadcast_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(locationAndBandwidthBroadcastR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "LocationAndBandwidthBroadcast-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case LocationAndBandwidthBroadcast_r17_SameAsSib1ConfiguredLocationAndBW:
		ie.SameAsSib1ConfiguredLocationAndBW = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case LocationAndBandwidthBroadcast_r17_LocationAndBandwidth:
		ie.LocationAndBandwidth = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 37949))
		if err != nil {
			return err
		}
		(*ie.LocationAndBandwidth) = v
	default:
		return &per.DecodeError{TypeName: "LocationAndBandwidthBroadcast-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
