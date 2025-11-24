package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RelayUE_Config_r17 struct {
	ThreshHighRelay_r17 *RSRP_Range `optional`
	ThreshLowRelay_r17  *RSRP_Range `optional`
	HystMaxRelay_r17    *Hysteresis `optional`
	HystMinRelay_r17    *Hysteresis `optional`
}

func (ie *SL_RelayUE_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ThreshHighRelay_r17 != nil, ie.ThreshLowRelay_r17 != nil, ie.HystMaxRelay_r17 != nil, ie.HystMinRelay_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ThreshHighRelay_r17 != nil {
		if err = ie.ThreshHighRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ThreshHighRelay_r17", err)
		}
	}
	if ie.ThreshLowRelay_r17 != nil {
		if err = ie.ThreshLowRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ThreshLowRelay_r17", err)
		}
	}
	if ie.HystMaxRelay_r17 != nil {
		if err = ie.HystMaxRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HystMaxRelay_r17", err)
		}
	}
	if ie.HystMinRelay_r17 != nil {
		if err = ie.HystMinRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HystMinRelay_r17", err)
		}
	}
	return nil
}

func (ie *SL_RelayUE_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var ThreshHighRelay_r17Present bool
	if ThreshHighRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ThreshLowRelay_r17Present bool
	if ThreshLowRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HystMaxRelay_r17Present bool
	if HystMaxRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HystMinRelay_r17Present bool
	if HystMinRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ThreshHighRelay_r17Present {
		ie.ThreshHighRelay_r17 = new(RSRP_Range)
		if err = ie.ThreshHighRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ThreshHighRelay_r17", err)
		}
	}
	if ThreshLowRelay_r17Present {
		ie.ThreshLowRelay_r17 = new(RSRP_Range)
		if err = ie.ThreshLowRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ThreshLowRelay_r17", err)
		}
	}
	if HystMaxRelay_r17Present {
		ie.HystMaxRelay_r17 = new(Hysteresis)
		if err = ie.HystMaxRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HystMaxRelay_r17", err)
		}
	}
	if HystMinRelay_r17Present {
		ie.HystMinRelay_r17 = new(Hysteresis)
		if err = ie.HystMinRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HystMinRelay_r17", err)
		}
	}
	return nil
}
