// login_steps.js

const { I } = inject();

module.exports = {
  login(username, password) {
    I.amOnPage('/Login');
    I.see('Login', 'mat-card-title');

    I.fillField('#usernameLogin', username);
    I.fillField('#passwordLogin', password);

    I.click('#loginButton');

    I.waitForElement('mat-spinner', 10);
    I.waitForInvisible('mat-spinner', 20);
    
    I.waitForElement('.swal2-popup', 15);
    I.seeElement('.swal2-success');
    I.click('.swal2-confirm');

    I.waitInUrl('', 20);
    I.seeInCurrentUrl('');
  }
};