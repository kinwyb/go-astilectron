package bootstrap

import (
	"github.com/asticode/go-astilectron"
)

// Options represents options
type Options struct {
	AstilectronOptions astilectron.Options
	CustomProvision    CustomProvision
	Debug              bool
	Homepage           string
	MenuOptions        []*astilectron.MenuItemOptions
	MessageHandler     MessageHandler
	OnWait             OnWait
	RestoreAssets      RestoreAssets
	TrayOptions        *astilectron.TrayOptions
	TrayMenuOptions    []*astilectron.MenuItemOptions
	WindowOptions      *astilectron.WindowOptions
}

// CustomProvision is a function that executes custom provisioning
type CustomProvision func(baseDirectoryPath string) error

// MessageHandler is a functions that handles messages
type MessageHandler func(w *astilectron.Window, m MessageIn)

// OnWait is a function that executes custom actions before waiting
type OnWait func(a *astilectron.Astilectron, w *astilectron.Window, m *astilectron.Menu, t *astilectron.Tray) error

// RestoreAssets is a function that restores assets namely the go-bindata's RestoreAssets method
type RestoreAssets func(dir, name string) error
