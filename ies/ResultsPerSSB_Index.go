package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerSSB_Index struct {
	Ssb_Index   SSB_Index            `madatory`
	Ssb_Results *MeasQuantityResults `optional`
}

func (ie *ResultsPerSSB_Index) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_Results != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Ssb_Index.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_Index", err)
	}
	if ie.Ssb_Results != nil {
		if err = ie.Ssb_Results.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_Results", err)
		}
	}
	return nil
}

func (ie *ResultsPerSSB_Index) Decode(r *uper.UperReader) error {
	var err error
	var Ssb_ResultsPresent bool
	if Ssb_ResultsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Ssb_Index.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_Index", err)
	}
	if Ssb_ResultsPresent {
		ie.Ssb_Results = new(MeasQuantityResults)
		if err = ie.Ssb_Results.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Results", err)
		}
	}
	return nil
}
