import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { BrowserModule } from '@angular/platform-browser';
import { LoginComponent } from './components/login/login/login.component';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms'
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';

@NgModule({
    declarations: [
        AppComponent,
        LoginComponent
    ],
    bootstrap: [AppComponent],
    imports: [AppRoutingModule,
        BrowserModule,
        ReactiveFormsModule],
    providers: [
        provideHttpClient(withInterceptorsFromDi()),
        CookieService
    ]
})
export class AppModule { }
