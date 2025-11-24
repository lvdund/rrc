package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapsIntraFreq_r16 struct {
	ServCellId_r16         ServCellIndex                                   `madatory`
	GapIndicationIntra_r16 NeedForGapsIntraFreq_r16_gapIndicationIntra_r16 `madatory`
}

func (ie *NeedForGapsIntraFreq_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ServCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellId_r16", err)
	}
	if err = ie.GapIndicationIntra_r16.Encode(w); err != nil {
		return utils.WrapError("Encode GapIndicationIntra_r16", err)
	}
	return nil
}

func (ie *NeedForGapsIntraFreq_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ServCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ServCellId_r16", err)
	}
	if err = ie.GapIndicationIntra_r16.Decode(r); err != nil {
		return utils.WrapError("Decode GapIndicationIntra_r16", err)
	}
	return nil
}
