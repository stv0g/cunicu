---
title: cunicu config reload
sidebar_label: config reload
sidebar_class_name: command-name
slug: /usage/man/config/reload
hide_title: true
keywords:
    - manpage
---

## cunicu config reload

Reload the configuration of the cunīcu daemon

```
cunicu config reload [flags]
```

### Options

```
  -h, --help   help for reload
```

### Options inherited from parent commands

```
  -q, --color string        Enable colorization of output (one of: auto, always, never) (default "auto")
  -l, --log-file string     path of a file to write logs to
  -d, --log-level string    log level (one of: debug, info, warn, error, dpanic, panic, and fatal) (default "info")
  -s, --rpc-socket string   Unix control and monitoring socket (default "/var/run/cunicu.sock")
  -v, --verbose int         verbosity level
```

### SEE ALSO

* [cunicu config](cunicu_config.md)	 - Manage configuration of a running cunīcu daemon.
