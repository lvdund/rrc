package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerCSI_RS_Index struct {
	Csi_RS_Index   CSI_RS_Index         `madatory`
	Csi_RS_Results *MeasQuantityResults `optional`
}

func (ie *ResultsPerCSI_RS_Index) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Csi_RS_Results != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Csi_RS_Index.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_RS_Index", err)
	}
	if ie.Csi_RS_Results != nil {
		if err = ie.Csi_RS_Results.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RS_Results", err)
		}
	}
	return nil
}

func (ie *ResultsPerCSI_RS_Index) Decode(r *uper.UperReader) error {
	var err error
	var Csi_RS_ResultsPresent bool
	if Csi_RS_ResultsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Csi_RS_Index.Decode(r); err != nil {
		return utils.WrapError("Decode Csi_RS_Index", err)
	}
	if Csi_RS_ResultsPresent {
		ie.Csi_RS_Results = new(MeasQuantityResults)
		if err = ie.Csi_RS_Results.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_Results", err)
		}
	}
	return nil
}
