import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { Transport } from '../models/transport.model';
import { TransportService } from './transport.service';
@Component({
  selector: 'app-transport-create',
  templateUrl: './add-transport.component.html'
})
export class AddTransportComponent {
  
  transport: Transport = new Transport();

  constructor(private router: Router, private transportService: TransportService) {

  }
  createTransport(): void {
    this.transportService.createTransport(this.transport)
        .subscribe( data => {
          alert("Транспорт добавлен!.");
          this.router.navigate(['/transport']);
        });

  };

}