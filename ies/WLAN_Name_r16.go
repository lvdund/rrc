package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type WLAN_Name_r16 struct {
	Value []byte `lb:1,ub:32,madatory`
}

func (ie *WLAN_Name_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("Encode WLAN_Name_r16", err)
	}
	return nil
}

func (ie *WLAN_Name_r16) Decode(r *aper.AperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("Decode WLAN_Name_r16", err)
	}
	ie.Value = v
	return nil
}
