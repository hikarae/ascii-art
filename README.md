# ascii-art
Make input image to black and white ascii-art
## Build
```
git clone https://github.com/hikarae/ascii-art.git
cd ascii-art
go mod download
// for Windows
go build -o asciiart.exe ./src/main.go
//for Linux
go build -o asciiart ./src/main.go
chmod +X asciiart
```
## Usage
--input-directory
set directory to input file (default .)
--input-file
set input image file name (default input.jpg)
--output-directory
set directory to output file (default .)
--output-file
set output image file name (default output.jpg)