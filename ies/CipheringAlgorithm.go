package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CipheringAlgorithm_Enum_nea0   aper.Enumerated = 0
	CipheringAlgorithm_Enum_nea1   aper.Enumerated = 1
	CipheringAlgorithm_Enum_nea2   aper.Enumerated = 2
	CipheringAlgorithm_Enum_nea3   aper.Enumerated = 3
	CipheringAlgorithm_Enum_spare4 aper.Enumerated = 4
	CipheringAlgorithm_Enum_spare3 aper.Enumerated = 5
	CipheringAlgorithm_Enum_spare2 aper.Enumerated = 6
	CipheringAlgorithm_Enum_spare1 aper.Enumerated = 7
)

type CipheringAlgorithm struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CipheringAlgorithm) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CipheringAlgorithm", err)
	}
	return nil
}

func (ie *CipheringAlgorithm) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CipheringAlgorithm", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
