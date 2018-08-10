+++
title = "Installing on Mac"
description = "Installing LogDisplayPlatform on Mac"
keywords = ["logdisplayplatform", "configuration", "documentation", "mac", "homebrew", "osx"]
type = "docs"
[menu.docs]
parent = "installation"
weight = 4
+++


# Installing on Mac

Installation can be done using [homebrew](http://brew.sh/)

Install latest stable:

```bash
brew update
brew install logdisplayplatform
```

To start logdisplayplatform look at the command printed after the homebrew install completes.

To upgrade use the reinstall command

```bash
brew update
brew reinstall logdisplayplatform
```

-------------

You can also install the latest unstable logdisplayplatform from git:


```bash
brew install --HEAD logdisplayplatform/logdisplayplatform/logdisplayplatform
```

To upgrade logdisplayplatform if you've installed from HEAD:

```bash
brew reinstall --HEAD logdisplayplatform/logdisplayplatform/logdisplayplatform
```

### Starting LogDisplayPlatform

To start LogDisplayPlatform using homebrew services first make sure homebrew/services is installed.

```bash
brew tap homebrew/services
```

Then start LogDisplayPlatform using:

```bash
brew services start logdisplayplatform
```


### Configuration

The Configuration file should be located at `/usr/local/etc/logdisplayplatform/logdisplayplatform.ini`.

### Logs

The log file should be located at `/usr/local/var/log/logdisplayplatform/logdisplayplatform.log`.

### Plugins

If you want to manually install a plugin place it here: `/usr/local/var/lib/logdisplayplatform/plugins`.

### Database

The default sqlite database is located at `/usr/local/var/lib/logdisplayplatform`

