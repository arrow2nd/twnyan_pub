module github.com/arrow2nd/twnyan

go 1.15

require (
	github.com/ChimeraCoder/anaconda v2.0.0+incompatible
	github.com/ChimeraCoder/tokenbucket v0.0.0-20131201223612-c5a927568de7 // indirect
	github.com/abiosoft/ishell v2.0.0+incompatible // indirect
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db // indirect
	github.com/azr/backoff v0.0.0-20160115115103-53511d3c7330 // indirect
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/dustin/go-jsonpointer v0.0.0-20160814072949-ba0abeacc3dc // indirect
	github.com/dustin/gojson v0.0.0-20160307161227-2e71ec9dd5ad // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/garyburd/go-oauth v0.0.0-20180319155456-bca2e7f09a17 // indirect
	github.com/gookit/color v1.3.6
	github.com/mattn/go-runewidth v0.0.10
	github.com/pkg/browser v0.0.0-20210115035449-ce105d075bb4
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
	gopkg.in/abiosoft/ishell.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db => github.com/arrow2nd/readline v0.0.0-20210131072404-666394925404
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 => github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
)
