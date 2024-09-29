import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { AuthService } from '../../../services/auth.service';

@Component({
  selector: 'aas-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit{
  constructor(
    private auth: AuthService
  ) {}

  formGroup = new FormGroup(
    {
      email: new FormControl('luytbq43@gmail.com'),
      password: new FormControl('password123')
    }
  );

  errorMsg = '';

  ngOnInit(): void {
    this.formGroup.controls.email.valueChanges.subscribe(value => this.clearError());
    this.formGroup.controls.password.valueChanges.subscribe(value => this.clearError());
  }
  clearError(): void {
    this.errorMsg = '';
  }

  submit() {
    this.auth.postLogin(this.formGroup.value).subscribe(
      res => {
        console.log('login', res)
        console.log(res.headers.get('Astrio-Auth-Token'))
        this.auth.saveToken(res.headers.get('Astrio-Auth-Token'));
      }, (error) => {
        this.errorMsg = error.error?.error?.message || 'something went wrong'
        console.error('login', error);
      }
    );
  }
}
