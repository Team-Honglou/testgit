import { LogDisplayPlatformDatasource } from './datasource';
import { QueryCtrl } from 'app/plugins/sdk';

class LogDisplayPlatformQueryCtrl extends QueryCtrl {
  static templateUrl = 'partials/query.editor.html';
}

class LogDisplayPlatformAnnotationsQueryCtrl {
  annotation: any;

  types = [{ text: 'Dashboard', value: 'dashboard' }, { text: 'Tags', value: 'tags' }];

  constructor() {
    this.annotation.type = this.annotation.type || 'tags';
    this.annotation.limit = this.annotation.limit || 100;
  }

  static templateUrl = 'partials/annotations.editor.html';
}

export {
  LogDisplayPlatformDatasource,
  LogDisplayPlatformDatasource as Datasource,
  LogDisplayPlatformQueryCtrl as QueryCtrl,
  LogDisplayPlatformAnnotationsQueryCtrl as AnnotationsQueryCtrl,
};
