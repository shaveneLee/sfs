function SfsCtrl($scope, $http) {
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
	var refresh = function() {
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
	})
	};

	//init layout data
	refresh();

	//save points
	$scope.saveData = function() {
		$scope.getSaveData();
		console.log($scope.SaveData)
		show_overlay('submitting');
		$http.post('/point/SavePoints', $scope.SaveData).
			error(function(data) {
					hide_overlay();
					alert(data.message);
					refresh();
			}).
			success(function(data) {
					console.log(data)
					hide_overlay();
					display_saved_msg();
					refresh();
			});
	}

	$scope.getSaveData = function() {
		$scope.SaveData = {};
		var key = 0;
		for (var k in $scope.ModelDatas) {
			for (var i in $scope.ModelDatas[k].point) {
				var model = $scope.ModelDatas[k].point[i];
				if (null == model.Name || '' == model.Name) {
					continue;
				}
				model.StartTime = $scope.ModelDatas[k].start_time;
				model.EndTime = $scope.ModelDatas[k].end_time
				$scope.SaveData[key++] = model
			}
		}
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

var display_saved_msg = function() {
	$('#fade_msg').fadeIn("slow");
	$('#fade_msg').fadeOut("slow");
}

var show_overlay = function(title) {
	if (undefined == title || null == title) {
		title = 'Loading';
	}
	$('#overlay_title').html(title);
	$('#overlay_modal').modal({
		keyboard: false,		
		backdrop: 'static',
	});
}

var hide_overlay = function() {
	$('#overlay_modal').modal('hide');
}
