package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16 struct {
	Rsrp_ResultEUTRA_r16 *RSRP_RangeEUTRA     `optional`
	Rsrq_ResultEUTRA_r16 *RSRQ_RangeEUTRA_r16 `optional`
}

func (ie *MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Rsrp_ResultEUTRA_r16 != nil, ie.Rsrq_ResultEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rsrp_ResultEUTRA_r16 != nil {
		if err = ie.Rsrp_ResultEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrp_ResultEUTRA_r16", err)
		}
	}
	if ie.Rsrq_ResultEUTRA_r16 != nil {
		if err = ie.Rsrq_ResultEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrq_ResultEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var Rsrp_ResultEUTRA_r16Present bool
	if Rsrp_ResultEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rsrq_ResultEUTRA_r16Present bool
	if Rsrq_ResultEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rsrp_ResultEUTRA_r16Present {
		ie.Rsrp_ResultEUTRA_r16 = new(RSRP_RangeEUTRA)
		if err = ie.Rsrp_ResultEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp_ResultEUTRA_r16", err)
		}
	}
	if Rsrq_ResultEUTRA_r16Present {
		ie.Rsrq_ResultEUTRA_r16 = new(RSRQ_RangeEUTRA_r16)
		if err = ie.Rsrq_ResultEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrq_ResultEUTRA_r16", err)
		}
	}
	return nil
}
