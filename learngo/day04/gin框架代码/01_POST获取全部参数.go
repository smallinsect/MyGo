package main

import "fmt"

func main() {
	logging.Debugf("c.Request.Method: %v", ctx.Request.Method)
	logging.Debugf("c.Request.ContentType: %v", ctx.ContentType())

	logging.Debugf("c.Request.Body: %v", ctx.Request.Body)
	ctx.Request.ParseForm()
	logging.Debugf("c.Request.Form: %v", ctx.Request.PostForm)

	for k, v := range ctx.Request.PostForm {
		logging.Debugf("k:%v\n", k)
		logging.Debugf("v:%v\n", v)
	}

	logging.Debugf("c.Request.ContentLength: %v", ctx.Request.ContentLength)
	data, _ := ioutil.ReadAll(ctx.Request.Body)

	logging.Debugf("c.Request.GetBody: %v", string(data))
}
