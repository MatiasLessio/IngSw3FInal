exports.config = {
    tests: './e2e/tests/*_test.js',
    output: './e2e/output',
    helpers: {
      Puppeteer: {
        url: 'https://frontend-5ynfwgsmkq-uc.a.run.app',
        show: true,
        windowSize: '1920x1080'
      }
    },
    include: {
      I: './e2e/steps_file.js'
    },
    bootstrap: null,
    mocha: {},
    name: 'frontend'
  }
  