package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandUTRA_FDD_r16_Enum_bandI      aper.Enumerated = 0
	SupportedBandUTRA_FDD_r16_Enum_bandII     aper.Enumerated = 1
	SupportedBandUTRA_FDD_r16_Enum_bandIII    aper.Enumerated = 2
	SupportedBandUTRA_FDD_r16_Enum_bandIV     aper.Enumerated = 3
	SupportedBandUTRA_FDD_r16_Enum_bandV      aper.Enumerated = 4
	SupportedBandUTRA_FDD_r16_Enum_bandVI     aper.Enumerated = 5
	SupportedBandUTRA_FDD_r16_Enum_bandVII    aper.Enumerated = 6
	SupportedBandUTRA_FDD_r16_Enum_bandVIII   aper.Enumerated = 7
	SupportedBandUTRA_FDD_r16_Enum_bandIX     aper.Enumerated = 8
	SupportedBandUTRA_FDD_r16_Enum_bandX      aper.Enumerated = 9
	SupportedBandUTRA_FDD_r16_Enum_bandXI     aper.Enumerated = 10
	SupportedBandUTRA_FDD_r16_Enum_bandXII    aper.Enumerated = 11
	SupportedBandUTRA_FDD_r16_Enum_bandXIII   aper.Enumerated = 12
	SupportedBandUTRA_FDD_r16_Enum_bandXIV    aper.Enumerated = 13
	SupportedBandUTRA_FDD_r16_Enum_bandXV     aper.Enumerated = 14
	SupportedBandUTRA_FDD_r16_Enum_bandXVI    aper.Enumerated = 15
	SupportedBandUTRA_FDD_r16_Enum_bandXVII   aper.Enumerated = 16
	SupportedBandUTRA_FDD_r16_Enum_bandXVIII  aper.Enumerated = 17
	SupportedBandUTRA_FDD_r16_Enum_bandXIX    aper.Enumerated = 18
	SupportedBandUTRA_FDD_r16_Enum_bandXX     aper.Enumerated = 19
	SupportedBandUTRA_FDD_r16_Enum_bandXXI    aper.Enumerated = 20
	SupportedBandUTRA_FDD_r16_Enum_bandXXII   aper.Enumerated = 21
	SupportedBandUTRA_FDD_r16_Enum_bandXXIII  aper.Enumerated = 22
	SupportedBandUTRA_FDD_r16_Enum_bandXXIV   aper.Enumerated = 23
	SupportedBandUTRA_FDD_r16_Enum_bandXXV    aper.Enumerated = 24
	SupportedBandUTRA_FDD_r16_Enum_bandXXVI   aper.Enumerated = 25
	SupportedBandUTRA_FDD_r16_Enum_bandXXVII  aper.Enumerated = 26
	SupportedBandUTRA_FDD_r16_Enum_bandXXVIII aper.Enumerated = 27
	SupportedBandUTRA_FDD_r16_Enum_bandXXIX   aper.Enumerated = 28
	SupportedBandUTRA_FDD_r16_Enum_bandXXX    aper.Enumerated = 29
	SupportedBandUTRA_FDD_r16_Enum_bandXXXI   aper.Enumerated = 30
	SupportedBandUTRA_FDD_r16_Enum_bandXXXII  aper.Enumerated = 31
)

type SupportedBandUTRA_FDD_r16 struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *SupportedBandUTRA_FDD_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode SupportedBandUTRA_FDD_r16", err)
	}
	return nil
}

func (ie *SupportedBandUTRA_FDD_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode SupportedBandUTRA_FDD_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
