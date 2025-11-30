package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistance_r17 struct {
	ReducedMaxBW_FR2_2_r17     *OverheatingAssistance_r17_reducedMaxBW_FR2_2_r17     `optional`
	ReducedMaxMIMO_LayersFR2_2 *OverheatingAssistance_r17_reducedMaxMIMO_LayersFR2_2 `optional`
}

func (ie *OverheatingAssistance_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ReducedMaxBW_FR2_2_r17 != nil, ie.ReducedMaxMIMO_LayersFR2_2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReducedMaxBW_FR2_2_r17 != nil {
		if err = ie.ReducedMaxBW_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxBW_FR2_2_r17", err)
		}
	}
	if ie.ReducedMaxMIMO_LayersFR2_2 != nil {
		if err = ie.ReducedMaxMIMO_LayersFR2_2.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxMIMO_LayersFR2_2", err)
		}
	}
	return nil
}

func (ie *OverheatingAssistance_r17) Decode(r *aper.AperReader) error {
	var err error
	var ReducedMaxBW_FR2_2_r17Present bool
	if ReducedMaxBW_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReducedMaxMIMO_LayersFR2_2Present bool
	if ReducedMaxMIMO_LayersFR2_2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ReducedMaxBW_FR2_2_r17Present {
		ie.ReducedMaxBW_FR2_2_r17 = new(OverheatingAssistance_r17_reducedMaxBW_FR2_2_r17)
		if err = ie.ReducedMaxBW_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxBW_FR2_2_r17", err)
		}
	}
	if ReducedMaxMIMO_LayersFR2_2Present {
		ie.ReducedMaxMIMO_LayersFR2_2 = new(OverheatingAssistance_r17_reducedMaxMIMO_LayersFR2_2)
		if err = ie.ReducedMaxMIMO_LayersFR2_2.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxMIMO_LayersFR2_2", err)
		}
	}
	return nil
}
