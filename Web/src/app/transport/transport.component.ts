import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { TransportService } from './transport.service';
@Component({
    selector: 'app-members',
    templateUrl: './transport.component.html',
    styleUrls: ['./transport.component.scss']
})
export class TransportComponent implements OnInit {
    accountData: any;
    transports: any[];
    transport: any;
    typeSearch: string

    constructor(private authService: AuthService, private router: Router,private transportService: TransportService) { }
    ngOnInit() {
      this.transportService.getTransports().subscribe((data:any[]) => {
        this.transports = data;
      })
        this.authService.getAccount().subscribe(
            (res: any) => {
                this.accountData = res;
            }, (err: any) => {
                this.router.navigateByUrl('/login');
            }
        );
    }

    deleteTransport(transport: any): void {
      this.transportService.deleteTransport(transport)
        .subscribe( () => {
          this.transports = this.transports.filter((u: any) => u !== transport);
        })
    };
    createTransport(): void {
      this.transportService.createTransport(this.transport)
        .subscribe( () => {
          alert("Транспорт добавлен!.");
          this.router.navigate(['/transport']);
        })
    };
    searchTransportByType(transport: any): void {
      if (this.typeSearch != "") {
        this.transportService.searchTransportByType(this.typeSearch)
        .subscribe( (data:any[]) => {
          this.transports = data
        
        })
      }
      else {
        this.transportService.getTransports().subscribe((data:any[]) => {
          this.transports = data;
      })
    }
  };
}