package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NRDC_Parameters_v15c0 struct {
	Pdcp_DuplicationSplitSRB *NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB `optional`
	Pdcp_DuplicationSplitDRB *NRDC_Parameters_v15c0_pdcp_DuplicationSplitDRB `optional`
}

func (ie *NRDC_Parameters_v15c0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdcp_DuplicationSplitSRB != nil, ie.Pdcp_DuplicationSplitDRB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdcp_DuplicationSplitSRB != nil {
		if err = ie.Pdcp_DuplicationSplitSRB.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_DuplicationSplitSRB", err)
		}
	}
	if ie.Pdcp_DuplicationSplitDRB != nil {
		if err = ie.Pdcp_DuplicationSplitDRB.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_DuplicationSplitDRB", err)
		}
	}
	return nil
}

func (ie *NRDC_Parameters_v15c0) Decode(r *aper.AperReader) error {
	var err error
	var Pdcp_DuplicationSplitSRBPresent bool
	if Pdcp_DuplicationSplitSRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_DuplicationSplitDRBPresent bool
	if Pdcp_DuplicationSplitDRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdcp_DuplicationSplitSRBPresent {
		ie.Pdcp_DuplicationSplitSRB = new(NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB)
		if err = ie.Pdcp_DuplicationSplitSRB.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_DuplicationSplitSRB", err)
		}
	}
	if Pdcp_DuplicationSplitDRBPresent {
		ie.Pdcp_DuplicationSplitDRB = new(NRDC_Parameters_v15c0_pdcp_DuplicationSplitDRB)
		if err = ie.Pdcp_DuplicationSplitDRB.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_DuplicationSplitDRB", err)
		}
	}
	return nil
}
