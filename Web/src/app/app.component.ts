import { Component, Injectable, OnInit } from '@angular/core';
import { AuthService } from './auth.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [AuthService]
})

export class AppComponent {
  title = 'Web';
  accountData: any;
  constructor(public authService: AuthService){}
}