package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BFD_RelaxationReportingConfig_r17 struct {
	Bfd_RelaxtionReportingProhibitTimer BFD_RelaxationReportingConfig_r17_bfd_RelaxtionReportingProhibitTimer `madatory`
}

func (ie *BFD_RelaxationReportingConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Bfd_RelaxtionReportingProhibitTimer.Encode(w); err != nil {
		return utils.WrapError("Encode Bfd_RelaxtionReportingProhibitTimer", err)
	}
	return nil
}

func (ie *BFD_RelaxationReportingConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Bfd_RelaxtionReportingProhibitTimer.Decode(r); err != nil {
		return utils.WrapError("Decode Bfd_RelaxtionReportingProhibitTimer", err)
	}
	return nil
}
