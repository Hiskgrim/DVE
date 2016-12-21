'use strict';

/**
 * @ngdoc function
 * @name dveApp.controller:AboutCtrl
 * @description
 * # AboutCtrl
 * Controller of the dveApp
 */
angular.module('dveApp')
  .controller('ResposgradoCtrl', function ($scope,$http) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    $http.get('http://localhost:8080/v1/categoria/').then(function(response){$scope.categorias=response.data});
  });
