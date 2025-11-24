package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_nothing uint64 = iota
	SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_Ssb_IndexServing_r16
	SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_Ssb_Ncell_r16
	SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_Dl_PRS_r16
)

type SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16 struct {
	Choice               uint64
	Ssb_IndexServing_r16 *SSB_Index
	Ssb_Ncell_r16        *SSB_InfoNcell_r16
	Dl_PRS_r16           *DL_PRS_Info_r16
}

func (ie *SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_Ssb_IndexServing_r16:
		if err = ie.Ssb_IndexServing_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_IndexServing_r16", err)
		}
	case SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_Ssb_Ncell_r16:
		if err = ie.Ssb_Ncell_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_Ncell_r16", err)
		}
	case SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_Dl_PRS_r16:
		if err = ie.Dl_PRS_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Dl_PRS_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_Ssb_IndexServing_r16:
		ie.Ssb_IndexServing_r16 = new(SSB_Index)
		if err = ie.Ssb_IndexServing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_IndexServing_r16", err)
		}
	case SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_Ssb_Ncell_r16:
		ie.Ssb_Ncell_r16 = new(SSB_InfoNcell_r16)
		if err = ie.Ssb_Ncell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Ncell_r16", err)
		}
	case SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16_Choice_Dl_PRS_r16:
		ie.Dl_PRS_r16 = new(DL_PRS_Info_r16)
		if err = ie.Dl_PRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_PRS_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
