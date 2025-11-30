package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PositionVelocity_r17 struct {
	PositionX_r17  PositionStateVector_r17 `madatory`
	PositionY_r17  PositionStateVector_r17 `madatory`
	PositionZ_r17  PositionStateVector_r17 `madatory`
	VelocityVX_r17 VelocityStateVector_r17 `madatory`
	VelocityVY_r17 VelocityStateVector_r17 `madatory`
	VelocityVZ_r17 VelocityStateVector_r17 `madatory`
}

func (ie *PositionVelocity_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.PositionX_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PositionX_r17", err)
	}
	if err = ie.PositionY_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PositionY_r17", err)
	}
	if err = ie.PositionZ_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PositionZ_r17", err)
	}
	if err = ie.VelocityVX_r17.Encode(w); err != nil {
		return utils.WrapError("Encode VelocityVX_r17", err)
	}
	if err = ie.VelocityVY_r17.Encode(w); err != nil {
		return utils.WrapError("Encode VelocityVY_r17", err)
	}
	if err = ie.VelocityVZ_r17.Encode(w); err != nil {
		return utils.WrapError("Encode VelocityVZ_r17", err)
	}
	return nil
}

func (ie *PositionVelocity_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.PositionX_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PositionX_r17", err)
	}
	if err = ie.PositionY_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PositionY_r17", err)
	}
	if err = ie.PositionZ_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PositionZ_r17", err)
	}
	if err = ie.VelocityVX_r17.Decode(r); err != nil {
		return utils.WrapError("Decode VelocityVX_r17", err)
	}
	if err = ie.VelocityVY_r17.Decode(r); err != nil {
		return utils.WrapError("Decode VelocityVY_r17", err)
	}
	if err = ie.VelocityVZ_r17.Decode(r); err != nil {
		return utils.WrapError("Decode VelocityVZ_r17", err)
	}
	return nil
}
