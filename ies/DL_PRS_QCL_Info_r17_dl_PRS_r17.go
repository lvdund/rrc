package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DL_PRS_QCL_Info_r17_dl_PRS_r17 struct {
	Qcl_DL_PRS_ResourceID_r17 NR_DL_PRS_ResourceID_r17 `madatory`
}

func (ie *DL_PRS_QCL_Info_r17_dl_PRS_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Qcl_DL_PRS_ResourceID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Qcl_DL_PRS_ResourceID_r17", err)
	}
	return nil
}

func (ie *DL_PRS_QCL_Info_r17_dl_PRS_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Qcl_DL_PRS_ResourceID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Qcl_DL_PRS_ResourceID_r17", err)
	}
	return nil
}
