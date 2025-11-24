package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_Resource_Mobility_associatedSSB struct {
	Ssb_Index        SSB_Index `madatory`
	IsQuasiColocated bool      `madatory`
}

func (ie *CSI_RS_Resource_Mobility_associatedSSB) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Ssb_Index.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_Index", err)
	}
	if err = w.WriteBoolean(ie.IsQuasiColocated); err != nil {
		return utils.WrapError("WriteBoolean IsQuasiColocated", err)
	}
	return nil
}

func (ie *CSI_RS_Resource_Mobility_associatedSSB) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Ssb_Index.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_Index", err)
	}
	var tmp_bool_IsQuasiColocated bool
	if tmp_bool_IsQuasiColocated, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean IsQuasiColocated", err)
	}
	ie.IsQuasiColocated = tmp_bool_IsQuasiColocated
	return nil
}
