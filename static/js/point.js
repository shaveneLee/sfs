function SfsCtrl($scope, $http) {
	$scope.data = "sarong";
	$scope.DefaultLines = 7;
	$scope.Rows = [];
	$scope.TypeList = TypeList
	$scope.TypeClass = {
		10: 'label label-primary',
		20: 'label label-default',
	}
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

	$scope.saveData = function() {
		alert('xx');
		var data = {
			name: 'sarong',
		}
		$http.post('/point/SavePoints', data).
			error(function(data) {
					console.log(data);
					alert('error');
			}).
			success(function(data) {
					console.log(data)
					alert('success')
			});
	}
}
