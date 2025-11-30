package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_cqi_Table_Enum_table1     aper.Enumerated = 0
	CSI_ReportConfig_cqi_Table_Enum_table2     aper.Enumerated = 1
	CSI_ReportConfig_cqi_Table_Enum_table3     aper.Enumerated = 2
	CSI_ReportConfig_cqi_Table_Enum_table4_r17 aper.Enumerated = 3
)

type CSI_ReportConfig_cqi_Table struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CSI_ReportConfig_cqi_Table) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportConfig_cqi_Table", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_cqi_Table) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportConfig_cqi_Table", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
