import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { TransportComponent } from './transport/transport.component';
import {TransportService} from './transport/transport.service';
import { AuthService } from './auth.service';
import { AuthInterceptorService } from './auth-interceptor.service';
import { CanActivateViaAuthGuard } from './can-activate-via-auth.guard';
import { AddTransportComponent } from './transport/add-transport.component';
import { DlDateTimeDateModule, DlDateTimePickerModule } from 'angular-bootstrap-datetimepicker';
import { UpdateTransportComponent } from './transport/update-transport.component';
const routes = [
    { path: 'login', component: LoginComponent },
    {
        path: 'transport',
        component: TransportComponent,
        canActivate: [
            CanActivateViaAuthGuard
        ]
    },
    { path: 'add', component: AddTransportComponent },
    { path: '', component: HomeComponent },
    { path: 'update/:id', component: UpdateTransportComponent },
];

@NgModule({
    declarations: [
        AppComponent,
        HomeComponent,
        LoginComponent,
        TransportComponent,
        AddTransportComponent,
        UpdateTransportComponent
    ],
    imports: [
        BrowserModule,
        FormsModule,
        ReactiveFormsModule,
        HttpClientModule,
        RouterModule.forRoot(routes),
        DlDateTimeDateModule,  // <--- Determines the data type of the model
        DlDateTimePickerModule,
    ],
    providers: [
        AuthService,
        {
            provide: HTTP_INTERCEPTORS,
            useClass: AuthInterceptorService,
            multi: true
        },
        CanActivateViaAuthGuard,
        TransportService,
        

    ],
    bootstrap: [AppComponent],
   
})
export class AppModule { }
