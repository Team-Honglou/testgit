+++
title = "Folder/Dashboard Search HTTP API "
description = "LogDisplayPlatform Folder/Dashboard Search HTTP API"
keywords = ["logdisplayplatform", "http", "documentation", "api", "search", "folder", "dashboard"]
aliases = ["/http_api/folder_dashboard_search/"]
type = "docs"
[menu.docs]
name = "Folder/dashboard search"
parent = "http_api"
+++

# Folder/Dashboard Search API

## Search folders and dashboards

`GET /api/search/`

Query parameters:

- **query** – Search Query
- **tag** – List of tags to search for
- **type** – Type to search for, `dash-folder` or `dash-db`
- **dashboardIds** – List of dashboard id's to search for
- **folderIds** – List of folder id's to search in for dashboards
- **starred** – Flag indicating if only starred Dashboards should be returned
- **limit** – Limit the number of returned results

**Example request for retrieving folders and dashboards of the general folder**:

```http
GET /api/search?folderIds=0&query=&starred=false HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
```

**Example response for retrieving folders and dashboards of the general folder**:

```http
HTTP/1.1 200
Content-Type: application/json

[
  {
    "id": 163,
    "uid": "000000163",
    "title": "Folder",
    "url": "/dashboards/f/000000163/folder",
    "type": "dash-folder",
    "tags": [],
    "isStarred": false,
    "uri":"db/folder" // deprecated in LogDisplayPlatform v5.0
  },
  {
    "id":1,
    "uid": "cIBgcSjkk",
    "title":"Production Overview",
    "url": "/d/cIBgcSjkk/production-overview",
    "type":"dash-db",
    "tags":[prod],
    "isStarred":true,
    "uri":"db/production-overview" // deprecated in LogDisplayPlatform v5.0
  }
]
```

**Example request searching for dashboards**:

```http
GET /api/search?query=Production%20Overview&starred=true&tag=prod HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
```

**Example response searching for dashboards**:

```http
HTTP/1.1 200
Content-Type: application/json

[
  {
    "id":1,
    "uid": "cIBgcSjkk",
    "title":"Production Overview",
    "url": "/d/cIBgcSjkk/production-overview",
    "type":"dash-db",
    "tags":[prod],
    "isStarred":true,
    "folderId": 2,
    "folderUid": "000000163",
    "folderTitle": "Folder",
    "folderUrl": "/dashboards/f/000000163/folder",
    "uri":"db/production-overview" // deprecated in LogDisplayPlatform v5.0
  }
]
```