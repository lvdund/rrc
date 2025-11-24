package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_MBSFN_SubframeConfig_subframeAllocation2_Choice_nothing uint64 = iota
	EUTRA_MBSFN_SubframeConfig_subframeAllocation2_Choice_OneFrame
	EUTRA_MBSFN_SubframeConfig_subframeAllocation2_Choice_FourFrames
)

type EUTRA_MBSFN_SubframeConfig_subframeAllocation2 struct {
	Choice     uint64
	OneFrame   uper.BitString `lb:2,ub:2,madatory`
	FourFrames uper.BitString `lb:8,ub:8,madatory`
}

func (ie *EUTRA_MBSFN_SubframeConfig_subframeAllocation2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation2_Choice_OneFrame:
		if err = w.WriteBitString(ie.OneFrame.Bytes, uint(ie.OneFrame.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			err = utils.WrapError("Encode OneFrame", err)
		}
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation2_Choice_FourFrames:
		if err = w.WriteBitString(ie.FourFrames.Bytes, uint(ie.FourFrames.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode FourFrames", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EUTRA_MBSFN_SubframeConfig_subframeAllocation2) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation2_Choice_OneFrame:
		var tmp_bs_OneFrame []byte
		var n_OneFrame uint
		if tmp_bs_OneFrame, n_OneFrame, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode OneFrame", err)
		}
		ie.OneFrame = uper.BitString{
			Bytes:   tmp_bs_OneFrame,
			NumBits: uint64(n_OneFrame),
		}
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation2_Choice_FourFrames:
		var tmp_bs_FourFrames []byte
		var n_FourFrames uint
		if tmp_bs_FourFrames, n_FourFrames, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode FourFrames", err)
		}
		ie.FourFrames = uper.BitString{
			Bytes:   tmp_bs_FourFrames,
			NumBits: uint64(n_FourFrames),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
