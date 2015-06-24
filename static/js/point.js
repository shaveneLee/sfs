function SfsCtrl($scope, $http) {
	$scope.data = "sarong";
	$scope.DefaultLines = 7;
	$scope.Rows = [];
	for(var i = 0; i <= $scope.DefaultLines; i++) {
		var row = {
			type: 10,
			name: '',
			hours: 0,
			stars: 0,
			score: 0,
		}
		$scope.Rows.push(row);	
	}
}
