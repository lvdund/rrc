package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ThreshX_Q struct {
	ThreshX_HighQ ReselectionThresholdQ `madatory`
	ThreshX_LowQ  ReselectionThresholdQ `madatory`
}

func (ie *ThreshX_Q) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ThreshX_HighQ.Encode(w); err != nil {
		return utils.WrapError("Encode ThreshX_HighQ", err)
	}
	if err = ie.ThreshX_LowQ.Encode(w); err != nil {
		return utils.WrapError("Encode ThreshX_LowQ", err)
	}
	return nil
}

func (ie *ThreshX_Q) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ThreshX_HighQ.Decode(r); err != nil {
		return utils.WrapError("Decode ThreshX_HighQ", err)
	}
	if err = ie.ThreshX_LowQ.Decode(r); err != nil {
		return utils.WrapError("Decode ThreshX_LowQ", err)
	}
	return nil
}
