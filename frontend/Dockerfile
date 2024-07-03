# Utiliza una imagen ligera de Node.js como base
FROM node:latest as build

WORKDIR /usr/local/frontend
COPY ./ /usr/local/frontend/

RUN npm install

#RUN npm install -g @angular/cli@16.2.0
#RUN npm install sweetalert2

RUN npm run build

FROM nginx:latest
COPY --from=build /usr/local/frontend/dist/frontend /usr/share/nginx/html

EXPOSE 4200
