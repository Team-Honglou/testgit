+++
title = "Other HTTP API "
description = "LogDisplayPlatform Other HTTP API"
keywords = ["logdisplayplatform", "http", "documentation", "api", "other"]
aliases = ["/http_api/other/"]
type = "docs"
[menu.docs]
name = "Other"
parent = "http_api"
+++


# Frontend Settings API

## Get Settings

`GET /api/frontend/settings`

**Example Request**:

```http
GET /api/frontend/settings HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{
  "allowOrgCreate":true,
  "appSubUrl":"",
  "buildInfo":{
    "buildstamp":xxxxxx,
    "commit":"vyyyy",
    "version":"zzzzz"
  },
  "datasources":{
    "datasourcename":{
      "index":"logdisplayplatform-dash",
      "meta":{
        "annotations":true,
        "module":"plugins/datasource/logdisplayplatform/datasource",
        "name":"LogDisplayPlatform",
        "partials":{
          "annotations":"app/plugins/datasource/logdisplayplatform/partials/annotations.editor.html",
          "config":"app/plugins/datasource/logdisplayplatform/partials/config.html"
        },
        "pluginType":"datasource",
        "serviceName":"LogDisplayPlatform",
        "type":"logdisplayplatformsearch"
      }
    }
  },
  "defaultDatasource": "LogDisplayPlatform"
}
```

# Login API

## Renew session based on remember cookie

`GET /api/login/ping`

**Example Request**:

```http
GET /api/login/ping HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"message": "Logged in"}
```