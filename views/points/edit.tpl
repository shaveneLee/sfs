<!doctype html>
<html ng-app>
<head>
<script>
	var TypeList = {{{.TypeList}}}
</script>
	<script src='/static/js/jquery-2.1.4.min.js'></script>
	<script src='/static/js/angular.min.js'></script>
	<script src='/static/js/bootstrap.min.js'></script>
	<link href='/static/css/bootstrap.min.css' rel='stylesheet' type='text/css'>
	<link href='/static/css/point.css' rel='stylesheet' type='text/css' >
	<script src='/static/js/point.js'></script>
</head>

<div class="alert alert-info" id="page_title">
	<h2>Summary & Forward</h2>
</div>
{{{.Str1}}}
<div ng-controller="SfsCtrl" id="point_box">
	<div>
		<a class="btn btn-primary" ng-click="saveData()">test</a>
	</div>
	<div class="panel panel-default col-md-6">
      <div class="panel-heading">This Week Summary</div>

      <!-- Table -->
      <table class="table table-bordered table-condensed table-hover">
        <thead>
          <tr>
            <th>#</th>
            <th>Name</th>
            <th><i class="glyphicon glyphicon-time"></i></th>
            <th><i class="glyphicon glyphicon-star-empty"></i></th>
            <th><i class="glyphicon glyphicon-usd"></i></th>
          </tr>
        </thead>
        <tbody>
          <tr ng-repeat="row in Rows">
			<td>
				<span ng-class="TypeClass[row.type]">{{TypeList[row.type]}}</span>
			</td>
            <td><input type="text" ng-model="row.name" class="form-control" /></td>
            <td><input type="text" ng-model="row.hours" class="form-control span1" /></td>
            <td><input type="text" ng-model="row.stars" class="form-control span1" /></td>
            <td><span class="label label-danger">{{row.hours * row.stars}}</span></td>
          </tr>
		
			<tr>
				<td colspan="4">&nbsp;</td>
				<td><span class="label label-danger">{{SumScore}}</span></td>
			</tr>
        </tbody>
      </table>
    </div>

	<div class="panel panel-default col-sm-6">
      <div class="panel-heading">Next Week Planning</div>

      <!-- Table -->
      <table class="table table-bordered table-condensed table-hover">
        <thead>
          <tr>
            <th>#</th>
            <th>Name</th>
            <th><i class="glyphicon glyphicon-time"></i></th>
            <th><i class="glyphicon glyphicon-star-empty"></i></th>
            <th><i class="glyphicon glyphicon-usd"></i></th>
          </tr>
        </thead>
        <tbody>
          <tr ng-repeat="row in Rows">
			<td>
				<span ng-if="10 == row.type" class="label label-primary">{{TypeList[row.type]}}</span>
				<span ng-if="20 == row.type" class="label label-default">{{TypeList[row.type]}}</span>
			</td>
            <td><input type="text" ng-model="row.name" class="form-control" /></td>
            <td><input type="text" ng-model="row.hours" class="form-control" /></td>
            <td><input type="text" ng-model="row.stars" class="form-control" /></td>
            <td><span class="label label-danger">{{row.hours * row.stars}}</span></td>
          </tr>
		
			<tr>
				<td colspan="4">&nbsp;</td>
				<td><span class="label label-danger">{{SumScore}}</span></td>
			</tr>
        </tbody>
      </table>
    </div>
  </div>

</div>
</html>
