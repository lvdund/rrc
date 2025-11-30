package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB17_IEs_r17_validityDuration_r17_Enum_t1       aper.Enumerated = 0
	SIB17_IEs_r17_validityDuration_r17_Enum_t2       aper.Enumerated = 1
	SIB17_IEs_r17_validityDuration_r17_Enum_t4       aper.Enumerated = 2
	SIB17_IEs_r17_validityDuration_r17_Enum_t8       aper.Enumerated = 3
	SIB17_IEs_r17_validityDuration_r17_Enum_t16      aper.Enumerated = 4
	SIB17_IEs_r17_validityDuration_r17_Enum_t32      aper.Enumerated = 5
	SIB17_IEs_r17_validityDuration_r17_Enum_t64      aper.Enumerated = 6
	SIB17_IEs_r17_validityDuration_r17_Enum_t128     aper.Enumerated = 7
	SIB17_IEs_r17_validityDuration_r17_Enum_t256     aper.Enumerated = 8
	SIB17_IEs_r17_validityDuration_r17_Enum_t512     aper.Enumerated = 9
	SIB17_IEs_r17_validityDuration_r17_Enum_infinity aper.Enumerated = 10
	SIB17_IEs_r17_validityDuration_r17_Enum_spare5   aper.Enumerated = 11
	SIB17_IEs_r17_validityDuration_r17_Enum_spare4   aper.Enumerated = 12
	SIB17_IEs_r17_validityDuration_r17_Enum_spare3   aper.Enumerated = 13
	SIB17_IEs_r17_validityDuration_r17_Enum_spare2   aper.Enumerated = 14
	SIB17_IEs_r17_validityDuration_r17_Enum_spare1   aper.Enumerated = 15
)

type SIB17_IEs_r17_validityDuration_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SIB17_IEs_r17_validityDuration_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SIB17_IEs_r17_validityDuration_r17", err)
	}
	return nil
}

func (ie *SIB17_IEs_r17_validityDuration_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SIB17_IEs_r17_validityDuration_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
