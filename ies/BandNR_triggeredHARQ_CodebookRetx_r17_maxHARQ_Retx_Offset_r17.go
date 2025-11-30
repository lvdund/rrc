package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n4  aper.Enumerated = 0
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n6  aper.Enumerated = 1
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n8  aper.Enumerated = 2
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n10 aper.Enumerated = 3
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n12 aper.Enumerated = 4
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n14 aper.Enumerated = 5
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n16 aper.Enumerated = 6
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n18 aper.Enumerated = 7
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n20 aper.Enumerated = 8
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n22 aper.Enumerated = 9
	BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17_Enum_n24 aper.Enumerated = 10
)

type BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17 struct {
	Value aper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17", err)
	}
	return nil
}

func (ie *BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode BandNR_triggeredHARQ_CodebookRetx_r17_maxHARQ_Retx_Offset_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
