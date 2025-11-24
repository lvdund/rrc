package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_nothing uint64 = iota
	MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_NotUsed
	MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_Rohc
)

type MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17 struct {
	Choice  uint64
	NotUsed uper.NULL `madatory`
	Rohc    *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_NotUsed:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode NotUsed", err)
		}
	case MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_Rohc:
		if err = ie.Rohc.Encode(w); err != nil {
			err = utils.WrapError("Encode Rohc", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_NotUsed:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode NotUsed", err)
		}
	case MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_Rohc:
		ie.Rohc = new(MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc)
		if err = ie.Rohc.Decode(r); err != nil {
			return utils.WrapError("Decode Rohc", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
