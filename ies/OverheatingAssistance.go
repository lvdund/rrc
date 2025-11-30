package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistance struct {
	ReducedMaxCCs            *ReducedMaxCCs_r16                              `optional`
	ReducedMaxBW_FR1         *ReducedMaxBW_FRx_r16                           `optional`
	ReducedMaxBW_FR2         *ReducedMaxBW_FRx_r16                           `optional`
	ReducedMaxMIMO_LayersFR1 *OverheatingAssistance_reducedMaxMIMO_LayersFR1 `optional`
	ReducedMaxMIMO_LayersFR2 *OverheatingAssistance_reducedMaxMIMO_LayersFR2 `optional`
}

func (ie *OverheatingAssistance) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ReducedMaxCCs != nil, ie.ReducedMaxBW_FR1 != nil, ie.ReducedMaxBW_FR2 != nil, ie.ReducedMaxMIMO_LayersFR1 != nil, ie.ReducedMaxMIMO_LayersFR2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReducedMaxCCs != nil {
		if err = ie.ReducedMaxCCs.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxCCs", err)
		}
	}
	if ie.ReducedMaxBW_FR1 != nil {
		if err = ie.ReducedMaxBW_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxBW_FR1", err)
		}
	}
	if ie.ReducedMaxBW_FR2 != nil {
		if err = ie.ReducedMaxBW_FR2.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxBW_FR2", err)
		}
	}
	if ie.ReducedMaxMIMO_LayersFR1 != nil {
		if err = ie.ReducedMaxMIMO_LayersFR1.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxMIMO_LayersFR1", err)
		}
	}
	if ie.ReducedMaxMIMO_LayersFR2 != nil {
		if err = ie.ReducedMaxMIMO_LayersFR2.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxMIMO_LayersFR2", err)
		}
	}
	return nil
}

func (ie *OverheatingAssistance) Decode(r *aper.AperReader) error {
	var err error
	var ReducedMaxCCsPresent bool
	if ReducedMaxCCsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReducedMaxBW_FR1Present bool
	if ReducedMaxBW_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReducedMaxBW_FR2Present bool
	if ReducedMaxBW_FR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReducedMaxMIMO_LayersFR1Present bool
	if ReducedMaxMIMO_LayersFR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReducedMaxMIMO_LayersFR2Present bool
	if ReducedMaxMIMO_LayersFR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ReducedMaxCCsPresent {
		ie.ReducedMaxCCs = new(ReducedMaxCCs_r16)
		if err = ie.ReducedMaxCCs.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxCCs", err)
		}
	}
	if ReducedMaxBW_FR1Present {
		ie.ReducedMaxBW_FR1 = new(ReducedMaxBW_FRx_r16)
		if err = ie.ReducedMaxBW_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxBW_FR1", err)
		}
	}
	if ReducedMaxBW_FR2Present {
		ie.ReducedMaxBW_FR2 = new(ReducedMaxBW_FRx_r16)
		if err = ie.ReducedMaxBW_FR2.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxBW_FR2", err)
		}
	}
	if ReducedMaxMIMO_LayersFR1Present {
		ie.ReducedMaxMIMO_LayersFR1 = new(OverheatingAssistance_reducedMaxMIMO_LayersFR1)
		if err = ie.ReducedMaxMIMO_LayersFR1.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxMIMO_LayersFR1", err)
		}
	}
	if ReducedMaxMIMO_LayersFR2Present {
		ie.ReducedMaxMIMO_LayersFR2 = new(OverheatingAssistance_reducedMaxMIMO_LayersFR2)
		if err = ie.ReducedMaxMIMO_LayersFR2.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxMIMO_LayersFR2", err)
		}
	}
	return nil
}
