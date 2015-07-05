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
    <!-- foreach output table -->
	<div class="panel panel-default col-md-6" ng-repeat="week in Weeks">
        <div class="panel-heading" ng-if="0 == $index">This Week Summary</div>
        <div class="panel-heading" ng-if="1 == $index">Next Week Planning</div>
        <!-- Table -->
        <table class="table table-bordered table-condensed table-hover">
            <thead>
                <tr>
                    <th><span class="label label-default">#{{ModelDatas[week]['start_time'] | date : 'dd'}} 
                        ~ {{ModelDatas[week]['end_date'] | date:'dd'}}</span></th>
                    <th>Name</th>
                    <th><i class="glyphicon glyphicon-time"></i></th>
                    <th><i class="glyphicon glyphicon-star-empty"></i></th>
                    <th><i class="glyphicon glyphicon-usd"></i></th>
                </tr>
            </thead>
            <tbody>
                <tr ng-repeat="row in ModelDatas[week]['point']">
                    <td class="point-type">
                        <span ng-class="TypeClass[row.Type]" ng-click="setType(row)">{{TypeList[row.Type]}}</span>
                    </td>
                    <td><input type="text" ng-model="row.Name" class="form-control" /></td>
                    <td><input type="text" ng-model="row.Hours" class="form-control span1" ng-change="sumScore(row, week)" /></td>
                    <td><input type="text" ng-model="row.Stars" class="form-control span1" ng-change="sumScore(row, week)" /></td>
                    <td><span class="label label-danger" >{{row.Points | number : 2}}</span></td>
                </tr>

                <tr>
                    <td><a class="btn btn-default btn-sm" ng-click="addRow(week)"><i class="glyphicon glyphicon-plus-sign"></i></a></td>
                    <td colspan="3" class="align-right">Sum:</td>
                    <td><span class="label label-danger" >{{SumScore[week] | number:2}}</span></td>
                </tr>
            </tbody>
        </table>
    </div>

    <div class="submit-btn">
        <a class="btn btn-primary" ng-click="saveData()">Submit</a>
    </div>
</div>
</html>
