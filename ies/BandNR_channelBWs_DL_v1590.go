package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_channelBWs_DL_v1590_Choice_nothing uint64 = iota
	BandNR_channelBWs_DL_v1590_Choice_Fr1
	BandNR_channelBWs_DL_v1590_Choice_Fr2
)

type BandNR_channelBWs_DL_v1590 struct {
	Choice uint64
	Fr1    *BandNR_channelBWs_DL_v1590_fr1
	Fr2    *BandNR_channelBWs_DL_v1590_fr2
}

func (ie *BandNR_channelBWs_DL_v1590) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandNR_channelBWs_DL_v1590_Choice_Fr1:
		if err = ie.Fr1.Encode(w); err != nil {
			err = utils.WrapError("Encode Fr1", err)
		}
	case BandNR_channelBWs_DL_v1590_Choice_Fr2:
		if err = ie.Fr2.Encode(w); err != nil {
			err = utils.WrapError("Encode Fr2", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandNR_channelBWs_DL_v1590) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandNR_channelBWs_DL_v1590_Choice_Fr1:
		ie.Fr1 = new(BandNR_channelBWs_DL_v1590_fr1)
		if err = ie.Fr1.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1", err)
		}
	case BandNR_channelBWs_DL_v1590_Choice_Fr2:
		ie.Fr2 = new(BandNR_channelBWs_DL_v1590_fr2)
		if err = ie.Fr2.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
