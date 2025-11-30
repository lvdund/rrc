package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PEI_Config_r17_po_NumPerPEI_r17_Enum_po1 aper.Enumerated = 0
	PEI_Config_r17_po_NumPerPEI_r17_Enum_po2 aper.Enumerated = 1
	PEI_Config_r17_po_NumPerPEI_r17_Enum_po4 aper.Enumerated = 2
	PEI_Config_r17_po_NumPerPEI_r17_Enum_po8 aper.Enumerated = 3
)

type PEI_Config_r17_po_NumPerPEI_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PEI_Config_r17_po_NumPerPEI_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PEI_Config_r17_po_NumPerPEI_r17", err)
	}
	return nil
}

func (ie *PEI_Config_r17_po_NumPerPEI_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PEI_Config_r17_po_NumPerPEI_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
