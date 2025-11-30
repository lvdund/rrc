package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_bwp_SameNumerology_Enum_upto2 aper.Enumerated = 0
	BandNR_bwp_SameNumerology_Enum_upto4 aper.Enumerated = 1
)

type BandNR_bwp_SameNumerology struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *BandNR_bwp_SameNumerology) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode BandNR_bwp_SameNumerology", err)
	}
	return nil
}

func (ie *BandNR_bwp_SameNumerology) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode BandNR_bwp_SameNumerology", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
