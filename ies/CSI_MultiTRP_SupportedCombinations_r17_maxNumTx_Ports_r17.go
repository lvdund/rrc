package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n2  aper.Enumerated = 0
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n4  aper.Enumerated = 1
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n8  aper.Enumerated = 2
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n12 aper.Enumerated = 3
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n16 aper.Enumerated = 4
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n24 aper.Enumerated = 5
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n32 aper.Enumerated = 6
)

type CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17", err)
	}
	return nil
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
