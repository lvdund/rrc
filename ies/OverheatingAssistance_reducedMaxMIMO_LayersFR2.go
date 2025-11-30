package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistance_reducedMaxMIMO_LayersFR2 struct {
	ReducedMIMO_LayersFR2_DL MIMO_LayersDL `madatory`
	ReducedMIMO_LayersFR2_UL MIMO_LayersUL `madatory`
}

func (ie *OverheatingAssistance_reducedMaxMIMO_LayersFR2) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ReducedMIMO_LayersFR2_DL.Encode(w); err != nil {
		return utils.WrapError("Encode ReducedMIMO_LayersFR2_DL", err)
	}
	if err = ie.ReducedMIMO_LayersFR2_UL.Encode(w); err != nil {
		return utils.WrapError("Encode ReducedMIMO_LayersFR2_UL", err)
	}
	return nil
}

func (ie *OverheatingAssistance_reducedMaxMIMO_LayersFR2) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ReducedMIMO_LayersFR2_DL.Decode(r); err != nil {
		return utils.WrapError("Decode ReducedMIMO_LayersFR2_DL", err)
	}
	if err = ie.ReducedMIMO_LayersFR2_UL.Decode(r); err != nil {
		return utils.WrapError("Decode ReducedMIMO_LayersFR2_UL", err)
	}
	return nil
}
