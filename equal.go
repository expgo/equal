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
	Receiver  string `value:"e"`     // receiver name
	Parameter string `value:"a"`     // parameter name
	Function  string `value:"Equal"` // function name
	Excludes  []string
	Fields    []string
}
