package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type WLAN_NameList_r16 struct {
	Value []WLAN_Name_r16 `lb:1,ub:maxWLAN_Name_r16,madatory`
}

func (ie *WLAN_NameList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*WLAN_Name_r16]([]*WLAN_Name_r16{}, aper.Constraint{Lb: 1, Ub: maxWLAN_Name_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode WLAN_NameList_r16", err)
	}
	return nil
}

func (ie *WLAN_NameList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*WLAN_Name_r16]([]*WLAN_Name_r16{}, aper.Constraint{Lb: 1, Ub: maxWLAN_Name_r16}, false)
	fn := func() *WLAN_Name_r16 {
		return new(WLAN_Name_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode WLAN_NameList_r16", err)
	}
	ie.Value = []WLAN_Name_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
