package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NonCellDefiningSSB_r17 struct {
	AbsoluteFrequencySSB_r17 ARFCN_ValueNR                               `madatory`
	Ssb_Periodicity_r17      *NonCellDefiningSSB_r17_ssb_Periodicity_r17 `optional`
	Ssb_TimeOffset_r17       *NonCellDefiningSSB_r17_ssb_TimeOffset_r17  `optional`
}

func (ie *NonCellDefiningSSB_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_Periodicity_r17 != nil, ie.Ssb_TimeOffset_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.AbsoluteFrequencySSB_r17.Encode(w); err != nil {
		return utils.WrapError("Encode AbsoluteFrequencySSB_r17", err)
	}
	if ie.Ssb_Periodicity_r17 != nil {
		if err = ie.Ssb_Periodicity_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_Periodicity_r17", err)
		}
	}
	if ie.Ssb_TimeOffset_r17 != nil {
		if err = ie.Ssb_TimeOffset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_TimeOffset_r17", err)
		}
	}
	return nil
}

func (ie *NonCellDefiningSSB_r17) Decode(r *uper.UperReader) error {
	var err error
	var Ssb_Periodicity_r17Present bool
	if Ssb_Periodicity_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_TimeOffset_r17Present bool
	if Ssb_TimeOffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.AbsoluteFrequencySSB_r17.Decode(r); err != nil {
		return utils.WrapError("Decode AbsoluteFrequencySSB_r17", err)
	}
	if Ssb_Periodicity_r17Present {
		ie.Ssb_Periodicity_r17 = new(NonCellDefiningSSB_r17_ssb_Periodicity_r17)
		if err = ie.Ssb_Periodicity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Periodicity_r17", err)
		}
	}
	if Ssb_TimeOffset_r17Present {
		ie.Ssb_TimeOffset_r17 = new(NonCellDefiningSSB_r17_ssb_TimeOffset_r17)
		if err = ie.Ssb_TimeOffset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_TimeOffset_r17", err)
		}
	}
	return nil
}
