cc_on_windows:
	go-bindata -o app/setup/bindata.go ./assets/... ./templates/...
	sed "s/package main/package setup/g" app/setup/bindata.go > app/setup/bindata.go.$$
	cd app/setup && move /Y bindata.go.$$ bindata.go
	set GOOS=linux& set GOARCH=amd64& go build -o bin/linux_x64
	set GOOS=windows& set GOARCH=amd64& go build -o bin/windows_x64.exe

cc_on_linux:
	go-bindata -o app/setup/bindata.go ./assets/... ./templates/...
	sed "s/package main/package setup/g" app/setup/bindata.go > app/setup/bindata.go.$$
	cd app/setup && move /Y bindata.go.$$ bindata.go
	set GOOS=linux& set GOARCH=amd64& go build -o bin/linux_x64
	set GOOS=windows& set GOARCH=amd64& go build -o bin/windows_x64.exe

dev_on_windows:
	set GOOS=linux& set GOARCH=amd64& go build -o bin/linux_x64
	set GOOS=windows& set GOARCH=amd64& go build -o bin/windows_x64.exe

renew_statics:
	go-bindata -o app/setup/bindata.go ./assets/... ./templates/...
	sed "s/package main/package setup/g" app/setup/bindata.go > app/setup/bindata.go.$$
	cd app/setup && move /Y bindata.go.$$ bindata.go