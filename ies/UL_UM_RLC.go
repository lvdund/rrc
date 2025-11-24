package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_UM_RLC struct {
	Sn_FieldLength *SN_FieldLengthUM `optional`
}

func (ie *UL_UM_RLC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sn_FieldLength != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sn_FieldLength != nil {
		if err = ie.Sn_FieldLength.Encode(w); err != nil {
			return utils.WrapError("Encode Sn_FieldLength", err)
		}
	}
	return nil
}

func (ie *UL_UM_RLC) Decode(r *uper.UperReader) error {
	var err error
	var Sn_FieldLengthPresent bool
	if Sn_FieldLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sn_FieldLengthPresent {
		ie.Sn_FieldLength = new(SN_FieldLengthUM)
		if err = ie.Sn_FieldLength.Decode(r); err != nil {
			return utils.WrapError("Decode Sn_FieldLength", err)
		}
	}
	return nil
}
