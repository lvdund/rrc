package ies

// SystemInformationBlockType19-r12-discConfig-r12 ::= SEQUENCE
type Systeminformationblocktype19R12DiscconfigR12 struct {
	DiscrxpoolR12       SlDiscrxpoollistR12
	DisctxpoolcommonR12 *SlDisctxpoollistR12
	DisctxpowerinfoR12  *SlDisctxpowerinfolistR12
	DiscsyncconfigR12   *SlSyncconfiglistR12
}
