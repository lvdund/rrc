package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ConnEstFailReport_r16 struct {
	MeasResultFailedCell_r16 MeasResultFailedCell_r16                        `madatory`
	LocationInfo_r16         *LocationInfo_r16                               `optional`
	MeasResultNeighCells_r16 *ConnEstFailReport_r16_measResultNeighCells_r16 `optional`
	NumberOfConnFail_r16     int64                                           `lb:1,ub:8,madatory`
	PerRAInfoList_r16        PerRAInfoList_r16                               `madatory`
	TimeSinceFailure_r16     TimeSinceFailure_r16                            `madatory`
}

func (ie *ConnEstFailReport_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.LocationInfo_r16 != nil, ie.MeasResultNeighCells_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MeasResultFailedCell_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultFailedCell_r16", err)
	}
	if ie.LocationInfo_r16 != nil {
		if err = ie.LocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LocationInfo_r16", err)
		}
	}
	if ie.MeasResultNeighCells_r16 != nil {
		if err = ie.MeasResultNeighCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultNeighCells_r16", err)
		}
	}
	if err = w.WriteInteger(ie.NumberOfConnFail_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfConnFail_r16", err)
	}
	if err = ie.PerRAInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PerRAInfoList_r16", err)
	}
	if err = ie.TimeSinceFailure_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TimeSinceFailure_r16", err)
	}
	return nil
}

func (ie *ConnEstFailReport_r16) Decode(r *aper.AperReader) error {
	var err error
	var LocationInfo_r16Present bool
	if LocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultNeighCells_r16Present bool
	if MeasResultNeighCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MeasResultFailedCell_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultFailedCell_r16", err)
	}
	if LocationInfo_r16Present {
		ie.LocationInfo_r16 = new(LocationInfo_r16)
		if err = ie.LocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LocationInfo_r16", err)
		}
	}
	if MeasResultNeighCells_r16Present {
		ie.MeasResultNeighCells_r16 = new(ConnEstFailReport_r16_measResultNeighCells_r16)
		if err = ie.MeasResultNeighCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultNeighCells_r16", err)
		}
	}
	var tmp_int_NumberOfConnFail_r16 int64
	if tmp_int_NumberOfConnFail_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfConnFail_r16", err)
	}
	ie.NumberOfConnFail_r16 = tmp_int_NumberOfConnFail_r16
	if err = ie.PerRAInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PerRAInfoList_r16", err)
	}
	if err = ie.TimeSinceFailure_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TimeSinceFailure_r16", err)
	}
	return nil
}
