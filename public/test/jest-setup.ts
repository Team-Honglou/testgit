import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import 'jquery';
import $ from 'jquery';
import 'angular';
import angular from 'angular';

angular.module('logdisplayplatform', ['ngRoute']);
angular.module('logdisplayplatform.services', ['ngRoute', '$strap.directives']);
angular.module('logdisplayplatform.panels', []);
angular.module('logdisplayplatform.controllers', []);
angular.module('logdisplayplatform.directives', []);
angular.module('logdisplayplatform.filters', []);
angular.module('logdisplayplatform.routes', ['ngRoute']);

jest.mock('app/core/core', () => ({}));
jest.mock('app/features/plugins/plugin_loader', () => ({}));

configure({ adapter: new Adapter() });

var global = <any>window;
global.$ = global.jQuery = $;
