package equal

import (
	"github.com/expgo/ag/api"
	"github.com/expgo/factory"
	"github.com/expgo/generic/list"
	"go/ast"
)

// @Singleton
type Factory struct {
}

func (f *Factory) Annotations() map[string][]api.AnnotationType {
	return map[string][]api.AnnotationType{
		AnnotationEqual.Val(): {api.AnnotationTypeType},
	}
}

func (f *Factory) New(typedAnnotations []*api.TypedAnnotation) (api.Generator, error) {
	var allEquals []*Equal

	for _, typedAnnotation := range typedAnnotations {
		for _, an := range typedAnnotation.Annotations.Annotations {
			equal := factory.New[Equal]()
			ts := typedAnnotation.Node.(*ast.TypeSpec)

			if err := an.To(equal); err != nil {
				return nil, err
			}

			equal.Name = ts.Name.Name

			structType, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}

			for _, field := range structType.Fields.List {
				for _, name := range field.Names {
					if !list.Contains(equal.Excludes, name.Name) {
						equal.Fields = append(equal.Fields, name.Name)
					}
				}
			}

			allEquals = append(allEquals, equal)
		}
	}

	if len(allEquals) > 0 {
		return NewGenerator(allEquals), nil
	}

	return nil, nil
}
