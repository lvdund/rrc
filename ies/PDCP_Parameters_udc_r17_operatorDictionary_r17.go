package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Parameters_udc_r17_operatorDictionary_r17 struct {
	VersionOfDictionary_r17 int64         `lb:0,ub:15,madatory`
	AssociatedPLMN_ID_r17   PLMN_Identity `madatory`
}

func (ie *PDCP_Parameters_udc_r17_operatorDictionary_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.VersionOfDictionary_r17, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger VersionOfDictionary_r17", err)
	}
	if err = ie.AssociatedPLMN_ID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode AssociatedPLMN_ID_r17", err)
	}
	return nil
}

func (ie *PDCP_Parameters_udc_r17_operatorDictionary_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_VersionOfDictionary_r17 int64
	if tmp_int_VersionOfDictionary_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger VersionOfDictionary_r17", err)
	}
	ie.VersionOfDictionary_r17 = tmp_int_VersionOfDictionary_r17
	if err = ie.AssociatedPLMN_ID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode AssociatedPLMN_ID_r17", err)
	}
	return nil
}
