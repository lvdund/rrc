package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_maxUplinkDutyCycle_FR2_Enum_n15  aper.Enumerated = 0
	BandNR_maxUplinkDutyCycle_FR2_Enum_n20  aper.Enumerated = 1
	BandNR_maxUplinkDutyCycle_FR2_Enum_n25  aper.Enumerated = 2
	BandNR_maxUplinkDutyCycle_FR2_Enum_n30  aper.Enumerated = 3
	BandNR_maxUplinkDutyCycle_FR2_Enum_n40  aper.Enumerated = 4
	BandNR_maxUplinkDutyCycle_FR2_Enum_n50  aper.Enumerated = 5
	BandNR_maxUplinkDutyCycle_FR2_Enum_n60  aper.Enumerated = 6
	BandNR_maxUplinkDutyCycle_FR2_Enum_n70  aper.Enumerated = 7
	BandNR_maxUplinkDutyCycle_FR2_Enum_n80  aper.Enumerated = 8
	BandNR_maxUplinkDutyCycle_FR2_Enum_n90  aper.Enumerated = 9
	BandNR_maxUplinkDutyCycle_FR2_Enum_n100 aper.Enumerated = 10
)

type BandNR_maxUplinkDutyCycle_FR2 struct {
	Value aper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *BandNR_maxUplinkDutyCycle_FR2) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode BandNR_maxUplinkDutyCycle_FR2", err)
	}
	return nil
}

func (ie *BandNR_maxUplinkDutyCycle_FR2) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode BandNR_maxUplinkDutyCycle_FR2", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
