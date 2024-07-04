Feature('Register')

Scenario('Test register', async ({ I }) => {
    I.amOnPage('/Register');
    I.see('Register', 'mat-card-title');

    const randomUsername = 'testuser' + Math.floor(Math.random() * 1000);
    I.fillField('#usernameRegister', randomUsername);
    I.fillField('#passwordRegister', '123');

    I.click('#registerButton');

    I.waitForElement('mat-spinner', 5); 
    I.waitForInvisible('mat-spinner', 10); 


    I.waitForElement('.swal2-popup', 10); 

    I.see('User successfully created', '.swal2-title');
    I.seeElement('.swal2-success');

    I.click('.swal2-confirm');

    I.waitInUrl('/Login', 10);
    I.seeInCurrentUrl('/Login');
});