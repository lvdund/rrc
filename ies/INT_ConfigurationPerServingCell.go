package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type INT_ConfigurationPerServingCell struct {
	ServingCellId ServCellIndex `madatory`
	PositionInDCI int64         `lb:0,ub:maxINT_DCI_PayloadSize_1,madatory`
}

func (ie *INT_ConfigurationPerServingCell) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ServingCellId.Encode(w); err != nil {
		return utils.WrapError("Encode ServingCellId", err)
	}
	if err = w.WriteInteger(ie.PositionInDCI, &aper.Constraint{Lb: 0, Ub: maxINT_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("WriteInteger PositionInDCI", err)
	}
	return nil
}

func (ie *INT_ConfigurationPerServingCell) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ServingCellId.Decode(r); err != nil {
		return utils.WrapError("Decode ServingCellId", err)
	}
	var tmp_int_PositionInDCI int64
	if tmp_int_PositionInDCI, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxINT_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("ReadInteger PositionInDCI", err)
	}
	ie.PositionInDCI = tmp_int_PositionInDCI
	return nil
}
