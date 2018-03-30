# Install image magic

Run following script
```
sudo yum install -y tar curl libpng-devel libjpeg-devel libxml2-devel \
  zlib-devel openjpeg-devel libtiff-devel gdk-pixbuf2-devel  \
  sqlite-devel cairo-devel glib2-devel \
  tar curl gtk-doc libxml2-devel libjpeg-turbo-devel \
  libpng-devel libtiff-devel libexif-devel libgsf-devel \
  lcms2-devel gobject-introspection-devel libwebp-devel
  bzip2-devel libXt-devel OpenEXR-devel ghostscript-devel openjpeg2-devel
wget https://www.imagemagick.org/download/linux/CentOS/x86_64/ImageMagick-devel-7.0.6-0.x86_64.rpm
sudo rpm -Uvh ImageMagick-devel-7.0.6-0.x86_64.rpm
```

# Install vip

Run following script
```
pushd /opt
sudo sh install_vip.sh
popd
```

# Install mozjpeg
Follow instruction: https://github.com/mozilla/mozjpeg