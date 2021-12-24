package evaluator

import (
	"github.com/petrostrak/the-monkey-programming-language/object"
)

var buildins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments; got %d, want 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported; got %s", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments; got %d, want 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY; got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments; got %d, want 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY; got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			legth := len(arr.Elements)
			if legth > 0 {
				return arr.Elements[legth-1]
			}

			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments; got %d, want 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY; got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			legth := len(arr.Elements)
			if legth > 0 {
				newElement := make([]object.Object, legth-1)
				copy(newElement, arr.Elements[1:legth])
				return &object.Array{Elements: newElement}
			}

			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments; got %d, want 2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY; got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			legth := len(arr.Elements)

			newElement := make([]object.Object, legth+1)
			copy(newElement, arr.Elements)
			newElement[legth] = args[1]

			return &object.Array{Elements: newElement}
		},
	},
}
