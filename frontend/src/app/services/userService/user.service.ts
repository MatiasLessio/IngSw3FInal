import { Injectable } from '@angular/core';
import { GlobalConstants } from '../../utilities/globalConstants';
import { Login } from '../../interfaces/login';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { Register } from '../../interfaces/register';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private loginEndpoint: string = GlobalConstants.ApiUrl+"Login";
  private registerEndpoint: string = GlobalConstants.ApiUrl+"Register";

  constructor(private http: HttpClient) { }

  Login(request : Login) :Observable<any>{
    console.log(this.loginEndpoint)
    return this.http.post(this.loginEndpoint, request);
  }
  
  Register(request : Register) :Observable<any>{
    return this.http.post(this.registerEndpoint, request);
  }

}
