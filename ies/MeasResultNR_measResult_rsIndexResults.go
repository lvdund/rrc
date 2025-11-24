package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_measResult_rsIndexResults struct {
	ResultsSSB_Indexes    *ResultsPerSSB_IndexList    `optional`
	ResultsCSI_RS_Indexes *ResultsPerCSI_RS_IndexList `optional`
}

func (ie *MeasResultNR_measResult_rsIndexResults) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ResultsSSB_Indexes != nil, ie.ResultsCSI_RS_Indexes != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ResultsSSB_Indexes != nil {
		if err = ie.ResultsSSB_Indexes.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsSSB_Indexes", err)
		}
	}
	if ie.ResultsCSI_RS_Indexes != nil {
		if err = ie.ResultsCSI_RS_Indexes.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsCSI_RS_Indexes", err)
		}
	}
	return nil
}

func (ie *MeasResultNR_measResult_rsIndexResults) Decode(r *uper.UperReader) error {
	var err error
	var ResultsSSB_IndexesPresent bool
	if ResultsSSB_IndexesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ResultsCSI_RS_IndexesPresent bool
	if ResultsCSI_RS_IndexesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ResultsSSB_IndexesPresent {
		ie.ResultsSSB_Indexes = new(ResultsPerSSB_IndexList)
		if err = ie.ResultsSSB_Indexes.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsSSB_Indexes", err)
		}
	}
	if ResultsCSI_RS_IndexesPresent {
		ie.ResultsCSI_RS_Indexes = new(ResultsPerCSI_RS_IndexList)
		if err = ie.ResultsCSI_RS_Indexes.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsCSI_RS_Indexes", err)
		}
	}
	return nil
}
