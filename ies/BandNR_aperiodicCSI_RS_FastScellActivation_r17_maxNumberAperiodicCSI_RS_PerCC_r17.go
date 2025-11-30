package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n8   aper.Enumerated = 0
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n16  aper.Enumerated = 1
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n32  aper.Enumerated = 2
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n48  aper.Enumerated = 3
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n64  aper.Enumerated = 4
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n128 aper.Enumerated = 5
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n255 aper.Enumerated = 6
)

type BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17", err)
	}
	return nil
}

func (ie *BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
