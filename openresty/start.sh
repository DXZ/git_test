wget https://openresty.org/download/openresty-1.9.7.5.tar.gz


tar -zxvf openresty-1.9.7.5.tar.gz

./configure --prefix=/database/seven/openresty \
            --with-luajit \
            --with-http_iconv_module \
            --with-http_postgres_module

make && make install