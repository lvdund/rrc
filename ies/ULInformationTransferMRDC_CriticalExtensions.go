package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ULInformationTransferMRDC_CriticalExtensions_Choice_nothing uint64 = iota
	ULInformationTransferMRDC_CriticalExtensions_Choice_C1
	ULInformationTransferMRDC_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type ULInformationTransferMRDC_CriticalExtensions struct {
	Choice                   uint64
	C1                       *ULInformationTransferMRDC_CriticalExtensions_C1
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *ULInformationTransferMRDC_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransferMRDC_CriticalExtensions_Choice_C1:
		if err = ie.C1.Encode(w); err != nil {
			err = utils.WrapError("Encode C1", err)
		}
	case ULInformationTransferMRDC_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ULInformationTransferMRDC_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransferMRDC_CriticalExtensions_Choice_C1:
		ie.C1 = new(ULInformationTransferMRDC_CriticalExtensions_C1)
		if err = ie.C1.Decode(r); err != nil {
			return utils.WrapError("Decode C1", err)
		}
	case ULInformationTransferMRDC_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
