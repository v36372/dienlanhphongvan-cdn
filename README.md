# CDN/IMGX - rsimgx

## Features
  - [Crop](#crop)
  - [Resize](#resize)
  - [Compress to progressive](#progressive)

## Prerequisites
- [libvips](https://github.com/jcupitt/libvips) 7.42+ or 8+ (8.4+ recommended)
- [ImageMagic](http://www.imagemagick.org/) (7.0.6+ recommended)
- [mozjpeg](https://github.com/mozilla/mozjpeg) (3.2+ recommended)
- Env
  ```sh
  $ export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/lib/pkgconfig:/usr/lib/pkgconfig:/usr/lib
  ```

---
## Build

### Development

- With docker, (PORT=3000)
```sh
$ make docker
$ sh run_dev.sh 
```

- With none docker ( have to setup all dependencies ), (PORT=8920)
```sh
$ go build -ldflags=-s server/main.go
```

### Production
```sh
$ make deploy
```

---
## API
### Request format

- Header:
    Key: `Content-Type`
    Value: `application/x-www-form-urlencoded`

- Body:
    Key: `file`
    Value: `file content`

- Query
    Key: `width`
    Type: `integer`

### Crop
Crop give image to given width
- Method: 
  `POST`
- API Endpoint
    `/v1/images/crop`

### Resize
Resize give image to given width and keep aspect ratio
- Method: 
  `POST`
- API Endpoint
    `/v1/images/resize`

### Compress
Convert given image to progressive jpeg
- Method: 
  `POST`
- API Endpoint
    `/v1/images/compress`
 
---
## Installation

### Manual
```sh
$ make
$ nohup ./bin/rsimgx > log/access.log &
```

### Service
```sh
$ sudo service rsimgx start
```
