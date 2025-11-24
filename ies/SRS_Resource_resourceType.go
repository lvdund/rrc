package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_resourceType_Choice_nothing uint64 = iota
	SRS_Resource_resourceType_Choice_Aperiodic
	SRS_Resource_resourceType_Choice_Semi_persistent
	SRS_Resource_resourceType_Choice_Periodic
)

type SRS_Resource_resourceType struct {
	Choice          uint64
	Aperiodic       interface{} `madatory`
	Semi_persistent *SRS_Resource_resourceType_semi_persistent
	Periodic        *SRS_Resource_resourceType_periodic
}

func (ie *SRS_Resource_resourceType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_resourceType_Choice_Aperiodic:
		// interface{} field of choice - nothing to encode
	case SRS_Resource_resourceType_Choice_Semi_persistent:
		if err = ie.Semi_persistent.Encode(w); err != nil {
			err = utils.WrapError("Encode Semi_persistent", err)
		}
	case SRS_Resource_resourceType_Choice_Periodic:
		if err = ie.Periodic.Encode(w); err != nil {
			err = utils.WrapError("Encode Periodic", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_Resource_resourceType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_resourceType_Choice_Aperiodic:
		// interface{} field of choice - nothing to decode
	case SRS_Resource_resourceType_Choice_Semi_persistent:
		ie.Semi_persistent = new(SRS_Resource_resourceType_semi_persistent)
		if err = ie.Semi_persistent.Decode(r); err != nil {
			return utils.WrapError("Decode Semi_persistent", err)
		}
	case SRS_Resource_resourceType_Choice_Periodic:
		ie.Periodic = new(SRS_Resource_resourceType_periodic)
		if err = ie.Periodic.Decode(r); err != nil {
			return utils.WrapError("Decode Periodic", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
