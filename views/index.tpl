<!DOCTYPE html>

<html>
<head>
    <title>Gogocar</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="mobile-web-app-capable" content="yes">

    <script src="bower_components/webcomponentsjs/webcomponents-lite.min.js"></script>
    <link href='http://fonts.googleapis.com/css?family=Noto+Sans' rel='stylesheet' type='text/css'>
    <link rel="import" href="bower_components/iron-flex-layout/classes/iron-flex-layout.html">
    <link rel="import" href="bower_components/paper-drawer-panel/paper-drawer-panel.html">
    <link rel="import" href="bower_components/paper-header-panel/paper-header-panel.html">
    <link rel="import" href="bower_components/paper-toolbar/paper-toolbar.html">
    <link rel="import" href="bower_components/iron-icons/iron-icons.html">
    <link rel="import" href="bower_components/iron-ajax/iron-ajax.html">
    <link rel="import" href="bower_components/paper-menu/paper-menu.html">
    <link rel="import" href="bower_components/paper-item/paper-item.html">
    <link rel="import" href="custom_components/clients-table.html">
    <link rel="import" href="custom_components/car-detail.html">
    <style is="custom-style">
        body{
            margin: 0px;
            font-family: 'Noto Sans', sans-serif;
        }
        <!--paper-toolbar {-->
            <!--&#45;&#45;paper-toolbar-background: red;-->
        <!--}-->
    </style>
    <script>
		HTMLImports.whenReady(function() {
			var clientsTable = document.querySelector("#clients-table");
			var clientsButton = document.querySelector("#client-button");
			var clientsAjax = document.querySelector("#getClientsAjax");
			
			clientsButton.addEventListener("click", function(e){
				clientsAjax.generateRequest();
			});
			
			clientsAjax.addEventListener("response", function(e){
				var data = e.detail.response;
				clientsTable.data = data;				
				//console.log(data);
			});

			var carDetail = document.querySelector("#car-detail");
			var detailButton = document.querySelector("#detail-button");
			var detailAjax = document.querySelector("#getDetailAjax");

			detailButton.addEventListener("click", function(e){
				detailAjax.generateRequest();
			});

			detailAjax.addEventListener("response", function(e){
				var data = e.detail.response;
				console.log("1" + data);
				carDetail.data = data;
				console.log(data);
			});
		})

    </script>
</head>

<body class="fullbleed vertical layout">
<paper-drawer-panel id="paperDrawerPanel" drawer-width="350px" responsive-width="700px">

    <paper-header-panel drawer>
        <paper-toolbar>
            <div>Gogocar</div>
        </paper-toolbar>
    <paper-menu>
        <paper-item id="client-button">Clients</paper-item>
        <paper-item id="detail-button">Detail</paper-item>
    </paper-menu>
    </paper-header-panel>

    <paper-header-panel main>
        <paper-toolbar>
            <paper-icon-button icon="menu" paper-drawer-toggle></paper-icon-button>
            <div>Client</div>
        </paper-toolbar>
        <clients-table id="clients-table"></clients-table>
        <car-detail id="car-detail"></car-detail>
    </paper-header-panel>

</paper-drawer-panel>

<iron-ajax id="getClientsAjax" url="/api/getclients" method="POST" handle-as="json"></iron-ajax>
<iron-ajax id="getDetailAjax" url="/api/getdetails" method="POST" handle-as="json"></iron-ajax>
</body>

</html>
