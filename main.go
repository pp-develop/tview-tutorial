package main

import (
    "github.com/rivo/tview"
)

func main() {
    // 新しいテキストボックスを作成
    textBox := tview.NewTextView().
        SetText("Hello, World!").
        SetTextAlign(tview.AlignCenter)

    // 新しいアプリケーションを作成
    app := tview.NewApplication()

    // テキストボックスをアプリケーションに追加
    app.SetRoot(textBox, true)

    // アプリケーションを実行
    if err := app.Run(); err != nil {
        panic(err)
    }
}
