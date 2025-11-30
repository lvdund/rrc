package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyC_supportedCodebookMode_Enum_mode1 aper.Enumerated = 0
	DummyC_supportedCodebookMode_Enum_mode2 aper.Enumerated = 1
	DummyC_supportedCodebookMode_Enum_both  aper.Enumerated = 2
)

type DummyC_supportedCodebookMode struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *DummyC_supportedCodebookMode) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode DummyC_supportedCodebookMode", err)
	}
	return nil
}

func (ie *DummyC_supportedCodebookMode) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode DummyC_supportedCodebookMode", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
