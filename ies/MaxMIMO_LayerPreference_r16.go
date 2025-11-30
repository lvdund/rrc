package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayerPreference_r16 struct {
	ReducedMaxMIMO_LayersFR1_r16 *MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR1_r16 `lb:1,ub:8,optional`
	ReducedMaxMIMO_LayersFR2_r16 *MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR2_r16 `lb:1,ub:8,optional`
}

func (ie *MaxMIMO_LayerPreference_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ReducedMaxMIMO_LayersFR1_r16 != nil, ie.ReducedMaxMIMO_LayersFR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReducedMaxMIMO_LayersFR1_r16 != nil {
		if err = ie.ReducedMaxMIMO_LayersFR1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxMIMO_LayersFR1_r16", err)
		}
	}
	if ie.ReducedMaxMIMO_LayersFR2_r16 != nil {
		if err = ie.ReducedMaxMIMO_LayersFR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxMIMO_LayersFR2_r16", err)
		}
	}
	return nil
}

func (ie *MaxMIMO_LayerPreference_r16) Decode(r *aper.AperReader) error {
	var err error
	var ReducedMaxMIMO_LayersFR1_r16Present bool
	if ReducedMaxMIMO_LayersFR1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReducedMaxMIMO_LayersFR2_r16Present bool
	if ReducedMaxMIMO_LayersFR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ReducedMaxMIMO_LayersFR1_r16Present {
		ie.ReducedMaxMIMO_LayersFR1_r16 = new(MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR1_r16)
		if err = ie.ReducedMaxMIMO_LayersFR1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxMIMO_LayersFR1_r16", err)
		}
	}
	if ReducedMaxMIMO_LayersFR2_r16Present {
		ie.ReducedMaxMIMO_LayersFR2_r16 = new(MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR2_r16)
		if err = ie.ReducedMaxMIMO_LayersFR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxMIMO_LayersFR2_r16", err)
		}
	}
	return nil
}
