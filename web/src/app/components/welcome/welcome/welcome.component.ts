import { Component, OnInit } from '@angular/core';
import { AuthService, User } from 'src/app/services/auth.service';

@Component({
  selector: 'aas-welcome',
  templateUrl: './welcome.component.html',
  styleUrl: './welcome.component.scss'
})
export class WelcomeComponent implements OnInit {
  user: User = {};
  isLoggedIn = false;

  constructor(
    private auth: AuthService,
  ) {}

  ngOnInit(): void {
    this.user = this.auth.user;
    this.isLoggedIn = this.auth.isLoggedIn;
  }

}
