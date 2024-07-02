import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";

@Injectable({
  providedIn: 'root',
})
export class BaseService {
  constructor(private http: HttpClient) { }

  makeGetRequest<T>(url: string, options?: any): Observable<any> {
    options = this.getOptions(options);
    return this.http.get<T>(url, options);
  }

  makePostRequest<T>(url: string, body: any, options?: any): Observable<any> {
    options = this.getOptions(options);
    return this.http.post<T>(url, body, options);
  }

  makePutRequest<T>(url: string, body: any, options?: any): Observable<any> {
    options = this.getOptions(options);
    return this.http.put<T>(url, body, options);
  }

  makeDeleteRequest<T>(url: string, options?: any): Observable<any> {
    options = this.getOptions(options);
    return this.http.delete<T>(url, options);
  }

  getOptions(options?: any) {
    if (!(options && options.headers)) {
      options = { headers: new HttpHeaders() };
    }
    const token = sessionStorage.getItem('token');
    if (token) {
      options.headers = options.headers.set('Authorization', `Bearer ${token}`);
      options.headers = options.headers.set('Content-Type', 'application/json');
    }
    return options;
  }
}
