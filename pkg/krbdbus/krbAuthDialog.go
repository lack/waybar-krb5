package krbdbus

import (
	"github.com/godbus/dbus/v5"
)

/*
    "<node>"
    "  <interface name='org.gnome.KrbAuthDialog'>"
    "    <method name='acquireTgt'>"
    "      <arg type='s' name='principal' direction='in' />"
    "      <arg type='b' name='success' direction='out'/>"
    "    </method>"
    "    <method name='destroyCCache'>"
    "      <arg type='b' name='success' direction='out'/>"
    "    </method>"
    "    <signal name='krb_tgt_acquired'>"
    "       <arg type='s' name='principal' direction ='out'/>"
    "       <arg type='u' name='expiry' direction ='out'/>"
    "    </signal>"
    "    <signal name='krb_tgt_renewed'>"
    "       <arg type='s' name='principal' direction ='out'/>"
    "       <arg type='u' name='expiry' direction ='out'/>"
    "    </signal>"
    "    <signal name='krb_tgt_expired'>"
    "       <arg type='s' name='principal' direction ='out'/>"
    "       <arg type='u' name='expiry' direction ='out'/>"
    "    </signal>"
    "  </interface>"
    "</node>";
*/
const (
	rootObject = "/org/gnome/KrbAuthDialog"

	interfaceName = "org.gnome.KrbAuthDialog"

	methodAcquireTgt    = "acquireTgt"
	methodDestroyCCache = "destroyCCache"

	signalTgtAcquired = "krb_tgt_acquired"
	signalTgtRenewed  = "krb_tgt_renewed"
	signalTgtExpired  = "krb_tgt_expired"
)

func RegisterDbusInterrupts() (chan *dbus.Signal, error) {
	sess, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	err = sess.AddMatchSignal(
		dbus.WithMatchObjectPath(rootObject),
		dbus.WithMatchInterface(interfaceName),
	)
	if err != nil {
		return nil, err
	}
	ch := make(chan *dbus.Signal, 10)
	sess.Signal(ch)
	return ch, nil
}
