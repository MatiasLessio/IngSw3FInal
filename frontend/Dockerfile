# Utiliza una imagen ligera de Node.js como base para construir la aplicación
FROM node:alpine AS build

# Establece el directorio de trabajo en /app
WORKDIR /app
COPY package*.json ./

COPY . ./

RUN npm install -g @angular/cli@16.2.0
RUN npm install sweetalert2
RUN npm install
# Compila la aplicación Angular
RUN npm run build -- --base-href /frontend/

EXPOSE 4200

CMD [ "node", "server.js" ]
