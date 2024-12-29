package main

import (
    "net/http"
    "github.com/webview/webview_go"
)

func main() {
    go func() {
        fs := http.FileServer(http.Dir("./svelte/src"))
        http.Handle("/", fs)
        http.ListenAndServe(":5173", nil)
    }()
    debug := true
    w := webview.New(debug)
    defer w.Destroy()
    w.SetTitle("To do list App")
    w.SetSize(900, 700,webview.HintNone)
    w.Navigate("http://localhost:5173/index.html") 
    w.Run()
}