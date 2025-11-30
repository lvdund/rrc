package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RA_Report_r16_raPurpose_r16_Enum_accessRelated             aper.Enumerated = 0
	RA_Report_r16_raPurpose_r16_Enum_beamFailureRecovery       aper.Enumerated = 1
	RA_Report_r16_raPurpose_r16_Enum_reconfigurationWithSync   aper.Enumerated = 2
	RA_Report_r16_raPurpose_r16_Enum_ulUnSynchronized          aper.Enumerated = 3
	RA_Report_r16_raPurpose_r16_Enum_schedulingRequestFailure  aper.Enumerated = 4
	RA_Report_r16_raPurpose_r16_Enum_noPUCCHResourceAvailable  aper.Enumerated = 5
	RA_Report_r16_raPurpose_r16_Enum_requestForOtherSI         aper.Enumerated = 6
	RA_Report_r16_raPurpose_r16_Enum_msg3RequestForOtherSI_r17 aper.Enumerated = 7
	RA_Report_r16_raPurpose_r16_Enum_spare8                    aper.Enumerated = 8
	RA_Report_r16_raPurpose_r16_Enum_spare7                    aper.Enumerated = 9
	RA_Report_r16_raPurpose_r16_Enum_spare6                    aper.Enumerated = 10
	RA_Report_r16_raPurpose_r16_Enum_spare5                    aper.Enumerated = 11
	RA_Report_r16_raPurpose_r16_Enum_spare4                    aper.Enumerated = 12
	RA_Report_r16_raPurpose_r16_Enum_spare3                    aper.Enumerated = 13
	RA_Report_r16_raPurpose_r16_Enum_spare2                    aper.Enumerated = 14
	RA_Report_r16_raPurpose_r16_Enum_spare1                    aper.Enumerated = 15
)

type RA_Report_r16_raPurpose_r16 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RA_Report_r16_raPurpose_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RA_Report_r16_raPurpose_r16", err)
	}
	return nil
}

func (ie *RA_Report_r16_raPurpose_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RA_Report_r16_raPurpose_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
