$ErrorActionPreference = "Stop"
Push-Location "$PSScriptRoot\build"
go mod tidy
& go run . $args
Pop-Location
exit $global:LASTEXITCODE
