package main

import (
	"go/types"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
)

func convertType(t types.Type) jen.Code {
	switch t := t.(type) {
	case *types.Named:
		return convertTypeNamed(t)

	case *types.Alias:
		return convertType(t.Rhs())

	case *types.Basic:
		return jen.Id(t.Name())

	case *types.Signature:
		return convertTypeSignature(t)

	case *types.Pointer:
		return jen.Op("*").Add(convertType(t.Elem()))

	case *types.Slice:
		return jen.Index().Add(convertType(t.Elem()))

	case *types.Map:
		return jen.Map(convertType(t.Key())).Add(convertType(t.Elem()))

	case *types.Struct:
		return convertTypeStruct(t)

	default:
		// Fallback for any other types
		return jen.Id(t.String())
	}
}

func convertTypeSignature(t *types.Signature) jen.Code {
	f := jen.Func().ParamsFunc(func(g *jen.Group) {
		for i := range t.Params().Len() {
			v := t.Params().At(i)

			g.Id(v.Name()).Add(convertType(v.Type()))
		}
	})

	if t.Results().Len() > 0 {
		f = f.ParamsFunc(func(g *jen.Group) {
			for i := range t.Results().Len() {
				v := t.Results().At(i)

				g.Add(convertType(v.Type()))
			}
		})
	}

	return f
}

func convertTypeStruct(t *types.Struct) jen.Code {
	fields := make([]jen.Code, 0, t.NumFields())

	for i := range t.NumFields() {
		field := t.Field(i)
		tag := t.Tag(i)
		fieldCode := jen.Id(field.Name()).Add(convertType(field.Type()))

		if tag != "" {
			tagMap := parseStructTags(tag)
			fieldCode = fieldCode.Tag(tagMap)
		}

		fields = append(fields, fieldCode)
	}

	return jen.Struct(fields...)
}

func convertTypeNamed(t *types.Named) jen.Code {
	pkg := t.Obj().Pkg()

	if pkg != nil && pkg.Path() != "" {
		// Create a qualified reference
		qual := jen.Qual(pkg.Path(), t.Obj().Name())

		// Handle type arguments if present
		typeArgs := t.TypeArgs()

		if typeArgs != nil && typeArgs.Len() > 0 {
			var args []jen.Code

			for i := range typeArgs.Len() {
				args = append(args, convertType(typeArgs.At(i)))
			}

			return qual.Types(args...)
		}

		return qual
	}

	return jen.Id(t.Obj().Name())
}

// parseStructTags parses a raw struct tag string into a map[string]string.
func parseStructTags(tag string) map[string]string {
	tags := make(map[string]string)

	// Simple state machine to parse tags
	for tag != "" {
		name, value, newTag, ok := parseStructTag(tag)
		if ok {
			tags[name] = value
		}

		tag = newTag
	}

	return tags
}

func parseStructTag(tag string) (name, value, rest string, ok bool) {
	tag = strings.TrimLeftFunc(tag, unicode.IsSpace)

	if tag == "" {
		return "", "", "", false
	}

	// Scan to colon
	i := 0
	for i < len(tag) && tag[i] != ':' {
		i++
	}

	if i >= len(tag) {
		return "", "", "", false
	}

	name = tag[:i]
	tag = tag[i+1:]

	// Scan to closing quote, handling escaped quotes
	if tag[0] != '"' {
		return "", "", tag, false
	}

	i = 1
	for i < len(tag) {
		if tag[i] == '"' && tag[i-1] != '\\' {
			break
		}

		i++
	}

	if i >= len(tag) {
		return "", "", tag, false
	}

	value = tag[1:i]
	tag = tag[i+1:]

	return name, value, tag, true
}
