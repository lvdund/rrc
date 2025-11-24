package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_Parameters_v1610 struct {
	Mac_ParametersFRX_Diff_r16 *MAC_ParametersFRX_Diff_r16 `optional`
}

func (ie *MAC_Parameters_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Mac_ParametersFRX_Diff_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mac_ParametersFRX_Diff_r16 != nil {
		if err = ie.Mac_ParametersFRX_Diff_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}

func (ie *MAC_Parameters_v1610) Decode(r *uper.UperReader) error {
	var err error
	var Mac_ParametersFRX_Diff_r16Present bool
	if Mac_ParametersFRX_Diff_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Mac_ParametersFRX_Diff_r16Present {
		ie.Mac_ParametersFRX_Diff_r16 = new(MAC_ParametersFRX_Diff_r16)
		if err = ie.Mac_ParametersFRX_Diff_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}
