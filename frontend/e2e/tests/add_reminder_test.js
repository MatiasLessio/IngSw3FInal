Feature('Home Page');

const loginSteps = require('./login_steps');

Scenario('Add Reminder', async ({ I }) => {
    loginSteps.login('test', '123');
    I.amOnPage('');

    
    I.see('Home / Reminders', 'h1');
    I.seeElement('.logout-container');

    const username = 'test'; I.see(`Hi ${username}`, 'h2');

    I.click('#addReminder');

    const title = 'Test Reminder';
    const description = 'This is a test reminder';
    I.fillField('#title', title);
    I.fillField('#description', description);

    I.click('#saveReminder');

    I.waitForElement('.swal2-popup', 15); 

    I.seeElement('.swal2-success');

    I.click('.swal2-confirm');

    await I.wait(1);
    await I.waitForElement('.alert-card', 10);

    I.click('.alert-card:last-child #deleteReminder');

    I.waitForElement('.swal2-popup', 15);
    await I.wait(1);
    I.click('.swal2-confirm');
    await I.wait(1);
    I.seeElement('.swal2-success');

    I.click('.swal2-confirm');

    I.waitInUrl('', 15);
    await I.wait(1);
});
