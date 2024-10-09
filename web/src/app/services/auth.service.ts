import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from '../../environments/environment';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(
    private http: HttpClient,
    private cookieService: CookieService,
  ) { }

  private serviceUrl: string = environment.aas.service_url;
  private _user: User = {
    email: GUEST
  };

  private _isloggedIn = false;

  get user(): User {
    return this._user;
  }

  set user(user: User) {
    this._user = user;
    this._isloggedIn = this._user.email != GUEST;
  }

  get isLoggedIn(): boolean {
    return this._isloggedIn
  }

  public postLogin(body: any): Observable<HttpResponse<any>> {
    return this.http.post(this.serviceUrl+'/users/login', body, {observe: 'response'});
  }

  public saveToken(token: string | null) {
    if (!token) return;
    token = token.replaceAll('Bearer ', '');
    token && this.cookieService.set('auth', token, 7, "/");
  }

}

export interface User {
  email?: string
}

export const GUEST = 'guest';