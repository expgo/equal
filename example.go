package equal

//go:generate ag --dev-plugin=github.com/expgo/equal

// @Equal(excludes={"SlotNumber", "Address", "OrderIndex"})
type DataMonitor struct {
	Name         string
	SlotNumber   int
	AddressType  string
	BlockAddress int
	Address      int
	BitIndex     int
	Scale        int
	OrderIndex   int
	DefaultValue string
	Description  string
	Identity     string
	Tag          string
}
