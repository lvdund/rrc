package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VarRLF_Report_r16 struct {
	Rlf_Report_r16        RLF_Report_r16         `madatory`
	Plmn_IdentityList_r16 PLMN_IdentityList2_r16 `madatory`
}

func (ie *VarRLF_Report_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Rlf_Report_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rlf_Report_r16", err)
	}
	if err = ie.Plmn_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_IdentityList_r16", err)
	}
	return nil
}

func (ie *VarRLF_Report_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Rlf_Report_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rlf_Report_r16", err)
	}
	if err = ie.Plmn_IdentityList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_IdentityList_r16", err)
	}
	return nil
}
