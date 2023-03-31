go build -o dist/mylib.dll -buildmode=c-shared mylib/main.go
dlltool -d mylib/exports.def -l dist/mylib.lib -D dist/mylib.dll
