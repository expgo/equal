package equal

/*
Annotation

	@Enum {
		Equal = "Equal"
	}
*/
type Annotation string

type Equal struct {
	Name      string
	Function  string `value:"Equal"` // function name
	Receiver  string `value:"e"`     // receiver name
	Parameter string // parameter name
	Excludes  []string
	Fields    []string
}
