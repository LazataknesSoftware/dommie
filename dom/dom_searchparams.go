package dom

import "syscall/js"

func GetArgument(key string) string { //LazataknesSoftware
	context := js.Global()
	return context.Get("URL").New(context.Get("location").Get("href")).Get("searchParams").Call("get",key).String()
}
