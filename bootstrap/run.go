package bootstrap

import (
	"path/filepath"

	"github.com/asticode/go-astilectron"
	"github.com/pkg/errors"
)

// Run runs the bootstrap
func Run(o Options) (err error) {
	// Create astilectron
	var a *astilectron.Astilectron
	if a, err = astilectron.New(o.AstilectronOptions); err != nil {
		return errors.Wrap(err, "creating new astilectron failed")
	}
	defer a.Close()
	a.HandleSignals()

	// Provision
	if err = provision(a.BaseDirectoryPath(), o.RestoreAssets, o.CustomProvision); err != nil {
		return errors.Wrap(err, "provisioning failed")
	}

	// Start
	if err = a.Start(); err != nil {
		return errors.Wrap(err, "starting astilectron failed")
	}

	// Debug
	if o.Debug {
		o.WindowOptions.Width = astilectron.PtrInt(*o.WindowOptions.Width + 700)
	}

	// Init window
	var w *astilectron.Window
	if w, err = a.NewWindow(filepath.Join(a.BaseDirectoryPath(), "resources", "app", o.Homepage), o.WindowOptions); err != nil {
		return errors.Wrap(err, "new window failed")
	}

	// Handle messages
	w.On(astilectron.EventNameWindowEventMessage, handleMessages(w, o.MessageHandler))

	// Create window
	if err = w.Create(); err != nil {
		return errors.Wrap(err, "creating window failed")
	}

	// Debug
	if o.Debug {
		if err = w.OpenDevTools(); err != nil {
			return errors.Wrap(err, "opening dev tools failed")
		}
	}

	// Menu
	var m *astilectron.Menu
	if len(o.MenuOptions) > 0 {
		// Init menu
		m = a.NewMenu(o.MenuOptions)

		// Create menu
		if err = m.Create(); err != nil {
			return errors.Wrap(err, "creating menu failed")
		}
	}

	// Tray
	var t *astilectron.Tray
	if o.TrayOptions != nil {
		// Init tray
		t = a.NewTray(o.TrayOptions)

		// Tray menu
		if len(o.TrayMenuOptions) > 0 {
			// Init menu
			var m = t.NewMenu(o.TrayMenuOptions)

			// Create menu
			if err = m.Create(); err != nil {
				return errors.Wrap(err, "creating tray menu failed")
			}
		}

		// Create tray
		if err = t.Create(); err != nil {
			return errors.Wrap(err, "creating tray failed")
		}
	}

	// On wait
	if o.OnWait != nil {
		if err = o.OnWait(a, w, m, t); err != nil {
			return errors.Wrap(err, "onwait failed")
		}
	}

	// Blocking pattern
	a.Wait()
	return
}
