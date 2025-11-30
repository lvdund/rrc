package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayerPreferenceFR2_2_r17 struct {
	ReducedMaxMIMO_LayersFR2_2_r17 *MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17 `lb:1,ub:8,optional`
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ReducedMaxMIMO_LayersFR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReducedMaxMIMO_LayersFR2_2_r17 != nil {
		if err = ie.ReducedMaxMIMO_LayersFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxMIMO_LayersFR2_2_r17", err)
		}
	}
	return nil
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17) Decode(r *aper.AperReader) error {
	var err error
	var ReducedMaxMIMO_LayersFR2_2_r17Present bool
	if ReducedMaxMIMO_LayersFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ReducedMaxMIMO_LayersFR2_2_r17Present {
		ie.ReducedMaxMIMO_LayersFR2_2_r17 = new(MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17)
		if err = ie.ReducedMaxMIMO_LayersFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxMIMO_LayersFR2_2_r17", err)
		}
	}
	return nil
}
