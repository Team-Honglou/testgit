# Plugin Development 

This document is not meant as complete guide for developing plugins but more as a changelog for changes in
LogDisplayPlatform that can impact plugin development. When ever you as plugin author encounter an issue with your plugin after
upgrading LogDisplayPlatform please check here before creating an issue. 

## Links

- [Datasource plugin written in typescript](https://github.com/logdisplayplatform/typescript-template-datasource)
- [Simple json dataource plugin](https://github.com/logdisplayplatform/simple-json-datasource)
- [Plugin development guide](http://docs.logdisplayplatform.org/plugins/developing/development/)
- [Webpack LogDisplayPlatform plugin template project](https://github.com/CorpGlory/logdisplayplatform-plugin-template-webpack)

## Changes in v4.6

This version of LogDisplayPlatform has big changes that will impact a limited set of plugins. We moved from systemjs to webpack
for built-in plugins & everything internal. External plugins still use systemjs but now with a limited 
set of LogDisplayPlatform components they can import. Plugins can depend on libs like lodash & moment and internal components 
like before using the same import paths. However since everything in LogDisplayPlatform is no longer accessible, a few plugins could encounter issues when importing a LogDisplayPlatform dependency. 

[List of exposed components plugins can import/require](https://github.com/logdisplayplatform/logdisplayplatform/blob/master/public/app/features/plugins/plugin_loader.ts#L48)

If you think we missed exposing a crucial lib or LogDisplayPlatform component let us know by opening an issue.  

### Deprecated components 

The angular directive `<spectrum-picker>` is now deprecated (will still work for a version more) but we recommend plugin authors
to upgrade to new `<color-picker color="ctrl.color" onChange="ctrl.onSparklineColorChange"></color-picker>`

