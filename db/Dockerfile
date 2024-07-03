FROM mysql:8



ENV HOST=0.0.0.0
# Variables de entorno para configurar la base de datos

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=ingsw3

COPY ingsw3.sql /docker-entrypoint-initdb.d/ingsw3.sql