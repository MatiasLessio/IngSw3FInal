Feature('Register')

Scenario('Test register', ({ I }) => {
    I.amOnPage('/Register');
    I.see('Register', 'mat-card-title');

    I.fillField('#usernameRegister', 'testuser');
    I.fillField('#passwordRegister', '123');

    I.click('#registerButton');

    I.waitForElement('mat-spinner', 5); 
    I.waitForInvisible('mat-spinner', 10); 


    I.waitForElement('.swal2-popup', 10); 

    I.see('User successfully created', '.swal2-title');
    I.seeElement('.swal2-success');

    I.click('.swal2-confirm');

    I.waitInUrl('/Login', 5);
    I.seeInCurrentUrl('/Login');
})