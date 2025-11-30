package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n1 aper.Enumerated = 0
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n2 aper.Enumerated = 1
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n3 aper.Enumerated = 2
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n4 aper.Enumerated = 3
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n7 aper.Enumerated = 4
)

type MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16", err)
	}
	return nil
}

func (ie *MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
