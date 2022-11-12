window.onload = function() {
    var conn;
    var rotationDegree = document.getElementById("rotation-degree");
    var captchaImage = document.getElementById("captcha-image");
    var groupID = document.getElementById("group-id");
    var captchaCount = document.getElementById("captcha-count");
    var captchaRestPlaceholder = document.getElementById("CaptchaRestPlaceholder");


    function displayCaptcha(item) {
        captchaImage.src = item.image;
        showSpinner(false);
        resetRotation()
    }

    function displayCounter(item) {
        captchaCount.innerHTML = item.captchaCount;
    }

    if (captchaRestPlaceholder.innerHTML != "{{.CaptchaJson}}"){
        var msgItem = JSON.parse(captchaRestPlaceholder.innerHTML);
        groupID.innerHTML = msgItem.groupId;
        displayCaptcha(msgItem);
    } else {
        var groupIDparam = get("groupID");
        if (groupIDparam != null){
            groupID.innerHTML = groupIDparam
        }
    }
    //temporary for parsing parameters
    function get(name){
        if(name=(new RegExp('[?&]'+encodeURIComponent(name)+'=([^&]*)')).exec(location.search))
            return decodeURIComponent(name[1]);
    }

    document.getElementById("form").onsubmit = function() {
        if (!conn) {
            return false;
        }
        var answer = {"groupID": groupID.innerHTML, "isRequest": false, "rotationDegree": rotationDegree.value}
        conn.send(JSON.stringify(answer));
        showSpinner(true);
        resetRotation();
        return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("wss://" + document.location.host + "/ws");
        conn.onclose = function(evt) {
            showSpinner(true);
        };
        conn.onmessage = function(evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                console.log(messages[i])
                var msgItem = JSON.parse(messages[i]);
                if (msgItem.groupId === groupID.innerHTML && msgItem.isRequest){
                    if (msgItem.solved == true) {
                        showConfirmation();
                    } else {
                        console.log(msgItem.groupId);
                        console.log(msgItem.image);
                        console.log(msgItem.captchaCount);

                        groupID.innerHTML = msgItem.groupId;
                        displayCaptcha(msgItem);
                        displayCounter(msgItem)
                    }
                }
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};