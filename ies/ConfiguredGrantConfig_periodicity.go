package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_periodicity_Enum_sym2       aper.Enumerated = 0
	ConfiguredGrantConfig_periodicity_Enum_sym7       aper.Enumerated = 1
	ConfiguredGrantConfig_periodicity_Enum_sym1x14    aper.Enumerated = 2
	ConfiguredGrantConfig_periodicity_Enum_sym2x14    aper.Enumerated = 3
	ConfiguredGrantConfig_periodicity_Enum_sym4x14    aper.Enumerated = 4
	ConfiguredGrantConfig_periodicity_Enum_sym5x14    aper.Enumerated = 5
	ConfiguredGrantConfig_periodicity_Enum_sym8x14    aper.Enumerated = 6
	ConfiguredGrantConfig_periodicity_Enum_sym10x14   aper.Enumerated = 7
	ConfiguredGrantConfig_periodicity_Enum_sym16x14   aper.Enumerated = 8
	ConfiguredGrantConfig_periodicity_Enum_sym20x14   aper.Enumerated = 9
	ConfiguredGrantConfig_periodicity_Enum_sym32x14   aper.Enumerated = 10
	ConfiguredGrantConfig_periodicity_Enum_sym40x14   aper.Enumerated = 11
	ConfiguredGrantConfig_periodicity_Enum_sym64x14   aper.Enumerated = 12
	ConfiguredGrantConfig_periodicity_Enum_sym80x14   aper.Enumerated = 13
	ConfiguredGrantConfig_periodicity_Enum_sym128x14  aper.Enumerated = 14
	ConfiguredGrantConfig_periodicity_Enum_sym160x14  aper.Enumerated = 15
	ConfiguredGrantConfig_periodicity_Enum_sym256x14  aper.Enumerated = 16
	ConfiguredGrantConfig_periodicity_Enum_sym320x14  aper.Enumerated = 17
	ConfiguredGrantConfig_periodicity_Enum_sym512x14  aper.Enumerated = 18
	ConfiguredGrantConfig_periodicity_Enum_sym640x14  aper.Enumerated = 19
	ConfiguredGrantConfig_periodicity_Enum_sym1024x14 aper.Enumerated = 20
	ConfiguredGrantConfig_periodicity_Enum_sym1280x14 aper.Enumerated = 21
	ConfiguredGrantConfig_periodicity_Enum_sym2560x14 aper.Enumerated = 22
	ConfiguredGrantConfig_periodicity_Enum_sym5120x14 aper.Enumerated = 23
	ConfiguredGrantConfig_periodicity_Enum_sym6       aper.Enumerated = 24
	ConfiguredGrantConfig_periodicity_Enum_sym1x12    aper.Enumerated = 25
	ConfiguredGrantConfig_periodicity_Enum_sym2x12    aper.Enumerated = 26
	ConfiguredGrantConfig_periodicity_Enum_sym4x12    aper.Enumerated = 27
	ConfiguredGrantConfig_periodicity_Enum_sym5x12    aper.Enumerated = 28
	ConfiguredGrantConfig_periodicity_Enum_sym8x12    aper.Enumerated = 29
	ConfiguredGrantConfig_periodicity_Enum_sym10x12   aper.Enumerated = 30
	ConfiguredGrantConfig_periodicity_Enum_sym16x12   aper.Enumerated = 31
	ConfiguredGrantConfig_periodicity_Enum_sym20x12   aper.Enumerated = 32
	ConfiguredGrantConfig_periodicity_Enum_sym32x12   aper.Enumerated = 33
	ConfiguredGrantConfig_periodicity_Enum_sym40x12   aper.Enumerated = 34
	ConfiguredGrantConfig_periodicity_Enum_sym64x12   aper.Enumerated = 35
	ConfiguredGrantConfig_periodicity_Enum_sym80x12   aper.Enumerated = 36
	ConfiguredGrantConfig_periodicity_Enum_sym128x12  aper.Enumerated = 37
	ConfiguredGrantConfig_periodicity_Enum_sym160x12  aper.Enumerated = 38
	ConfiguredGrantConfig_periodicity_Enum_sym256x12  aper.Enumerated = 39
	ConfiguredGrantConfig_periodicity_Enum_sym320x12  aper.Enumerated = 40
	ConfiguredGrantConfig_periodicity_Enum_sym512x12  aper.Enumerated = 41
	ConfiguredGrantConfig_periodicity_Enum_sym640x12  aper.Enumerated = 42
	ConfiguredGrantConfig_periodicity_Enum_sym1280x12 aper.Enumerated = 43
	ConfiguredGrantConfig_periodicity_Enum_sym2560x12 aper.Enumerated = 44
)

type ConfiguredGrantConfig_periodicity struct {
	Value aper.Enumerated `lb:0,ub:44,madatory`
}

func (ie *ConfiguredGrantConfig_periodicity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 44}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_periodicity", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_periodicity) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 44}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_periodicity", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
