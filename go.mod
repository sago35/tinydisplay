module github.com/sago35/tinydisplay

go 1.16

require (
	fyne.io/fyne/v2 v2.1.2
	tinygo.org/x/drivers v0.19.0
	tinygo.org/x/tinydraw v0.0.0-20200416172542-c30d6d84353c
	tinygo.org/x/tinyfont v0.2.1
)

replace tinygo.org/x/drivers => ../../tinygo-org/drivers
