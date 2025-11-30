package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSetupComplete_IEs_guami_Type_Enum_native aper.Enumerated = 0
	RRCSetupComplete_IEs_guami_Type_Enum_mapped aper.Enumerated = 1
)

type RRCSetupComplete_IEs_guami_Type struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *RRCSetupComplete_IEs_guami_Type) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode RRCSetupComplete_IEs_guami_Type", err)
	}
	return nil
}

func (ie *RRCSetupComplete_IEs_guami_Type) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode RRCSetupComplete_IEs_guami_Type", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
