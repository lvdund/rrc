package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NotificationMessageSidelink_r17_CriticalExtensions_Choice_nothing uint64 = iota
	NotificationMessageSidelink_r17_CriticalExtensions_Choice_NotificationMessageSidelink_r17
	NotificationMessageSidelink_r17_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type NotificationMessageSidelink_r17_CriticalExtensions struct {
	Choice                          uint64
	NotificationMessageSidelink_r17 *NotificationMessageSidelink_r17_IEs
	CriticalExtensionsFuture        interface{} `madatory`
}

func (ie *NotificationMessageSidelink_r17_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NotificationMessageSidelink_r17_CriticalExtensions_Choice_NotificationMessageSidelink_r17:
		if err = ie.NotificationMessageSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode NotificationMessageSidelink_r17", err)
		}
	case NotificationMessageSidelink_r17_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NotificationMessageSidelink_r17_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NotificationMessageSidelink_r17_CriticalExtensions_Choice_NotificationMessageSidelink_r17:
		ie.NotificationMessageSidelink_r17 = new(NotificationMessageSidelink_r17_IEs)
		if err = ie.NotificationMessageSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NotificationMessageSidelink_r17", err)
		}
	case NotificationMessageSidelink_r17_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
