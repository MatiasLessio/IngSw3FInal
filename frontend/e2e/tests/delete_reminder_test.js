Feature('Home Page');

const loginSteps = require('./login_steps')
const addExampleReminder = require('./add_example_reminder')

Scenario('Delete Reminder', async ({ I }) => {
    loginSteps.login('test', '123');
    I.amOnPage('');

    I.see('Home / Reminders', 'h1');
    I.seeElement('.logout-container');

    const username = 'test';
    I.see(`Hi ${username}`, 'h2');
    await I.wait(1);
    addExampleReminder.addExampleReminder();

    await I.wait(1);
    await I.waitForElement('.alert-card', 10);

    I.click('.alert-card:last-child #deleteReminder');

    I.waitForElement('.swal2-popup', 15);
    await I.wait(1);
    I.click('.swal2-confirm');
    await I.wait(2);
    I.seeElement('.swal2-success');

    I.click('.swal2-confirm');

});
