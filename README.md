# Browser Remote

Extension allows to bind any buttons on web pages and then call them remotely.
Remote controls appear on `yourhost.com/` or `localhost:8082/` by default as you bind buttons through the extension.

You need to install both server and extension as extension communicates with server while server porvides web ui.

Binding elements is done through context menu and can be managed through the extension widnow.

If it is hosted remotely you need to specify address in both `config.yaml` and extension.
There's no "user" concept in the implementation so server is supposed to be used per-user.

Server runs with:
```shell
./browser-remote-server -config path/to/config
```

WIP!