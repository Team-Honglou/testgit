// const context = require.context('./', true, /_specs\.ts/);
// context.keys().forEach(context);
// module.exports = context;

import 'babel-polyfill';
import 'jquery';
import angular from 'angular';
import 'angular-mocks';
import 'app/app';

// configure enzyme
import Enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
Enzyme.configure({ adapter: new Adapter() });

angular.module('logdisplayplatform', ['ngRoute']);
angular.module('logdisplayplatform.services', ['ngRoute', '$strap.directives']);
angular.module('logdisplayplatform.panels', []);
angular.module('logdisplayplatform.controllers', []);
angular.module('logdisplayplatform.directives', []);
angular.module('logdisplayplatform.filters', []);
angular.module('logdisplayplatform.routes', ['ngRoute']);

const context = (<any>require).context('../', true, /specs\.(tsx?|js)/);
for (let key of context.keys()) {
  context(key);
}




