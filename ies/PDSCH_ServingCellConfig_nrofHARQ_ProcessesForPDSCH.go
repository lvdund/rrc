package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n2  aper.Enumerated = 0
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n4  aper.Enumerated = 1
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n6  aper.Enumerated = 2
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n10 aper.Enumerated = 3
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n12 aper.Enumerated = 4
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n16 aper.Enumerated = 5
)

type PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH", err)
	}
	return nil
}

func (ie *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
