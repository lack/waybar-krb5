package krbmit

import (
	"strings"
	"time"

	"github.com/lack/gokrb5"
)

type KrbStatus struct {
	HasTicket bool
	Principal string
	Expired   bool
	Remaining time.Duration
	Renewal   time.Duration
}

func Poll() (KrbStatus, error) {
	ks := KrbStatus{}
	krb5, err := gokrb5.InitContext()
	if err != nil {
		return ks, err
	}

	cc, err := krb5.CcDefault()
	if err != nil {
		return ks, err
	}
	defer cc.Close()

	p, err := cc.GetPrincipal()
	if err != nil {
		return ks, err
	}
	ks.Principal = strings.Join(p.Name(), "::") + "@" + p.Realm()

	tgt, err := cc.FindTgt()
	if err != nil || tgt == nil {
		return ks, err
	}
	ks.HasTicket = true
	ks.Expired = tgt.Expired()
	if !ks.Expired {
		ks.Remaining = time.Until(tgt.End())
		ks.Renewal = time.Until(tgt.RenewUntil())
	}
	return ks, nil
}
