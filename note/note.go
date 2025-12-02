package note

import (
	"fmt"
	"path"
	"strings"

	"github.com/Krzysztofz01/fns/utils"
)

type Type int

const (
	None Type = iota
	Plain
	Markdown
)

type Note interface {
	GetPath() string
	GetName() string
	GetSearchVector() string
	GetType() Type
}

type note struct {
	Path         string
	Name         string
	SearchVector string
	Type         Type
}

func (n *note) GetType() Type {
	return n.Type
}

func (n *note) GetSearchVector() string {
	return n.SearchVector
}

func (n *note) GetPath() string {
	return n.Path
}

func (n *note) GetName() string {
	return n.Name
}

func NewNote(p string) (Note, error) {
	var t Type
	switch strings.ToLower(path.Ext(p)) {
	case ExtPlain:
		t = Plain
	case ExtMarkdown:
		t = Markdown
	default:
		return nil, fmt.Errorf("note: invalid unsupported note type")
	}

	var (
		nameBuilder         *strings.Builder = new(strings.Builder)
		searchVectorBuilder *strings.Builder = new(strings.Builder)
	)

	for index, token := range parseNameTokens(p) {
		if index > 0 {
			nameBuilder.WriteString(NameSpaceSeparator)
			searchVectorBuilder.WriteString(NameSpaceSeparator)
		}

		nameBuilder.WriteString(utils.Capitalize(token))
		searchVectorBuilder.WriteString(strings.ToLower(token))
	}

	return &note{
		Path:         p,
		Name:         nameBuilder.String(),
		SearchVector: searchVectorBuilder.String(),
		Type:         t,
	}, nil
}

func parseNameTokens(p string) []string {
	p = strings.TrimSuffix(path.Base(p), path.Ext(p))

	for _, separator := range pathNameSeparators {
		p = strings.ReplaceAll(p, separator, NameSpaceSeparator)
	}

	return strings.Split(p, NameSpaceSeparator)
}
