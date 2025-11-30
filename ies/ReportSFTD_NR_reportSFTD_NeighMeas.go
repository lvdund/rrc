package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportSFTD_NR_reportSFTD_NeighMeas_Enum_true aper.Enumerated = 0
)

type ReportSFTD_NR_reportSFTD_NeighMeas struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *ReportSFTD_NR_reportSFTD_NeighMeas) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode ReportSFTD_NR_reportSFTD_NeighMeas", err)
	}
	return nil
}

func (ie *ReportSFTD_NR_reportSFTD_NeighMeas) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode ReportSFTD_NR_reportSFTD_NeighMeas", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
