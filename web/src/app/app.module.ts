import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { BrowserModule } from '@angular/platform-browser';
import { LoginComponent } from './components/login/login/login.component';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms'
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';
import { RouterModule } from '@angular/router';
import { WelcomeComponent } from './components/welcome/welcome/welcome.component';

@NgModule({
    declarations: [
        AppComponent,
        LoginComponent,
        WelcomeComponent
    ],
    bootstrap: [AppComponent],
    imports: [
        AppRoutingModule,
        BrowserModule,
        ReactiveFormsModule,
        RouterModule,
    ],
    providers: [
        provideHttpClient(withInterceptorsFromDi()),
        CookieService
    ]
})
export class AppModule { }
