FROM golang:1.8.3-alpine3.6

# Environment Variables
ARG LIBVIPS_VERSION=8.5.6
ARG MOZJPEG_VERSION="v3.2"

# Install dependencies
RUN apk add --update --no-cache \
    zlib libxml2 libxslt glib libexif lcms2 fftw ca-certificates \
    giflib libpng libwebp orc tiff poppler-glib librsvg libjpeg-turbo pkgconfig wget && \

    apk add --no-cache --virtual .build-dependencies autoconf automake build-base \
    git libtool nasm zlib-dev libxml2-dev libxslt-dev glib-dev \
    libexif-dev lcms2-dev fftw-dev giflib-dev libpng-dev libwebp-dev orc-dev tiff-dev \
    poppler-dev librsvg-dev && \

# Install imagemagic
    apk add --no-cache imagemagick && \

# Install mozjpeg
    cd /tmp && \
    git clone git://github.com/mozilla/mozjpeg.git && \
    cd /tmp/mozjpeg && \
    git checkout ${MOZJPEG_VERSION} && \
    autoreconf -fiv && ./configure --prefix=/usr && make install && \

# Install libvips
    wget -O- https://github.com/jcupitt/libvips/releases/download/v${LIBVIPS_VERSION}/vips-${LIBVIPS_VERSION}.tar.gz | tar xzC /tmp && \
    cd /tmp/vips-${LIBVIPS_VERSION} && \
    ./configure --prefix=/usr \
                --without-python \
                --without-gsf \
                --enable-debug=no \
                --disable-dependency-tracking \
                --disable-static \
                --enable-silent-rules && \
    make -s install-strip && \
    cd / && \

# Cleanup
    rm -rf /tmp/vips-${LIBVIPS_VERSION} && \
    rm -rf /tmp/mozjpeg && \
#   apk del --purge .build-dependencies && \
    rm -rf /var/cache/apk/* && \

# Install dependencies
    go get github.com/codegangsta/gin

COPY ./* /go/src/dienlanhphongvan-cdn/

WORKDIR $GOPATH/src/dienlanhphongvan-cdn 

CMD  pwd &&ls

EXPOSE 3000 8920



