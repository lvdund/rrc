package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TMGI_r17 struct {
	Plmn_Id_r17   TMGI_r17_plmn_Id_r17 `lb:1,ub:maxPLMN,madatory`
	ServiceId_r17 []byte               `lb:3,ub:3,madatory`
}

func (ie *TMGI_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Plmn_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_Id_r17", err)
	}
	if err = w.WriteOctetString(ie.ServiceId_r17, &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteOctetString ServiceId_r17", err)
	}
	return nil
}

func (ie *TMGI_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Plmn_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_Id_r17", err)
	}
	var tmp_os_ServiceId_r17 []byte
	if tmp_os_ServiceId_r17, err = r.ReadOctetString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadOctetString ServiceId_r17", err)
	}
	ie.ServiceId_r17 = tmp_os_ServiceId_r17
	return nil
}
