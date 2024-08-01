return {
	trigger_files = { "css", "js" }, -- List of file extensions to trigger actions
	update_file = "cmd/main.go",
	build_command = "go build -o tmp/bin/main cmd/main.go",
}
