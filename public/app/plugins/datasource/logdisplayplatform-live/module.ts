import { LogDisplayPlatformStreamDS } from './datasource';
import { QueryCtrl } from 'app/plugins/sdk';

class LogDisplayPlatformQueryCtrl extends QueryCtrl {
  static templateUrl = 'partials/query.editor.html';
}

export { LogDisplayPlatformStreamDS as Datasource, LogDisplayPlatformQueryCtrl as QueryCtrl };
