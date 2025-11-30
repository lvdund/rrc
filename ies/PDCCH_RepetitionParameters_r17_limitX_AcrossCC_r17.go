package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n4      aper.Enumerated = 0
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n8      aper.Enumerated = 1
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n16     aper.Enumerated = 2
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n32     aper.Enumerated = 3
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n44     aper.Enumerated = 4
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n64     aper.Enumerated = 5
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n128    aper.Enumerated = 6
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n256    aper.Enumerated = 7
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n512    aper.Enumerated = 8
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_nolimit aper.Enumerated = 9
)

type PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17 struct {
	Value aper.Enumerated `lb:0,ub:9,madatory`
}

func (ie *PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17", err)
	}
	return nil
}

func (ie *PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
