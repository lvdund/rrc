package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForNCSG_IntraFreq_r17 struct {
	ServCellId_r17         ServCellIndex                                    `madatory`
	GapIndicationIntra_r17 NeedForNCSG_IntraFreq_r17_gapIndicationIntra_r17 `madatory`
}

func (ie *NeedForNCSG_IntraFreq_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ServCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellId_r17", err)
	}
	if err = ie.GapIndicationIntra_r17.Encode(w); err != nil {
		return utils.WrapError("Encode GapIndicationIntra_r17", err)
	}
	return nil
}

func (ie *NeedForNCSG_IntraFreq_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ServCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ServCellId_r17", err)
	}
	if err = ie.GapIndicationIntra_r17.Decode(r); err != nil {
		return utils.WrapError("Decode GapIndicationIntra_r17", err)
	}
	return nil
}
