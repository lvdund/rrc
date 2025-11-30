package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Config_drb_discardTimer_Enum_ms10     aper.Enumerated = 0
	PDCP_Config_drb_discardTimer_Enum_ms20     aper.Enumerated = 1
	PDCP_Config_drb_discardTimer_Enum_ms30     aper.Enumerated = 2
	PDCP_Config_drb_discardTimer_Enum_ms40     aper.Enumerated = 3
	PDCP_Config_drb_discardTimer_Enum_ms50     aper.Enumerated = 4
	PDCP_Config_drb_discardTimer_Enum_ms60     aper.Enumerated = 5
	PDCP_Config_drb_discardTimer_Enum_ms75     aper.Enumerated = 6
	PDCP_Config_drb_discardTimer_Enum_ms100    aper.Enumerated = 7
	PDCP_Config_drb_discardTimer_Enum_ms150    aper.Enumerated = 8
	PDCP_Config_drb_discardTimer_Enum_ms200    aper.Enumerated = 9
	PDCP_Config_drb_discardTimer_Enum_ms250    aper.Enumerated = 10
	PDCP_Config_drb_discardTimer_Enum_ms300    aper.Enumerated = 11
	PDCP_Config_drb_discardTimer_Enum_ms500    aper.Enumerated = 12
	PDCP_Config_drb_discardTimer_Enum_ms750    aper.Enumerated = 13
	PDCP_Config_drb_discardTimer_Enum_ms1500   aper.Enumerated = 14
	PDCP_Config_drb_discardTimer_Enum_infinity aper.Enumerated = 15
)

type PDCP_Config_drb_discardTimer struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PDCP_Config_drb_discardTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PDCP_Config_drb_discardTimer", err)
	}
	return nil
}

func (ie *PDCP_Config_drb_discardTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PDCP_Config_drb_discardTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
