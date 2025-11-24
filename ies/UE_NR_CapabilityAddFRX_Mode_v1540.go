package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_CapabilityAddFRX_Mode_v1540 struct {
	Ims_ParametersFRX_Diff *IMS_ParametersFRX_Diff `optional`
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ims_ParametersFRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ims_ParametersFRX_Diff != nil {
		if err = ie.Ims_ParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode Ims_ParametersFRX_Diff", err)
		}
	}
	return nil
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1540) Decode(r *uper.UperReader) error {
	var err error
	var Ims_ParametersFRX_DiffPresent bool
	if Ims_ParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ims_ParametersFRX_DiffPresent {
		ie.Ims_ParametersFRX_Diff = new(IMS_ParametersFRX_Diff)
		if err = ie.Ims_ParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode Ims_ParametersFRX_Diff", err)
		}
	}
	return nil
}
