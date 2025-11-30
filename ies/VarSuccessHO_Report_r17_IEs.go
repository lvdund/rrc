package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VarSuccessHO_Report_r17_IEs struct {
	SuccessHO_Report_r17  SuccessHO_Report_r17   `madatory`
	Plmn_IdentityList_r17 PLMN_IdentityList2_r16 `madatory`
}

func (ie *VarSuccessHO_Report_r17_IEs) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.SuccessHO_Report_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SuccessHO_Report_r17", err)
	}
	if err = ie.Plmn_IdentityList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_IdentityList_r17", err)
	}
	return nil
}

func (ie *VarSuccessHO_Report_r17_IEs) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.SuccessHO_Report_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SuccessHO_Report_r17", err)
	}
	if err = ie.Plmn_IdentityList_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_IdentityList_r17", err)
	}
	return nil
}
