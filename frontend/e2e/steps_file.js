module.exports = function() {
    return actor({
  
      // Define custom steps here, use 'this' to access default methods of I.
      // It is recommended to place a general 'login' function here.
      login: function(username, password) {
        this.amOnPage('/Login');
        this.fillField('Username', username);
        this.fillField('Password', password);
        this.click('Login');
      }
    });
  }