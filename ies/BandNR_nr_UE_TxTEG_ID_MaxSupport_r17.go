package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_nr_UE_TxTEG_ID_MaxSupport_r17_Enum_n1 aper.Enumerated = 0
	BandNR_nr_UE_TxTEG_ID_MaxSupport_r17_Enum_n2 aper.Enumerated = 1
	BandNR_nr_UE_TxTEG_ID_MaxSupport_r17_Enum_n3 aper.Enumerated = 2
	BandNR_nr_UE_TxTEG_ID_MaxSupport_r17_Enum_n4 aper.Enumerated = 3
	BandNR_nr_UE_TxTEG_ID_MaxSupport_r17_Enum_n6 aper.Enumerated = 4
	BandNR_nr_UE_TxTEG_ID_MaxSupport_r17_Enum_n8 aper.Enumerated = 5
)

type BandNR_nr_UE_TxTEG_ID_MaxSupport_r17 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *BandNR_nr_UE_TxTEG_ID_MaxSupport_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode BandNR_nr_UE_TxTEG_ID_MaxSupport_r17", err)
	}
	return nil
}

func (ie *BandNR_nr_UE_TxTEG_ID_MaxSupport_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode BandNR_nr_UE_TxTEG_ID_MaxSupport_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
