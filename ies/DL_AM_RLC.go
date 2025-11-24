package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DL_AM_RLC struct {
	Sn_FieldLength   *SN_FieldLengthAM `optional`
	T_Reassembly     T_Reassembly      `madatory`
	T_StatusProhibit T_StatusProhibit  `madatory`
}

func (ie *DL_AM_RLC) Encode(w *uper.UperWriter) error {
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
	if err = ie.T_Reassembly.Encode(w); err != nil {
		return utils.WrapError("Encode T_Reassembly", err)
	}
	if err = ie.T_StatusProhibit.Encode(w); err != nil {
		return utils.WrapError("Encode T_StatusProhibit", err)
	}
	return nil
}

func (ie *DL_AM_RLC) Decode(r *uper.UperReader) error {
	var err error
	var Sn_FieldLengthPresent bool
	if Sn_FieldLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sn_FieldLengthPresent {
		ie.Sn_FieldLength = new(SN_FieldLengthAM)
		if err = ie.Sn_FieldLength.Decode(r); err != nil {
			return utils.WrapError("Decode Sn_FieldLength", err)
		}
	}
	if err = ie.T_Reassembly.Decode(r); err != nil {
		return utils.WrapError("Decode T_Reassembly", err)
	}
	if err = ie.T_StatusProhibit.Decode(r); err != nil {
		return utils.WrapError("Decode T_StatusProhibit", err)
	}
	return nil
}
