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

<div ng-controller="SfsCtrl">
<pre>{{data}}</pre>
	<div class="panel panel-default col-md-6">
      <div class="panel-heading">Panel heading</div>

      <!-- Table -->
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
