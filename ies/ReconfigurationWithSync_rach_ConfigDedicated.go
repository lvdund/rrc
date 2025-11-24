package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReconfigurationWithSync_rach_ConfigDedicated_Choice_nothing uint64 = iota
	ReconfigurationWithSync_rach_ConfigDedicated_Choice_Uplink
	ReconfigurationWithSync_rach_ConfigDedicated_Choice_SupplementaryUplink
)

type ReconfigurationWithSync_rach_ConfigDedicated struct {
	Choice              uint64
	Uplink              *RACH_ConfigDedicated
	SupplementaryUplink *RACH_ConfigDedicated
}

func (ie *ReconfigurationWithSync_rach_ConfigDedicated) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReconfigurationWithSync_rach_ConfigDedicated_Choice_Uplink:
		if err = ie.Uplink.Encode(w); err != nil {
			err = utils.WrapError("Encode Uplink", err)
		}
	case ReconfigurationWithSync_rach_ConfigDedicated_Choice_SupplementaryUplink:
		if err = ie.SupplementaryUplink.Encode(w); err != nil {
			err = utils.WrapError("Encode SupplementaryUplink", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ReconfigurationWithSync_rach_ConfigDedicated) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReconfigurationWithSync_rach_ConfigDedicated_Choice_Uplink:
		ie.Uplink = new(RACH_ConfigDedicated)
		if err = ie.Uplink.Decode(r); err != nil {
			return utils.WrapError("Decode Uplink", err)
		}
	case ReconfigurationWithSync_rach_ConfigDedicated_Choice_SupplementaryUplink:
		ie.SupplementaryUplink = new(RACH_ConfigDedicated)
		if err = ie.SupplementaryUplink.Decode(r); err != nil {
			return utils.WrapError("Decode SupplementaryUplink", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
