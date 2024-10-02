import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { AuthService } from '../../../services/auth.service';
import { ActivatedRoute, Router } from '@angular/router';
import { first } from 'rxjs';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'aas-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  constructor(
    private auth: AuthService,
    private router: Router,
    private activatedRoute: ActivatedRoute,
  ) { }

  formGroup = new FormGroup(
    {
      email: new FormControl(''),
      password: new FormControl('')
    }
  );

  errorMsg = '';
  backUrl = '';

  ngOnInit(): void {
    this.formGroup.controls.email.valueChanges.subscribe(value => this.clearError());
    this.formGroup.controls.password.valueChanges.subscribe(value => this.clearError());

    this.activatedRoute.queryParams.pipe(first()).subscribe(params => {
      const pUrl = decodeURIComponent(params['back']); 
      if (this.isValidURL(pUrl)) {
        this.backUrl = pUrl;
      }
    })
  }

  clearError(): void {
    this.errorMsg = '';
  }

  submit() {
    this.auth.postLogin(this.formGroup.value).subscribe(
      res => {
        this.auth.saveToken(res.headers.get('Astrio-Auth-Token'));

        this.auth.user = {
          email: res.body.email
        };

        if (this.backUrl) {
          window.location.href = this.backUrl;
        } else {
          this.router.navigate(['welcome']);
        }
      }, (error: HttpErrorResponse) => {
        this.errorMsg = error.error?.error?.message || 'something went wrong'
      }
    );
  }

  isValidURL(url: string): boolean {
    try {
      new URL(url);
      return true;
    } catch (e) {
      return false;
    }
  }
}