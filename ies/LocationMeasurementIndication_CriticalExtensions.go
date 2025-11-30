package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LocationMeasurementIndication_CriticalExtensions_Choice_nothing uint64 = iota
	LocationMeasurementIndication_CriticalExtensions_Choice_LocationMeasurementIndication
	LocationMeasurementIndication_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type LocationMeasurementIndication_CriticalExtensions struct {
	Choice                        uint64
	LocationMeasurementIndication *LocationMeasurementIndication_IEs
	CriticalExtensionsFuture      interface{} `madatory`
}

func (ie *LocationMeasurementIndication_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationMeasurementIndication_CriticalExtensions_Choice_LocationMeasurementIndication:
		if err = ie.LocationMeasurementIndication.Encode(w); err != nil {
			err = utils.WrapError("Encode LocationMeasurementIndication", err)
		}
	case LocationMeasurementIndication_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *LocationMeasurementIndication_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationMeasurementIndication_CriticalExtensions_Choice_LocationMeasurementIndication:
		ie.LocationMeasurementIndication = new(LocationMeasurementIndication_IEs)
		if err = ie.LocationMeasurementIndication.Decode(r); err != nil {
			return utils.WrapError("Decode LocationMeasurementIndication", err)
		}
	case LocationMeasurementIndication_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
