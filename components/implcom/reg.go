package implcom

import (
	"strings"

	"github.com/starter-go/application/components"
)

////////////////////////////////////////////////////////////////////////////////

type registrationNormalizer struct {
	r *components.Registration
}

func (inst *registrationNormalizer) GetID() components.ID {
	r := inst.r
	idstr := r.ID.String()
	idstr = strings.TrimSpace(idstr)
	if idstr == "" {
		idstr = "no-id"
	}
	return components.ID(idstr)
}

func (inst *registrationNormalizer) GetAliases(cid components.ID) []string {
	id := cid.String()
	r := inst.r
	text := r.Aliases
	list := inst.parseStringArray(text, func(s string) bool {
		return (s != "" && s != id)
	})
	return list
}

func (inst *registrationNormalizer) GetClasses() []string {
	r := inst.r
	text := r.Classes
	list := inst.parseStringArray(text, func(s string) bool {
		return (s != "")
	})
	return list
}

func (inst *registrationNormalizer) parseStringArray(text string, accept func(string) bool) []string {
	const (
		space = string(' ')
		tab   = "\t"
		nl    = "\n"
	)
	text = strings.ReplaceAll(text, tab, nl)
	text = strings.ReplaceAll(text, space, nl)
	items := strings.Split(text, nl)
	dst := make([]string, 0)
	for _, item := range items {
		item = strings.TrimSpace(item)
		if accept(item) {
			dst = append(dst, item)
		}
	}
	return dst
}

func (inst *registrationNormalizer) GetScope() components.Scope {
	text := inst.r.Scope
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	if text == "prototype" {
		return components.ScopePrototype
	}
	return components.ScopeSingleton
}

func (inst *registrationNormalizer) makeFactory() components.Factory {
	factory := &comFactory{
		r: inst.r,
	}
	return factory
}

////////////////////////////////////////////////////////////////////////////////
