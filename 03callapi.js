handbookControllers.controller(
  'DirectoryCtrl',
  ['$scope', '$http',
   function($scope, $http) {
     $http.get('/api/1/people.json').success(function(data) {
       $scope.people = data;
   });
   }
  ]
);
