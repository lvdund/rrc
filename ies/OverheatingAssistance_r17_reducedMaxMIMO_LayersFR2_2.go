package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistance_r17_reducedMaxMIMO_LayersFR2_2 struct {
	ReducedMIMO_LayersFR2_2_DL MIMO_LayersDL `madatory`
	ReducedMIMO_LayersFR2_2_UL MIMO_LayersUL `madatory`
}

func (ie *OverheatingAssistance_r17_reducedMaxMIMO_LayersFR2_2) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ReducedMIMO_LayersFR2_2_DL.Encode(w); err != nil {
		return utils.WrapError("Encode ReducedMIMO_LayersFR2_2_DL", err)
	}
	if err = ie.ReducedMIMO_LayersFR2_2_UL.Encode(w); err != nil {
		return utils.WrapError("Encode ReducedMIMO_LayersFR2_2_UL", err)
	}
	return nil
}

func (ie *OverheatingAssistance_r17_reducedMaxMIMO_LayersFR2_2) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ReducedMIMO_LayersFR2_2_DL.Decode(r); err != nil {
		return utils.WrapError("Decode ReducedMIMO_LayersFR2_2_DL", err)
	}
	if err = ie.ReducedMIMO_LayersFR2_2_UL.Decode(r); err != nil {
		return utils.WrapError("Decode ReducedMIMO_LayersFR2_2_UL", err)
	}
	return nil
}
