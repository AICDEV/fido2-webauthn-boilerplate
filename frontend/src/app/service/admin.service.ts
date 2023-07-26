import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AppStateService } from './appstate.service';
import { Injectable } from '@angular/core';


@Injectable({
  providedIn: 'root',
})
export class AdminService {
  constructor(
    private httpClient: HttpClient,
    private appStateService: AppStateService,
  ) {}

  public fetchProtectedData(): Observable<{ name: string }> {
    const bearerToken = this.appStateService.getRawToken();

    if(!bearerToken) {
      throw new Error('No bearer token in state');
    }

    // sample service for protected resource
    return this.httpClient.get<{ name: string }>('https://fido.workshop/api/v1/member/name', {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.appStateService.getRawToken()}`,
      },
    });
  }
}
