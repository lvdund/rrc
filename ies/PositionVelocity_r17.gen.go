// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PositionVelocity-r17 (line 8267).

var positionVelocityR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "positionX-r17"},
		{Name: "positionY-r17"},
		{Name: "positionZ-r17"},
		{Name: "velocityVX-r17"},
		{Name: "velocityVY-r17"},
		{Name: "velocityVZ-r17"},
	},
}

type PositionVelocity_r17 struct {
	PositionX_r17  PositionStateVector_r17
	PositionY_r17  PositionStateVector_r17
	PositionZ_r17  PositionStateVector_r17
	VelocityVX_r17 VelocityStateVector_r17
	VelocityVY_r17 VelocityStateVector_r17
	VelocityVZ_r17 VelocityStateVector_r17
}

func (ie *PositionVelocity_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(positionVelocityR17Constraints)
	_ = seq

	// 1. positionX-r17: ref
	{
		if err := ie.PositionX_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. positionY-r17: ref
	{
		if err := ie.PositionY_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. positionZ-r17: ref
	{
		if err := ie.PositionZ_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. velocityVX-r17: ref
	{
		if err := ie.VelocityVX_r17.Encode(e); err != nil {
			return err
		}
	}

	// 5. velocityVY-r17: ref
	{
		if err := ie.VelocityVY_r17.Encode(e); err != nil {
			return err
		}
	}

	// 6. velocityVZ-r17: ref
	{
		if err := ie.VelocityVZ_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PositionVelocity_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(positionVelocityR17Constraints)
	_ = seq

	// 1. positionX-r17: ref
	{
		if err := ie.PositionX_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. positionY-r17: ref
	{
		if err := ie.PositionY_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. positionZ-r17: ref
	{
		if err := ie.PositionZ_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. velocityVX-r17: ref
	{
		if err := ie.VelocityVX_r17.Decode(d); err != nil {
			return err
		}
	}

	// 5. velocityVY-r17: ref
	{
		if err := ie.VelocityVY_r17.Decode(d); err != nil {
			return err
		}
	}

	// 6. velocityVZ-r17: ref
	{
		if err := ie.VelocityVZ_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
