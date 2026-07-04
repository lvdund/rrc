// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CGI-InfoEUTRA (line 5912).

var cGIInfoEUTRAConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cgi-info-EPC", Optional: true},
		{Name: "cgi-info-5GC", Optional: true},
		{Name: "freqBandIndicator"},
		{Name: "multiBandInfoList", Optional: true},
		{Name: "freqBandIndicatorPriority", Optional: true},
	},
}

var cGIInfoEUTRACgiInfo5GCConstraints = per.SizeRange(1, common.MaxPLMN)

const (
	CGI_InfoEUTRA_FreqBandIndicatorPriority_True = 0
)

var cGIInfoEUTRAFreqBandIndicatorPriorityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CGI_InfoEUTRA_FreqBandIndicatorPriority_True},
}

var cGIInfoEUTRACgiInfoEPCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cgi-info-EPC-legacy"},
		{Name: "cgi-info-EPC-list", Optional: true},
	},
}

var cGIInfoEUTRACgiInfoEPCCgiInfoEPCListConstraints = per.SizeRange(1, common.MaxPLMN)

type CGI_InfoEUTRA struct {
	Cgi_Info_EPC *struct {
		Cgi_Info_EPC_Legacy CellAccessRelatedInfo_EUTRA_EPC
		Cgi_Info_EPC_List   []CellAccessRelatedInfo_EUTRA_EPC
	}
	Cgi_Info_5GC              []CellAccessRelatedInfo_EUTRA_5GC
	FreqBandIndicator         FreqBandIndicatorEUTRA
	MultiBandInfoList         *MultiBandInfoListEUTRA
	FreqBandIndicatorPriority *int64
}

func (ie *CGI_InfoEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGIInfoEUTRAConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cgi_Info_EPC != nil, ie.Cgi_Info_5GC != nil, ie.MultiBandInfoList != nil, ie.FreqBandIndicatorPriority != nil}); err != nil {
		return err
	}

	// 2. cgi-info-EPC: sequence
	{
		if ie.Cgi_Info_EPC != nil {
			c := ie.Cgi_Info_EPC
			cGIInfoEUTRACgiInfoEPCSeq := e.NewSequenceEncoder(cGIInfoEUTRACgiInfoEPCConstraints)
			if err := cGIInfoEUTRACgiInfoEPCSeq.EncodePreamble([]bool{c.Cgi_Info_EPC_List != nil}); err != nil {
				return err
			}
			if err := c.Cgi_Info_EPC_Legacy.Encode(e); err != nil {
				return err
			}
			if c.Cgi_Info_EPC_List != nil {
				seqOf := e.NewSequenceOfEncoder(cGIInfoEUTRACgiInfoEPCCgiInfoEPCListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.Cgi_Info_EPC_List))); err != nil {
					return err
				}
				for i := range c.Cgi_Info_EPC_List {
					if err := c.Cgi_Info_EPC_List[i].Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. cgi-info-5GC: sequence-of
	{
		if ie.Cgi_Info_5GC != nil {
			seqOf := e.NewSequenceOfEncoder(cGIInfoEUTRACgiInfo5GCConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cgi_Info_5GC))); err != nil {
				return err
			}
			for i := range ie.Cgi_Info_5GC {
				if err := ie.Cgi_Info_5GC[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. freqBandIndicator: ref
	{
		if err := ie.FreqBandIndicator.Encode(e); err != nil {
			return err
		}
	}

	// 5. multiBandInfoList: ref
	{
		if ie.MultiBandInfoList != nil {
			if err := ie.MultiBandInfoList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. freqBandIndicatorPriority: enumerated
	{
		if ie.FreqBandIndicatorPriority != nil {
			if err := e.EncodeEnumerated(*ie.FreqBandIndicatorPriority, cGIInfoEUTRAFreqBandIndicatorPriorityConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CGI_InfoEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGIInfoEUTRAConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cgi-info-EPC: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Cgi_Info_EPC = &struct {
				Cgi_Info_EPC_Legacy CellAccessRelatedInfo_EUTRA_EPC
				Cgi_Info_EPC_List   []CellAccessRelatedInfo_EUTRA_EPC
			}{}
			c := ie.Cgi_Info_EPC
			cGIInfoEUTRACgiInfoEPCSeq := d.NewSequenceDecoder(cGIInfoEUTRACgiInfoEPCConstraints)
			if err := cGIInfoEUTRACgiInfoEPCSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.Cgi_Info_EPC_Legacy.Decode(d); err != nil {
					return err
				}
			}
			if cGIInfoEUTRACgiInfoEPCSeq.IsComponentPresent(1) {
				seqOf := d.NewSequenceOfDecoder(cGIInfoEUTRACgiInfoEPCCgiInfoEPCListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Cgi_Info_EPC_List = make([]CellAccessRelatedInfo_EUTRA_EPC, n)
				for i := int64(0); i < n; i++ {
					if err := c.Cgi_Info_EPC_List[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 3. cgi-info-5GC: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(cGIInfoEUTRACgiInfo5GCConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cgi_Info_5GC = make([]CellAccessRelatedInfo_EUTRA_5GC, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cgi_Info_5GC[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. freqBandIndicator: ref
	{
		if err := ie.FreqBandIndicator.Decode(d); err != nil {
			return err
		}
	}

	// 5. multiBandInfoList: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MultiBandInfoList = new(MultiBandInfoListEUTRA)
			if err := ie.MultiBandInfoList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. freqBandIndicatorPriority: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cGIInfoEUTRAFreqBandIndicatorPriorityConstraints)
			if err != nil {
				return err
			}
			ie.FreqBandIndicatorPriority = &idx
		}
	}

	return nil
}
