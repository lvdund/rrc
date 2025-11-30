package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedConfigFR2_r17 struct {
	HighSpeedMeasFlagFR2_r17              *HighSpeedConfigFR2_r17_highSpeedMeasFlagFR2_r17              `optional`
	HighSpeedDeploymentTypeFR2_r17        *HighSpeedConfigFR2_r17_highSpeedDeploymentTypeFR2_r17        `optional`
	HighSpeedLargeOneStepUL_TimingFR2_r17 *HighSpeedConfigFR2_r17_highSpeedLargeOneStepUL_TimingFR2_r17 `optional`
}

func (ie *HighSpeedConfigFR2_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.HighSpeedMeasFlagFR2_r17 != nil, ie.HighSpeedDeploymentTypeFR2_r17 != nil, ie.HighSpeedLargeOneStepUL_TimingFR2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.HighSpeedMeasFlagFR2_r17 != nil {
		if err = ie.HighSpeedMeasFlagFR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedMeasFlagFR2_r17", err)
		}
	}
	if ie.HighSpeedDeploymentTypeFR2_r17 != nil {
		if err = ie.HighSpeedDeploymentTypeFR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedDeploymentTypeFR2_r17", err)
		}
	}
	if ie.HighSpeedLargeOneStepUL_TimingFR2_r17 != nil {
		if err = ie.HighSpeedLargeOneStepUL_TimingFR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedLargeOneStepUL_TimingFR2_r17", err)
		}
	}
	return nil
}

func (ie *HighSpeedConfigFR2_r17) Decode(r *aper.AperReader) error {
	var err error
	var HighSpeedMeasFlagFR2_r17Present bool
	if HighSpeedMeasFlagFR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighSpeedDeploymentTypeFR2_r17Present bool
	if HighSpeedDeploymentTypeFR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighSpeedLargeOneStepUL_TimingFR2_r17Present bool
	if HighSpeedLargeOneStepUL_TimingFR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if HighSpeedMeasFlagFR2_r17Present {
		ie.HighSpeedMeasFlagFR2_r17 = new(HighSpeedConfigFR2_r17_highSpeedMeasFlagFR2_r17)
		if err = ie.HighSpeedMeasFlagFR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedMeasFlagFR2_r17", err)
		}
	}
	if HighSpeedDeploymentTypeFR2_r17Present {
		ie.HighSpeedDeploymentTypeFR2_r17 = new(HighSpeedConfigFR2_r17_highSpeedDeploymentTypeFR2_r17)
		if err = ie.HighSpeedDeploymentTypeFR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedDeploymentTypeFR2_r17", err)
		}
	}
	if HighSpeedLargeOneStepUL_TimingFR2_r17Present {
		ie.HighSpeedLargeOneStepUL_TimingFR2_r17 = new(HighSpeedConfigFR2_r17_highSpeedLargeOneStepUL_TimingFR2_r17)
		if err = ie.HighSpeedLargeOneStepUL_TimingFR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedLargeOneStepUL_TimingFR2_r17", err)
		}
	}
	return nil
}
