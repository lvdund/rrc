package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_1        aper.Enumerated = 0
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_2        aper.Enumerated = 1
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_3        aper.Enumerated = 2
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_4        aper.Enumerated = 3
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_5        aper.Enumerated = 4
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_6        aper.Enumerated = 5
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_7        aper.Enumerated = 6
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_8        aper.Enumerated = 7
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_1        aper.Enumerated = 8
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_2        aper.Enumerated = 9
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_3        aper.Enumerated = 10
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_4        aper.Enumerated = 11
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_5        aper.Enumerated = 12
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_6        aper.Enumerated = 13
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_7        aper.Enumerated = 14
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_8        aper.Enumerated = 15
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_9        aper.Enumerated = 16
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_10       aper.Enumerated = 17
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_11       aper.Enumerated = 18
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_12       aper.Enumerated = 19
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_13       aper.Enumerated = 20
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_14       aper.Enumerated = 21
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_15       aper.Enumerated = 22
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_16       aper.Enumerated = 23
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_17       aper.Enumerated = 24
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_18       aper.Enumerated = 25
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_19       aper.Enumerated = 26
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_20       aper.Enumerated = 27
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_21       aper.Enumerated = 28
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_22       aper.Enumerated = 29
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_23       aper.Enumerated = 30
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType3_1        aper.Enumerated = 31
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType4_1        aper.Enumerated = 32
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType5_1        aper.Enumerated = 33
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_1        aper.Enumerated = 34
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_2        aper.Enumerated = 35
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_3        aper.Enumerated = 36
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_9_v1710  aper.Enumerated = 37
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_10_v1710 aper.Enumerated = 38
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_24_v1710 aper.Enumerated = 39
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_25_v1710 aper.Enumerated = 40
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_4_v1710  aper.Enumerated = 41
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_5_v1710  aper.Enumerated = 42
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_6_v1710  aper.Enumerated = 43
)

type PosSIB_ReqInfo_r16_posSibType_r16 struct {
	Value aper.Enumerated `lb:0,ub:36,madatory`
}

func (ie *PosSIB_ReqInfo_r16_posSibType_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 36}, false); err != nil {
		return utils.WrapError("Encode PosSIB_ReqInfo_r16_posSibType_r16", err)
	}
	return nil
}

func (ie *PosSIB_ReqInfo_r16_posSibType_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 36}, false); err != nil {
		return utils.WrapError("Decode PosSIB_ReqInfo_r16_posSibType_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
