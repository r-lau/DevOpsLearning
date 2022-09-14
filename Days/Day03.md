## Day 03


## Linux

Brief knowldge of CentOS and YUM(particular for ansible)

Package manager in centOS is RPM(Red Hat Package Manager)

```systemctl start <service>``` is the prefer way to start a service
```systemctl enable <service>``` configured a service to start at startup

Configuring application as a service
Adding a .service file to ```/etc/systemd/system``` with the following content:
``` [Content] \n ExecStart=<command to invoke the application>```

Can then use ```systemctl daemon-reload``` to let the system know there's a new service configured followed by ```systemctl start <new service>```

