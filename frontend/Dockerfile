# Utiliza una imagen ligera de Node.js como base para construir la aplicación
FROM node:alpine AS build

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el archivo package.json y package-lock.json (si existe)
COPY package*.json ./

# Instala las dependencias
RUN npm install
RUN npm install -g @angular/cli@16.2.0
RUN npm install sweetalert2
RUN npm install codeceptjs puppeteer


# Copia el resto de los archivos de la aplicación
COPY . .

# Compila la aplicación Angular
RUN npm run build

# Configura Express.js para servir los archivos estáticos
FROM node:alpine
WORKDIR /app
COPY --from=build /app/dist/frontend ./public

# Instala Express.js
RUN npm install express

# Configura Express.js para manejar todas las rutas
COPY server.js .

# Expone el puerto 4200
EXPOSE 4200

# Inicia el servidor Express.js
CMD [ "node", "server.js" ]
