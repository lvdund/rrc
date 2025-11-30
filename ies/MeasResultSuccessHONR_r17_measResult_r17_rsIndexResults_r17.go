package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17 struct {
	ResultsSSB_Indexes_r17    *ResultsPerSSB_IndexList    `optional`
	ResultsCSI_RS_Indexes_r17 *ResultsPerCSI_RS_IndexList `optional`
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ResultsSSB_Indexes_r17 != nil, ie.ResultsCSI_RS_Indexes_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ResultsSSB_Indexes_r17 != nil {
		if err = ie.ResultsSSB_Indexes_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsSSB_Indexes_r17", err)
		}
	}
	if ie.ResultsCSI_RS_Indexes_r17 != nil {
		if err = ie.ResultsCSI_RS_Indexes_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsCSI_RS_Indexes_r17", err)
		}
	}
	return nil
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17) Decode(r *aper.AperReader) error {
	var err error
	var ResultsSSB_Indexes_r17Present bool
	if ResultsSSB_Indexes_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ResultsCSI_RS_Indexes_r17Present bool
	if ResultsCSI_RS_Indexes_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ResultsSSB_Indexes_r17Present {
		ie.ResultsSSB_Indexes_r17 = new(ResultsPerSSB_IndexList)
		if err = ie.ResultsSSB_Indexes_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsSSB_Indexes_r17", err)
		}
	}
	if ResultsCSI_RS_Indexes_r17Present {
		ie.ResultsCSI_RS_Indexes_r17 = new(ResultsPerCSI_RS_IndexList)
		if err = ie.ResultsCSI_RS_Indexes_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsCSI_RS_Indexes_r17", err)
		}
	}
	return nil
}
