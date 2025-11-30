package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PHR_Config_phr_Tx_PowerFactorChange_Enum_dB1      aper.Enumerated = 0
	PHR_Config_phr_Tx_PowerFactorChange_Enum_dB3      aper.Enumerated = 1
	PHR_Config_phr_Tx_PowerFactorChange_Enum_dB6      aper.Enumerated = 2
	PHR_Config_phr_Tx_PowerFactorChange_Enum_infinity aper.Enumerated = 3
)

type PHR_Config_phr_Tx_PowerFactorChange struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PHR_Config_phr_Tx_PowerFactorChange) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PHR_Config_phr_Tx_PowerFactorChange", err)
	}
	return nil
}

func (ie *PHR_Config_phr_Tx_PowerFactorChange) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PHR_Config_phr_Tx_PowerFactorChange", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
