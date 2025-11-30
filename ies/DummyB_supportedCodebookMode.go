package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyB_supportedCodebookMode_Enum_mode1         aper.Enumerated = 0
	DummyB_supportedCodebookMode_Enum_mode1AndMode2 aper.Enumerated = 1
)

type DummyB_supportedCodebookMode struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DummyB_supportedCodebookMode) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DummyB_supportedCodebookMode", err)
	}
	return nil
}

func (ie *DummyB_supportedCodebookMode) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DummyB_supportedCodebookMode", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
