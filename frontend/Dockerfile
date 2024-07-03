# Usar una imagen base de Node.js para construir la aplicación
FROM node:16 AS build

# Crear el directorio de trabajo
WORKDIR /app

# Copiar el package.json y el package-lock.json
COPY package*.json ./

# Instalar las dependencias
RUN npm install

# Copiar el resto del código de la aplicación
COPY . .

# Construir la aplicación Angular
RUN npm run build --prod

# Usar una imagen base de Nginx para servir la aplicación
FROM nginx:1.21

# Copiar los archivos construidos desde la etapa anterior
COPY --from=build /app/dist/your-angular-app /usr/share/nginx/html

# Exponer el puerto 4200
EXPOSE 4200

# Comando por defecto para ejecutar Nginx
CMD ["nginx", "-g", "daemon off;"]
