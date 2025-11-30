package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17_Enum_n1 aper.Enumerated = 0
	BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17_Enum_n2 aper.Enumerated = 1
	BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17_Enum_n3 aper.Enumerated = 2
	BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17_Enum_n4 aper.Enumerated = 3
)

type BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17", err)
	}
	return nil
}

func (ie *BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
