<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1,maximum-scale=1, user-scalable=no">
		<meta content="telephone=no" name="format-detection" />
		<title>API Doc Tool</title>

		<link rel="icon" type="image/png" href="/static/assets/favicon.png">
		<script type="text/javascript" src="/static/assets/js/jquery.min.js"></script>

		<style>
			@media (min-width: 992px) .container {
				width: 970px;
			}

			@media (min-width: 768px) .container {
				width: 750px;
			}

			#search_container {
				background-color: white;
			}

			#search_container .row {
				padding: 0;
			}

			#search {
				-webkit-tap-highlight-color: transparent;
			}

			.col-xs-12 {
				width: 60%;
				margin: 0 auto;
			}

			#search form {
				height: 320px;
				display: block;
				position: relative;
				-webkit-user-select: none;
				cursor: text;
			}

			#search {
				-webkit-tap-highlight-color: transparent;
			}

			#search form fieldset {
				position: absolute;
				bottom: 16px;
				width: 100%;
			}

			fieldset {
				padding: 0;
				margin: 0;
				border: 0;
			}

			#search form input[type=search] {
				color: #fb0006;
				display: block;
				margin: 0;
				padding: 0;
				width: 100%;
				height: 1.2em;
				border: 0px solid #98989f;
				border-bottom: 1px solid grey;
				background: transparent;
				font-size: 3em;
				-webkit-appearance: none;
				-webkit-tap-highlight-color: transparent;
			}

			input[type="search"] {
				-webkit-box-sizing: border-box;
				-moz-box-sizing: border-box;
				box-sizing: border-box;
			}

			#search form input[type=search]:focus {
				outline: 0;
			}

			#search form span.amount {
				color: gray;
				font-size: 2em;
				position: absolute;
				right: 1.5em;
				top: 0.5em;
			}

			#search form span#search_loupe {
				width: 34px;
				height: 35px;
				position: absolute;
				background: url(/static/assets/img/search_loupe.png);
				background-size: 34px 35px;
				right: 0.8em;
				top: 1.1em;
			}

			#search form input[type=search]::-webkit-search-cancel-button {
				z-index: 9;
				position: absolute;
				right: 0.3em;
				top: 0.5em;
				width: 32px;
				height: 32px;
				margin: 0;
				padding: 0;
				background: url(/static/assets/img/cancel.svg);
				-webkit-appearance: none;
				cursor: pointer
			}

			#search form input[type=search]::-webkit-search-cancel-button:active {
				background-position: 0 0
			}

			.list-group {
				padding-left: 0;
				margin-bottom: 20px;
			}

			ol,
			ul {
				margin-top: 0;
				margin-bottom: 10px;
			}

			.list-group-item:last-child {
				margin-bottom: 0;
				border-bottom-right-radius: 4px;
				border-bottom-left-radius: 4px;
			}

			.list-group-item:first-child {
				border-top-left-radius: 4px;
				border-top-right-radius: 4px;
			}

			.list-group-item {
				position: relative;
				display: block;
				padding: 10px 15px;
				margin-bottom: -1px;
				background-color: #fff;
				border: 1px solid #ddd;
			}

			* {
				-webkit-box-sizing: border-box;
				-moz-box-sizing: border-box;
				box-sizing: border-box;
			}



			.search_result {
				width: 60%;
				margin: 0 auto;
				margin-top: 50px;
			}
		</style>

	</head>

	<body>

		<div class="container" id="search_container">
			<section class="row">
				<div class="col-xs-12" id="search">
					<form onkeypress="return event.keyCode != 13;" class="empty">
						<fieldset>
							<input oninput="search()" autocomplete="off" autofocus="autofocus" class="query" id="search_txt" placeholder="SEARCH*" type="search" />

							<span class="amount"></span>
							<span id="search_loupe"></span>
						</fieldset>
					</form>
				</div>
			</section>
		</div>
		<div class="search_result">
			<ul class="list-group">

			</ul>

		</div>

		<!--<button type="submit" class="btn btn-default" onclick="query()">查询</button>
			<ul class="list-group">

			</ul>-->

		<script>
			function search() {
				if($("#search_txt").val().length>0){
					$("#search_loupe").hide()
				}
				else{
					$("#search_loupe").show();
					$(".list-group").html("");
					return;
				}

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