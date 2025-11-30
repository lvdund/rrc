package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n1  aper.Enumerated = 0
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n2  aper.Enumerated = 1
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n3  aper.Enumerated = 2
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n4  aper.Enumerated = 3
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n5  aper.Enumerated = 4
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n6  aper.Enumerated = 5
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n8  aper.Enumerated = 6
	RadioLinkMonitoringConfig_beamFailureInstanceMaxCount_Enum_n10 aper.Enumerated = 7
)

type RadioLinkMonitoringConfig_beamFailureInstanceMaxCount struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RadioLinkMonitoringConfig_beamFailureInstanceMaxCount) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RadioLinkMonitoringConfig_beamFailureInstanceMaxCount", err)
	}
	return nil
}

func (ie *RadioLinkMonitoringConfig_beamFailureInstanceMaxCount) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RadioLinkMonitoringConfig_beamFailureInstanceMaxCount", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
