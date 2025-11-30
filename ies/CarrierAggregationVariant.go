package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CarrierAggregationVariant struct {
	Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD        *CarrierAggregationVariant_fr1fdd_FR1TDD_CA_SpCellOnFR1FDD        `optional`
	Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD        *CarrierAggregationVariant_fr1fdd_FR1TDD_CA_SpCellOnFR1TDD        `optional`
	Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD        *CarrierAggregationVariant_fr1fdd_FR2TDD_CA_SpCellOnFR1FDD        `optional`
	Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD        *CarrierAggregationVariant_fr1fdd_FR2TDD_CA_SpCellOnFR2TDD        `optional`
	Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD        *CarrierAggregationVariant_fr1tdd_FR2TDD_CA_SpCellOnFR1TDD        `optional`
	Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD        *CarrierAggregationVariant_fr1tdd_FR2TDD_CA_SpCellOnFR2TDD        `optional`
	Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD *CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD `optional`
	Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD *CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD `optional`
	Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD *CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD `optional`
}

func (ie *CarrierAggregationVariant) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD != nil, ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD != nil, ie.Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD != nil, ie.Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD != nil, ie.Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD != nil, ie.Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD != nil, ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD != nil, ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD != nil, ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD != nil {
		if err = ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD != nil {
		if err = ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if ie.Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD != nil {
		if err = ie.Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if ie.Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD != nil {
		if err = ie.Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	if ie.Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD != nil {
		if err = ie.Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if ie.Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD != nil {
		if err = ie.Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	if ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD != nil {
		if err = ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD != nil {
		if err = ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD != nil {
		if err = ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	return nil
}

func (ie *CarrierAggregationVariant) Decode(r *aper.AperReader) error {
	var err error
	var Fr1fdd_FR1TDD_CA_SpCellOnFR1FDDPresent bool
	if Fr1fdd_FR1TDD_CA_SpCellOnFR1FDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1fdd_FR1TDD_CA_SpCellOnFR1TDDPresent bool
	if Fr1fdd_FR1TDD_CA_SpCellOnFR1TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1fdd_FR2TDD_CA_SpCellOnFR1FDDPresent bool
	if Fr1fdd_FR2TDD_CA_SpCellOnFR1FDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1fdd_FR2TDD_CA_SpCellOnFR2TDDPresent bool
	if Fr1fdd_FR2TDD_CA_SpCellOnFR2TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1tdd_FR2TDD_CA_SpCellOnFR1TDDPresent bool
	if Fr1tdd_FR2TDD_CA_SpCellOnFR1TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1tdd_FR2TDD_CA_SpCellOnFR2TDDPresent bool
	if Fr1tdd_FR2TDD_CA_SpCellOnFR2TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDDPresent bool
	if Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDDPresent bool
	if Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDDPresent bool
	if Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Fr1fdd_FR1TDD_CA_SpCellOnFR1FDDPresent {
		ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_CA_SpCellOnFR1FDD)
		if err = ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if Fr1fdd_FR1TDD_CA_SpCellOnFR1TDDPresent {
		ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_CA_SpCellOnFR1TDD)
		if err = ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if Fr1fdd_FR2TDD_CA_SpCellOnFR1FDDPresent {
		ie.Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD = new(CarrierAggregationVariant_fr1fdd_FR2TDD_CA_SpCellOnFR1FDD)
		if err = ie.Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if Fr1fdd_FR2TDD_CA_SpCellOnFR2TDDPresent {
		ie.Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD = new(CarrierAggregationVariant_fr1fdd_FR2TDD_CA_SpCellOnFR2TDD)
		if err = ie.Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	if Fr1tdd_FR2TDD_CA_SpCellOnFR1TDDPresent {
		ie.Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD = new(CarrierAggregationVariant_fr1tdd_FR2TDD_CA_SpCellOnFR1TDD)
		if err = ie.Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if Fr1tdd_FR2TDD_CA_SpCellOnFR2TDDPresent {
		ie.Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD = new(CarrierAggregationVariant_fr1tdd_FR2TDD_CA_SpCellOnFR2TDD)
		if err = ie.Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	if Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDDPresent {
		ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD)
		if err = ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDDPresent {
		ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD)
		if err = ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDDPresent {
		ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD)
		if err = ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	return nil
}
