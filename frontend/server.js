const express = require('express');
const app = express();
const path = require('path');

// Configura Express para servir los archivos estáticos desde la carpeta "public"
app.use(express.static(path.join(__dirname, 'public')));

// Configura Express para manejar todas las rutas
app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, 'public/index.html'));
});

// Escucha en el puerto 4200
app.listen(4200, () => {
  console.log('Servidor Express escuchando en el puerto 4200');
});
