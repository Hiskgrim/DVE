'use strict';

/**
 * @ngdoc function
 * @name dveApp.controller:AboutCtrl
 * @description
 * # AboutCtrl
 * Controller of the dveApp
 */
angular.module('dveApp')
  .controller('RespregradoCtrl', function ($scope,$http) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    var path = "http://127.0.0.1:8080/v1/";
    $http.get(path+'categoria').then(function(response){$scope.categorias=response.data});
    //$http.post(path+'categoria',{id: 5, nombre: "Ejemplo"});

  });
