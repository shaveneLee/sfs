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
		for (var k in data.index_keys) {
			var sunday = data.index_keys[k];
			var points = data.points[sunday].point
			for (var i in points) {
				var row = {
					type: points[i]['Type'],
					name: points[i]['Name'],
					hours: points[i]['Hours'],
					stars: points[i]['Stars'],
					score: points[i]['Hours'] * points[i]['Stars']
				}
				$scope.Rows.push(row);	
			}
		}
	});
/*
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
	*/

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
