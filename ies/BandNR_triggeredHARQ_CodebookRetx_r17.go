package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_triggeredHARQ_CodebookRetx_r17 struct {
	MinHARQ_Retx_Offset_r17 BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17 `madatory`
	MaxHARQ_Retx_Offset_r17 BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17 `madatory`
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MinHARQ_Retx_Offset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MinHARQ_Retx_Offset_r17", err)
	}
	if err = ie.MaxHARQ_Retx_Offset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxHARQ_Retx_Offset_r17", err)
	}
	return nil
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MinHARQ_Retx_Offset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MinHARQ_Retx_Offset_r17", err)
	}
	if err = ie.MaxHARQ_Retx_Offset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxHARQ_Retx_Offset_r17", err)
	}
	return nil
}
