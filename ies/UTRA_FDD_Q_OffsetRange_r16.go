package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_24 aper.Enumerated = 0
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_22 aper.Enumerated = 1
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_20 aper.Enumerated = 2
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_18 aper.Enumerated = 3
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_16 aper.Enumerated = 4
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_14 aper.Enumerated = 5
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_12 aper.Enumerated = 6
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_10 aper.Enumerated = 7
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_8  aper.Enumerated = 8
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_6  aper.Enumerated = 9
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_5  aper.Enumerated = 10
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_4  aper.Enumerated = 11
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_3  aper.Enumerated = 12
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_2  aper.Enumerated = 13
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB_1  aper.Enumerated = 14
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB0   aper.Enumerated = 15
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB1   aper.Enumerated = 16
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB2   aper.Enumerated = 17
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB3   aper.Enumerated = 18
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB4   aper.Enumerated = 19
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB5   aper.Enumerated = 20
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB6   aper.Enumerated = 21
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB8   aper.Enumerated = 22
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB10  aper.Enumerated = 23
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB12  aper.Enumerated = 24
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB14  aper.Enumerated = 25
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB16  aper.Enumerated = 26
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB18  aper.Enumerated = 27
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB20  aper.Enumerated = 28
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB22  aper.Enumerated = 29
	UTRA_FDD_Q_OffsetRange_r16_Enum_dB24  aper.Enumerated = 30
)

type UTRA_FDD_Q_OffsetRange_r16 struct {
	Value aper.Enumerated `lb:0,ub:30,madatory`
}

func (ie *UTRA_FDD_Q_OffsetRange_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Encode UTRA_FDD_Q_OffsetRange_r16", err)
	}
	return nil
}

func (ie *UTRA_FDD_Q_OffsetRange_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Decode UTRA_FDD_Q_OffsetRange_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
