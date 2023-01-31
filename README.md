## ScreenshotTool
A small CLI tool for taking screenshots

## Example
```
$ ./go-screenshot-tool --path=/tmp/
/tmp/Screenshot_1675176230048015076_1600x900.jpg
```

That's it!

## Available options
 - `-delay`, delay screenshot in seconds, 1,2,3 etc. Example: `-delay 5` or `-delay=5`
 - `-display`, number of display, 1,2,3 etc or 0 (for all). Example: `-display 2` or `-display=2`
 - `-path`, path to save screenshot. Example: `-path /tmp/` or `-path=/tmp/`
 - `-quality`, 0/100 quality of image. Example: `-quality 70` or `-quality=70`

 ## Build
  1. `git clone https://github.com/infnetdanpro/go-screenshot-tool`
  2. `go build -ldflags "-s -w"`
