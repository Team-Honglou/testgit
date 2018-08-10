+++
title = "Upgrading"
description = "Upgrading LogDisplayPlatform guide"
keywords = ["logdisplayplatform", "configuration", "documentation", "upgrade"]
type = "docs"
[menu.docs]
name = "Upgrading"
identifier = "upgrading"
parent = "installation"
weight = 10
+++

# Upgrading LogDisplayPlatform

We recommend everyone to upgrade LogDisplayPlatform often to stay up to date with the latest fixes and enhancements.
In order make this a reality LogDisplayPlatform upgrades are backward compatible and the upgrade process is simple & quick.

Upgrading is generally always safe (between many minor and one major version) and dashboards and graphs will look the same. There can be minor breaking changes in some edge cases which are usually outlined in the [Release Notes](https://community.logdisplayplatform.com/c/releases) and [Changelog](https://github.com/logdisplayplatform/logdisplayplatform/blob/master/CHANGELOG.md)

## Database Backup

Before upgrading it can be a good idea to backup your LogDisplayPlatform database. This will ensure that you can always rollback to your previous version. During startup, LogDisplayPlatform will automatically migrate the database schema (if there are changes or new tables). Sometimes this can cause issues if you later want to downgrade.

#### sqlite

If you use sqlite you only need to make a backup of your `logdisplayplatform.db` file. This is usually located at `/var/lib/logdisplayplatform/logdisplayplatform.db` on unix system.
If you are unsure what database you use and where it is stored check you logdisplayplatform configuration file. If you
installed logdisplayplatform to custom location using a binary tar/zip it is usually in `<logdisplayplatform_install_dir>/data`.

#### mysql

```bash
backup:
> mysqldump -u root -p[root_password] [logdisplayplatform] > logdisplayplatform_backup.sql

restore:
> mysql -u root -p logdisplayplatform < logdisplayplatform_backup.sql
```

#### postgres

```bash
backup:
> pg_dump logdisplayplatform > logdisplayplatform_backup

restore:
> psql logdisplayplatform < logdisplayplatform_backup
```

### Ubuntu / Debian

If you installed logdisplayplatform by downloading a debian package (`.deb`) you can just follow the same installation guide
and execute the same `dpkg -i` command but with the new package. It will upgrade your LogDisplayPlatform install.

If you used our APT repository:

```bash
sudo apt-get update
sudo apt-get install logdisplayplatform
```

#### Upgrading from binary tar file

If you downloaded the binary tar package you can just download and extract a new package
and overwrite all your existing files. But this might overwrite your config changes. We
recommend you place your config changes in a file named  `<logdisplayplatform_install_dir>/conf/custom.ini`
as this will make upgrades easier without risking losing your config changes.

### Centos / RHEL

If you installed logdisplayplatform by downloading a rpm package you can just follow the same installation guide
and execute the same `yum install` or `rpm -i` command but with the new package. It will upgrade your LogDisplayPlatform install.

If you used our YUM repository:

```bash
sudo yum update logdisplayplatform
```

### Docker

This just an example, details depend on how you configured your logdisplayplatform container.
```bash
docker pull logdisplayplatform
docker stop my-logdisplayplatform-container
docker rm my-logdisplayplatform-container
docker run --name=my-logdisplayplatform-container --restart=always -v /var/lib/logdisplayplatform:/var/lib/logdisplayplatform
```

### Windows

If you downloaded the windows binary package you can just download a newer package and extract
to the same location (and overwrite the existing files). This might overwrite your config changes. We
recommend you place your config changes in a file named  `<logdisplayplatform_install_dir>/conf/custom.ini`
as this will make upgrades easier without risking losing your config changes.

## Upgrading from 1.x

[Migrating from 1.x to 2.x]({{< relref "installation/migrating_to2.md" >}})

## Upgrading from 2.x

We are not aware of any issues upgrading directly from 2.x to 4.x but to be on the safe side go via 3.x => 4.x.

## Upgrading to v5.0

The dashboard grid layout engine has changed. All dashboards will be automatically upgraded to new
positioning system when you load them in v5. Dashboards saved in v5 will not work in older versions of LogDisplayPlatform. Some
external panel plugins might need to be updated to work properly.

For more details on the new panel positioning system, [click here]({{< relref "reference/dashboard.md#panel-size-position" >}})
