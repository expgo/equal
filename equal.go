package equal

/*
Annotation

	@Enum {
		Equal = "Equal"
	}
*/
type Annotation string

type Equal struct {
	Name          string
	ReceiverName  string `value:"e"`
	ParameterName string `value:"a"`
	FunctionName  string `value:"Equal"`
	Excludes      []string
	Fields        []string
}
