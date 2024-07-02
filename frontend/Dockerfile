# Utiliza una imagen ligera de Node.js como base
FROM node:alpine

# Expone el puerto en el que el servidor de desarrollo de Angular escucha
EXPOSE 4200

# Establece el directorio de trabajo en /frontend
WORKDIR /frontend

# Copia los archivos de tu aplicación al contenedor
COPY . .

# Instala las dependencias globales y locales necesarias
RUN npm install -g @angular/cli@16.2.0
RUN npm install sweetalert2
RUN npm install

# Comando para iniciar el servidor de desarrollo de Angular y exponerlo externamente
CMD ["ng", "serve", "--host", "0.0.0.0"]
