package main

import (
	"fmt"
	"time"

	"github.com/lack/waybar-krb5/pkg/krbdbus"
	"github.com/lack/waybar-krb5/pkg/krbmit"

	waybar "github.com/lack/gowaybarplug"
)

const (
	Auth     = "authenticated"
	Unauth   = "unauthenticated"
	Expiring = "expiring"
)

var ExpiryWarning = "1h"

func round(t, toNearest time.Duration) time.Duration {
	return t.Round(time.Second).Round(toNearest)
}

func loop(interval time.Duration) {
	wb := waybar.NewUpdater()
	expiryWarning, err := time.ParseDuration(ExpiryWarning)
	if err != nil {
		panic(err)
	}
	dbusSignal, _ := krbdbus.RegisterDbusInterrupts()
	for {
		status := waybar.Status{
			Text: "krb5",
		}
		ks, err := krbmit.Poll()
		status.Alt = Unauth
		status.Class = []string{Unauth}
		if ks.Principal == "" {
			status.Tooltip = "Uninitialized"
		} else {
			status.Tooltip = ks.Principal
		}
		if ks.HasTicket {
			if ks.Expired {
				status.Tooltip += ": Expired"
			} else {
				status.Class = []string{Auth}
				status.Alt = Auth
				status.Tooltip += ": Valid"
				status.Tooltip += fmt.Sprintf("\n\nExpires in %s\nRenewable for %s", round(ks.Remaining, interval), round(ks.Renewal, interval))
				if ks.Remaining < expiryWarning {
					status.Class = append(status.Class, Expiring)
				}
			}
		} else {
			status.Tooltip += ": No ticket"
		}
		if err != nil {
			status.Tooltip += fmt.Sprintf("\n\nError: %s", err)
		}
		wb.Status <- &status
		// Block until poll interval or interrupted by a dbus signal
		select {
		case <-dbusSignal:
		case <-time.After(interval):
		}
	}
}

func main() {
	loop(30 * time.Second)
}
