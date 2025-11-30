package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PollByte_Enum_kB1      aper.Enumerated = 0
	PollByte_Enum_kB2      aper.Enumerated = 1
	PollByte_Enum_kB5      aper.Enumerated = 2
	PollByte_Enum_kB8      aper.Enumerated = 3
	PollByte_Enum_kB10     aper.Enumerated = 4
	PollByte_Enum_kB15     aper.Enumerated = 5
	PollByte_Enum_kB25     aper.Enumerated = 6
	PollByte_Enum_kB50     aper.Enumerated = 7
	PollByte_Enum_kB75     aper.Enumerated = 8
	PollByte_Enum_kB100    aper.Enumerated = 9
	PollByte_Enum_kB125    aper.Enumerated = 10
	PollByte_Enum_kB250    aper.Enumerated = 11
	PollByte_Enum_kB375    aper.Enumerated = 12
	PollByte_Enum_kB500    aper.Enumerated = 13
	PollByte_Enum_kB750    aper.Enumerated = 14
	PollByte_Enum_kB1000   aper.Enumerated = 15
	PollByte_Enum_kB1250   aper.Enumerated = 16
	PollByte_Enum_kB1500   aper.Enumerated = 17
	PollByte_Enum_kB2000   aper.Enumerated = 18
	PollByte_Enum_kB3000   aper.Enumerated = 19
	PollByte_Enum_kB4000   aper.Enumerated = 20
	PollByte_Enum_kB4500   aper.Enumerated = 21
	PollByte_Enum_kB5000   aper.Enumerated = 22
	PollByte_Enum_kB5500   aper.Enumerated = 23
	PollByte_Enum_kB6000   aper.Enumerated = 24
	PollByte_Enum_kB6500   aper.Enumerated = 25
	PollByte_Enum_kB7000   aper.Enumerated = 26
	PollByte_Enum_kB7500   aper.Enumerated = 27
	PollByte_Enum_mB8      aper.Enumerated = 28
	PollByte_Enum_mB9      aper.Enumerated = 29
	PollByte_Enum_mB10     aper.Enumerated = 30
	PollByte_Enum_mB11     aper.Enumerated = 31
	PollByte_Enum_mB12     aper.Enumerated = 32
	PollByte_Enum_mB13     aper.Enumerated = 33
	PollByte_Enum_mB14     aper.Enumerated = 34
	PollByte_Enum_mB15     aper.Enumerated = 35
	PollByte_Enum_mB16     aper.Enumerated = 36
	PollByte_Enum_mB17     aper.Enumerated = 37
	PollByte_Enum_mB18     aper.Enumerated = 38
	PollByte_Enum_mB20     aper.Enumerated = 39
	PollByte_Enum_mB25     aper.Enumerated = 40
	PollByte_Enum_mB30     aper.Enumerated = 41
	PollByte_Enum_mB40     aper.Enumerated = 42
	PollByte_Enum_infinity aper.Enumerated = 43
	PollByte_Enum_spare20  aper.Enumerated = 44
	PollByte_Enum_spare19  aper.Enumerated = 45
	PollByte_Enum_spare18  aper.Enumerated = 46
	PollByte_Enum_spare17  aper.Enumerated = 47
	PollByte_Enum_spare16  aper.Enumerated = 48
	PollByte_Enum_spare15  aper.Enumerated = 49
	PollByte_Enum_spare14  aper.Enumerated = 50
	PollByte_Enum_spare13  aper.Enumerated = 51
	PollByte_Enum_spare12  aper.Enumerated = 52
	PollByte_Enum_spare11  aper.Enumerated = 53
	PollByte_Enum_spare10  aper.Enumerated = 54
	PollByte_Enum_spare9   aper.Enumerated = 55
	PollByte_Enum_spare8   aper.Enumerated = 56
	PollByte_Enum_spare7   aper.Enumerated = 57
	PollByte_Enum_spare6   aper.Enumerated = 58
	PollByte_Enum_spare5   aper.Enumerated = 59
	PollByte_Enum_spare4   aper.Enumerated = 60
	PollByte_Enum_spare3   aper.Enumerated = 61
	PollByte_Enum_spare2   aper.Enumerated = 62
	PollByte_Enum_spare1   aper.Enumerated = 63
)

type PollByte struct {
	Value aper.Enumerated `lb:0,ub:63,madatory`
}

func (ie *PollByte) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Encode PollByte", err)
	}
	return nil
}

func (ie *PollByte) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Decode PollByte", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
