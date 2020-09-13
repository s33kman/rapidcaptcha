# RapidCaptcha

This programm displays Captcha challenges over a webapp which can be submitted by 3rd party applications. The Captcha
challenge can be then solved manually be a real human. The answer to the challenge will be send back to the 3rd party client
application so that those apps can solve Captchas even in headless modes. This programm has been implemented using the (https://github.com/gorilla/websocket) package.

## Running a RapidCaptcha server

You can download, build and run the server
using the following commands.

    $ go get github.com/bullrox/rapidcaptcha-server
    $ cd `go list -f '{{.Dir}}' github.com/bullrox/rapidcaptcha-server`
    $ go run *.go

Open http://localhost:80/ in your browser and you are good to go.