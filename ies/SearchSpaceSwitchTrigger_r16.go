package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceSwitchTrigger_r16 struct {
	ServingCellId_r16 ServCellIndex `madatory`
	PositionInDCI_r16 int64         `lb:0,ub:maxSFI_DCI_PayloadSize_1,madatory`
}

func (ie *SearchSpaceSwitchTrigger_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ServingCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ServingCellId_r16", err)
	}
	if err = w.WriteInteger(ie.PositionInDCI_r16, &aper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("WriteInteger PositionInDCI_r16", err)
	}
	return nil
}

func (ie *SearchSpaceSwitchTrigger_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ServingCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ServingCellId_r16", err)
	}
	var tmp_int_PositionInDCI_r16 int64
	if tmp_int_PositionInDCI_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("ReadInteger PositionInDCI_r16", err)
	}
	ie.PositionInDCI_r16 = tmp_int_PositionInDCI_r16
	return nil
}
