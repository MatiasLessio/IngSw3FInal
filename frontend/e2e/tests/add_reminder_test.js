Feature('Home Page');

const loginSteps = require('./login_steps');

Scenario('Add Reminder', async ({ I }) => {
    loginSteps.login('test', '123');
    // Abrir la página principal
    I.amOnPage('');

    // Verificar elementos visibles en la página principal
    I.see('Home / Reminders', 'h1');
    I.seeElement('.logout-container');

    // Simular inicio de sesión (suponiendo que ya esté logueado)
    const username = 'test'; // Puedes cambiar esto según tu configuración de usuario
    I.see(`Hi ${username}`, 'h2');

    // Hacer clic en el botón para agregar un nuevo recordatorio
    I.click('#addReminder');

    // Verificar que se haya abierto el diálogo de agregar recordatorio
    I.waitForVisible('.mat-dialog-content', 5); // Esperar a que el diálogo esté visible

    // Llenar el formulario del recordatorio
    const title = 'Test Reminder';
    const description = 'This is a test reminder';
    I.fillField('#title', title);
    I.fillField('#description', description);

    // Hacer clic en el botón de guardar en el diálogo
    I.click('#saveReminder');

    // Esperar a que el diálogo se cierre
    I.waitForInvisible('.mat-dialog-content', 5);

    // Verificar que el nuevo recordatorio aparezca en la lista de recordatorios
    I.waitForVisible('.alert-card', 5); // Esperar a que al menos un recordatorio esté visible

    // Verificar el contenido del último recordatorio agregado
    I.see(title, '.alert-header');
    I.see(description, '.alert-description');

    // Opcional: Realizar otras verificaciones como la eliminación de un recordatorio, etc.

    // Finalizar sesión simulada
    I.click('#logOut');

    // Verificar que se haya redirigido a la página de inicio de sesión
    I.waitInUrl('/login', 10);
    I.seeInCurrentUrl('/login');
});
