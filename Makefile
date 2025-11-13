smoketest: FORCE
	mkdir -p out
	tinygo build -o ./out/wioterminal_basic.uf2          --target wioterminal --size short ./examples/basic
	tinygo build -o ./out/wioterminal_buttons.uf2        --target wioterminal --size short ./examples/buttons
	tinygo build -o ./out/wioterminal_displays.uf2       --target wioterminal --size short ./examples/displays
	tinygo build -o ./out/wioterminal_pyportal_boing.uf2 --target wioterminal --size short ./examples/pyportal_boing
	tinygo build -o ./out/wioterminal_tinydraw.uf2       --target wioterminal --size short ./examples/tinydraw
	#tinygo build -o ./out/wioterminal_touch_paint.uf2    --target wioterminal --size short ./examples/touch_paint
	tinygo build -o ./out/wioterminal_unicode_font.uf2   --target wioterminal --size short ./examples/unicode_font
	go build -o out/basic          ./examples/basic
	go build -o out/buttons        ./examples/buttons
	go build -o out/displays       ./examples/displays
	go build -o out/pyportal_boing ./examples/pyportal_boing
	go build -o out/tinydraw       ./examples/tinydraw
	go build -o out/touch_paint    ./examples/touch_paint
	go build -o out/unicode_font   ./examples/unicode_font

FORCE:

