package dom

import "syscall/js"

func ReadFile(input js.Value, what2do func(string)) {
	if input.Get("value").String() == "" {
		return;
	}
	fileReader := js.Global().Get("FileReader").New()
	var contents string
	fileReader.Set("onload", js.FuncOf(func(this js.Value, args []js.Value)any{
		contents = fileReader.Get("result").String()
    what2do(contents)
		return nil
	}))
	fileReader.Call("readAsText",input.Get("files").Index(0))
}
