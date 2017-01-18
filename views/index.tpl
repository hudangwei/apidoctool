<!DOCTYPE html>
<html>

	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1,maximum-scale=1, user-scalable=no">
		<meta content="telephone=no" name="format-detection" />
		<title>API Doc Tool</title>

        <link rel="icon" type="image/png" href="/static/assets/favicon.png">
		<link rel="stylesheet" href="/static/assets/css/bootstrap.min.css" />
		<script type="text/javascript" src="/static/assets/js/jquery.min.js"></script>
		<script type="text/javascript" src="/static/assets/js/bootstrap.min.js"></script>

		<style>
		</style>

	</head>

	<body>
		<div class="container">
				<div class="form-group">
					<input type="text" id="search_txt" class="form-control" placeholder="输入关键字搜索API文档">
				</div>
				<button type="submit" class="btn btn-default" onclick="query()">搜搜搜!</button>
				<ul class="list-group">

				</ul>
		</div>

		<script>

			function query(){
				$.ajax({
					type: "get",
					url: "search",
					body:{search:$("#search_txt").val()},
					dataType: "json",
					success: function(data) {
						if(data.err) {
							console.log(err)
						} else {
							fullData(data.data);
						}
					}
				});
			}

			function fullData(data){
				$(".list-group").html("");
				var str = '<li class="list-group-item">'+data+'</li>';
                $(".list-group").append(str);
			}

		</script>
	</body>

</html>