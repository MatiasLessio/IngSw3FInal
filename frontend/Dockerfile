# Build stage
FROM node:alpine AS build

WORKDIR /app
COPY package*.json ./
COPY . ./

RUN npm install -g @angular/cli@16.2.0
RUN npm install sweetalert2
RUN npm install
RUN npm run build

# Production stage
FROM nginx:alpine
COPY --from=build /app/dist/frontend /usr/share/nginx/html

# Copia el archivo de configuración de nginx
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
