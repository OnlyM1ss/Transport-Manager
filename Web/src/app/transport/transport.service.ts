import {Injectable} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { TransportGroupType } from '../models/transportGroupType';
import {Transport} from '../models/transport.model'
const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable()
export class TransportService {

  constructor(private authService: AuthService, private router: Router,private http:HttpClient) {}

  private transportUrl = 'http://localhost:801'; // TODO add in env

  public getTransports() {
    return this.http.get<any[]>(this.transportUrl + "/transport");
  }

  public getTransport(id: string) {
    return this.http.get(this.transportUrl + "/transport/"+ id);
  }

  public deleteTransport(transport : any) {
    return this.http.delete(this.transportUrl + "/transport/"+ transport.id);
  }

  public createTransport(transport : any) {
    return this.http.post(this.transportUrl + "/transport", transport);
  }
public searchTransportByType(type: string)  {
  return this.http.get<any[]>(this.transportUrl + "/transport/"+ type);
}
  public updateTransport(transport) {
    return this.http.put<Transport>(this.transportUrl + "/transport", transport);
  }
  public getGroupTypes()  {
    return this.http.get<TransportGroupType[]>(this.transportUrl + "/transport");
  }
}
