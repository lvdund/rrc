package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_n2     aper.Enumerated = 0
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_n4     aper.Enumerated = 1
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_n6     aper.Enumerated = 2
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_n12    aper.Enumerated = 3
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_spare4 aper.Enumerated = 4
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_spare3 aper.Enumerated = 5
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_spare2 aper.Enumerated = 6
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_spare1 aper.Enumerated = 7
)

type NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17", err)
	}
	return nil
}

func (ie *NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
