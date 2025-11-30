package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringPerCat struct {
	AccessCategory          int64                   `lb:1,ub:maxAccessCat_1,madatory`
	Uac_barringInfoSetIndex UAC_BarringInfoSetIndex `madatory`
}

func (ie *UAC_BarringPerCat) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.AccessCategory, &aper.Constraint{Lb: 1, Ub: maxAccessCat_1}, false); err != nil {
		return utils.WrapError("WriteInteger AccessCategory", err)
	}
	if err = ie.Uac_barringInfoSetIndex.Encode(w); err != nil {
		return utils.WrapError("Encode Uac_barringInfoSetIndex", err)
	}
	return nil
}

func (ie *UAC_BarringPerCat) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_AccessCategory int64
	if tmp_int_AccessCategory, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxAccessCat_1}, false); err != nil {
		return utils.WrapError("ReadInteger AccessCategory", err)
	}
	ie.AccessCategory = tmp_int_AccessCategory
	if err = ie.Uac_barringInfoSetIndex.Decode(r); err != nil {
		return utils.WrapError("Decode Uac_barringInfoSetIndex", err)
	}
	return nil
}
