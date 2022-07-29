# waybar-krb5

Waybar plugin to display Kerberos TGT ticket status

## Installation

Expects you have krb-auth-dialog installed, as well as the MIT kerberos libries.

```
go install github.com/lack/waybar-krb5@latest
```

## Configuration

In `$XDG_CONFIG_HOME/waybar/config`
```json
{
    // ... other waybar configuration
    "custom/krb5": {
        "format": "{} {icon}",
        "return-type": "json",
        "exec": "$GOPATH/bin/waybar-krb5",
        "format-icons": {
            "authenticated": "",
            "unauthenticated": ""
        },
        "on-click-right": "krb5-auth-dialog",
        "on-click": "dbus-send --print-reply --type=method_call --dest=org.gnome.KrbAuthDialog /org/gnome
/KrbAuthDialog org.gnome.KrbAuthDialog.acquireTgt string:''"
    }
}
```

In `$XDG_CONFIG_HOME/waybar/style.css`
```css
#custom-krb5 {
    background-color: rgba(0x9b, 0x59, 0xb6, 0.8);
    box-shadow: inset 0 -3px rgba(0x9b, 0x59, 0xb6, 1.0);
    color: #000000;
    padding: 0 4px;
}

#custom-krb5.authenticated {
    background-color: rgba(0x2e, 0xcc, 0x71, 0.8);
    box-shadow: inset 0 -3px rgba(0x2e, 0xcc, 0x71, 1.0);
}

#custom-krb5.expiring {
    background-color: rgba(0xf0, 0x93, 0x2b, 0.8);
    box-shadow: inset 0 -3px rgba(0xf0, 0x93, 0x2b, 1.0);
}

#custom-krb5.unauthenticated {
    background-color: rgba(0xeb, 0x4d, 0x4b, 0.8);
    box-shadow: inset 0 -3px rgba(0xeb, 0x4d, 0x4b, 1.0);
}
```
