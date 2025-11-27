package note

const (
	ExtPlain    = ".txt"
	ExtMarkdown = ".md"
)

var targetExtensions []string = []string{ExtPlain, ExtMarkdown}

const (
	NameSpaceSeparator = " "
	NameMinusSeparator = "-"
	NameDashSeparator  = "_"
)

var pathNameSeparators []string = []string{NameMinusSeparator, NameDashSeparator}
