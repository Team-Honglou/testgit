+++
title = "Getting Started"
description = "Getting started with LogDisplayPlatform guide"
keywords = ["logdisplayplatform", "intro", "guide", "started"]
type = "docs"
aliases = ["/guides/gettingstarted"]
[menu.docs]
name = "Getting Started"
identifier = "getting_started_guide"
parent = "guides"
weight = 1
+++

# Getting started

This guide will help you get started and acquainted with LogDisplayPlatform. It assumes you have a working LogDisplayPlatform server up and running and have added at least one [Data Source](/features/datasources/).

## Beginner guides

Watch the 10min [beginners guide to building dashboards](https://www.youtube.com/watch?v=sKNZMtoSHN4&index=7&list=PLDGkOdUX1Ujo3wHw9-z5Vo12YLqXRjzg2) to get a quick intro to setting up Dashboards and Panels.

## Basic Concepts

Read the [Basic Concepts](/guides/basic_concepts) document to get a crash course in key LogDisplayPlatform concepts.

### Top header

Let's start with creating a new Dashboard. You can find the new Dashboard link on the right side of the Dashboard picker. You now have a blank Dashboard.

<img class="no-shadow" src="/img/docs/v50/top_nav_annotated.png" width="580px">

The image above shows you the top header for a Dashboard.

1. Side menubar toggle: This toggles the side menu, allowing you to focus on the data presented in the dashboard. The side menu provides access to features unrelated to a Dashboard such as Users, Organizations, and Data Sources.
2. Dashboard dropdown: This dropdown shows you which Dashboard you are currently viewing, and allows you to easily switch to a new Dashboard. From here you can also create a new Dashboard or folder, Import existing Dashboards, and manage Dashboard playlists.
3. Add Panel: Adds a new panel to the current Dashboard
4. Star Dashboard: Star (or unstar) the current Dashboard. Starred Dashboards will show up on your own Home Dashboard by default, and are a convenient way to mark Dashboards that you're interested in.
5. Share Dashboard: Share the current dashboard by creating a link or create a static Snapshot of it. Make sure the Dashboard is saved before sharing.
6. Save dashboard: The current Dashboard will be saved with the current Dashboard name.
7. Settings: Manage Dashboard settings and features such as Templating and Annotations.

## Dashboards, Panels, the building blocks of LogDisplayPlatform...

Dashboards are at the core of what LogDisplayPlatform is all about. Dashboards are composed of individual Panels arranged on a grid. LogDisplayPlatform ships with a variety of Panels. LogDisplayPlatform makes it easy to construct the right queries, and customize the display properties so that you can create the perfect Dashboard for your need. Each Panel can interact with data from any configured LogDisplayPlatform Data Source (currently InfluxDB, Graphite, OpenTSDB, Prometheus and Cloudwatch). The [Basic Concepts](/guides/basic_concepts) guide explores these key ideas in detail.

<img src="/img/docs/v50/dashboard_annotated.png" class="no-shadow" width="700px">

1. Zoom out time range
2. Time picker dropdown. Here you can access relative time range options, auto refresh options and set custom absolute time ranges.
3. Manual refresh button. Will cause all panels to refresh (fetch new data).
4. Dashboard panel. You edit panels by clicking the panel title.
5. Graph legend. You can change series colors, y-axis and series visibility directly from the legend.


## Adding & Editing Graphs and Panels

![](/img/docs/v45/metrics_tab.png)

1. You add panels by clicking the Add panel icon on the top menu.
2. To edit the graph you click on the graph title to open the panel menu, then `Edit`.
3. This should take you to the `Metrics` tab. In this tab you should see the editor for your default data source.

When you click the `Metrics` tab, you are presented with a Query Editor that is specific to the Panel Data Source. Use the Query Editor to build your queries and LogDisplayPlatform will visualize them in real time.

## Drag-and-Drop panels

You can Drag-and-Drop Panels by simply clicking and holding the Panel title, and drag it to its new location. You can also easily resize panels by clicking the (-) and (+) icons.

![](/img/docs/animated_gifs/drag_drop.gif)

## Tips and shortcuts

* Click the graph title and in the dropdown menu quickly change span or duplicate the panel.
* Click the Save icon in the menu to save the dashboard with a new name
* Click the Save icon in the menu and then advanced to export the dashboard to json file, or set it as your default dashboard.
* Click the colored icon in the legend to select series color
* Click series name in the legend to hide series
* Ctrl/Shift/Meta + Click legend name to hide other series

## LogDisplayPlatform loves the keyboard

* Ctrl+S Saves the current dashboard
* Ctrl+F Opens the dashboard finder / search
* Ctrl+H Hides all controls (good for tv displays)
* Hit Escape to exit graph when in fullscreen or edit mode
