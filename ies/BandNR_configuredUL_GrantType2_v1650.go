package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_configuredUL_GrantType2_v1650_Enum_supported aper.Enumerated = 0
)

type BandNR_configuredUL_GrantType2_v1650 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *BandNR_configuredUL_GrantType2_v1650) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode BandNR_configuredUL_GrantType2_v1650", err)
	}
	return nil
}

func (ie *BandNR_configuredUL_GrantType2_v1650) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode BandNR_configuredUL_GrantType2_v1650", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
