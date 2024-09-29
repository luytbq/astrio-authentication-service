import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from '../../environments/environment.development';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(
    private http: HttpClient,
    private cookieService: CookieService,
  ) { }

  serviceUrl: string = environment.aas_service.url;

  public postLogin(body: any): Observable<HttpResponse<any>> {
    return this.http.post(this.serviceUrl+'/login', body, {observe: 'response'});
  }

  public saveToken(token: string | null) {
    if (!token) return;
    token = token.replaceAll('Bearer ', '');
    token && this.cookieService.set('auth', token, 7, "/");
  }
}
