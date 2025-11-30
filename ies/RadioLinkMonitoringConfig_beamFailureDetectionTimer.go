package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd1  aper.Enumerated = 0
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd2  aper.Enumerated = 1
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd3  aper.Enumerated = 2
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd4  aper.Enumerated = 3
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd5  aper.Enumerated = 4
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd6  aper.Enumerated = 5
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd8  aper.Enumerated = 6
	RadioLinkMonitoringConfig_beamFailureDetectionTimer_Enum_pbfd10 aper.Enumerated = 7
)

type RadioLinkMonitoringConfig_beamFailureDetectionTimer struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RadioLinkMonitoringConfig_beamFailureDetectionTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RadioLinkMonitoringConfig_beamFailureDetectionTimer", err)
	}
	return nil
}

func (ie *RadioLinkMonitoringConfig_beamFailureDetectionTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RadioLinkMonitoringConfig_beamFailureDetectionTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
