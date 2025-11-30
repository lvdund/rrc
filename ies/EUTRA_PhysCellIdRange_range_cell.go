package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_PhysCellIdRange_range_cell_Enum_n4     aper.Enumerated = 0
	EUTRA_PhysCellIdRange_range_cell_Enum_n8     aper.Enumerated = 1
	EUTRA_PhysCellIdRange_range_cell_Enum_n12    aper.Enumerated = 2
	EUTRA_PhysCellIdRange_range_cell_Enum_n16    aper.Enumerated = 3
	EUTRA_PhysCellIdRange_range_cell_Enum_n24    aper.Enumerated = 4
	EUTRA_PhysCellIdRange_range_cell_Enum_n32    aper.Enumerated = 5
	EUTRA_PhysCellIdRange_range_cell_Enum_n48    aper.Enumerated = 6
	EUTRA_PhysCellIdRange_range_cell_Enum_n64    aper.Enumerated = 7
	EUTRA_PhysCellIdRange_range_cell_Enum_n84    aper.Enumerated = 8
	EUTRA_PhysCellIdRange_range_cell_Enum_n96    aper.Enumerated = 9
	EUTRA_PhysCellIdRange_range_cell_Enum_n128   aper.Enumerated = 10
	EUTRA_PhysCellIdRange_range_cell_Enum_n168   aper.Enumerated = 11
	EUTRA_PhysCellIdRange_range_cell_Enum_n252   aper.Enumerated = 12
	EUTRA_PhysCellIdRange_range_cell_Enum_n504   aper.Enumerated = 13
	EUTRA_PhysCellIdRange_range_cell_Enum_spare2 aper.Enumerated = 14
	EUTRA_PhysCellIdRange_range_cell_Enum_spare1 aper.Enumerated = 15
)

type EUTRA_PhysCellIdRange_range_cell struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *EUTRA_PhysCellIdRange_range_cell) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode EUTRA_PhysCellIdRange_range_cell", err)
	}
	return nil
}

func (ie *EUTRA_PhysCellIdRange_range_cell) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode EUTRA_PhysCellIdRange_range_cell", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
