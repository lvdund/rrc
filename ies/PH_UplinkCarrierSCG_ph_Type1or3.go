package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PH_UplinkCarrierSCG_ph_Type1or3_Enum_type1 aper.Enumerated = 0
	PH_UplinkCarrierSCG_ph_Type1or3_Enum_type3 aper.Enumerated = 1
)

type PH_UplinkCarrierSCG_ph_Type1or3 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PH_UplinkCarrierSCG_ph_Type1or3) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PH_UplinkCarrierSCG_ph_Type1or3", err)
	}
	return nil
}

func (ie *PH_UplinkCarrierSCG_ph_Type1or3) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PH_UplinkCarrierSCG_ph_Type1or3", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
