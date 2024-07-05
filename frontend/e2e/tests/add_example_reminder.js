const { I } = inject();

module.exports = {
    addExampleReminder() {
        I.click('#addReminder');
        const title = 'Test Reminder for update/delete';
        const description = 'This is a test reminder and it will be updated/deleted soon';
        I.fillField('#title', title);
        I.fillField('#description', description);

        I.click('#saveReminder');
        I.waitForElement('.swal2-popup', 15);

        I.seeElement('.swal2-success');

        I.click('.swal2-confirm');
    }
}