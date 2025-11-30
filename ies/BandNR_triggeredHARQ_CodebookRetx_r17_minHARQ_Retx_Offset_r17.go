package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17_Enum_n_7 aper.Enumerated = 0
	BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17_Enum_n_5 aper.Enumerated = 1
	BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17_Enum_n_3 aper.Enumerated = 2
	BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17_Enum_n_1 aper.Enumerated = 3
	BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17_Enum_n1  aper.Enumerated = 4
)

type BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17", err)
	}
	return nil
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode BandNR_triggeredHARQ_CodebookRetx_r17_minHARQ_Retx_Offset_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
