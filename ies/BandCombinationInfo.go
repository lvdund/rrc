package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationInfo struct {
	BandCombinationIndex   BandCombinationIndex   `madatory`
	AllowedFeatureSetsList []FeatureSetEntryIndex `lb:1,ub:maxFeatureSetsPerBand,madatory`
}

func (ie *BandCombinationInfo) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.BandCombinationIndex.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationIndex", err)
	}
	tmp_AllowedFeatureSetsList := utils.NewSequence[*FeatureSetEntryIndex]([]*FeatureSetEntryIndex{}, aper.Constraint{Lb: 1, Ub: maxFeatureSetsPerBand}, false)
	for _, i := range ie.AllowedFeatureSetsList {
		tmp_AllowedFeatureSetsList.Value = append(tmp_AllowedFeatureSetsList.Value, &i)
	}
	if err = tmp_AllowedFeatureSetsList.Encode(w); err != nil {
		return utils.WrapError("Encode AllowedFeatureSetsList", err)
	}
	return nil
}

func (ie *BandCombinationInfo) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.BandCombinationIndex.Decode(r); err != nil {
		return utils.WrapError("Decode BandCombinationIndex", err)
	}
	tmp_AllowedFeatureSetsList := utils.NewSequence[*FeatureSetEntryIndex]([]*FeatureSetEntryIndex{}, aper.Constraint{Lb: 1, Ub: maxFeatureSetsPerBand}, false)
	fn_AllowedFeatureSetsList := func() *FeatureSetEntryIndex {
		return new(FeatureSetEntryIndex)
	}
	if err = tmp_AllowedFeatureSetsList.Decode(r, fn_AllowedFeatureSetsList); err != nil {
		return utils.WrapError("Decode AllowedFeatureSetsList", err)
	}
	ie.AllowedFeatureSetsList = []FeatureSetEntryIndex{}
	for _, i := range tmp_AllowedFeatureSetsList.Value {
		ie.AllowedFeatureSetsList = append(ie.AllowedFeatureSetsList, *i)
	}
	return nil
}
