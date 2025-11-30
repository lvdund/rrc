package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700_Enum_n32 aper.Enumerated = 0
)

type PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700", err)
	}
	return nil
}

func (ie *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
