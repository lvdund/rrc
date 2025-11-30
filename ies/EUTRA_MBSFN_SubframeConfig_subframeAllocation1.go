package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_nothing uint64 = iota
	EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_OneFrame
	EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_FourFrames
)

type EUTRA_MBSFN_SubframeConfig_subframeAllocation1 struct {
	Choice     uint64
	OneFrame   aper.BitString `lb:6,ub:6,madatory`
	FourFrames aper.BitString `lb:24,ub:24,madatory`
}

func (ie *EUTRA_MBSFN_SubframeConfig_subframeAllocation1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_OneFrame:
		if err = w.WriteBitString(ie.OneFrame.Bytes, uint(ie.OneFrame.NumBits), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			err = utils.WrapError("Encode OneFrame", err)
		}
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_FourFrames:
		if err = w.WriteBitString(ie.FourFrames.Bytes, uint(ie.FourFrames.NumBits), &aper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			err = utils.WrapError("Encode FourFrames", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EUTRA_MBSFN_SubframeConfig_subframeAllocation1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_OneFrame:
		var tmp_bs_OneFrame []byte
		var n_OneFrame uint
		if tmp_bs_OneFrame, n_OneFrame, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode OneFrame", err)
		}
		ie.OneFrame = aper.BitString{
			Bytes:   tmp_bs_OneFrame,
			NumBits: uint64(n_OneFrame),
		}
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_FourFrames:
		var tmp_bs_FourFrames []byte
		var n_FourFrames uint
		if tmp_bs_FourFrames, n_FourFrames, err = r.ReadBitString(&aper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode FourFrames", err)
		}
		ie.FourFrames = aper.BitString{
			Bytes:   tmp_bs_FourFrames,
			NumBits: uint64(n_FourFrames),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
