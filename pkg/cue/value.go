package cue

import (
	"fmt"
	"strconv"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/build"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/parser"
	"github.com/pkg/errors"
)

// Value is an object with cue.context and vendors
type Value struct {
	v cue.Value
	r *cue.Context
}

// NewValue new a value
func NewValue(s string, tagTempl string) (*Value, error) {
	builder := &build.Instance{}

	file, err := parser.ParseFile("-", s, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	if err := builder.AddSyntax(file); err != nil {
		return nil, err
	}
	return newValue(builder, tagTempl)
}

func newValue(builder *build.Instance, tagTempl string) (*Value, error) {
	r := cuecontext.New()
	inst := r.BuildInstance(builder)
	val := new(Value)
	val.r = r
	val.v = inst

	// do not check val.Err() error here, because the value may be filled later
	return val, nil
}

// LookupValue reports the value at a path starting from val
func (val *Value) LookupValue(paths ...string) (*Value, error) {
	v := val.v.LookupPath(FieldPath(paths...))
	if !v.Exists() {
		return nil, errors.Errorf("failed to lookup value: var(path=%s) not exist", strings.Join(paths, "."))
	}
	return &Value{
		v: v,
		r: val.r,
	}, nil
}

// FieldPath return the cue path of the given paths
func FieldPath(paths ...string) cue.Path {
	s := makePath(paths...)
	if isNumber(s) {
		return cue.MakePath(cue.Str(s))
	}
	return cue.ParsePath(s)
}

// makePath creates a Path from a sequence of string.
func makePath(paths ...string) string {
	mergedPath := ""
	if len(paths) == 0 {
		return mergedPath
	}
	mergedPath = paths[0]
	if mergedPath == "" || (len(paths) == 1 && (strings.Contains(mergedPath, ".") || strings.Contains(mergedPath, "[") || isNumber(mergedPath))) {
		return unquoteString(paths[0])
	}
	if !strings.HasPrefix(mergedPath, "_") && !strings.HasPrefix(mergedPath, "#") {
		mergedPath = fmt.Sprintf("\"%s\"", unquoteString(mergedPath))
	}
	for _, p := range paths[1:] {
		p = unquoteString(p)
		if !strings.HasPrefix(p, "#") {
			mergedPath += fmt.Sprintf("[\"%s\"]", p)
		} else {
			mergedPath += fmt.Sprintf(".%s", p)
		}
	}
	return mergedPath
}

func unquoteString(s string) string {
	if unquote, err := strconv.Unquote(s); err == nil {
		return unquote
	}
	return s
}

func isNumber(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

// Error return value's error information.
func (val *Value) Error() error {
	v := val.CueValue()
	if !v.Exists() {
		return errors.New("empty value")
	}
	if err := val.v.Err(); err != nil {
		return err
	}
	var gerr error
	v.Walk(func(value cue.Value) bool {
		if err := value.Eval().Err(); err != nil {
			gerr = err
			return false
		}
		return true
	}, nil)
	return gerr
}

// CueValue return cue.Value
func (val *Value) CueValue() cue.Value {
	return val.v
}
