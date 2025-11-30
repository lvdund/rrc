package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PollPDU_Enum_p4       aper.Enumerated = 0
	PollPDU_Enum_p8       aper.Enumerated = 1
	PollPDU_Enum_p16      aper.Enumerated = 2
	PollPDU_Enum_p32      aper.Enumerated = 3
	PollPDU_Enum_p64      aper.Enumerated = 4
	PollPDU_Enum_p128     aper.Enumerated = 5
	PollPDU_Enum_p256     aper.Enumerated = 6
	PollPDU_Enum_p512     aper.Enumerated = 7
	PollPDU_Enum_p1024    aper.Enumerated = 8
	PollPDU_Enum_p2048    aper.Enumerated = 9
	PollPDU_Enum_p4096    aper.Enumerated = 10
	PollPDU_Enum_p6144    aper.Enumerated = 11
	PollPDU_Enum_p8192    aper.Enumerated = 12
	PollPDU_Enum_p12288   aper.Enumerated = 13
	PollPDU_Enum_p16384   aper.Enumerated = 14
	PollPDU_Enum_p20480   aper.Enumerated = 15
	PollPDU_Enum_p24576   aper.Enumerated = 16
	PollPDU_Enum_p28672   aper.Enumerated = 17
	PollPDU_Enum_p32768   aper.Enumerated = 18
	PollPDU_Enum_p40960   aper.Enumerated = 19
	PollPDU_Enum_p49152   aper.Enumerated = 20
	PollPDU_Enum_p57344   aper.Enumerated = 21
	PollPDU_Enum_p65536   aper.Enumerated = 22
	PollPDU_Enum_infinity aper.Enumerated = 23
	PollPDU_Enum_spare8   aper.Enumerated = 24
	PollPDU_Enum_spare7   aper.Enumerated = 25
	PollPDU_Enum_spare6   aper.Enumerated = 26
	PollPDU_Enum_spare5   aper.Enumerated = 27
	PollPDU_Enum_spare4   aper.Enumerated = 28
	PollPDU_Enum_spare3   aper.Enumerated = 29
	PollPDU_Enum_spare2   aper.Enumerated = 30
	PollPDU_Enum_spare1   aper.Enumerated = 31
)

type PollPDU struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *PollPDU) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode PollPDU", err)
	}
	return nil
}

func (ie *PollPDU) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode PollPDU", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
