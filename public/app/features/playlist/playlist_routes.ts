import angular from 'angular';

/** @ngInject */
function logdisplayplatformRoutes($routeProvider) {
  $routeProvider
    .when('/playlists', {
      templateUrl: 'public/app/features/playlist/partials/playlists.html',
      controllerAs: 'ctrl',
      controller: 'PlaylistsCtrl',
    })
    .when('/playlists/create', {
      templateUrl: 'public/app/features/playlist/partials/playlist.html',
      controllerAs: 'ctrl',
      controller: 'PlaylistEditCtrl',
    })
    .when('/playlists/edit/:id', {
      templateUrl: 'public/app/features/playlist/partials/playlist.html',
      controllerAs: 'ctrl',
      controller: 'PlaylistEditCtrl',
    })
    .when('/playlists/play/:id', {
      templateUrl: 'public/app/features/playlist/partials/playlists.html',
      controllerAs: 'ctrl',
      controller: 'PlaylistsCtrl',
      resolve: {
        init: function(playlistSrv, $route) {
          let playlistId = $route.current.params.id;
          playlistSrv.start(playlistId);
        },
      },
    });
}

angular.module('logdisplayplatform.routes').config(logdisplayplatformRoutes);