package evaluator

import "github.com/petrostrak/the-monkey-programming-language/object"

var buildins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return NULL
		},
	},
}
