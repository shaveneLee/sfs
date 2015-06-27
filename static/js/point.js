function SfsCtrl($scope, $http) {
	$scope.data = "sarong";
	$scope.DefaultLines = 7;
	$scope.Rows = [];
	$scope.TypeClass = {
		10: 'label label-primary',
		20: 'label label-default',
	}

	//get type list options.
	$http.get('/point/GetTypeList').error(function(data) {
		alert('get type list error.')
	}).success(function(data) {
		$scope.TypeList = data;
	});

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
