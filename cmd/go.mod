module github.com/Qovery/e2e-app/cmd

go 1.16

replace (
	github.com/Qovery/e2e-app/pkg => ../pkg
	github.com/Qovery/e2e-app/utils => ../utils
)

require github.com/Qovery/e2e-app/pkg v0.0.0-00010101000000-000000000000
