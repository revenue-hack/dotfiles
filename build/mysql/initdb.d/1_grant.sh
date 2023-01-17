#!/bin/bash

MYSQL="mysql -uroot -p${MYSQL_ROOT_PASSWORD}"

echo "\
GRANT ALL ON *.* TO 'root'@'%';
FLUSH PRIVILEGES;
" | ${MYSQL}
