package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersSidelink_r17 struct {
	Drx_OnSidelink_r17 *MAC_ParametersSidelink_r17_drx_OnSidelink_r17 `optional`
}

func (ie *MAC_ParametersSidelink_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Drx_OnSidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Drx_OnSidelink_r17 != nil {
		if err = ie.Drx_OnSidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_OnSidelink_r17", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersSidelink_r17) Decode(r *uper.UperReader) error {
	var err error
	var Drx_OnSidelink_r17Present bool
	if Drx_OnSidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Drx_OnSidelink_r17Present {
		ie.Drx_OnSidelink_r17 = new(MAC_ParametersSidelink_r17_drx_OnSidelink_r17)
		if err = ie.Drx_OnSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_OnSidelink_r17", err)
		}
	}
	return nil
}
