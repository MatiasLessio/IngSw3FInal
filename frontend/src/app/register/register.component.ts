import { Component } from '@angular/core';
import { MaterialModule } from '../material.module';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Register } from '../interfaces/register';
import { UserService } from '../services/userService/user.service';
import Swal from 'sweetalert2'
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-register',
  standalone: true,
  imports: [MaterialModule, ReactiveFormsModule, CommonModule],
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent {
  loginForm : FormGroup = new FormGroup({
    username : new FormControl('',[Validators.required, Validators.maxLength(12)]),
    password : new FormControl('' , [Validators.required, Validators.minLength(3)])
  });
  constructor(private _service : UserService){}
  
  loadingRequest :boolean = false;
  register(){
    const registerRequest :Register = {
      password : this.loginForm.get('password')?.value,
      username: this.loginForm.get('username')?.value
    }
    this.loadingRequest = true;
    this._service.Register(registerRequest).subscribe((response : Register)=>{
      this.loadingRequest = false;
      
      Swal.fire({
        heightAuto: false,
        title: 'User successfully created',
        icon: 'success'
      }).then(()=>{
        window.location.href = '/Login';
      });
    }, (error)=>{
      this.loadingRequest = false;
      Swal.fire({
        heightAuto: false,
        title: 'Oops',
        text: 'Something wrong happened',
        icon: 'error'
      })
    })

  }
}
