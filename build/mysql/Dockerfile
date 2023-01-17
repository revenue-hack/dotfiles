FROM mysql:8.0.30

RUN /bin/cp -f /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

CMD ["/entrypoint.sh", "mysqld", "--innodb_use_native_aio=0"]
