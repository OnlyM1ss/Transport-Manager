import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { TransportService } from './transport.service';

@Component({
  selector: 'app-transport-update',
  templateUrl: './update-transport.component.html',
})
export class UpdateTransportComponent implements OnInit {

  transport: any = {};

  constructor(private router: Router, private route: ActivatedRoute, private transportService: TransportService) {
    
  }

  ngOnInit() {
    this.transportService.getTransport(this.route.snapshot.params['id'])
      .subscribe(data => {
        this.transport = data;
      });
  };

  updatetransport(): void {
    this.transportService.updateTransport(this.transport)
      .subscribe(data => {
        alert("transport updated successfully.");
        this.router.navigate(['/transport']);
      });

  };

}