package printer

import (
	"fmt"
	"github.com/pterm/pterm"
	"os"
	"sync"
)

const (
	selectMinEntries       = 5
	selectMaxHeightPadding = 2
)

var (
	instance     Printer
	instanceOnce sync.Once
)

func GetPrinter() Printer {
	instanceOnce.Do(func() {
		instance = new(printer)
	})

	return instance
}

type Printer interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	ErrorStdErr(msg string, args ...any)
	FuzzySelect(msg string, values []string) (string, error)
	Print(msg string)
	Printf(msg string, args ...any)
}

type printer struct {
}

func (p printer) ErrorStdErr(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	_, _ = fmt.Fprintln(os.Stderr, msg)
}

func (p printer) Info(msg string, args ...any) {
	pterm.DefaultBasicText.WithStyle(&pterm.ThemeDefault.InfoMessageStyle).Printfln(msg, args...)
}

func (p printer) Warn(msg string, args ...any) {
	pterm.DefaultBasicText.WithStyle(&pterm.ThemeDefault.WarningMessageStyle).Printfln(msg, args...)
}

func (p printer) Error(msg string, args ...any) {
	pterm.DefaultBasicText.WithStyle(&pterm.ThemeDefault.ErrorMessageStyle).Printfln(msg, args...)
}

func (p printer) FuzzySelect(msg string, values []string) (string, error) {
	_, height, err := pterm.GetTerminalSize()
	if err != nil {
		return "", fmt.Errorf("printer: failed to access the terminal size: %w", err)
	}

	selectHeight := height - selectMaxHeightPadding
	if selectHeight < selectMinEntries {
		selectHeight = selectMinEntries
	}

	interactiveSelect := pterm.DefaultInteractiveSelect.WithDefaultText(msg).WithOptions(values).WithMaxHeight(selectHeight)
	if value, err := interactiveSelect.Show(); err != nil {
		return "", fmt.Errorf("printer: failed the fuzzy select a value: %w", err)
	} else {
		return value, nil
	}
}

func (p printer) Print(msg string) {
	pterm.DefaultBasicText.Print(msg)
}

func (p printer) Printf(msg string, args ...any) {
	pterm.DefaultBasicText.Printf(msg, args...)
}
