'use strict';

/**
 * @ngdoc overview
 * @name dveApp
 * @description
 * # dveApp
 *
 * Main module of the application.
 */
angular
  .module('dveApp', [
    'ngAnimate',
    'ngCookies',
    'ngResource',
    'ngRoute',
    'ngSanitize',
    'ngTouch',
    'afOAuth2',
    'treeControl',
    'ngMaterial',
    'ui.grid',
    'ui.grid.edit',
    'ui.grid.rowEdit',
    'ui.grid.cellNav',
    'ui.grid.treeView',
    'ui.grid.selection',
    'ui.grid.exporter'
  ])
    .config(['$locationProvider','$routeProvider', function($locationProvider, $routeProvider) {
      $locationProvider.hashPrefix("");
      $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl',
        controllerAs: 'main'
      })
      .when('/about', {
        templateUrl: 'views/about.html',
        controller: 'AboutCtrl',
        controllerAs: 'about'
      })
      .when('/respregrado', {
        templateUrl: 'views/respregrado.html',
        controller: 'RespregradoCtrl',
        controllerAs: 'respregrado'
      })
      .when('/resposgrado', {
        templateUrl: 'views/resposgrado.html',
        controller: 'ResposgradoCtrl',
        controllerAs: 'resposgrado'
      })
      .when('/hispregrado', {
        templateUrl: 'views/hispregrado.html',
        controller: 'HispregradoCtrl',
        controllerAs: 'hispregrado'
      })
      .when('/hisposgrado', {
        templateUrl: 'views/hisposgrado.html',
        controller: 'HisposgradoCtrl',
        controllerAs: 'hisposgrado'
      })
      .otherwise({
        redirectTo: '/'
      });
  }]);
