package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasResult_r16 struct {
	Sl_ResultDMRS_r16 *SL_MeasQuantityResult_r16 `optional`
}

func (ie *SL_MeasResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ResultDMRS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_ResultDMRS_r16 != nil {
		if err = ie.Sl_ResultDMRS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ResultDMRS_r16", err)
		}
	}
	return nil
}

func (ie *SL_MeasResult_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_ResultDMRS_r16Present bool
	if Sl_ResultDMRS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_ResultDMRS_r16Present {
		ie.Sl_ResultDMRS_r16 = new(SL_MeasQuantityResult_r16)
		if err = ie.Sl_ResultDMRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ResultDMRS_r16", err)
		}
	}
	return nil
}
