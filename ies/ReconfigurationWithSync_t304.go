package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReconfigurationWithSync_t304_Enum_ms50    aper.Enumerated = 0
	ReconfigurationWithSync_t304_Enum_ms100   aper.Enumerated = 1
	ReconfigurationWithSync_t304_Enum_ms150   aper.Enumerated = 2
	ReconfigurationWithSync_t304_Enum_ms200   aper.Enumerated = 3
	ReconfigurationWithSync_t304_Enum_ms500   aper.Enumerated = 4
	ReconfigurationWithSync_t304_Enum_ms1000  aper.Enumerated = 5
	ReconfigurationWithSync_t304_Enum_ms2000  aper.Enumerated = 6
	ReconfigurationWithSync_t304_Enum_ms10000 aper.Enumerated = 7
)

type ReconfigurationWithSync_t304 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *ReconfigurationWithSync_t304) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode ReconfigurationWithSync_t304", err)
	}
	return nil
}

func (ie *ReconfigurationWithSync_t304) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode ReconfigurationWithSync_t304", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
