import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from '../../environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(
    private http: HttpClient,
  ) { }

  serviceUrl: string = environment.aas_service.url;

  public postLogin(body: any): Observable<any> {
    return this.http.post(this.serviceUrl+'/login', body, {observe: 'response'});
  }

  public saveToken(token: string) {
    
  }
}
