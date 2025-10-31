package ies

// PUSCH-ConfigDedicatedScell-v1530-uci-OnPUSCH-r15 ::= CHOICE
const (
	PuschConfigdedicatedscellV1530UciOnpuschR15ChoiceNothing = iota
	PuschConfigdedicatedscellV1530UciOnpuschR15ChoiceRelease
	PuschConfigdedicatedscellV1530UciOnpuschR15ChoiceSetup
)

type PuschConfigdedicatedscellV1530UciOnpuschR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PuschConfigdedicatedscellV1530UciOnpuschR15Setup
}
