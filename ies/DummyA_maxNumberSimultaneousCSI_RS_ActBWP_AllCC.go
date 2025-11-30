package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n5  aper.Enumerated = 0
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n6  aper.Enumerated = 1
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n7  aper.Enumerated = 2
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n8  aper.Enumerated = 3
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n9  aper.Enumerated = 4
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n10 aper.Enumerated = 5
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n12 aper.Enumerated = 6
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n14 aper.Enumerated = 7
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n16 aper.Enumerated = 8
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n18 aper.Enumerated = 9
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n20 aper.Enumerated = 10
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n22 aper.Enumerated = 11
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n24 aper.Enumerated = 12
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n26 aper.Enumerated = 13
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n28 aper.Enumerated = 14
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n30 aper.Enumerated = 15
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n32 aper.Enumerated = 16
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n34 aper.Enumerated = 17
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n36 aper.Enumerated = 18
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n38 aper.Enumerated = 19
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n40 aper.Enumerated = 20
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n42 aper.Enumerated = 21
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n44 aper.Enumerated = 22
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n46 aper.Enumerated = 23
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n48 aper.Enumerated = 24
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n50 aper.Enumerated = 25
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n52 aper.Enumerated = 26
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n54 aper.Enumerated = 27
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n56 aper.Enumerated = 28
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n58 aper.Enumerated = 29
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n60 aper.Enumerated = 30
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n62 aper.Enumerated = 31
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n64 aper.Enumerated = 32
)

type DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC struct {
	Value aper.Enumerated `lb:0,ub:32,madatory`
}

func (ie *DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
		return utils.WrapError("Encode DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	return nil
}

func (ie *DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
		return utils.WrapError("Decode DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
