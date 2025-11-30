package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VarRA_Report_r16 struct {
	Ra_ReportList_r16     RA_ReportList_r16     `madatory`
	Plmn_IdentityList_r16 PLMN_IdentityList_r16 `madatory`
}

func (ie *VarRA_Report_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ra_ReportList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_ReportList_r16", err)
	}
	if err = ie.Plmn_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_IdentityList_r16", err)
	}
	return nil
}

func (ie *VarRA_Report_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ra_ReportList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ra_ReportList_r16", err)
	}
	if err = ie.Plmn_IdentityList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_IdentityList_r16", err)
	}
	return nil
}
