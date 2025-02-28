package dom

import "syscall/js"

// ----------------------------------------------------------------------------
func ShowAlert(msg string) {
	js.Global().Call("alert", msg)
}

func ShowConfirm(q string) bool { //LazataknesSoftware
	re := js.Global().Call("confirm", q).Bool()
	return re
}

func ShowPrompt(query string) string { //LazataknesSoftware
	res := js.Global().Call("prompt", query).String()
	return res
}
