package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReestabUE_Identity struct {
	C_RNTI     RNTI_Value `madatory`
	PhysCellId PhysCellId `madatory`
	ShortMAC_I ShortMAC_I `madatory`
}

func (ie *ReestabUE_Identity) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.C_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode C_RNTI", err)
	}
	if err = ie.PhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId", err)
	}
	if err = ie.ShortMAC_I.Encode(w); err != nil {
		return utils.WrapError("Encode ShortMAC_I", err)
	}
	return nil
}

func (ie *ReestabUE_Identity) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.C_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode C_RNTI", err)
	}
	if err = ie.PhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId", err)
	}
	if err = ie.ShortMAC_I.Decode(r); err != nil {
		return utils.WrapError("Decode ShortMAC_I", err)
	}
	return nil
}
