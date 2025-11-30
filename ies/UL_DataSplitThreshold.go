package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_DataSplitThreshold_Enum_b0       aper.Enumerated = 0
	UL_DataSplitThreshold_Enum_b100     aper.Enumerated = 1
	UL_DataSplitThreshold_Enum_b200     aper.Enumerated = 2
	UL_DataSplitThreshold_Enum_b400     aper.Enumerated = 3
	UL_DataSplitThreshold_Enum_b800     aper.Enumerated = 4
	UL_DataSplitThreshold_Enum_b1600    aper.Enumerated = 5
	UL_DataSplitThreshold_Enum_b3200    aper.Enumerated = 6
	UL_DataSplitThreshold_Enum_b6400    aper.Enumerated = 7
	UL_DataSplitThreshold_Enum_b12800   aper.Enumerated = 8
	UL_DataSplitThreshold_Enum_b25600   aper.Enumerated = 9
	UL_DataSplitThreshold_Enum_b51200   aper.Enumerated = 10
	UL_DataSplitThreshold_Enum_b102400  aper.Enumerated = 11
	UL_DataSplitThreshold_Enum_b204800  aper.Enumerated = 12
	UL_DataSplitThreshold_Enum_b409600  aper.Enumerated = 13
	UL_DataSplitThreshold_Enum_b819200  aper.Enumerated = 14
	UL_DataSplitThreshold_Enum_b1228800 aper.Enumerated = 15
	UL_DataSplitThreshold_Enum_b1638400 aper.Enumerated = 16
	UL_DataSplitThreshold_Enum_b2457600 aper.Enumerated = 17
	UL_DataSplitThreshold_Enum_b3276800 aper.Enumerated = 18
	UL_DataSplitThreshold_Enum_b4096000 aper.Enumerated = 19
	UL_DataSplitThreshold_Enum_b4915200 aper.Enumerated = 20
	UL_DataSplitThreshold_Enum_b5734400 aper.Enumerated = 21
	UL_DataSplitThreshold_Enum_b6553600 aper.Enumerated = 22
	UL_DataSplitThreshold_Enum_infinity aper.Enumerated = 23
	UL_DataSplitThreshold_Enum_spare8   aper.Enumerated = 24
	UL_DataSplitThreshold_Enum_spare7   aper.Enumerated = 25
	UL_DataSplitThreshold_Enum_spare6   aper.Enumerated = 26
	UL_DataSplitThreshold_Enum_spare5   aper.Enumerated = 27
	UL_DataSplitThreshold_Enum_spare4   aper.Enumerated = 28
	UL_DataSplitThreshold_Enum_spare3   aper.Enumerated = 29
	UL_DataSplitThreshold_Enum_spare2   aper.Enumerated = 30
	UL_DataSplitThreshold_Enum_spare1   aper.Enumerated = 31
)

type UL_DataSplitThreshold struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *UL_DataSplitThreshold) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode UL_DataSplitThreshold", err)
	}
	return nil
}

func (ie *UL_DataSplitThreshold) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode UL_DataSplitThreshold", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
