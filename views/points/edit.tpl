<!doctype html>
<html ng-app>
<head>
	<script src='/static/js/jquery-2.1.4.min.js'></script>
	<script src='/static/js/angular.min.js'></script>
	<script src='/static/js/bootstrap.min.js'></script>
	<link href='/static/css/bootstrap.min.css' rel='stylesheet' type='text/css'>
	<link href='/static/css/point.css'>
	<script src='/static/js/point.js'></script>
</head>

<div class="alert alert-info">
	<h2>Summary & Forward</h2>
</div>

<div ng-controller="SfsCtrl" class="point-data">
<pre>{{data}}</pre>
	<div class="panel panel-default col-md-6">
      <div class="panel-heading">Panel heading</div>

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
			<td><span class="label label-primary">Planning</span></td>
<!--
			<th><span class="label label-danger">Unexpect</span></th>
-->
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

	<div class="panel panel-default col-sm-6">
      
      <div class="panel-heading">Panel heading</div>

      
      <table class="table">
        <thead>
          <tr>
            <th>#</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Username</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <th scope="row">1</th>
            <td>Mark</td>
            <td>Otto</td>
            <td>@mdo</td>
          </tr>
          <tr>
            <th scope="row">2</th>
            <td>Jacob</td>
            <td>Thornton</td>
            <td>@fat</td>
          </tr>
          <tr>
            <th scope="row">3</th>
            <td>Larry</td>
            <td>the Bird</td>
            <td>@twitter</td>
          </tr>
        </tbody>
      </table>
    </div>

</div>
</html>
