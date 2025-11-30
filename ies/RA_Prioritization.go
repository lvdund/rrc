package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RA_Prioritization struct {
	PowerRampingStepHighPriority RA_Prioritization_powerRampingStepHighPriority `madatory`
	ScalingFactorBI              *RA_Prioritization_scalingFactorBI             `optional`
}

func (ie *RA_Prioritization) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ScalingFactorBI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PowerRampingStepHighPriority.Encode(w); err != nil {
		return utils.WrapError("Encode PowerRampingStepHighPriority", err)
	}
	if ie.ScalingFactorBI != nil {
		if err = ie.ScalingFactorBI.Encode(w); err != nil {
			return utils.WrapError("Encode ScalingFactorBI", err)
		}
	}
	return nil
}

func (ie *RA_Prioritization) Decode(r *aper.AperReader) error {
	var err error
	var ScalingFactorBIPresent bool
	if ScalingFactorBIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PowerRampingStepHighPriority.Decode(r); err != nil {
		return utils.WrapError("Decode PowerRampingStepHighPriority", err)
	}
	if ScalingFactorBIPresent {
		ie.ScalingFactorBI = new(RA_Prioritization_scalingFactorBI)
		if err = ie.ScalingFactorBI.Decode(r); err != nil {
			return utils.WrapError("Decode ScalingFactorBI", err)
		}
	}
	return nil
}
