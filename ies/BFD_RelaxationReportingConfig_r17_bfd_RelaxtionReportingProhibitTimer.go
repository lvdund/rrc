package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s0       aper.Enumerated = 0
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s0dot5   aper.Enumerated = 1
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s1       aper.Enumerated = 2
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s2       aper.Enumerated = 3
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s5       aper.Enumerated = 4
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s10      aper.Enumerated = 5
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s20      aper.Enumerated = 6
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s30      aper.Enumerated = 7
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s60      aper.Enumerated = 8
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s90      aper.Enumerated = 9
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s120     aper.Enumerated = 10
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s300     aper.Enumerated = 11
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_s600     aper.Enumerated = 12
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_infinity aper.Enumerated = 13
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_spare2   aper.Enumerated = 14
	BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer_Enum_spare1   aper.Enumerated = 15
)

type BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer", err)
	}
	return nil
}

func (ie *BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
