package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_Q_OffsetRange_Enum_dB_24 aper.Enumerated = 0
	EUTRA_Q_OffsetRange_Enum_dB_22 aper.Enumerated = 1
	EUTRA_Q_OffsetRange_Enum_dB_20 aper.Enumerated = 2
	EUTRA_Q_OffsetRange_Enum_dB_18 aper.Enumerated = 3
	EUTRA_Q_OffsetRange_Enum_dB_16 aper.Enumerated = 4
	EUTRA_Q_OffsetRange_Enum_dB_14 aper.Enumerated = 5
	EUTRA_Q_OffsetRange_Enum_dB_12 aper.Enumerated = 6
	EUTRA_Q_OffsetRange_Enum_dB_10 aper.Enumerated = 7
	EUTRA_Q_OffsetRange_Enum_dB_8  aper.Enumerated = 8
	EUTRA_Q_OffsetRange_Enum_dB_6  aper.Enumerated = 9
	EUTRA_Q_OffsetRange_Enum_dB_5  aper.Enumerated = 10
	EUTRA_Q_OffsetRange_Enum_dB_4  aper.Enumerated = 11
	EUTRA_Q_OffsetRange_Enum_dB_3  aper.Enumerated = 12
	EUTRA_Q_OffsetRange_Enum_dB_2  aper.Enumerated = 13
	EUTRA_Q_OffsetRange_Enum_dB_1  aper.Enumerated = 14
	EUTRA_Q_OffsetRange_Enum_dB0   aper.Enumerated = 15
	EUTRA_Q_OffsetRange_Enum_dB1   aper.Enumerated = 16
	EUTRA_Q_OffsetRange_Enum_dB2   aper.Enumerated = 17
	EUTRA_Q_OffsetRange_Enum_dB3   aper.Enumerated = 18
	EUTRA_Q_OffsetRange_Enum_dB4   aper.Enumerated = 19
	EUTRA_Q_OffsetRange_Enum_dB5   aper.Enumerated = 20
	EUTRA_Q_OffsetRange_Enum_dB6   aper.Enumerated = 21
	EUTRA_Q_OffsetRange_Enum_dB8   aper.Enumerated = 22
	EUTRA_Q_OffsetRange_Enum_dB10  aper.Enumerated = 23
	EUTRA_Q_OffsetRange_Enum_dB12  aper.Enumerated = 24
	EUTRA_Q_OffsetRange_Enum_dB14  aper.Enumerated = 25
	EUTRA_Q_OffsetRange_Enum_dB16  aper.Enumerated = 26
	EUTRA_Q_OffsetRange_Enum_dB18  aper.Enumerated = 27
	EUTRA_Q_OffsetRange_Enum_dB20  aper.Enumerated = 28
	EUTRA_Q_OffsetRange_Enum_dB22  aper.Enumerated = 29
	EUTRA_Q_OffsetRange_Enum_dB24  aper.Enumerated = 30
)

type EUTRA_Q_OffsetRange struct {
	Value aper.Enumerated `lb:0,ub:30,madatory`
}

func (ie *EUTRA_Q_OffsetRange) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Encode EUTRA_Q_OffsetRange", err)
	}
	return nil
}

func (ie *EUTRA_Q_OffsetRange) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Decode EUTRA_Q_OffsetRange", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
