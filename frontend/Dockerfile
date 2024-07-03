# Utiliza una imagen ligera de Node.js como base para construir la aplicación
FROM node:alpine AS build

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos de tu aplicación al contenedor
COPY package*.json ./
RUN npm install
COPY . .

# Compila la aplicación Angular
RUN npm run build --prod

# Utiliza una imagen ligera de Nginx para servir la aplicación
FROM nginx:alpine

# Copia los archivos compilados al directorio de Nginx
COPY --from=build /app/dist/frontend /usr/share/nginx/html

# Expone el puerto en el que el servidor de Nginx escucha
EXPOSE 4200

# Comando para iniciar Nginx
CMD ["nginx", "-g", "daemon off;"]
