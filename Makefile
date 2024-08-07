# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
live/templ:
	templ generate --watch --proxy="http://localhost:3030" --open-browser=false -v

# run air to detect any go file changes to re-build and re-run the server.
live/server:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build -o tmp/bin/main cmd/main.go" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true 

# run tailwindcss to generate the styles.css bundle in watch mode.
# live/tailwind:
# 	npx tailwindcss -i ./static/styles/input.css -o ./static/styles/styles.css --minify --watch

# run esbuild to generate the index.js bundle in watch mode.
# live/esbuild:
# 	npx esbuild js/index.ts --bundle --outdir=assets/ --watch

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
live/sync_assets:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "assets" \
	--build.include_ext "js,css,png,jpg,jpeg,gif,svg"
	# --build.include_dir "static" \
	# --build.include_dir "static/styles" \


# start all 5 watch processes in parallel.
live: 
	# make -j2 live/templ live/server 
	make -j3 live/templ live/server live/sync_assets
	# make -j5 live/templ live/server live/sync_assets
	# make -j5 live/templ live/server live/tailwind live/esbuild live/sync_assets
