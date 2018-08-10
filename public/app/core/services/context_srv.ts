import config from 'app/core/config';
import _ from 'lodash';
import coreModule from 'app/core/core_module';
import store from 'app/core/store';

export class User {
  isLogDisplayPlatformAdmin: any;
  isSignedIn: any;
  orgRole: any;
  orgId: number;
  timezone: string;
  helpFlags1: number;
  lightTheme: boolean;
  hasEditPermissionInFolders: boolean;

  constructor() {
    if (config.bootData.user) {
      _.extend(this, config.bootData.user);
    }
  }
}

export class ContextSrv {
  pinned: any;
  version: any;
  user: User;
  isSignedIn: any;
  isLogDisplayPlatformAdmin: any;
  isEditor: any;
  sidemenu: any;
  sidemenuSmallBreakpoint = false;
  hasEditPermissionInFolders: boolean;

  constructor() {
    this.sidemenu = store.getBool('logdisplayplatform.sidemenu', true);

    if (!config.buildInfo) {
      config.buildInfo = {};
    }
    if (!config.bootData) {
      config.bootData = { user: {}, settings: {} };
    }

    this.version = config.buildInfo.version;
    this.user = new User();
    this.isSignedIn = this.user.isSignedIn;
    this.isLogDisplayPlatformAdmin = this.user.isLogDisplayPlatformAdmin;
    this.isEditor = this.hasRole('Editor') || this.hasRole('Admin');
    this.hasEditPermissionInFolders = this.user.hasEditPermissionInFolders;
  }

  hasRole(role) {
    return this.user.orgRole === role;
  }

  isLogDisplayPlatformVisible() {
    return !!(document.visibilityState === undefined || document.visibilityState === 'visible');
  }

  toggleSideMenu() {
    this.sidemenu = !this.sidemenu;
    store.set('logdisplayplatform.sidemenu', this.sidemenu);
  }
}

var contextSrv = new ContextSrv();
export { contextSrv };

coreModule.factory('contextSrv', function() {
  return contextSrv;
});
