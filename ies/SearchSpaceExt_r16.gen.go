// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SearchSpaceExt-r16 (line 14445).

var searchSpaceExtR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "controlResourceSetId-r16", Optional: true},
		{Name: "searchSpaceType-r16", Optional: true},
		{Name: "searchSpaceGroupIdList-r16", Optional: true},
		{Name: "freqMonitorLocations-r16", Optional: true},
	},
}

var searchSpaceExtR16SearchSpaceGroupIdListR16Constraints = per.SizeRange(1, 2)

var searchSpaceExtR16FreqMonitorLocationsR16Constraints = per.FixedSize(5)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dci-Format2-4-r16", Optional: true},
		{Name: "dci-Format2-5-r16", Optional: true},
		{Name: "dci-Format2-6-r16", Optional: true},
	},
}

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofCandidates-CI-r16"},
	},
}

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "aggregationLevel1-r16", Optional: true},
		{Name: "aggregationLevel2-r16", Optional: true},
		{Name: "aggregationLevel4-r16", Optional: true},
		{Name: "aggregationLevel8-r16", Optional: true},
		{Name: "aggregationLevel16-r16", Optional: true},
	},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel1_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel1_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel1_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel1_r16_N2},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel2_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel2_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel2_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel2_r16_N2},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel4_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel4_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel4R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel4_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel4_r16_N2},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel8_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel8_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel8R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel8_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel8_r16_N2},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel16_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel16_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel16R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel16_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_4_r16_NrofCandidates_CI_r16_AggregationLevel16_r16_N2},
}

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofCandidates-IAB-r16"},
	},
}

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "aggregationLevel1-r16", Optional: true},
		{Name: "aggregationLevel2-r16", Optional: true},
		{Name: "aggregationLevel4-r16", Optional: true},
		{Name: "aggregationLevel8-r16", Optional: true},
		{Name: "aggregationLevel16-r16", Optional: true},
	},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel1_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel1_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel1_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel1_r16_N2},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel2_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel2_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel2_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel2_r16_N2},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel4_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel4_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel4R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel4_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel4_r16_N2},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel8_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel8_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel8R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel8_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel8_r16_N2},
}

const (
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel16_r16_N1 = 0
	SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel16_r16_N2 = 1
)

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel16R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel16_r16_N1, SearchSpaceExt_r16_SearchSpaceType_r16_Common_r16_Dci_Format2_5_r16_NrofCandidates_IAB_r16_AggregationLevel16_r16_N2},
}

var searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat26R16Constraints = per.SequenceConstraints{
	Extensible: true,
}

type SearchSpaceExt_r16 struct {
	ControlResourceSetId_r16 *ControlResourceSetId_r16
	SearchSpaceType_r16      *struct {
		Common_r16 struct {
			Dci_Format2_4_r16 *struct {
				NrofCandidates_CI_r16 struct {
					AggregationLevel1_r16  *int64
					AggregationLevel2_r16  *int64
					AggregationLevel4_r16  *int64
					AggregationLevel8_r16  *int64
					AggregationLevel16_r16 *int64
				}
			}
			Dci_Format2_5_r16 *struct {
				NrofCandidates_IAB_r16 struct {
					AggregationLevel1_r16  *int64
					AggregationLevel2_r16  *int64
					AggregationLevel4_r16  *int64
					AggregationLevel8_r16  *int64
					AggregationLevel16_r16 *int64
				}
			}
			Dci_Format2_6_r16 *struct{}
		}
	}
	SearchSpaceGroupIdList_r16 []int64
	FreqMonitorLocations_r16   *per.BitString
}

func (ie *SearchSpaceExt_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(searchSpaceExtR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ControlResourceSetId_r16 != nil, ie.SearchSpaceType_r16 != nil, ie.SearchSpaceGroupIdList_r16 != nil, ie.FreqMonitorLocations_r16 != nil}); err != nil {
		return err
	}

	// 2. controlResourceSetId-r16: ref
	{
		if ie.ControlResourceSetId_r16 != nil {
			if err := ie.ControlResourceSetId_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. searchSpaceType-r16: sequence
	{
		if ie.SearchSpaceType_r16 != nil {
			c := ie.SearchSpaceType_r16
			{
				c := &c.Common_r16
				searchSpaceExtR16SearchSpaceTypeR16CommonR16Seq := e.NewSequenceEncoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16Constraints)
				if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16Seq.EncodeExtensionBit(false); err != nil {
					return err
				}
				if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16Seq.EncodePreamble([]bool{c.Dci_Format2_4_r16 != nil, c.Dci_Format2_5_r16 != nil, c.Dci_Format2_6_r16 != nil}); err != nil {
					return err
				}
				if c.Dci_Format2_4_r16 != nil {
					c := (*c.Dci_Format2_4_r16)
					searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16Seq := e.NewSequenceEncoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16Constraints)
					if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
					{
						c := &c.NrofCandidates_CI_r16
						searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Seq := e.NewSequenceEncoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Constraints)
						if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Seq.EncodePreamble([]bool{c.AggregationLevel1_r16 != nil, c.AggregationLevel2_r16 != nil, c.AggregationLevel4_r16 != nil, c.AggregationLevel8_r16 != nil, c.AggregationLevel16_r16 != nil}); err != nil {
							return err
						}
						if c.AggregationLevel1_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel1_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel1R16Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel2_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel2_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel2R16Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel4_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel4_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel4R16Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel8_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel8_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel8R16Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel16_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel16_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel16R16Constraints); err != nil {
								return err
							}
						}
					}
				}
				if c.Dci_Format2_5_r16 != nil {
					c := (*c.Dci_Format2_5_r16)
					searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16Seq := e.NewSequenceEncoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16Constraints)
					if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
					{
						c := &c.NrofCandidates_IAB_r16
						searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Seq := e.NewSequenceEncoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Constraints)
						if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Seq.EncodePreamble([]bool{c.AggregationLevel1_r16 != nil, c.AggregationLevel2_r16 != nil, c.AggregationLevel4_r16 != nil, c.AggregationLevel8_r16 != nil, c.AggregationLevel16_r16 != nil}); err != nil {
							return err
						}
						if c.AggregationLevel1_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel1_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel1R16Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel2_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel2_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel2R16Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel4_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel4_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel4R16Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel8_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel8_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel8R16Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel16_r16 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel16_r16), searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel16R16Constraints); err != nil {
								return err
							}
						}
					}
				}
				if c.Dci_Format2_6_r16 != nil {
					searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat26R16Seq := e.NewSequenceEncoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat26R16Constraints)
					if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat26R16Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. searchSpaceGroupIdList-r16: sequence-of
	{
		if ie.SearchSpaceGroupIdList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(searchSpaceExtR16SearchSpaceGroupIdListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SearchSpaceGroupIdList_r16))); err != nil {
				return err
			}
			for i := range ie.SearchSpaceGroupIdList_r16 {
				if err := e.EncodeInteger(ie.SearchSpaceGroupIdList_r16[i], per.Constrained(0, 1)); err != nil {
					return err
				}
			}
		}
	}

	// 5. freqMonitorLocations-r16: bit-string
	{
		if ie.FreqMonitorLocations_r16 != nil {
			if err := e.EncodeBitString(*ie.FreqMonitorLocations_r16, searchSpaceExtR16FreqMonitorLocationsR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SearchSpaceExt_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(searchSpaceExtR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. controlResourceSetId-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ControlResourceSetId_r16 = new(ControlResourceSetId_r16)
			if err := ie.ControlResourceSetId_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. searchSpaceType-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.SearchSpaceType_r16 = &struct {
				Common_r16 struct {
					Dci_Format2_4_r16 *struct {
						NrofCandidates_CI_r16 struct {
							AggregationLevel1_r16  *int64
							AggregationLevel2_r16  *int64
							AggregationLevel4_r16  *int64
							AggregationLevel8_r16  *int64
							AggregationLevel16_r16 *int64
						}
					}
					Dci_Format2_5_r16 *struct {
						NrofCandidates_IAB_r16 struct {
							AggregationLevel1_r16  *int64
							AggregationLevel2_r16  *int64
							AggregationLevel4_r16  *int64
							AggregationLevel8_r16  *int64
							AggregationLevel16_r16 *int64
						}
					}
					Dci_Format2_6_r16 *struct{}
				}
			}{}
			c := ie.SearchSpaceType_r16
			{
				c := &c.Common_r16
				searchSpaceExtR16SearchSpaceTypeR16CommonR16Seq := d.NewSequenceDecoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16Constraints)
				if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16Seq.DecodeExtensionBit(); err != nil {
					return err
				}
				if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16Seq.DecodePreamble(); err != nil {
					return err
				}
				if searchSpaceExtR16SearchSpaceTypeR16CommonR16Seq.IsComponentPresent(0) {
					c.Dci_Format2_4_r16 = &struct {
						NrofCandidates_CI_r16 struct {
							AggregationLevel1_r16  *int64
							AggregationLevel2_r16  *int64
							AggregationLevel4_r16  *int64
							AggregationLevel8_r16  *int64
							AggregationLevel16_r16 *int64
						}
					}{}
					c := (*c.Dci_Format2_4_r16)
					searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16Seq := d.NewSequenceDecoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16Constraints)
					if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16Seq.DecodeExtensionBit(); err != nil {
						return err
					}
					{
						c := &c.NrofCandidates_CI_r16
						searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Seq := d.NewSequenceDecoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Constraints)
						if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Seq.DecodePreamble(); err != nil {
							return err
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Seq.IsComponentPresent(0) {
							c.AggregationLevel1_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel1R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel1_r16) = v
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Seq.IsComponentPresent(1) {
							c.AggregationLevel2_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel2R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel2_r16) = v
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Seq.IsComponentPresent(2) {
							c.AggregationLevel4_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel4R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel4_r16) = v
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Seq.IsComponentPresent(3) {
							c.AggregationLevel8_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel8R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel8_r16) = v
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16Seq.IsComponentPresent(4) {
							c.AggregationLevel16_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat24R16NrofCandidatesCIR16AggregationLevel16R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel16_r16) = v
						}
					}
				}
				if searchSpaceExtR16SearchSpaceTypeR16CommonR16Seq.IsComponentPresent(1) {
					c.Dci_Format2_5_r16 = &struct {
						NrofCandidates_IAB_r16 struct {
							AggregationLevel1_r16  *int64
							AggregationLevel2_r16  *int64
							AggregationLevel4_r16  *int64
							AggregationLevel8_r16  *int64
							AggregationLevel16_r16 *int64
						}
					}{}
					c := (*c.Dci_Format2_5_r16)
					searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16Seq := d.NewSequenceDecoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16Constraints)
					if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16Seq.DecodeExtensionBit(); err != nil {
						return err
					}
					{
						c := &c.NrofCandidates_IAB_r16
						searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Seq := d.NewSequenceDecoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Constraints)
						if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Seq.DecodePreamble(); err != nil {
							return err
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Seq.IsComponentPresent(0) {
							c.AggregationLevel1_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel1R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel1_r16) = v
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Seq.IsComponentPresent(1) {
							c.AggregationLevel2_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel2R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel2_r16) = v
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Seq.IsComponentPresent(2) {
							c.AggregationLevel4_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel4R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel4_r16) = v
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Seq.IsComponentPresent(3) {
							c.AggregationLevel8_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel8R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel8_r16) = v
						}
						if searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16Seq.IsComponentPresent(4) {
							c.AggregationLevel16_r16 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat25R16NrofCandidatesIABR16AggregationLevel16R16Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel16_r16) = v
						}
					}
				}
				if searchSpaceExtR16SearchSpaceTypeR16CommonR16Seq.IsComponentPresent(2) {
					c.Dci_Format2_6_r16 = &struct{}{}
					searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat26R16Seq := d.NewSequenceDecoder(searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat26R16Constraints)
					if err := searchSpaceExtR16SearchSpaceTypeR16CommonR16DciFormat26R16Seq.DecodeExtensionBit(); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. searchSpaceGroupIdList-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(searchSpaceExtR16SearchSpaceGroupIdListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SearchSpaceGroupIdList_r16 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				ie.SearchSpaceGroupIdList_r16[i] = v
			}
		}
	}

	// 5. freqMonitorLocations-r16: bit-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeBitString(searchSpaceExtR16FreqMonitorLocationsR16Constraints)
			if err != nil {
				return err
			}
			ie.FreqMonitorLocations_r16 = &v
		}
	}

	return nil
}
