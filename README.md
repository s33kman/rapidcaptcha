# RapidCaptcha

RapidCaptcha is listening for incoming Captcha challenges over websockets and displays those challenges in a  web application. Captcha challenges can be submitted by any 3rd party application as long as it is websocket capable. As soon as the challenge has been processed by RapidCaptcha, it will display the challenge in the browser which is then manually solvable over the webapp. The answer to the challenge will be send back to the 3rd party client
application so that even headless applications can solve Captchas from now on. This programm has been implemented using the (https://github.com/gorilla/websocket) package.

## Running a RapidCaptcha server

You can download, build and run the server
using the following commands.

    $ go get github.com/bullrox/rapidcaptcha-server
    $ cd `go list -f '{{.Dir}}' github.com/bullrox/rapidcaptcha-server`
    $ go run *.go

Open http://localhost:80/ in your browser and you are good to go.