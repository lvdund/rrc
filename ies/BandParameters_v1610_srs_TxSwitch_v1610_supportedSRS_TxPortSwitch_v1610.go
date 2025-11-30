package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610_Enum_t1r1_t1r2                aper.Enumerated = 0
	BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610_Enum_t1r1_t1r2_t1r4           aper.Enumerated = 1
	BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610_Enum_t1r1_t1r2_t2r2_t2r4      aper.Enumerated = 2
	BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610_Enum_t1r1_t1r2_t2r2_t1r4_t2r4 aper.Enumerated = 3
	BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610_Enum_t1r1_t2r2                aper.Enumerated = 4
	BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610_Enum_t1r1_t2r2_t4r4           aper.Enumerated = 5
)

type BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610", err)
	}
	return nil
}

func (ie *BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
