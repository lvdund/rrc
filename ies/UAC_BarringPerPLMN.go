package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringPerPLMN struct {
	Plmn_IdentityIndex    int64                                     `lb:1,ub:maxPLMN,madatory`
	Uac_ACBarringListType *UAC_BarringPerPLMN_uac_ACBarringListType `lb:maxAccessCat_1,ub:maxAccessCat_1,optional`
}

func (ie *UAC_BarringPerPLMN) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Uac_ACBarringListType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Plmn_IdentityIndex, &uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("WriteInteger Plmn_IdentityIndex", err)
	}
	if ie.Uac_ACBarringListType != nil {
		if err = ie.Uac_ACBarringListType.Encode(w); err != nil {
			return utils.WrapError("Encode Uac_ACBarringListType", err)
		}
	}
	return nil
}

func (ie *UAC_BarringPerPLMN) Decode(r *uper.UperReader) error {
	var err error
	var Uac_ACBarringListTypePresent bool
	if Uac_ACBarringListTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Plmn_IdentityIndex int64
	if tmp_int_Plmn_IdentityIndex, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("ReadInteger Plmn_IdentityIndex", err)
	}
	ie.Plmn_IdentityIndex = tmp_int_Plmn_IdentityIndex
	if Uac_ACBarringListTypePresent {
		ie.Uac_ACBarringListType = new(UAC_BarringPerPLMN_uac_ACBarringListType)
		if err = ie.Uac_ACBarringListType.Decode(r); err != nil {
			return utils.WrapError("Decode Uac_ACBarringListType", err)
		}
	}
	return nil
}
