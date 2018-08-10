+++
title = "Installing on Debian / Ubuntu"
description = "Install guide for LogDisplayPlatform"
keywords = ["logdisplayplatform", "installation", "documentation"]
type = "docs"
aliases = ["/installation/installation/debian"]
[menu.docs]
name = "Installing on Ubuntu / Debian"
identifier = "debian"
parent = "installation"
weight = 1
+++

# Installing on Debian / Ubuntu

Description | Download
------------ | -------------
Stable for Debian-based Linux | [logdisplayplatform_5.1.3_amd64.deb](https://s3-us-west-2.amazonaws.com/logdisplayplatform-releases/release/logdisplayplatform_5.1.3_amd64.deb)
<!--
Beta for Debian-based Linux | [logdisplayplatform_5.1.0-beta1_amd64.deb](https://s3-us-west-2.amazonaws.com/logdisplayplatform-releases/release/logdisplayplatform_5.1.0-beta1_amd64.deb)
-->

Read [Upgrading LogDisplayPlatform]({{< relref "installation/upgrading.md" >}}) for tips and guidance on updating an existing
installation.

## Install Stable


```bash
wget https://s3-us-west-2.amazonaws.com/logdisplayplatform-releases/release/logdisplayplatform_5.1.3_amd64.deb
sudo apt-get install -y adduser libfontconfig
sudo dpkg -i logdisplayplatform_5.1.3_amd64.deb
```

<!-- ## Install Latest Beta
```bash
wget https://s3-us-west-2.amazonaws.com/logdisplayplatform-releases/release/logdisplayplatform_5.1.0-beta1_amd64.deb
sudo apt-get install -y adduser libfontconfig
sudo dpkg -i logdisplayplatform_5.1.0-beta1_amd64.deb
``` -->

## APT Repository

Add the following line to your `/etc/apt/sources.list` file.

```bash
deb https://packagecloud.io/logdisplayplatform/stable/debian/ stretch main
```

Use the above line even if you are on Ubuntu or another Debian version.
There is also a testing repository if you want beta or release
candidates.

```bash
deb https://packagecloud.io/logdisplayplatform/testing/debian/ stretch main
```

Then add the [Package Cloud](https://packagecloud.io/logdisplayplatform) key. This
allows you to install signed packages.

```bash
curl https://packagecloud.io/gpg.key | sudo apt-key add -
```

Update your Apt repositories and install LogDisplayPlatform

```bash
sudo apt-get update
sudo apt-get install logdisplayplatform
```

On some older versions of Ubuntu and Debian you may need to install the
`apt-transport-https` package which is needed to fetch packages over
HTTPS.

```bash
sudo apt-get install -y apt-transport-https
```

## Package details

- Installs binary to `/usr/sbin/logdisplayplatform-server`
- Installs Init.d script to `/etc/init.d/logdisplayplatform-server`
- Creates default file (environment vars) to `/etc/default/logdisplayplatform-server`
- Installs configuration file to `/etc/logdisplayplatform/logdisplayplatform.ini`
- Installs systemd service (if systemd is available) name `logdisplayplatform-server.service`
- The default configuration sets the log file at `/var/log/logdisplayplatform/logdisplayplatform.log`
- The default configuration specifies an sqlite3 db at `/var/lib/logdisplayplatform/logdisplayplatform.db`
- Installs HTML/JS/CSS and other LogDisplayPlatform files at `/usr/share/logdisplayplatform`

## Start the server (init.d service)

Start LogDisplayPlatform by running:

```bash
sudo service logdisplayplatform-server start
```

This will start the `logdisplayplatform-server` process as the `logdisplayplatform` user,
which was created during the package installation. The default HTTP port
is `3000` and default user and group is `admin`.

To configure the LogDisplayPlatform server to start at boot time:

```bash
sudo update-rc.d logdisplayplatform-server defaults
```

## Start the server (via systemd)

To start the service using systemd:

```bash
systemctl daemon-reload
systemctl start logdisplayplatform-server
systemctl status logdisplayplatform-server
```

Enable the systemd service so that LogDisplayPlatform starts at boot.

```bash
sudo systemctl enable logdisplayplatform-server.service
```

## Environment file

The systemd service file and init.d script both use the file located at
`/etc/default/logdisplayplatform-server` for environment variables used when
starting the back-end. Here you can override log directory, data
directory and other variables.

### Logging

By default LogDisplayPlatform will log to `/var/log/logdisplayplatform`

### Database

The default configuration specifies a sqlite3 database located at
`/var/lib/logdisplayplatform/logdisplayplatform.db`. Please backup this database before
upgrades. You can also use MySQL or Postgres as the LogDisplayPlatform database, as detailed on [the configuration page]({{< relref "configuration.md#database" >}}).

## Configuration

The configuration file is located at `/etc/logdisplayplatform/logdisplayplatform.ini`.  Go the
[Configuration]({{< relref "configuration.md" >}}) page for details on all
those options.

### Adding data sources

- [Graphite]({{< relref "features/datasources/graphite.md" >}})
- [InfluxDB]({{< relref "features/datasources/influxdb.md" >}})
- [OpenTSDB]({{< relref "features/datasources/opentsdb.md" >}})
- [Prometheus]({{< relref "features/datasources/prometheus.md" >}})

## Installing from binary tar file

Download [the latest `.tar.gz` file](https://logdisplayplatform.com/get) and
extract it.  This will extract into a folder named after the version you
downloaded. This folder contains all files required to run LogDisplayPlatform.  There are
no init scripts or install scripts in this package.

To configure LogDisplayPlatform add a configuration file named `custom.ini` to the
`conf` folder and override any of the settings defined in
`conf/defaults.ini`.

Start LogDisplayPlatform by executing `./bin/logdisplayplatform-server web`. The `logdisplayplatform-server`
binary needs the working directory to be the root install directory (where the
binary and the `public` folder is located).
