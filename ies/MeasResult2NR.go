package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2NR struct {
	SsbFrequency              *ARFCN_ValueNR    `optional`
	RefFreqCSI_RS             *ARFCN_ValueNR    `optional`
	MeasResultServingCell     *MeasResultNR     `optional`
	MeasResultNeighCellListNR *MeasResultListNR `optional`
}

func (ie *MeasResult2NR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SsbFrequency != nil, ie.RefFreqCSI_RS != nil, ie.MeasResultServingCell != nil, ie.MeasResultNeighCellListNR != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SsbFrequency != nil {
		if err = ie.SsbFrequency.Encode(w); err != nil {
			return utils.WrapError("Encode SsbFrequency", err)
		}
	}
	if ie.RefFreqCSI_RS != nil {
		if err = ie.RefFreqCSI_RS.Encode(w); err != nil {
			return utils.WrapError("Encode RefFreqCSI_RS", err)
		}
	}
	if ie.MeasResultServingCell != nil {
		if err = ie.MeasResultServingCell.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultServingCell", err)
		}
	}
	if ie.MeasResultNeighCellListNR != nil {
		if err = ie.MeasResultNeighCellListNR.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultNeighCellListNR", err)
		}
	}
	return nil
}

func (ie *MeasResult2NR) Decode(r *uper.UperReader) error {
	var err error
	var SsbFrequencyPresent bool
	if SsbFrequencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RefFreqCSI_RSPresent bool
	if RefFreqCSI_RSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultServingCellPresent bool
	if MeasResultServingCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultNeighCellListNRPresent bool
	if MeasResultNeighCellListNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SsbFrequencyPresent {
		ie.SsbFrequency = new(ARFCN_ValueNR)
		if err = ie.SsbFrequency.Decode(r); err != nil {
			return utils.WrapError("Decode SsbFrequency", err)
		}
	}
	if RefFreqCSI_RSPresent {
		ie.RefFreqCSI_RS = new(ARFCN_ValueNR)
		if err = ie.RefFreqCSI_RS.Decode(r); err != nil {
			return utils.WrapError("Decode RefFreqCSI_RS", err)
		}
	}
	if MeasResultServingCellPresent {
		ie.MeasResultServingCell = new(MeasResultNR)
		if err = ie.MeasResultServingCell.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultServingCell", err)
		}
	}
	if MeasResultNeighCellListNRPresent {
		ie.MeasResultNeighCellListNR = new(MeasResultListNR)
		if err = ie.MeasResultNeighCellListNR.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultNeighCellListNR", err)
		}
	}
	return nil
}
