# json-formatter-in-clipboard

ファイルを作らずにJSONをフォーマットする  
Format JSON without creating a file  
  
使うとしたら、exeなど各々ビルドしてもらい、ビルドしたファイルを適当なショートカットキーに割り当てるのが良いと思う。  
１時間程度で作ったものなので適当です。  
If you want to use it, I think it is better to have each build such as exe and assign the built file to an appropriate shortcut key.  
It is suitable because it was made in about an hour.  
  
 ## Motivation

クリップボードにあるJSONをフォーマットして書き換える物があると便利だなと思ったから  
I thought it would be convenient to have something to format and rewrite the JSON on the clipboard.  

## Development environment

```bash
$ go version
go version go1.15.6 linux/amd64
```

```cmd
> go version
go version go1.17 windows/amd64
```