import { Component } from '@angular/core';
import { MaterialModule } from '../material.module';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Login } from '../interfaces/login';
import { CommonModule } from '@angular/common';
import { UserService } from '../services/userService/user.service';
import { LoginResponse } from '../interfaces/login-response';
import Swal from 'sweetalert2';
import { GlobalConstants } from '../utilities/globalConstants';
@Component({
  selector: 'app-login',
  standalone: true,
  imports: [MaterialModule, ReactiveFormsModule, CommonModule],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  loginForm : FormGroup = new FormGroup({
    username : new FormControl('',[Validators.required, Validators.maxLength(12)]),
    password : new FormControl('' , [Validators.required, Validators.minLength(3)])
  });
  constructor(private _service : UserService){}
  
  loadingRequest :boolean = false;
  login(){
    const loginRequest :Login = {
      password : this.loginForm.get('password')?.value,
      username: this.loginForm.get('username')?.value
    }
    this.loadingRequest = true;
    this._service.Login(loginRequest).subscribe((response : LoginResponse)=>{
      this.loadingRequest = false;
      sessionStorage.setItem('token', response.token)
      sessionStorage.setItem('username', response.username)
      Swal.fire({
        heightAuto: false,
        title: 'Welcome Back ' + response.username,
        icon: 'success'
      }).then(()=>{
        window.location.href = '';
      });
    }, (error)=>{
      this.loadingRequest = false;
      Swal.fire({
        heightAuto: false,
        title: 'Oops',
        text: 'Wrong username or password',
        icon: 'error'
      })
    })

  }
}
