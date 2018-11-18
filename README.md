# About

Imgmin is cli utility to minify and optimize images for the web.


## How to install 

To install this tool you must have go installed. Detailed information about installing go can be found [here](https://golang.org/doc/install)

Make sure that `$GOPATH` is set by running `go env`.

after that run the following command to install the package:

`go get -v github.com/thetinygoat/imgmin`

add the `$GOPATH/bin` directory to your `$PATH` and you are good to go.

## Usage

To use imgmin simply cd into the folder containing images and run `$ imgmin` and bam! its done.

## Config

the config file can be found in `~/.config/imgmin/config.json`.

it is a simple json file and can be edited to specify the quality of compression.

```
[
    {
        "type": "jpg",
        "compression": 50
    },
    {
        "type": "png",
        "compression": -2
    }
]

```

#### Jpeg

 `"compression": 50` describes by what value the jpeg images are compressed, 50 in this case. it ranges from `0 - 100`.

#### Png

```
    -2      best speed
    -3      best compression(slower)

```

values can be set according to needs.



