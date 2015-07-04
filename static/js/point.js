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
		$scope.SumScore = {};
		for(var k in $scope.ModelDatas) {
			if (null == $scope.ModelDatas[k]['point']) {
				$scope.ModelDatas[k]['point'] = {};
				var len = 0;
			} else {
				var len = $scope.ModelDatas[k]['point'].length;
			}
			if (len < $scope.DefaultLines) {
				for (len; len <= $scope.DefaultLines; len++) {
					var row = {
						Name: '',
						Hours: 0,
						Stars: 0,
						Type: 10
					}
					if (len > 4) { //the left point set unexpect type
						row.Type = 20;
					}
					$scope.ModelDatas[k]['point'][len] = row;
				}
			}
		}

		//sum score
		for(var k in $scope.ModelDatas) {
			$scope.SumScore[k] = 0;
			for (var i in $scope.ModelDatas[k]['point']) {
				if (undefined == $scope.ModelDatas[k]['point'][i].Points) {
					$scope.ModelDatas[k]['point'][i].Points = 0;
				}
				$scope.SumScore[k] += $scope.ModelDatas[k]['point'][i].Points;
			}
		}
	});

	//save points
	$scope.saveData = function() {
		$http.post('/point/SavePoints', {Title: "xxxxf"}).
			error(function(data) {
					console.log(data);
					alert('error');
			}).
			success(function(data) {
					console.log(data)
					alert('success')
			});
	}

	//auto sum points
	$scope.sumScore = function(row, week) {
		row.Points = row.Hours * row.Stars;
		$scope.SumScore[week] = 0;
		for(var i in $scope.ModelDatas[week]['point']) {
			if (undefined == $scope.ModelDatas[week]['point'][i].Points || 0 == $scope.ModelDatas[week]['point'][i].Points) {
				return;
			}
			$scope.SumScore[week] += $scope.ModelDatas[week]['point'][i].Points;
		}
	}

	//add a new row
	$scope.addRow = function(week) {
		var row = {
			Name: '',
			Hours: 0,
			Stars: 0,
			Type: 20
		}
		var len = $scope.ModelDatas[week]['point'].length;
		$scope.ModelDatas[week]['point'][len] = row;
	}

	//inverse point type
	$scope.setType = function(row) {
		if (20 == row.Type) {
			row.Type = 10;
		} else {
			row.Type = 20;
		}
	}
}
