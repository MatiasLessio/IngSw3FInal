exports.config = {
    tests: './tests/*_test.js',
    output: './output',
    helpers: {
        Playwright: {
            url: 'https://frontend-5ynfwgsmkq-uc.a.run.app', // Cambia la URL base según tu entorno
            show: true, // Opcional: muestra el navegador durante las pruebas
            browser: 'chromium' // Puedes cambiar a 'firefox' o 'webkit' según tus necesidades
        }
    },
    include: {
        I: './steps_file.js'
    },
    bootstrap: null,
    mocha: {},
    name: 'frontend'
}
