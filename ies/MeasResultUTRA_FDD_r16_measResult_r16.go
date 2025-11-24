package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultUTRA_FDD_r16_measResult_r16 struct {
	Utra_FDD_RSCP_r16 *int64 `lb:-5,ub:91,optional`
	Utra_FDD_EcN0_r16 *int64 `lb:0,ub:49,optional`
}

func (ie *MeasResultUTRA_FDD_r16_measResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Utra_FDD_RSCP_r16 != nil, ie.Utra_FDD_EcN0_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Utra_FDD_RSCP_r16 != nil {
		if err = w.WriteInteger(*ie.Utra_FDD_RSCP_r16, &uper.Constraint{Lb: -5, Ub: 91}, false); err != nil {
			return utils.WrapError("Encode Utra_FDD_RSCP_r16", err)
		}
	}
	if ie.Utra_FDD_EcN0_r16 != nil {
		if err = w.WriteInteger(*ie.Utra_FDD_EcN0_r16, &uper.Constraint{Lb: 0, Ub: 49}, false); err != nil {
			return utils.WrapError("Encode Utra_FDD_EcN0_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultUTRA_FDD_r16_measResult_r16) Decode(r *uper.UperReader) error {
	var err error
	var Utra_FDD_RSCP_r16Present bool
	if Utra_FDD_RSCP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Utra_FDD_EcN0_r16Present bool
	if Utra_FDD_EcN0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Utra_FDD_RSCP_r16Present {
		var tmp_int_Utra_FDD_RSCP_r16 int64
		if tmp_int_Utra_FDD_RSCP_r16, err = r.ReadInteger(&uper.Constraint{Lb: -5, Ub: 91}, false); err != nil {
			return utils.WrapError("Decode Utra_FDD_RSCP_r16", err)
		}
		ie.Utra_FDD_RSCP_r16 = &tmp_int_Utra_FDD_RSCP_r16
	}
	if Utra_FDD_EcN0_r16Present {
		var tmp_int_Utra_FDD_EcN0_r16 int64
		if tmp_int_Utra_FDD_EcN0_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 49}, false); err != nil {
			return utils.WrapError("Decode Utra_FDD_EcN0_r16", err)
		}
		ie.Utra_FDD_EcN0_r16 = &tmp_int_Utra_FDD_EcN0_r16
	}
	return nil
}
