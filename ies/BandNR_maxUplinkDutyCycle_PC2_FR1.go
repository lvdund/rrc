package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n60  aper.Enumerated = 0
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n70  aper.Enumerated = 1
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n80  aper.Enumerated = 2
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n90  aper.Enumerated = 3
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n100 aper.Enumerated = 4
)

type BandNR_maxUplinkDutyCycle_PC2_FR1 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *BandNR_maxUplinkDutyCycle_PC2_FR1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode BandNR_maxUplinkDutyCycle_PC2_FR1", err)
	}
	return nil
}

func (ie *BandNR_maxUplinkDutyCycle_PC2_FR1) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode BandNR_maxUplinkDutyCycle_PC2_FR1", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
