package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17_Enum_supported aper.Enumerated = 0
)

type BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17", err)
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
