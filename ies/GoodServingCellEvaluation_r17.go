package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type GoodServingCellEvaluation_r17 struct {
	Offset_r17 *GoodServingCellEvaluation_r17_offset_r17 `optional`
}

func (ie *GoodServingCellEvaluation_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Offset_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Offset_r17 != nil {
		if err = ie.Offset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Offset_r17", err)
		}
	}
	return nil
}

func (ie *GoodServingCellEvaluation_r17) Decode(r *aper.AperReader) error {
	var err error
	var Offset_r17Present bool
	if Offset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Offset_r17Present {
		ie.Offset_r17 = new(GoodServingCellEvaluation_r17_offset_r17)
		if err = ie.Offset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Offset_r17", err)
		}
	}
	return nil
}
