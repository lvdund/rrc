package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_eutra_nr      aper.Enumerated = 0
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_nr            aper.Enumerated = 1
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_other         aper.Enumerated = 2
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_utra_nr_other aper.Enumerated = 3
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_nr_other      aper.Enumerated = 4
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_spare3        aper.Enumerated = 5
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_spare2        aper.Enumerated = 6
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_spare1        aper.Enumerated = 7
)

type AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC", err)
	}
	return nil
}

func (ie *AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
