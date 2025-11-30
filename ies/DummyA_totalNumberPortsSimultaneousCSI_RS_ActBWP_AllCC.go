package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p8   aper.Enumerated = 0
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p12  aper.Enumerated = 1
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p16  aper.Enumerated = 2
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p24  aper.Enumerated = 3
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p32  aper.Enumerated = 4
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p40  aper.Enumerated = 5
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p48  aper.Enumerated = 6
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p56  aper.Enumerated = 7
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p64  aper.Enumerated = 8
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p72  aper.Enumerated = 9
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p80  aper.Enumerated = 10
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p88  aper.Enumerated = 11
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p96  aper.Enumerated = 12
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p104 aper.Enumerated = 13
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p112 aper.Enumerated = 14
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p120 aper.Enumerated = 15
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p128 aper.Enumerated = 16
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p136 aper.Enumerated = 17
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p144 aper.Enumerated = 18
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p152 aper.Enumerated = 19
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p160 aper.Enumerated = 20
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p168 aper.Enumerated = 21
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p176 aper.Enumerated = 22
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p184 aper.Enumerated = 23
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p192 aper.Enumerated = 24
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p200 aper.Enumerated = 25
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p208 aper.Enumerated = 26
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p216 aper.Enumerated = 27
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p224 aper.Enumerated = 28
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p232 aper.Enumerated = 29
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p240 aper.Enumerated = 30
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p248 aper.Enumerated = 31
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p256 aper.Enumerated = 32
)

type DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC struct {
	Value aper.Enumerated `lb:0,ub:32,madatory`
}

func (ie *DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
		return utils.WrapError("Encode DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	return nil
}

func (ie *DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
		return utils.WrapError("Decode DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
