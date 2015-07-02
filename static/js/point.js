function SfsCtrl($scope, $http) {
	$scope.data = "sarong";
	$scope.MaxShowLastWeeks = 1;
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

	//get layout data
	$http.get('/point/GetWeekPointsJson/' + $scope.MaxShowLastWeeks).error(function(data) {
		alert('get data error.')
	}).success(function(data) {
		console.log(data)
		$scope.Weeks = data.index_keys;
		$scope.ModelDatas = data.points;
		for(var k in $scope.ModelDatas) {
			if ($scope.ModelDatas[k]['point'].length < $scope.DefaultLines) {
				var len = $scope.ModelDatas[k]['point'].length;
				for (len; len <= $scope.DefaultLines; len++) {
					var row = {
						Name: '',
						Hours: 0,
						Stars: 0,
						Type: 10	
					}
					$scope.ModelDatas[k]['point'][len] = row;
				}
			}
		}
	});

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
